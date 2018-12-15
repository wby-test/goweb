package main

import (
	"fmt"
	"net/http"
)

//cookie.go
//set cookie
func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie {
		Name:"first_cookie",
		Value:"wby",
		HttpOnly:true,
	}

	c2 := http.Cookie{
		Name:"second_cookie",
		Value:"wyr",
		HttpOnly:true,
	}

	//w.Header().Set("Set-Cookie", c1.String())
	//w.Header().Add("Set-Cookie", c2.String())
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	h := r.Header["Cookie"]
	fmt.Fprintln(w, h)
}

func main() {
	server := http.Server{
		Addr:"127.0.0.1:8080",
	}

	http.HandleFunc("/cookie", setCookie)
	http.HandleFunc("/getCookie", getCookie)
	server.ListenAndServe()
}