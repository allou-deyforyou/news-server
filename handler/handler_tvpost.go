package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"news/internal"
	"news/internal/store/tvsource"
)

func (h *Handler) TvPost(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	params := internal.Params(r.Form)
	language, _ := params.String("language")
	country, _ := params.String("country")

	tvQuery := h.TvSource.Query()
	if len(language) != 0 {
		tvQuery = tvQuery.Where(tvsource.Language(language))
	}
	if len(country) != 0 {
		tvQuery = tvQuery.Where(tvsource.Country(country))
	}
	response := tvQuery.Where(tvsource.Status(true)).AllX(ctx)
	json.NewEncoder(w).Encode(response)
}
