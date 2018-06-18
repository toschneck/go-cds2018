package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func main() {
	log.Printf("Booting ...")

	r := mux.NewRouter()
	r.HandleFunc("/home", rootHandler())


	log.Fatal(http.ListenAndServe(":8080", r));
}

func rootHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Helloooo"))
	}
}
