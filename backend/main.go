package main

import (
	socket "backend/appsocket"
	"backend/constants"
	"backend/redisdb"
	"backend/router"
	"log"
	d "runtime/debug"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in main\n", string(d.Stack()))
		}
		if constants.SocketInitialised {
			socket.Close()
		}
		if constants.DBInitialised {
			if dbCloseErr := redisdb.Close(); dbCloseErr != nil {
				log.Println("db close err:", dbCloseErr)
			}
		}
	}()
	redisdb.Init()
	socket.Init()
	router.Init()
}
