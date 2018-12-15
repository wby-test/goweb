package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("hahahhahhah")
	c := http.Cookie{
		Name:"flash",
		Value:base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
}

func showMessgae(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "No Message found")
		}
	} else {
		rc := http.Cookie{
			Name : "flash",
			MaxAge:-1,
			Expires:time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val,_ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
}

func main() {
	serve := http.Server{
		Addr:"127.0.0.1:8080",
	}

	http.HandleFunc("/setMessage", setMessage)
	http.HandleFunc("/showMessage", showMessgae)

	serve.ListenAndServe()
}