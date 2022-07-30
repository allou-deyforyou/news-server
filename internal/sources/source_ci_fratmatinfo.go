package sources

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"news/internal/storage"
	"news/internal/storage/custom"

	"github.com/PuerkitoBio/goquery"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const FratmatInfoName = "Fratmat Info"

type FratmatInfoSource struct {
	*storage.Source
	*http.Client
}

func NewFratmatInfoSource(source *storage.Source) *FratmatInfoSource {
	return &FratmatInfoSource{
		Client: http.DefaultClient,
		Source: source,
	}
}

/// ArticleFeaturedPostList
//////////////
func (src *FratmatInfoSource) ArticleFeaturedPostList(ctx context.Context) []*custom.ArticlePost {
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

func (src *FratmatInfoSource) articleFeaturedPostList(selector *custom.SourcePostSelector, document *Element) []*custom.ArticlePost {
	result := make([]*custom.ArticlePost, 0)
	document.ForEach(selector.List[0],
		func(_ int, element *Element) {
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			link := element.ChildAttribute(selector.Link[0], selector.Link[1])
			title := element.ChildText(selector.Title[0])
			date := element.ChildText(selector.Date[0])

			image = ParseURL(src.URL, image)
			dateTime, _ := ParseTime(date)

			result = append(result, &custom.ArticlePost{
				Date:   timestamppb.New(dateTime),
				Description: src.Description,
				Source: src.Name,
				Logo:   src.Logo,
				Image:  image,
				Title:  title,
				Link:   link,
			})
		})
	return result
}

/// ArticleCategoryPostList
////////////////
func (src *FratmatInfoSource) ArticleCategoryPostList(ctx context.Context, category string, page int) []*custom.ArticlePost {
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

func (src *FratmatInfoSource) articleCategoryPostList(document *Element) []*custom.ArticlePost {
	return src.articleFeaturedPostList(src.ArticleCategoryPostSelector, document)
}

/// ArticleContent
///////////////
func (src *FratmatInfoSource) ArticleContent(ctx context.Context, link string) *custom.ArticlePost {
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

func (src *FratmatInfoSource) articleContent(document *Element) *custom.ArticlePost {
	selector := src.ArticleContentSelector
	selection := document.Selection.Find(selector.Content[0])
	selection.Find(selector.Content[1]).Remove()
	content := strings.Join(strings.Fields(NewElement(selection).OuterHtml()), " ")
	return &custom.ArticlePost{Content: content}
}

/// MediaLivePostList
//////////////
func (src *FratmatInfoSource) MediaLivePostList(ctx context.Context) []*custom.MediaPost {
	return nil
}

/// MediaCategoryPostList
//////////////
func (src *FratmatInfoSource) MediaCategoryPostList(ctx context.Context, category string, page int) []*custom.MediaPost {
	response, err := RodGetRequest(fmt.Sprintf("%s%s", src.URL, fmt.Sprintf(src.MediaCategoryPostURL, src.MediaCategories[category], page)))
	if err != nil {
		log.Println(err)
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response)
	if err != nil {
		log.Println(err)
		return nil
	}
	return src.mediaCategoryPostList(NewElement(document.Selection))
}

func (src *FratmatInfoSource) mediaCategoryPostList(document *Element) []*custom.MediaPost {
	selector := src.MediaCategoryPostSelector
	result := make([]*custom.MediaPost, 0)
	document.ForEach(selector.List[0],
		func(_ int, element *Element) {
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			link := element.ChildAttribute(selector.Link[0], selector.Link[1])
			title := element.ChildText(selector.Title[0])
			date := element.ChildText(selector.Date[0])

			image = ParseURL(src.URL, image)
			dateTime, _ := ParseTime(date)

			result = append(result, &custom.MediaPost{
				Date:   timestamppb.New(dateTime),
				Description: src.Description,
				Source: src.Name,
				Logo:   src.Logo,
				Image:  image,
				Title:  title,
				Link:   link,
			})
		})
	return result
}

/// MediaContent
//////////////
func (src *FratmatInfoSource) MediaContent(ctx context.Context, link string) *custom.MediaPost {
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
	return src.mediaContent(NewElement(document.Selection))
}

func (src *FratmatInfoSource) mediaContent(document *Element) *custom.MediaPost {
	selector := src.MediaContentSelector

	content := document.ChildAttribute(selector.Content[0], selector.Content[1])
	return &custom.MediaPost{
		Type:    custom.MediaPost_YOUTUBE,
		Content: content,
	}
}
