package sources

import (
	"context"
	"errors"
	"news/internal/storage"
	"news/internal/storage/custom"
)

type Source interface {
	ArticleCategoryPostList(ctx context.Context, category string, page int) []*custom.ArticlePost
	ArticleContent(ctx context.Context, link string) *custom.ArticlePost
	ArticleFeaturedPostList(ctx context.Context) []*custom.ArticlePost

	MediaCategoryPostList(ctx context.Context, category string, page int) []*custom.MediaPost
	MediaContent(ctx context.Context, link string) *custom.MediaPost
	MediaLivePostList(ctx context.Context) []*custom.MediaPost
}

func ParseSourceList(sources []*storage.Source) (result []Source) {
	for _, source := range sources {
		value, err := ParseSource(source.Name, source)
		if err == nil {
			result = append(result, value)
		}
	}
	return
}

func ParseSource(name string, source *storage.Source) (Source, error) {
	switch name {
	/// International
	case AfricaNewsName:
		return NewAfricaNewsSource(source), nil
	case France24Name:
		return NewFrance24Source(source), nil
	case RFIName:
		return NewRFISource(source), nil
	case BBCName:
		return NewBBCSource(source), nil

	/// CIV
	case YecloName:
		return NewYecloSource(source), nil
	case RTIInfoName:
		return NewRTIInfoSource(source), nil
	case FratmatInfoName:
		return NewFratmatInfoSource(source), nil
	default:
		return nil, errors.New("no-found")
	}
}
