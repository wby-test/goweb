package main

import (
	"fmt"
	"net/http"
)

func hello(r http.ResponseWriter, w *http.Request) {
	h := w.Header["Connection"]
	t := w.Header.Get("Connection")
	fmt.Fprintf(r, "hello")
	fmt.Fprintln(r)
	fmt.Fprintln(r, t)
	fmt.Fprintln(r)
	fmt.Fprintln(r, h)
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

	server.ListenAndServe()
}

