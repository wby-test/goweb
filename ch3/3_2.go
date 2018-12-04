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
	/*
	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: nil,
	}

	server.ListenAndServe()
	*/
	handler := MyHandle{}
	server := http.Server{
		Addr:"127.0.0.1:12345",
		Handler:&handler,
	}

	server.ListenAndServe()
}
