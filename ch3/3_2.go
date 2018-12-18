package main

import (
	"fmt"
	"net/http"
)

type MyHandle struct {}

func (h *MyHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "just do it")
}

func main() {
	handler := MyHandle{}
	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: &handler,
	}

	server.ListenAndServe()
}
