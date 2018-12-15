package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	fileHeader := r.MultipartForm.File["uploaded"][0]
	file, err := fileHeader.Open()
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func process(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("uploaded")
	if err == nil {
		date, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(date))
		}
	}
	fmt.Fprintln(w, "hhhhh")
}
func main() {
	server := http.Server{
		Addr:"127.0.0.1:8080",
	}

	http.HandleFunc("/upload", upload)
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
