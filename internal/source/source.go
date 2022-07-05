package source

import (
	"context"
	"errors"
	"news/internal/store"
	"news/internal/store/schema"
)

type NewsSource interface {
	CategoryPost(ctx context.Context, category string, page int) []*schema.NewsPost
	NewsArticle(ctx context.Context, link string) *schema.NewsArticle
	LatestPost(ctx context.Context) []*schema.NewsPost
}

func ParseListNewsSource(sources []*store.NewsSource) (result []NewsSource) {
	for _, source := range sources {
		value, err := ParseNewsSource(source.Name, source)
		if err == nil {
			result = append(result, value)
		}
	}
	return
}

func ParseNewsSource(name string, source *store.NewsSource) (NewsSource, error) {
	switch name {
	case "Fratmat Info":
		return NewFratmatInfoSource(source), nil
	case "Abidjan.Net":
		return NewAbidjanNetSource(source), nil
	case "AfrikMag":
		return NewAfrikMagSource(source), nil
	// case "France 24":
	// 	return NewFrance24Source(source), nil
	case "RFI":
		return NewRfiSource(source), nil
	default:
		return nil, errors.New("no-found")
	}
}
