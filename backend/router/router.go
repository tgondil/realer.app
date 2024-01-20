package router

import (
	"backend/appmiddleware"
	"backend/request_handler/auth_handler"
	"log"
	//"backend/auth0"
	"backend/request_handler"
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// How to use: if there is an error, do not write to response writer, just return the error
func CustomHandler(f func(http.ResponseWriter, *http.Request) (e error, statusCode int)) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err, statusCode := f(writer, request)
		if err != nil {
			log.Println("error in ", request.RequestURI, err)
			http.Error(writer, err.Error(), statusCode)
		}
	}
}

func Init() {
	app := chi.NewRouter()

	app.Use(middleware.Logger)
	app.Use(middleware.Recoverer)
	app.Use(cors.Handler(cors.Options{
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions},
		AllowedHeaders: []string{"*"},
	}))
	app.Use(middleware.Timeout(20 * time.Second))
	app.Use(middleware.RequestSize(55 << 20)) // 55 MB

	app.Get("/", CustomHandler(func(writer http.ResponseWriter, request *http.Request) (error, int) {
		//request.URL.Host = strings.Split((r.URL.Host), ".")[0]
		writer.Header().Set("Content-Type", "text/html")
		_, err := writer.Write([]byte("<CENTER><H1>Hello world</H1></CENTER>"))
		return err, 400
	}))

	/*app.Get("/getCurrentUTCTime", CustomHandler(func(writer http.ResponseWriter, request *http.Request) (error, int) {
		_, err := writer.Write([]byte(time.Now().UTC().Format(time.DateTime)))
		return err,400
	}))*/

	app.Get("/files/+", CustomHandler(func(writer http.ResponseWriter, request *http.Request) (error, int) {
		return request_handler.GetFiles(writer, request)
	}))

	app.Post("/login", CustomHandler(auth_handler.Login))
	app.Post("/signUp", CustomHandler(auth_handler.Signup))

	//user auth middleware
	app.Group(func(r chi.Router) {
		r.Use(appmiddleware.AuthMiddleware)
		r.Post("/chatMessages/:otherPersonID", CustomHandler(request_handler.GetSingleChatMessages))
		r.Post("/chats", CustomHandler(request_handler.GetChats))
		r.Post("/sendMessageWithFile", CustomHandler(request_handler.SendMessageWithFile))
		r.Post("/addReaction/:messageID", CustomHandler(request_handler.AddReactionToMessage))
	})

	server := &http.Server{
		Addr:              ":8080",
		Handler:           app,
		ReadHeaderTimeout: time.Second * 5,
		WriteTimeout:      0,
		IdleTimeout:       0,
	}
	serverCtx, serverCancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c,
		syscall.SIGINT,
		syscall.SIGKILL,
		syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT,
	)
	go func() {
		<-c
		shutdownCtx, secondCancel := context.WithTimeout(serverCtx, time.Second*5)
		go func() {
			<-shutdownCtx.Done()
			if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
				log.Fatalln("server shutdown timed out")
			}
		}()

		shutdownErr := server.Shutdown(shutdownCtx)
		if shutdownErr != nil {
			log.Fatalln("server shutdown error:", shutdownErr)
		}
		secondCancel()
		serverCancel()
		log.Println("server shutdown complete")
	}()

	const startupPrint = `------------------------------------------------------------------------------------------------------------
	Starting server...
------------------------------------------------------------------------------------------------------------`

	fmt.Println(startupPrint)

	if listenErr := server.ListenAndServe(); listenErr != nil && !errors.Is(listenErr, http.ErrServerClosed) {
		log.Panicln(listenErr)
	}
	<-serverCtx.Done()
}
