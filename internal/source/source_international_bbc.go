package source

import (
	"context"
	"fmt"
	"html"
	"log"
	"net/http"
	"strings"

	"news/internal/store"
	"news/internal/store/schema"
	"news/internal/util"

	"github.com/PuerkitoBio/goquery"
)

const BBCName = "BBC"

type BBCSource struct {
	*store.NewsSource
	*http.Client
}

func NewBBCSource(source *store.NewsSource) *BBCSource {
	return &BBCSource{
		Client:     http.DefaultClient,
		NewsSource: source,
	}
}

/// LatestPost
///
///
func (src *BBCSource) LatestPost(ctx context.Context) []*schema.NewsPost {
	response, err := util.RodNavigate(fmt.Sprintf("%s%s", src.URL, *src.LatestPostURL), true)
	if err != nil {
		log.Println(err)
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response)
	if err != nil {
		log.Println(err)
		return nil
	}
	return src.latestPost(util.NewElement(document.Selection))
}

func (src *BBCSource) latestPost(document *util.Element) []*schema.NewsPost {
	selector := src.LatestPostSelector
	result := make([]*schema.NewsPost, 0)
	document.ForEach(selector.List[0], func(i int, element *util.Element) {
		// category := element.ChildText(selector.Category[0])
		image := element.ChildAttribute(selector.Image[0], selector.Image[1])
		link := element.ChildAttribute(selector.Link[0], selector.Link[1])
		date := element.ChildAttribute(selector.Date[0], selector.Date[1])

		title := element.ChildText(selector.Title[0])

		if len(image) != 0 {
			rawImage := strings.Split(image, ",")
			image = strings.Fields(rawImage[len(rawImage)-1])[0]

			image = util.ParseURL(src.URL, image)
			link = util.ParseURL(src.URL, link)
			date, _ = util.ParseTime(date)

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
func (src *BBCSource) CategoryPost(ctx context.Context, category string, page int) []*schema.NewsPost {
	category, err := util.ParseCategorySource(src.NewsSource, category)
	if err != nil {
		log.Println(err)
		return nil
	}
	response, err := util.RodNavigate(fmt.Sprintf("%s%s", src.URL, fmt.Sprintf(*src.CategoryPostURL, category, page)), true)
	if err != nil {
		log.Println(err)
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response)
	if err != nil {
		log.Println(err)
		return nil
	}
	return src.categoryPost(util.NewElement(document.Selection))
}

func (src *BBCSource) categoryPost(document *util.Element) []*schema.NewsPost {
	selector := src.CategoryPostSelector
	result := make([]*schema.NewsPost, 0)
	document.ForEach(selector.List[0], func(i int, element *util.Element) {
		// category := element.ChildText(selector.Category[0])
		image := element.ChildAttribute(selector.Image[0], selector.Image[1])
		link := element.ChildAttribute(selector.Link[0], selector.Link[1])
		date := element.ChildAttribute(selector.Date[0], selector.Date[1])

		title := element.ChildText(selector.Title[0])

		if len(image) != 0 {
			rawImage := strings.Split(image, ",")
			image = strings.Fields(rawImage[len(rawImage)-1])[0]

			image = util.ParseURL(src.URL, image)
			link = util.ParseURL(src.URL, link)
			date, _ = util.ParseTime(date)

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

/// PostArticle
///////////////
func (src *BBCSource) NewsArticle(ctx context.Context, link string) *schema.NewsArticle {
	response, err := util.RodNavigate(link)
	if err != nil {
		log.Println(err)
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response)
	if err != nil {
		log.Println(err)
		return nil
	}
	return src.newsArticle(util.NewElement(document.Selection))
}

func (src *BBCSource) newsArticle(document *util.Element) *schema.NewsArticle {
	selector := src.ArticleSelector
	document.Selection.Find(selector.Description[1]).Each(func(i int, s *goquery.Selection) {
		data, _ := s.Html()
		s.ReplaceWithHtml(html.UnescapeString(data))
	})
	contents := document.ChildrenOuterHtmls(selector.Description[0])
	description := strings.Join(contents, "")
	return &schema.NewsArticle{
		Description: description,
	}
}
