package main

import (
	"net/http"
	"log"
	"github.com/toschneck/go-cds2018/internal/router"
)

func main() {
	log.Printf("Booting ...")

	log.Fatal(http.ListenAndServe(":8080", router.NewBLRouter()));
}

