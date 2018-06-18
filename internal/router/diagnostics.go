package router

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

func NewDiagnoticRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/healthz	", handleOK())
	r.HandleFunc("/readyz	", handleOK())
	return r
}
func handleOK() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, http.StatusText(http.StatusOK))
	}
}
