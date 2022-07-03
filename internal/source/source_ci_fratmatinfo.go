package source

import (
	"context"
	"fmt"
	"net/http"
	"news/internal/store"
	"news/internal/store/schema"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type FratmatInfoSource struct {
	*store.NewsSource
	*http.Client
}

func NewFratmatInfoSource(source *store.NewsSource) *FratmatInfoSource {
	return &FratmatInfoSource{
		Client:     http.DefaultClient,
		NewsSource: source,
	}
}

/// NewsLatest
//////////////
func (src *FratmatInfoSource) LatestPost(ctx context.Context) []*schema.NewsPost {
	response, err := rodGetRequest(fmt.Sprintf("%s%s", src.URL, *src.LatestPostURL), "main")
	if err != nil {
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response)
	if err != nil {
		return nil
	}
	return src.latestPost(NewElement(document.Selection))
}

func (src *FratmatInfoSource) latestPost(document *Element) []*schema.NewsPost {
	selector := src.LatestPostSelector
	filmList := make([]*schema.NewsPost, 0)
	document.ForEach(selector.List[0],
		func(i int, element *Element) {
			// category := element.ChildText(selector.Category[0])
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			link := element.ChildAttribute(selector.Link[0], selector.Link[1])
			title := element.ChildText(selector.Title[0])
			date := element.ChildText(selector.Date[0])

			image = parseURL(src.URL, image)
			date, _ = parseTime(date)

			filmList = append(filmList, &schema.NewsPost{
				Source: src.Name,
				Logo:   src.Logo,
				Image:  image,
				Title:  title,
				Link:   link,
				Date:   date,
			})
		})
	return filmList
}

/// NewsCategory
////////////////
func (src *FratmatInfoSource) CategoryPost(ctx context.Context, category string, page int) []*schema.NewsPost {
	category, err := parseCategorySource(src.NewsSource, category)
	if err != nil {
		return nil
	}
	response, err := rodGetRequest(fmt.Sprintf("%s%s", src.URL, fmt.Sprintf(*src.CategoryPostURL, category, page)), "main")
	if err != nil {
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response)
	if err != nil {
		return nil
	}
	return src.categoryPost(NewElement(document.Selection))
}

func (src *FratmatInfoSource) categoryPost(document *Element) []*schema.NewsPost {
	selector := src.CategoryPostSelector
	filmList := make([]*schema.NewsPost, 0)
	document.ForEach(selector.List[0],
		func(i int, element *Element) {
			// category := element.ChildText(selector.Category[0])
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			link := element.ChildAttribute(selector.Link[0], selector.Link[1])
			title := element.ChildText(selector.Title[0])
			date := element.ChildText(selector.Date[0])

			image = parseURL(src.URL, image)
			date, _ = parseTime(date)

			filmList = append(filmList, &schema.NewsPost{
				Source: src.Name,
				Logo:   src.Logo,
				Image:  image,
				Title:  title,
				Link:   link,
				Date:   date,
			})
		})
	return filmList
}

/// PostArticle
///////////////
func (src *FratmatInfoSource) NewsArticle(ctx context.Context, link string) *schema.NewsArticle {
	response, err := rodGetRequest(link, "main")
	if err != nil {
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response)
	if err != nil {
		return nil
	}
	return src.newsArticle(NewElement(document.Selection))
}

func (src *FratmatInfoSource) newsArticle(document *Element) *schema.NewsArticle {
	selector := src.ArticleSelector
	description := document.ChildContent(selector.Description[0])
	description = strings.Join(strings.Fields(description), " ")
	return &schema.NewsArticle{
		Description: description,
	}
}
