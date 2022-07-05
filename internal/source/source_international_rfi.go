package source

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"news/internal/store"
	"news/internal/store/schema"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type RfiSource struct {
	*store.NewsSource
	*http.Client
}

func NewRfiSource(source *store.NewsSource) *RfiSource {
	return &RfiSource{
		Client:     http.DefaultClient,
		NewsSource: source,
	}
}

/// LatestPost
///
///
func (src *RfiSource) LatestPost(ctx context.Context) []*schema.NewsPost {
	response, err := rodGetRequest(fmt.Sprintf("%s%s", src.URL, *src.LatestPostURL))
	if err != nil {
		log.Println(err)
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response)
	if err != nil {
		log.Println(err)
		return nil
	}
	return src.latestPost(NewElement(document.Selection))
}

func (src *RfiSource) latestPost(document *Element) []*schema.NewsPost {
	selector := src.LatestPostSelector
	result := make([]*schema.NewsPost, 0)
	document.ForEach(selector.List[0], func(i int, element *Element) {
		// category := element.ChildText(selector.Category[0])
		image := element.ChildAttribute(selector.Image[0], selector.Image[1])
		link := element.ChildAttribute(selector.Link[0], selector.Link[1])
		title := element.ChildText(selector.Title[0])

		if len(title) != 0 {
			rawImage := strings.Split(image, ",")
			image = strings.Fields(rawImage[len(rawImage)-1])[0]

			date := strings.Split(path.Base(link), "-")[0]
			date = fmt.Sprintf("%v-%v-%v", string(date[:4]), string(date[4:6]), string(date[6:8]))

			image = parseURL(src.URL, image)
			link = parseURL(src.URL, link)
			date, _ = parseTime(date)

			result = append(result, &schema.NewsPost{
				Source: src.Name,
				Logo:   src.Logo,
				Image:  image,
				Title:  title,
				Link:   link,
				Date:   date,
			})
		}
	})
	return result
}

/// NewsCategory
////////////////
func (src *RfiSource) CategoryPost(ctx context.Context, category string, page int) []*schema.NewsPost {
	category, err := parseCategorySource(src.NewsSource, category)
	if err != nil {
		log.Println(err)
		return nil
	}
	response, err := rodGetRequest(fmt.Sprintf("%s%s", src.URL, fmt.Sprintf(*src.CategoryPostURL, category, page)))
	if err != nil {
		log.Println(err)
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response)
	if err != nil {
		log.Println(err)
		return nil
	}
	return src.categoryPost(NewElement(document.Selection))
}

func (src *RfiSource) categoryPost(document *Element) []*schema.NewsPost {
	selector := src.CategoryPostSelector
	result := make([]*schema.NewsPost, 0)
	document.ForEach(selector.List[0], func(i int, element *Element) {
		// category := element.ChildText(selector.Category[0])
		image := element.ChildAttribute(selector.Image[0], selector.Image[1])
		link := element.ChildAttribute(selector.Link[0], selector.Link[1])
		title := element.ChildText(selector.Title[0])

		rawImage := strings.Split(image, ",")
		image = strings.Fields(rawImage[len(rawImage)-1])[0]

		date := strings.Split(path.Base(link), "-")[0]
		date = fmt.Sprintf("%v %v %v", string(date[:4]), string(date[4:6]), string(date[6:8]))

		image = parseURL(src.URL, image)
		link = parseURL(src.URL, link)
		date, _ = parseTime(date)

		result = append(result, &schema.NewsPost{
			Source: src.Name,
			Logo:   src.Logo,
			Image:  image,
			Title:  title,
			Link:   link,
			Date:   date,
		})
	})
	return result
}

/// PostArticle
///////////////
func (src *RfiSource) NewsArticle(ctx context.Context, link string) *schema.NewsArticle {
	response, err := rodGetRequest(link)
	if err != nil {
		log.Println(err)
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response)
	if err != nil {
		log.Println(err)
		return nil
	}
	return src.newsArticle(NewElement(document.Selection))
}

func (src *RfiSource) newsArticle(document *Element) *schema.NewsArticle {
	selector := src.ArticleSelector
	contents := document.ChildrenOuterHtmls(selector.Description[0])
	description := strings.Join(contents, "")
	return &schema.NewsArticle{
		Description: description,
	}
}
