package main

import (
	"html/template"
	"net/http"
	"time"
)

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func bindFunc(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate" : formatDate}
	t := template.New("tmplTime.html").Funcs(funcMap)
	t, _ = t.ParseFiles("src/goweb/ch5/tmplTime.html")
	t.Execute(w, time.Now())
}

func main() {
	server := http.Server{
		Addr:"127.0.0.1:8080",
	}

	http.HandleFunc("/bindFunc", bindFunc)
	server.ListenAndServe()
}