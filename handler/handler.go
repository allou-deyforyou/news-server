package handler

import (
	"net/http"
	"time"

	"news/internal/store"
)

const timeout =  time.Minute;

type Handler struct {
	*store.Client
	*http.ServeMux
}

func NewHandler(client *store.Client) *Handler {
	return &Handler{
		ServeMux: http.NewServeMux(),
		Client:   client,
	}
}

func ParseHandler(h http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		h.ServeHTTP(w, r)
	})
}
