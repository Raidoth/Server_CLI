package main

import (
	"log"
	"server/service"
)

func main() {

	log.Println("Service start")
	service.Run()
	log.Println("Service end")
}
