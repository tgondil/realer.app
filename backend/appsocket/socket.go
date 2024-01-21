package appsocket

import (
	"backend/constants"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zishang520/engine.io/v2/types"
	"github.com/zishang520/socket.io/v2/socket"
	"log"
	"strings"
)

var io *socket.Server
var httpServer *types.HttpServer

func Init() {
	const prefix = "Socket"
	httpServer = types.NewWebServer(nil)
	io = socket.NewServer(httpServer, nil)

	_ = io.On("connection", func(clients ...any) {
		defer func() {
			if r := recover(); r != nil {
				log.Println("Socket recover", r)
			}
		}()
		client := clients[0].(*socket.Socket)
		_ = client.On("join_subscription", func(data ...any) {
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
			_, containsPersonId := payload["personId"]

			ok = ok && containsPersonId

			if !(ok && token.Valid) {
				log.Println(prefix, "Invalid token")
				leaveAllAndDisconnect(client)
				return
			}

			room := fmt.Sprint(payload["personId"])
			client.Join(socket.Room(room))
		})
		_ = client.On("disconnect", func(...any) {
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
