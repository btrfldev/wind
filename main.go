package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Wind is starting...")
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", ping)
	mux.HandleFunc("GET /call", call)

	fmt.Println("API is running on 0.0.0.0:2358")
	if err := http.ListenAndServe(":2358", mux); err != nil {
		panic(err)
	}
}

func call(w http.ResponseWriter, r *http.Request) {
	r.URL.Query().Get("mod")
	fmt.Println(r.URL.Path)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
