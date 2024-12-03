package api

import "time"

type Api struct {
	version string
	startup time.Time
}

func NewAPI(v string) *Api {
	return &Api{
		version: v,
		startup: time.Now(),
	}
}
