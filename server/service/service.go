package service

import (
	"log"
	"server/service/server"
)

func Run() {
	log.Println("Server start")
	server.Run()
	log.Println("Server end")
}
