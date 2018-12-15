package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func randTime(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("/project/goProject/src/goweb/ch5/tmpl1.html"))
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}

func dealRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("/project/goProject/src/goweb/ch5/tmpl2.html"))
	daysOfWeek := []string{}//{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	t.Execute(w, daysOfWeek)
}

func setAction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("/project/goProject/src/goweb/ch5/tmpl3.html"))
	t.Execute(w, "hhhhh")
}

//include action
func include(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("/project/goProject/src/goweb/ch5/tmpl4.html", "/project/goProject/src/goweb/ch5/tmpl5.html"))
	t.Execute(w, "hello")
}

func main() {
	serve := http.Server{
		Addr:"127.0.0.1:8080",
	}

	http.HandleFunc("/randTime", randTime)
	http.HandleFunc("/range", dealRange)
	http.HandleFunc("/setAction", setAction)
	http.HandleFunc("/include", include)

	serve.ListenAndServe()
}