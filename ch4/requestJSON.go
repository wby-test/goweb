package main

import (
	json2 "encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	User string
	Threads []string
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>Go Web Programming</title></head>
<body><h1>Hello World</h1></body>
</html>`
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service, try next door")
}

//是不是这样带出cookie的?
func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://baidu.com")
	w.WriteHeader(302)
}

//write json
func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:"wbb",
		Threads:[]string{"fk", "hhh", "biubiubiu"},
	}

	json, _ := json2.Marshal(post)
	w.Write(json)
}

func main() {
	server := http.Server{
		Addr:"127.0.0.1:8080",
	}

	http.HandleFunc("/writeExample", writeExample)
	http.HandleFunc("/writeHeader", writeHeaderExample)
	http.HandleFunc("/header", headerExample)
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}

