package main

import (
	"fmt"
	"net/http"
)

func processes(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
	fmt.Fprintln(w, r.Form["hello"])
	fmt.Fprintln(w, r.Form["thread"])
	fmt.Fprintln(w, r.Form["post"])
}
func main()  {
	server := http.Server{
		Addr:"127.0.0.1:8080",
	}
	http.HandleFunc("/process", processes)
	server.ListenAndServe()
}