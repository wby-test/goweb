package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Method
	fmt.Fprintln(w, h)
	fmt.Fprintln(w, r.Header)
	fmt.Fprintln(w, r.Body)
	fmt.Fprintln(w, r.Host)
}

func main() {
	server := http.Server{
		Addr:              "127.0.0.1:9090",
		Handler:           nil,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	http.HandleFunc("/headers", headers)
	server.ListenAndServe()
}
