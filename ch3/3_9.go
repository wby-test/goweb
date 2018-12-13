package main 

import (
	"reflect"
	"runtime"
	"net/http"
	"fmt"
)

//日志记录， 安全检查， 错误处理 ，使用串联技术处理分离代码中的横切关注点

func Hello(r http.ResponseWriter, w *http.Request) {
	fmt.Fprintf(r, "hello")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("handle func called - " + name)
		h(w, r)
	}
}

func main() {
	server := http.Server {
		Addr : "localhost:8080",
	}
	http.HandleFunc("/hello", log(Hello))
	server.ListenAndServe()
}