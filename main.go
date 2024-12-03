package main

import (
	"math/rand/v2"
	"net/http"

	"github.com/barealek/komma/pkg/mwm"
)

func authmiddleware(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// auth logic
		if rand.Float32() < 0.5 {
			http.Error(w, "not allowed nig", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func block(_ http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// auth logic
		http.Error(w, "blocked", http.StatusUnauthorized)
	})
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}

func main() {
	n := http.NewServeMux()
	d := mwm.Mwm(n)

	d.HandleFunc("/", root, authmiddleware, block)

	http.ListenAndServe(":8080", d)
}
