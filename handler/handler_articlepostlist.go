package handler

import (
	"context"
	"net/http"
	"news/internal/sources"
	"news/internal/storage/custom"
	"news/internal/storage/source"
	"sync"
)

func (h *Handler) ArticlePostList(w http.ResponseWriter, r *http.Request) {
	context, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()
	/// Filter
	params := Params(r.Form)
	language, country := params.StringX("language", "fr"), params.StringX("country", "ci")
	sourceQuery := h.Source.Query().Where(
		source.And(
			source.Language(language),
			source.Status(true),
			source.Or(
				source.Country("international"),
				source.Country(country),
			),
		),
	)
	sourceList := sources.ParseSourceList(sourceQuery.AllX(context))
	/// Fetch
	category, page := params.StringX("category", "featured"), params.IntX("page", 1)
	posts := make([]*custom.ArticlePost, 0)
	group := new(sync.WaitGroup)
	for _, s := range sourceList {
		group.Add(1)
		go func(source sources.Source) {
			defer RecoverFunc(group)
			switch category {
			case "featured":
				posts = append(posts, source.ArticleFeaturedPostList(context)...)
			default:
				posts = append(posts, source.ArticleCategoryPostList(context, category, page)...)
			}
			group.Done()
		}(s)
	}
	group.Wait()
	/// Response
	posts = Remove(posts, func(a, b *custom.ArticlePost) bool { return a.Link == b.Link })
	JsonEncode(w, &custom.ArticlePostListResponse{Data: Shuffle(posts)})
}
