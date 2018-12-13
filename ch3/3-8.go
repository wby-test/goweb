package main 

import (
	"net/http"
	"fmt"
)

func hello(r http.ResponseWriter, w *http.Request) {
	fmt.Fprintf(r, "hello")
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