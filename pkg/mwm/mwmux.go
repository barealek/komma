package mwm

import (
	"net/http"
	"slices"
)

type mwm struct {
	*http.ServeMux
}

func (m *mwm) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request), mws ...func(http.HandlerFunc) http.HandlerFunc) {
	var h func(http.ResponseWriter, *http.Request) = handler

	slices.Reverse(mws)

	for _, mw := range mws {
		h = mw(h)
	}

	m.ServeMux.HandleFunc(pattern, h)
}

func Mwm(mux *http.ServeMux) mwm {
	return mwm{ServeMux: mux}
}
