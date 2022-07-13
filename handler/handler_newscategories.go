package handler

import (
	"context"
	"net/http"
	"news/internal"
	"news/internal/store/newscategories"
	"news/internal/store/schema"
)

func (h *Handler) NewsCateggories(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		timeout,
	)
	defer cancel()

	params := internal.Params(r.Form)
	language, _ := params.String("language")

	data := h.NewsCategories.Query().Where(newscategories.And(newscategories.Status(true), newscategories.Language(language))).OnlyX(ctx)

	internal.ProtoEncode(w, &schema.NewsCategoryResponse{
		ArticleCategories: data.ArticleCategories,
		TvCategories:      data.TvCategories,
	})
}
