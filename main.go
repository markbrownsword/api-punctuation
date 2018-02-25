package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health/ping", ping).Methods("GET")

	http.ListenAndServe(":8080", r)
}
