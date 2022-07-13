package handler

import (
	"context"
	"net/http"
	"news/internal"
	"news/internal/store/newstvsource"
	"news/internal/store/schema"
)

func (h *Handler) NewsTvPost(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		timeout,
	)
	defer cancel()

	params := internal.Params(r.Form)
	language, _ := params.String("language")
	country, _ := params.String("country")

	tvQuery := h.NewsTvSource.Query()
	if len(language) != 0 {
		tvQuery = tvQuery.Where(newstvsource.Language(language))
	}
	if len(country) != 0 {
		tvQuery = tvQuery.Where(newstvsource.Country(country))
	}
	data := tvQuery.Where(newstvsource.Status(true)).AllX(ctx)
	response := internal.ConvertToNewsTvPost(data)

	response = internal.Shuffle(response)
	internal.ProtoEncode(w, &schema.NewsTvPostResponse{
		Data: response,
	})
}
