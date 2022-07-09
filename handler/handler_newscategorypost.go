package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"news/internal"
	"news/internal/source"
	"news/internal/store/newssource"
	"news/internal/store/schema"

	"sync"
)

func (h *Handler) NewsCategoryPost(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		timeout,
	)
	defer cancel()

	params := internal.Params(r.Form)
	name, _ := params.String("name")
	page, _ := params.Int("page")

	newsSources := h.NewsSource.Query().Where(newssource.Status(true)).AllX(ctx)
	sources := source.ParseListNewsSource(newsSources)

	response := make([]*schema.NewsPost, 0)
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
			posts := source.CategoryPost(ctx, name, page)
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
