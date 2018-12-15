package main

import (
	"net/http"
	"html/template"
)

func process(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("/project/goProject/src/goweb/ch5/tmpl.html"))
	t.Execute(w, "wby hhhh")
}

func main() {
	serve := http.Server{
		Addr:"127.0.0.1:8080",
	}

	http.HandleFunc("/process", process)

	serve.ListenAndServe()
}