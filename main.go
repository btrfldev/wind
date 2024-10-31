package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/btrfldev/wind/config"
	"github.com/btrfldev/wind/component/run"
)

func main() {
	fmt.Println("Wind is starting...")

	cfg := config.GetConfig()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", ping)
	mux.HandleFunc("GET /call", call)

	fmt.Println("API is running on 0.0.0.0:" + cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, mux); err != nil {
		panic(err)
	}
}

func call(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	comp := query.Get("comp")
	//fmt.Println(r.URL.Path, query)
	start := time.Now()
	res, err := run.Invoke(comp, "./components/"+comp+".wasm", map[string]string{})
	finish := time.Now()
	duration := finish.Sub(start)
	fmt.Printf("duration.Seconds(): %v\n", duration.Seconds())
	if err != nil {
		w.Write([]byte("Not Found"))
	} else {
		w.Write([]byte(res))
	}
	
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
