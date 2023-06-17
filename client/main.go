package main

import (
	"client/service"
	"log"
)

func main() {
	log.Println("Service sender start")
	service.Run()
	log.Println("Service sender end")
}
