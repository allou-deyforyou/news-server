package sources

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"

	"news/internal/storage"
	"news/internal/storage/custom"

	"github.com/PuerkitoBio/goquery"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const France24Name = "France 24"

type France24Source struct {
	*storage.Source
	*http.Client
}

func NewFrance24Source(source *storage.Source) *France24Source {
	return &France24Source{
		Client: http.DefaultClient,
		Source: source,
	}
}

/// ArticleFeaturedPostList
func (src *France24Source) ArticleFeaturedPostList(ctx context.Context) []*custom.ArticlePost {
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

func (src *France24Source) articleFeaturedPostList(selector *custom.SourcePostSelector, document *Element) []*custom.ArticlePost {
	result := make([]*custom.ArticlePost, 0)
	document.ForEach(selector.List[0], func(_ int, element *Element) {
		image := element.ChildAttribute(selector.Image[0], selector.Image[1])
		link := element.ChildAttribute(selector.Link[0], selector.Link[1])
		title := element.ChildText(selector.Title[0])
		if len(image) != 0 {
			rawImage := strings.Split(image, ",")
			image = strings.Fields(rawImage[len(rawImage)-1])[0]

			date := strings.Split(path.Base(link), "-")[0]
			date = fmt.Sprintf("%v-%v-%v", string(date[:4]), string(date[4:6]), string(date[6:8]))

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
func (src *France24Source) ArticleCategoryPostList(ctx context.Context, category string, page int) []*custom.ArticlePost {
	if page != 1 {
		return nil
	}
	response, err := RodGetRequest(fmt.Sprintf("%s%s", src.URL, fmt.Sprintf(src.ArticleCategoryPostURL, src.ArticleCategories[category])))
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

func (src *France24Source) articleCategoryPostList(document *Element) []*custom.ArticlePost {
	return src.articleFeaturedPostList(src.ArticleCategoryPostSelector, document)
}

/// ArticleContent
func (src *France24Source) ArticleContent(ctx context.Context, link string) *custom.ArticlePost {
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

func (src *France24Source) articleContent(document *Element) *custom.ArticlePost {
	selector := src.ArticleContentSelector
	content := strings.Join(document.ChildrenOuterHtmls(selector.Content[0]), "")
	return &custom.ArticlePost{Content: content}
}

/// MediaLivePostList
//////////////
func (src *France24Source) MediaLivePostList(ctx context.Context) []*custom.MediaPost {
	return nil
}

/// MediaCategoryPostList
//////////////
func (src *France24Source) MediaCategoryPostList(ctx context.Context, category string, page int) []*custom.MediaPost {
	return nil
}

/// MediaContent
//////////////
func (src *France24Source) MediaContent(ctx context.Context, link string) *custom.MediaPost {
	return nil
}
