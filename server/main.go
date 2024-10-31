package main

import (
	"context"
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"time"

	"github.com/btrfldev/wind/config"
	"github.com/tetratelabs/wazero"
)

//go:embed registry.html
var RegistryPage string

func main() {
	fmt.Println("Wind is starting...")

	cfg := config.GetConfig()

	s := NewServer(cfg.Port)

	fmt.Println("API is running on 0.0.0.0:" + cfg.Port)
	if err := s.Start(); err != nil {
		panic(err)
	}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", s.ping)
	mux.HandleFunc("GET /call", s.Call)
	mux.HandleFunc("/registry", s.Registry)
	mux.HandleFunc("/ui/registry", s.UIRegistry)

	return http.ListenAndServe(":"+s.Port, mux)
}

func (s *Server) Call(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	compName := query.Get("comp")
	if compName == "" {
		http.Error(w, "empty comp", http.StatusNotFound)
		return
	}
	fmt.Println(compName)

	comp, err := s.ComponentStorage.Get(compName) //Memory.Get(compName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	start := time.Now()
	//res, err :=run.Invoke(comp, "./components/"+comp+".wasm", map[string]string{})
	res, err := comp.Invoke(map[string]string{})
	finish := time.Now()
	duration := finish.Sub(start)
	fmt.Printf("Invoked in: %v\n", duration.Seconds())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Write([]byte(res))
	}

}

func (s *Server) Registry(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		compFileForm, _, err := r.FormFile("compFile")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		compName := r.FormValue("compName")
		if compName == "" {
			http.Error(w, "No name Component!", http.StatusBadRequest)
			return
		}

		defer compFileForm.Close()

		compFile, err := io.ReadAll(compFileForm)
		if err != nil {
			http.Error(w, "can`t read file", http.StatusInternalServerError)
			return
		}

		start := time.Now()
		err = s.ComponentStorage.Register(wazero.NewRuntime(context.Background()), compName, compFile, r.Context())
		finish := time.Now()
		duration := finish.Sub(start)
		fmt.Printf("Compiled in: %v\n", duration.Seconds())
		if err != nil {
			http.Error(w, "can`t register component", http.StatusInternalServerError)
			return
		} else {
			w.Write([]byte(compName))
		}
	}
}

func (s *Server) UIRegistry(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.New("registry").Parse(RegistryPage)
	if err != nil {
		http.Error(w, "can`t parse html", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "can`t render html", http.StatusInternalServerError)
		return
	}

}

func (s *Server) ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
