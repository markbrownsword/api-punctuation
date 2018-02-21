package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", hello).Methods("GET")

	http.ListenAndServe(":8080", r)
}

//	http://localhost:8080/?name=mark
func hello(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	param := query.Get("name")
	result := fmt.Sprintf("%v %v", "hello", param)

	fmt.Fprint(w, result)
}
