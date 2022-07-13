package handler

import (
	"context"
	"net/http"

	"news/internal"
	"news/internal/source"
	"news/internal/store/newsarticlesource"
	"news/internal/store/schema"
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

	movieSource := h.NewsArticleSource.Query().Where(newsarticlesource.And(newsarticlesource.Status(true), newsarticlesource.Name(name))).OnlyX(ctx)
	source, _ := source.ParseNewsSource(movieSource.Name, movieSource)

	response := source.NewsArticle(ctx, link)
	internal.ProtoEncode(w, &schema.NewsArticleResponse{
		Data: response,
	})
}
