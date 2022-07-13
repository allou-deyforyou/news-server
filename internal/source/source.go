package source

import (
	"context"
	"errors"

	"news/internal/source/trash"
	"news/internal/store"
	"news/internal/store/schema"
)

type NewsSource interface {
	CategoryPost(ctx context.Context, category string, page int) []*schema.NewsArticlePost
	NewsArticle(ctx context.Context, link string) *schema.NewsArticlePost
	LatestPost(ctx context.Context) []*schema.NewsArticlePost
}

func ParseListNewsSource(sources []*store.NewsArticleSource) (result []NewsSource) {
	for _, source := range sources {
		value, err := ParseNewsSource(source.Name, source)
		if err == nil {
			result = append(result, value)
		}
	}
	return
}

func ParseNewsSource(name string, source *store.NewsArticleSource) (NewsSource, error) {
	switch name {
	// Trash Low Quality Image
	case trash.FratmatInfoName:
		return trash.NewFratmatInfoSource(source), nil
	case trash.AbidjanNetName:
		return trash.NewAbidjanNetSource(source), nil
	case trash.AfrikMagName:
		return trash.NewAfrikMagSource(source), nil

	// High Quality Image
	case AfricaNewsName:
		return NewAfricaNewsSource(source), nil
	case France24Name:
		return NewFrance24Source(source), nil
	case RFIName:
		return NewRFISource(source), nil
	case BBCName:
		return NewBBCSource(source), nil
	case YecloName:
		return NewYecloSource(source), nil
	case RTIInfoName:
		return NewRTIInfoSource(source), nil
	default:
		return nil, errors.New("no-found")
	}
}
