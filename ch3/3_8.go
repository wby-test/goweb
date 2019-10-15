package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	h := r.Header["Connection"]
	t := r.Header.Get("Connection")
	fmt.Fprintf(w, "hello")
	fmt.Fprintln(w)
	fmt.Fprintln(w, t)
	fmt.Fprintln(w)
	fmt.Fprintln(w, h)
}

func test(w http.ResponseWriter, r *http.Request) {
	h := r.Header["User-Agent"]
	fmt.Fprintln(w, h)
	fmt.Fprintln(w, len(r.Header))
}

func world(r http.ResponseWriter, w *http.Request) {
	fmt.Fprintf(r, "world")
}

//handle 将处理器函数转换为处理器
func main() {
	server := http.Server {
		Addr : "127.0.0.1:8080",
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)
	http.HandleFunc("/test", test)

	server.ListenAndServe()
}