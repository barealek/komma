package server

import (
	"fmt"
	"net/http"
	"strings"

	_ "embed"

	"github.com/barealek/komma/api"
	"github.com/barealek/komma/config"
)

//go:embed VERSION
var version string

type Server struct{}

func NewServer(cfg *config.Config) *http.Server {
	api := api.NewAPI(strings.TrimSuffix(version, "\r\n"))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: api.RegisterRoutes(),
	}

	return server
}
