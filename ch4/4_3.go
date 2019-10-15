package main

import (
	"fmt"
	"net/http"
)


func body(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		length := r.ContentLength
		body := make([]byte, length)
		r.Body.Read(body)
		fmt.Fprintf(w, string(body))
	}
}

func main() {
	server := http.Server{
		Addr:"127.0.0.1:8080",
	}

	http.HandleFunc("/body", body)
	server.ListenAndServe()

}