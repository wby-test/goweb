package main

import (
	"fmt"
	"net/http"
)
//log 将真正的逻辑处理函数wrap在中间，记录程序执行过程中的日志

type HelloHandler struct{}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello" )
}

func log (h http.Handler) http.Handler {
	return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("handler called - %T\n", h)
		h.ServeHTTP(w, r)
	})
}

func protected(h http.Handler) http.Handler {
	return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {
		//TODO: 完成登陆验证
		h.ServeHTTP(w, r)
	})
}

func main() {
	server := http.Server{
		Addr:"127.0.0.1:8080",
	}
	hello := HelloHandler{}
	http.Handle("/hello", protected(log(hello)))
	server.ListenAndServe()
}