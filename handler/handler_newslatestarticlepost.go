package handler

import (
	"context"
	"log"
	"net/http"

	"news/internal"
	"news/internal/source"
	"news/internal/store/newsarticlesource"
	"news/internal/store/schema"

	"sync"
)

func (h *Handler) NewsLatestArticlePost(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		timeout,
	)
	defer cancel()

	newsSources := h.NewsArticleSource.Query().Where(newsarticlesource.Status(true)).AllX(ctx)
	sources := source.ParseListNewsSource(newsSources)

	response := make([]*schema.NewsArticlePost, 0)
	group := new(sync.WaitGroup)
	for _, s := range sources {
		group.Add(1)
		go func(source source.NewsSource) {
			defer func() {
				if r := recover(); r != nil {
					group.Done()
					log.Println("Recovered in f", r)
				}
			}()
			posts := source.LatestPost(ctx)
			response = append(response, posts...)
			group.Done()
		}(s)
	}
	group.Wait()

	response = internal.Remove(response,
		func(a, b *schema.NewsArticlePost) bool {
			return a.Link == b.Link
		})
	response = internal.Shuffle(response)
	internal.ProtoEncode(w, &schema.NewsArticlePostResponse{
		Data: response,
	})
}
