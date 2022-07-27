package handler

import (
	"context"
	"net/http"
	"news/internal/storage/categories"
	"news/internal/storage/custom"
)

func (h *Handler) NewsCategories(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	/// Filter
	params := Params(r.Form)
	language := params.StringX("language", "fr")
	gender := params.StringX("gender", "article")
	data := h.Client.Categories.Query().Where(categories.And(categories.Status(true), categories.Language(language))).FirstX(ctx)
	/// Fetch
	var categories []*custom.NewsCategory
	values := data.MediaCategories
	if gender == "article" {
		values = data.ArticleCategories
	}
	for name, value := range values {
		categories = append(categories, &custom.NewsCategory{
			Value: value,
			Name:  name,
		})
	}
	/// Response
	ProtoEncode(w, &custom.NewsCategoryListResponse{Data: categories})
}
