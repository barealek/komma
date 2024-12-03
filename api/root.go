package api

import (
	"encoding/json"
	"net/http"
	"time"
)

func (a *Api) HandleRoot(w http.ResponseWriter, r *http.Request) {
	rsp := map[string]any{
		"version": a.version,
		"uptime":  time.Since(a.startup).Seconds(),
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rsp)
}
