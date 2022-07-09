package source

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"news/internal/source/sutil"
	"news/internal/store"
	"news/internal/store/schema"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const France24Name = "France 24"

type France24Source struct {
	*store.NewsSource
	*http.Client
}

func NewFrance24Source(source *store.NewsSource) *France24Source {
	return &France24Source{
		Client:     http.DefaultClient,
		NewsSource: source,
	}
}

/// LatestPost
///
///
func (src *France24Source) LatestPost(ctx context.Context) []*schema.NewsPost {
	response, err := sutil.RodGetRequest(fmt.Sprintf("%s%s", src.URL, *src.LatestPostURL))
	if err != nil {
		log.Println(err)
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response)
	if err != nil {
		log.Println(err)
		return nil
	}
	return src.latestPost(sutil.NewElement(document.Selection))
}

func (src *France24Source) latestPost(document *sutil.Element) []*schema.NewsPost {
	selector := src.LatestPostSelector
	result := make([]*schema.NewsPost, 0)
	document.ForEach(selector.List[0], func(i int, element *sutil.Element) {
		// category := element.ChildText(selector.Category[0])
		image := element.ChildAttribute(selector.Image[0], selector.Image[1])
		link := element.ChildAttribute(selector.Link[0], selector.Link[1])
		title := element.ChildText(selector.Title[0])
		if len(image) != 0 {
			rawImage := strings.Split(image, ",")
			image = strings.Fields(rawImage[len(rawImage)-1])[0]

			date := strings.Split(path.Base(link), "-")[0]
			date = fmt.Sprintf("%v-%v-%v", string(date[:4]), string(date[4:6]), string(date[6:8]))

			image = sutil.ParseURL(src.URL, image)
			link = sutil.ParseURL(src.URL, link)
			date, _ = sutil.ParseTime(date)

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
func (src *France24Source) CategoryPost(ctx context.Context, category string, page int) []*schema.NewsPost {
	if page != 1 {
		return nil
	}
	category, err := sutil.ParseCategorySource(src.NewsSource, category)
	if err != nil {
		log.Println(err)
		return nil
	}
	response, err := sutil.RodGetRequest(fmt.Sprintf("%s%s", src.URL, fmt.Sprintf(*src.CategoryPostURL, category)))
	if err != nil {
		log.Println(err)
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response)
	if err != nil {
		log.Println(err)
		return nil
	}
	return src.categoryPost(sutil.NewElement(document.Selection))
}

func (src *France24Source) categoryPost(document *sutil.Element) []*schema.NewsPost {
	selector := src.CategoryPostSelector
	result := make([]*schema.NewsPost, 0)
	document.ForEach(selector.List[0], func(i int, element *sutil.Element) {
		// category := element.ChildText(selector.Category[0])
		image := element.ChildAttribute(selector.Image[0], selector.Image[1])
		link := element.ChildAttribute(selector.Link[0], selector.Link[1])
		title := element.ChildText(selector.Title[0])
		if len(image) != 0 {

			rawImage := strings.Split(image, ",")
			image = strings.Fields(rawImage[len(rawImage)-1])[0]

			date := strings.Split(path.Base(link), "-")[0]
			date = fmt.Sprintf("%v-%v-%v", string(date[:4]), string(date[4:6]), string(date[6:8]))

			image = sutil.ParseURL(src.URL, image)
			link = sutil.ParseURL(src.URL, link)
			date, _ = sutil.ParseTime(date)

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
func (src *France24Source) NewsArticle(ctx context.Context, link string) *schema.NewsArticle {
	response, err := sutil.RodGetRequest(link)
	if err != nil {
		log.Println(err)
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response)
	if err != nil {
		log.Println(err)
		return nil
	}
	return src.newsArticle(sutil.NewElement(document.Selection))
}

func (src *France24Source) newsArticle(document *sutil.Element) *schema.NewsArticle {
	selector := src.ArticleSelector
	contents := document.ChildrenOuterHtmls(selector.Description[0])
	description := strings.Join(contents, "")
	return &schema.NewsArticle{
		Description: description,
	}
}
