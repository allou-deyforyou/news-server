package handler

import (
	"context"
	"net/http"
	"news/internal/sources"
	"news/internal/storage/custom"
	entSource "news/internal/storage/source"
)

func (h *Handler) ArticleContent(w http.ResponseWriter, r *http.Request) {
	context, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()
	/// Filter
	params := Params(r.Form)
	link, linkError := params.String("link")
	if linkError != nil {
		panic(linkError)
	}
	source, sourceError := params.String("source")
	if sourceError != nil {
		panic(sourceError)
	}
	queryValue := h.Source.Query().Where(entSource.Name(source)).OnlyX(context)
	sourceParsed, parseError := sources.ParseSource(queryValue.Name, queryValue)
	if parseError != nil {
		panic(parseError)
	}
	// Fetch
	post := sourceParsed.ArticleContent(context, link)
	// Response
	JsonEncode(w, &custom.ArticlePostResponse{Data: post})
}
