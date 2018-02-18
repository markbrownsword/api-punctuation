package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

//	http://localhost:8080/?name=mark
func hello(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	param := query.Get("name")
	result := fmt.Sprintf("%v %v", "hello", param)

	fmt.Fprint(w, result)
}
