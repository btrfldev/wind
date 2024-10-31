package main

import (
	"github.com/btrfldev/wind/component"
	//"github.com/btrfldev/wind/storage"
)

type Server struct {
	Port   string
	//Memory storage.MemoryStore[string, component.Component]
	ComponentStorage component.ComponentStorage
}

func NewServer(port string) *Server {
	return &Server{
		Port:   port,
		//Memory: storage.MemoryStore[string, component.Component]{},
		ComponentStorage: *component.NewComponentStorage(),
	}
}

