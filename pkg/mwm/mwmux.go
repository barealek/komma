package mwm

import (
	"net/http"
	"slices"
)

type mwm struct {
	*http.ServeMux
}

// HandleFunc registers the handler function for the given pattern with the provided middleware functions.
// The middleware functions are applied in the order they are provided.
//
// pattern: The URL pattern to match.
//
// handler: The http.HandlerFunc to handle the request.
//
// mws: A variadic list of middleware functions to apply to the handler. The first element of the list will be the first invoked middleware.
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
