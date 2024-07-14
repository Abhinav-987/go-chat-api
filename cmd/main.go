package main

import (
	"log"

	"github.com/Abhinav-987/go-chat/db"
	"github.com/Abhinav-987/go-chat/internal/user"
	"github.com/Abhinav-987/go-chat/internal/ws"
	"github.com/Abhinav-987/go-chat/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	userRp := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRp)
	userHandler := user.NewHandler(userSvc)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	router.Start("0.0.0.0:8080")

}
