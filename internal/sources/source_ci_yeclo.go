package sources

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"path"
	"regexp"
	"strings"

	"news/internal/storage"
	"news/internal/storage/custom"

	"github.com/PuerkitoBio/goquery"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const YecloName = "Yeclo"

type YecloSource struct {
	*storage.Source
	*http.Client
}

func NewYecloSource(source *storage.Source) *YecloSource {
	return &YecloSource{
		Client: http.DefaultClient,
		Source: source,
	}
}

/// ArticleFeaturedPostList
///
func (src *YecloSource) ArticleFeaturedPostList(ctx context.Context) []*custom.ArticlePost {
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

func (src *YecloSource) articleFeaturedPostList(selector *custom.SourcePostSelector, document *Element) []*custom.ArticlePost {
	result := make([]*custom.ArticlePost, 0)
	document.ForEach(selector.List[0], func(_ int, element *Element) {
		image := element.ChildAttribute(selector.Image[0], selector.Image[1])
		link := element.ChildAttribute(selector.Link[0], selector.Link[1])
		date := element.ChildAttribute(selector.Date[0], selector.Date[1])
		title := element.ChildText(selector.Title[0])

		if len(image) != 0 {
			if len(regexp.MustCompile(`-\d{3}x\d{3}.`).FindString(image)) != 0 {
				rawPath := strings.Split(image, "-")
				image = strings.ReplaceAll(image, fmt.Sprintf("-%v", rawPath[len(rawPath)-1]), path.Ext(image))
			}

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
////////////////
func (src *YecloSource) ArticleCategoryPostList(ctx context.Context, category string, page int) []*custom.ArticlePost {
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

func (src *YecloSource) articleCategoryPostList(document *Element) []*custom.ArticlePost {
	return src.articleFeaturedPostList(src.ArticleCategoryPostSelector, document)
}

/// ArticleContent
///////////////
func (src *YecloSource) ArticleContent(ctx context.Context, link string) *custom.ArticlePost {
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

func (src *YecloSource) articleContent(document *Element) *custom.ArticlePost {
	selector := src.ArticleContentSelector
	content := strings.Join(document.ChildrenOuterHtmls(selector.Content[0]), "")
	return &custom.ArticlePost{Content: content}
}


/// MediaLivePostList
//////////////
func (src *YecloSource) MediaLivePostList(ctx context.Context) []*custom.MediaPost {
	return nil
}

/// MediaCategoryPostList
//////////////
func (src *YecloSource) MediaCategoryPostList(ctx context.Context, category string, page int) []*custom.MediaPost {
	return nil
}

/// MediaContent
//////////////
func (src *YecloSource) MediaContent(ctx context.Context, link string) *custom.MediaPost {
	return nil
}