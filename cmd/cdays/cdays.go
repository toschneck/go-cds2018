package main

import (
	"net/http"
	"log"
	"github.com/toschneck/go-cds2018/internal/router"
	"os"
	"github.com/toschneck/go-cds2018/internal/version"
)

func main() {
	log.Printf("Booting ...")
	log.Printf("Version: %v, BuildTime: %v, Commit: %v", version.Release, version.BuildTime, version.Commit)

	appPort := os.Getenv("PORT")
	if len(appPort) == 0 {
		appPort = ":8080"
	}
	log.Printf("use port%v", appPort)
	log.Fatal(http.ListenAndServe(appPort, router.NewBLRouter()));
}
