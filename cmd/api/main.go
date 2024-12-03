package main

import (
	"log"

	"github.com/barealek/komma/config"
	"github.com/barealek/komma/pkg/must"
	"github.com/barealek/komma/server"
)

func main() {
	// ctx := context.Background()
	cfg := must.Must(config.NewConfig())
	s := server.NewServer(cfg)

	if err := s.ListenAndServe(); err != nil {
		log.Println("error starting server")
	}
}
