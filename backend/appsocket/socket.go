package appsocket

import (
	"backend/constants"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zishang520/engine.io/v2/types"
	"github.com/zishang520/socket.io/v2/socket"
	"log"
	"net/http"
	"strings"
)

var io *socket.Server
var httpServer *types.HttpServer

func Init() {
	const prefix = "Socket"
	app := chi.NewRouter()
	app.Use(cors.Handler(cors.Options{
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions},
		AllowedHeaders: []string{"*"},
	}))

	httpServer = types.NewWebServer(app)
	io = socket.NewServer(httpServer, nil)

	_ = io.On("connection", func(clients ...any) {
		log.Println(prefix, "connection", clients)
		defer func() {
			if r := recover(); r != nil {
				log.Println("Socket recover", r)
			}
		}()
		client := clients[0].(*socket.Socket)
		_ = client.On("join_subscription", func(data ...any) {
			log.Println(prefix, "join_subscription", data)
			if len(data) != 1 {
				leaveAllAndDisconnect(client)
				return
			}
			d, ok := (data[0]).(string)
			if !ok {
				return
			}
			defer func() {
				if r := recover(); r != nil {
					log.Println(prefix, "recovering in connection method", r)
				}
			}()
			token, err := jwt.Parse(
				strings.TrimSpace(d),
				func(token *jwt.Token) (interface{}, error) {
					if _, ok = token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, errors.New(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
					}
					return []byte("boilermake"), nil
				},
			)

			if err != nil && !errors.Is(err, jwt.ErrTokenExpired) {
				log.Println(prefix, err)
				leaveAllAndDisconnect(client)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if issuer, err := claims.GetIssuer(); err != nil || issuer != "Boilermake" {
				log.Println(prefix, err)
				leaveAllAndDisconnect(client)
				return
			}
			payload := claims["pld"].(map[string]any)
			personIDAny, containsPersonId := payload["personId"]

			personIDFloat, containsPersonId := personIDAny.(float64)

			ok = ok && containsPersonId

			if !(ok && token.Valid) {
				log.Println(prefix, "Invalid token")
				leaveAllAndDisconnect(client)
				return
			}

			room := fmt.Sprint(int64(personIDFloat))
			log.Println(prefix, "joining room", room)
			client.Join(socket.Room(room))
		})
		_ = client.On("disconnect", func(...any) {
			log.Println(prefix, "disconnect")
			leaveAllAndDisconnect(client)
		})
	})
	fmt.Println("socket http started")
	httpServer.Listen(":4000", nil)
	constants.SocketInitialised = true
}

func Close() {
	log.Println("closing socket")
	io.Close(nil)
	if err := httpServer.Close(nil); err != nil {
		log.Println("err in closing socket", err)
	}
	log.Println("closed socket")
}

func Broadcast(personIDs []int64, endPoint string, message ...any) {
	if len(personIDs) == 0 {
		io.Emit(endPoint, message)
		return
	}
	for _, room := range personIDs {
		if err := io.To(socket.Room(fmt.Sprint(room))).Emit(endPoint, message); err != nil {
			log.Println("Socket broadcast to", room, err)
		}
	}
}

func leaveAllAndDisconnect(client *socket.Socket) {
	for _, r := range client.Rooms().Keys() {
		client.Leave(r)
	}
	client.Disconnect(true)
}
