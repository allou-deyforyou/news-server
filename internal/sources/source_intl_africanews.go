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

const AfricaNewsName = "Africa News"

type AfricaNewsSource struct {
	*storage.Source
	*http.Client
}

func NewAfricaNewsSource(source *storage.Source) *AfricaNewsSource {
	return &AfricaNewsSource{
		Client: http.DefaultClient,
		Source: source,
	}
}

/// ArticleFeaturedPostList
func (src *AfricaNewsSource) ArticleFeaturedPostList(ctx context.Context) []*custom.ArticlePost {
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

func (src *AfricaNewsSource) articleFeaturedPostList(selector *custom.SourcePostSelector, document *Element) []*custom.ArticlePost {
	result := make([]*custom.ArticlePost, 0)
	document.ForEach(selector.List[0], func(_ int, element *Element) {
		// category := element.ChildText(selector.Category[0])
		image := element.ChildAttribute(selector.Image[0], selector.Image[1])
		link := element.ChildAttribute(selector.Link[0], selector.Link[1])
		date := element.ChildAttribute(selector.Date[0], selector.Date[1])
		title := element.ChildText(selector.Title[0])

		if len(image) != 0 {
			dir, file := path.Split(image)
			imageBase := strings.Split(path.Base(file), "_")
			image = fmt.Sprintf("%s738x415_%s", dir, strings.Join(imageBase[1:], "_"))

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
func (src *AfricaNewsSource) ArticleCategoryPostList(ctx context.Context, category string, page int) []*custom.ArticlePost {
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

func (src *AfricaNewsSource) articleCategoryPostList(document *Element) []*custom.ArticlePost {
	return src.articleFeaturedPostList(src.ArticleCategoryPostSelector, document)
}

/// ArticleContent
func (src *AfricaNewsSource) ArticleContent(ctx context.Context, link string) *custom.ArticlePost {
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

func (src *AfricaNewsSource) articleContent(document *Element) *custom.ArticlePost {
	selector := src.ArticleContentSelector
	content := strings.Join(document.ChildrenOuterHtmls(selector.Content[0]), "")
	return &custom.ArticlePost{Content: content}
}

/// MediaLivePostList
//////////////
func (src *AfricaNewsSource) MediaLivePostList(ctx context.Context) []*custom.MediaPost {
	return nil
}

/// MediaCategoryPostList
//////////////
func (src *AfricaNewsSource) MediaCategoryPostList(ctx context.Context, category string, page int) []*custom.MediaPost {
	return nil
}

/// MediaContent
//////////////
func (src *AfricaNewsSource) MediaContent(ctx context.Context, link string) *custom.MediaPost {
	return nil
}
