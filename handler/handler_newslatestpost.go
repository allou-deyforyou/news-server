package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"news/internal"
	"news/internal/source"
	"news/internal/store/newssource"
	"news/internal/store/schema"

	"sync"
)

func (h *Handler) NewsLatestPost(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		timeout,
	)
	defer cancel()

	newsSources := h.NewsSource.Query().Where(newssource.Status(true)).AllX(ctx)
	sources := source.ParseListNewsSource(newsSources)

	response := make([]*schema.NewsPost, 0)
	group := new(sync.WaitGroup)
	for _, s := range sources {
		group.Add(1)
		go func(source source.NewsSource) {
			posts := source.LatestPost(ctx)
			response = append(response, posts...)
			group.Done()
		}(s)
	}
	group.Wait()

	response = internal.Remove(response, func(a, b *schema.NewsPost) bool {
		return a.Link == b.Link
	})
	response = internal.Shuffle(response)
	json.NewEncoder(w).Encode(response)
}
