package main

import (
	"net/http"
	"log"
	"github.com/toschneck/go-cds2018/internal/router"
	"os"
	"github.com/toschneck/go-cds2018/internal/version"
	"os/signal"
	"context"
	"time"
	"syscall"
)

func main() {
	log.Printf("Booting ...")
	log.Printf("Version: %v, BuildTime: %v, Commit: %v", version.Release, version.BuildTime, version.Commit)
	appPort := os.Getenv("PORT")
	if len(appPort) == 0 {
		appPort = "8080"
	}
	diagPort := os.Getenv("DIAGNOSE_PORT")
	if len(diagPort) == 0 {
		diagPort = "8088"
	}

	var blSever, diagServer http.Server
	srvErrors := make(chan error, 2)
	go func() {
		log.Printf("use app port: %v", appPort)
		r := router.NewBLRouter()
		blSever = http.Server{
			Addr:    ":" + appPort,
			Handler: r,
		}
		srvErrors <- blSever.ListenAndServe()
	}()

	go func() {
		log.Printf("use diagnostic port: %v", diagPort)
		r := router.NewDiagnoticRouter()
		diagServer = http.Server{
			Addr:    ":" + diagPort,
			Handler: r,
		}
		srvErrors <- diagServer.ListenAndServe()
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGINT)

	select {
	case killSignal := <-interrupt:
		log.Printf("Got %s signal", killSignal)
	case serverErr := <-srvErrors:
		log.Printf("Get server error: %v", serverErr)
	}
	{
		ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()
		log.Println("shutdwon app server")
		blSever.Shutdown(ctx)
	}
	{
		ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()
		log.Println("shutdwon diagnostic server")
		diagServer.Shutdown(ctx)

	}

}
