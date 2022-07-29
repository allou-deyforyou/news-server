package handler

import (
	"context"
	"net/http"
	"news/internal/sources"
	"news/internal/storage"
	"news/internal/storage/custom"
	"news/internal/storage/source"
	"sync"
)

func (h *Handler) ArticlePostList(w http.ResponseWriter, r *http.Request) {
	context, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()
	/// Filter
	params := Params(r.Form)
	language, country, category, page := params.StringX("language", "fr"), params.StringX("country", "ci"), params.StringX("category", custom.FeaturedArticleCategory), params.IntX("page", 1)
	sourceQuery := h.Source.Query().Where(source.And(source.Language(language), source.Status(true), source.Or(source.Country(custom.InternationalArticleCategory), source.Country(country))))
	values := Filter(sourceQuery.AllX(context), func(source *storage.Source) bool { _, ok := source.ArticleCategories[category]; return ok })
	sourceList := sources.ParseSourceList(values)
	/// Fetch
	posts := make([]*custom.ArticlePost, 0)
	group := new(sync.WaitGroup)
	for _, s := range sourceList {
		group.Add(1)
		go func(source sources.Source) {
			defer RecoverFunc(group)
			switch category {
			case custom.FeaturedArticleCategory:
				if page == 1 {
					posts = append(posts, source.ArticleFeaturedPostList(context)...)
				}
			default:
				posts = append(posts, source.ArticleCategoryPostList(context, category, page)...)
			}
			group.Done()
		}(s)
	}
	group.Wait()
	/// Response
	posts = Remove(posts, func(a, b *custom.ArticlePost) bool { return a.Link == b.Link || a.Title == b.Title })
	ProtoEncode(w, &custom.ArticlePostListResponse{Data: Shuffle(posts)})
}
