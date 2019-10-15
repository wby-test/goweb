package main

import juu   a

func process(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "(1)", r.FormValue("hello"))
	fmt.Fprintln(w, "(2)", r.PostFormValue("hello"))
	fmt.Fprintln(w, "(1)", r.PostForm)
	fmt.Fprintln(w, "(1)", r.MultipartForm)
}

func main()  {
	server := http.Server{
		Addr:"127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}