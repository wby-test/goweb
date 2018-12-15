package main

import (
	"html/template"
	"net/http"
)

func contextAware(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("src/goweb/ch5/aware.html")
	content := `I asked: <i> "what is up?" </i>`
	t.Execute(w, content)
}
func main() {
	server := http.Server{
		Addr:"127.0.0.1:8080",
	}

	http.HandleFunc("/aware", contextAware)
	server.ListenAndServe()
}