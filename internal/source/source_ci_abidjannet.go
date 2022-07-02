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

type AbidjanNetSource struct {
	*store.NewsSource
	*http.Client
}

func NewAbidjanNetSource(source *store.NewsSource) *AbidjanNetSource {
	return &AbidjanNetSource{
		Client:     http.DefaultClient,
		NewsSource: source,
	}
}

/// LatestPost
///
///
func (src *AbidjanNetSource) LatestPost(ctx context.Context) []*schema.NewsPost {
	response, err := rodGetRequest(fmt.Sprintf("%s%s", src.URL, *src.LatestPostURL), "body")
	if err != nil {
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response)
	if err != nil {
		return nil
	}
	return src.latestPost(NewElement(document.Selection))
}

func (src *AbidjanNetSource) latestPost(document *Element) []*schema.NewsPost {
	selector := src.LatestPostSelector
	filmList := make([]*schema.NewsPost, 0)

	elementCallBack := func(element *Element) {
		// category := element.ChildText(selector.Category[0])
		image := element.ChildAttribute(selector.Image[0], selector.Image[1])
		title := element.ChildText(selector.Title[0])
		link := element.Attribute(selector.Link[0])
		date := element.ChildText(selector.Date[0])
		if len(image) == 0 {
			image = element.ChildAttribute(selector.Image[0], selector.Image[2])
		}
		image = parseURL(src.URL, image)
		link = parseURL(src.URL, link)
		if strings.Contains(image, "defaut-cover-photo.svg") {
			image = ""
		}

		value := strings.Split(date, "-")
		date = strings.TrimSpace(value[len(value)-1])

		if !strings.Contains(strings.Join(value, ""), "Fraternité Matin") && len(image) != 0 {
			filmList = append(filmList, &schema.NewsPost{
				Source: src.Name,
				Logo:   src.Logo,
				Image:  image,
				Title:  title,
				Link:   link,
				Date:   date,
			})
		}
	}

	elementCallBack(NewElement(document.Selection.Find(selector.List[0])))

	document.ForEach(selector.List[0],
		func(i int, element *Element) {
			elementCallBack(element)
		})
	return filmList
}

/// CategoryPost
///
///
func (src *AbidjanNetSource) CategoryPost(ctx context.Context, category string, page int) []*schema.NewsPost {
	category, err := parseCategorySource(src.NewsSource, category)
	if err != nil {
		return nil
	}
	response, err := rodGetRequest(fmt.Sprintf("%s%s", src.URL, fmt.Sprintf(*src.CategoryPostURL, category, page)), "body")
	if err != nil {
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response)
	if err != nil {
		return nil
	}
	return src.categoryPost(NewElement(document.Selection))
}

func (src *AbidjanNetSource) categoryPost(document *Element) []*schema.NewsPost {
	selector := src.LatestPostSelector
	filmList := make([]*schema.NewsPost, 0)
	document.ForEach(selector.List[0],
		func(i int, element *Element) {
			// category := element.ChildText(selector.Category[0])
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			title := element.ChildText(selector.Title[0])
			link := element.Attribute(selector.Link[0])
			date := element.ChildText(selector.Date[0])
			image = parseURL(src.URL, image)
			link = parseURL(src.URL, link)
			if strings.Contains(image, "defaut-cover-photo.svg") {
				image = ""
			}
			value := strings.Split(date, "-")
			date = strings.TrimSpace(value[len(value)-1])

			if !strings.Contains(strings.Join(value, ""), "Fraternité Matin") && len(image) != 0 {
				filmList = append(filmList, &schema.NewsPost{
					Source: src.Name,
					Logo:   src.Logo,
					Image: image,
					Title: title,
					Link:  link,
					Date:  date,
				})
			}
		})
	return filmList
}

/// NewsArticle
///
///

func (src *AbidjanNetSource) NewsArticle(ctx context.Context, link string) *schema.NewsArticle {
	response, err := rodGetRequest(link, "body")
	if err != nil {
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response)
	if err != nil {
		return nil
	}
	return src.newsArticle(NewElement(document.Selection))
}

func (src *AbidjanNetSource) newsArticle(document *Element) *schema.NewsArticle {
	selector := src.ArticleSelector
	description := document.ChildContent(selector.Description[0])
	description = strings.Join(strings.Fields(description), " ")
	return &schema.NewsArticle{
		Description: description,
	}
}
