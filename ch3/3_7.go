package main

import (
	"fmt"
	"net/http"
)

//将一个handle注册进入server中，
//1、在结构体中直接定义
//2、使用Handle函数
type jsdHandler struct {}

type jsdtHandler struct {}

func (h *jsdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "just do it:wby")
}

func (h *jsdtHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "just do it too:wby")
}

func main() {
	jsd := jsdHandler{}
	jsdt := jsdtHandler{}

	server := http.Server{
		Addr:"127.0.0.1:8080",
	}

	http.Handle("/jsd", &jsd)
	http.Handle("/jsdt", &jsdt)

	server.ListenAndServe()
}
