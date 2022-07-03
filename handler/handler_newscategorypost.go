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
	"time"
)

func (h *Handler) NewsCategoryPost(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
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
			posts := source.CategoryPost(ctx, name, page)
			response = append(response, posts...)
			group.Done()
		}(s)
	}
	group.Wait()

	response = internal.Shuffle(response)
	json.NewEncoder(w).Encode(response)
}
