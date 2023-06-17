package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"server/service/server/handlefunc"
	"server/service/server/logs"
	"syscall"
)

const (
	apiSubString = "/api/substring"
	port         = 9000
)

func Run() {
	server := &http.Server{
		Addr: fmt.Sprintf(":%v", port),
	}

	http.HandleFunc(apiSubString, logs.LoggingHTTP(http.HandlerFunc(handlefunc.SubString)))

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка запуска сервера: %v", err)
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Ошибка остановки сервера: %v", err)
	}

}
