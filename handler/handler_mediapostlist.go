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

func (h *Handler) MediaPostList(w http.ResponseWriter, r *http.Request) {
	context, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()
	/// Filter
	params := Params(r.Form)
	language, country, category, page := params.StringX("language", "fr"), params.StringX("country", "ci"), params.StringX("category", "live"), params.IntX("page", 1)
	sourceQuery := h.Source.Query().Where(source.And(source.Language(language), source.Status(true), source.Or(source.Country(custom.InternationalMediaCategory), source.Country(country))))
	values := Filter(sourceQuery.AllX(context), func(source *storage.Source) bool { _, ok := source.MediaCategories[category]; return ok })
	sourceList := sources.ParseSourceList(values)
	/// Fetch
	posts := make([]*custom.MediaPost, 0)
	switch category {
	case "live":
		if page == 1 {
			posts = append(posts, sources.GetMediaPostList(h.Client, context)...)
		}
	default:
		group := new(sync.WaitGroup)
		for _, s := range sourceList {
			group.Add(1)
			go func(source sources.Source) {
				defer RecoverFunc(group)
				posts = append(posts, source.MediaCategoryPostList(context, category, page)...)
				group.Done()
			}(s)
		}
		group.Wait()
	}
	/// Response
	posts = Remove(posts, func(a, b *custom.MediaPost) bool { return a.Link == b.Link && a.Content == b.Content })
	ProtoEncode(w, &custom.MediaPostListResponse{Data: Shuffle(posts)})
}
