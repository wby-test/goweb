package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)


//权限密码验证
func helloMutex(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello, %s %s\n", p.ByName("name"), p.ByName("password"))
}

func main() {
	mux := httprouter.New()
	mux.GET("/hello/:name/:password", helloMutex)

	server := http.Server{
		Addr:"127.0.0.1:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
