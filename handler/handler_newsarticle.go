package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"news/internal"
	"news/internal/source"
	"news/internal/store/newssource"
)

func (h *Handler) NewsArticle(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		timeout,
	)
	defer cancel()

	params := internal.Params(r.Form)
	name, _ := params.String("source")
	link, _ := params.String("link")

	movieSource := h.NewsSource.Query().Where(newssource.And(newssource.Status(true), newssource.Name(name))).OnlyX(ctx)
	source, _ := source.ParseNewsSource(movieSource.Name, movieSource)

	response := source.NewsArticle(ctx, link)
	json.NewEncoder(w).Encode(response)
}
