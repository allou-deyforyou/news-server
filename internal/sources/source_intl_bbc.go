package sources

import (
	"context"
	"fmt"
	"html"
	"log"
	"net/http"
	"strings"

	"news/internal/storage"
	"news/internal/storage/custom"

	"github.com/PuerkitoBio/goquery"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const BBCName = "BBC"

type BBCSource struct {
	*storage.Source
	*http.Client
}

func NewBBCSource(source *storage.Source) *BBCSource {
	return &BBCSource{
		Client: http.DefaultClient,
		Source: source,
	}
}

/// ArticleFeaturedPostList
func (src *BBCSource) ArticleFeaturedPostList(ctx context.Context) []*custom.ArticlePost {
	response, err := RodGetRequest(fmt.Sprintf("%s%s", src.URL, src.ArticleFeaturedPostURL))
	if err != nil {
		log.Println(err)
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response)
	if err != nil {
		log.Println(err)
		return nil
	}
	return src.articleFeaturedPostList(src.ArticleFeaturedPostSelector, NewElement(document.Selection))
}

func (src *BBCSource) articleFeaturedPostList(selector *custom.SourcePostSelector, document *Element) []*custom.ArticlePost {
	result := make([]*custom.ArticlePost, 0)
	document.ForEach(selector.List[0], func(_ int, element *Element) {
		image := element.ChildAttribute(selector.Image[0], selector.Image[1])
		link := element.ChildAttribute(selector.Link[0], selector.Link[1])
		date := element.ChildAttribute(selector.Date[0], selector.Date[1])
		title := element.ChildText(selector.Title[0])

		if len(image) != 0 {
			rawImage := strings.Split(image, ",")
			image = strings.Fields(rawImage[len(rawImage)-1])[0]

			image = ParseURL(src.URL, image)
			link = ParseURL(src.URL, link)
			dateTime, _ := ParseTime(date)

			result = append(result, &custom.ArticlePost{
				Date:   timestamppb.New(dateTime),
				Source: src.Name,
				Logo:   src.Logo,
				Image:  image,
				Title:  title,
				Link:   link,
			})
		}
	})
	return result
}

/// ArticleCategoryPostList
func (src *BBCSource) ArticleCategoryPostList(ctx context.Context, category string, page int) []*custom.ArticlePost {
	response, err := RodGetRequest(fmt.Sprintf("%s%s", src.URL, fmt.Sprintf(src.ArticleCategoryPostURL, src.ArticleCategories[category], page)))
	if err != nil {
		log.Println(err)
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response)
	if err != nil {
		log.Println(err)
		return nil
	}
	return src.articleCategoryPostList(NewElement(document.Selection))
}

func (src *BBCSource) articleCategoryPostList(document *Element) []*custom.ArticlePost {
	return src.articleFeaturedPostList(src.ArticleCategoryPostSelector, document)
}

/// ArticleContent
func (src *BBCSource) ArticleContent(ctx context.Context, link string) *custom.ArticlePost {
	response, err := RodGetRequest(link)
	if err != nil {
		log.Println(err)
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response)
	if err != nil {
		log.Println(err)
		return nil
	}
	return src.articleContent(NewElement(document.Selection))
}

func (src *BBCSource) articleContent(document *Element) *custom.ArticlePost {
	selector := src.ArticleContentSelector
	document.ForEach(selector.Content[1], func(i int, e *Element) {
		e.Selection.ReplaceWithHtml(html.UnescapeString(e.InnerHtml()))
	})
	content := strings.Join(document.ChildrenOuterHtmls(selector.Content[0]), "")
	return &custom.ArticlePost{Content: content}
}


/// MediaLivePostList
//////////////
func (src *BBCSource) MediaLivePostList(ctx context.Context) []*custom.MediaPost {
	return nil
}

/// MediaCategoryPostList
//////////////
func (src *BBCSource) MediaCategoryPostList(ctx context.Context, category string, page int) []*custom.MediaPost {
	return nil
}

/// MediaContent
//////////////
func (src *BBCSource) MediaContent(ctx context.Context, link string) *custom.MediaPost {
	return nil
}