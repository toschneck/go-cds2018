package main

import (
	"net/http"
	"log"
	"github.com/toschneck/go-cds2018/internal/router"
	"os"
)

func main() {
	log.Printf("Booting ...")

	appPort := os.Getenv("PORT")
	if (len(appPort) == 0){
		appPort = ":8080"
	}

	log.Fatal(http.ListenAndServe(appPort, router.NewBLRouter()));
}

