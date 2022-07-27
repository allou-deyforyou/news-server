package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	"news/internal/storage"
	"news/internal/storage/migrate"

	"entgo.io/ent/dialect"
)

const timeout = time.Minute

type Handler struct {
	*storage.Client
	*http.ServeMux
}

func NewHandler(client *storage.Client) *Handler {
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

func NewEntClient() *storage.Client {
	client, err := storage.Open(dialect.SQLite, "yola.db?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed connecting to sqlite: %v", err)
	}
	ctx := context.Background()
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
