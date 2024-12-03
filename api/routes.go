package api

import (
	"net/http"

	"github.com/barealek/komma/pkg/mwm"
)

func (a *Api) RegisterRoutes() http.Handler {
	r := mwm.Mwm(http.NewServeMux())

	r.HandleFunc("GET /", a.HandleRoot)

	return r
}
