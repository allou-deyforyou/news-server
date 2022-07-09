package trash

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"news/internal/source/sutil"
	"news/internal/store"
	"news/internal/store/schema"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const FratmatInfoName = "Fratmat Info"

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

func (src *FratmatInfoSource) latestPost(document *sutil.Element) []*schema.NewsPost {
	selector := src.LatestPostSelector
	filmList := make([]*schema.NewsPost, 0)
	document.ForEach(selector.List[0],
		func(i int, element *sutil.Element) {
			// category := element.ChildText(selector.Category[0])
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			link := element.ChildAttribute(selector.Link[0], selector.Link[1])
			title := element.ChildText(selector.Title[0])
			date := element.ChildText(selector.Date[0])

			image = sutil.ParseURL(src.URL, image)
			date, _ = sutil.ParseTime(date)

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
	category, err := sutil.ParseCategorySource(src.NewsSource, category)
	if err != nil {
		log.Println(err)
		return nil
	}
	response, err := sutil.RodGetRequest(fmt.Sprintf("%s%s", src.URL, fmt.Sprintf(*src.CategoryPostURL, category, page)))
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

func (src *FratmatInfoSource) categoryPost(document *sutil.Element) []*schema.NewsPost {
	selector := src.CategoryPostSelector
	filmList := make([]*schema.NewsPost, 0)
	document.ForEach(selector.List[0],
		func(i int, element *sutil.Element) {
			// category := element.ChildText(selector.Category[0])
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			link := element.ChildAttribute(selector.Link[0], selector.Link[1])
			title := element.ChildText(selector.Title[0])
			date := element.ChildText(selector.Date[0])

			image = sutil.ParseURL(src.URL, image)
			date, _ = sutil.ParseTime(date)

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

func (src *FratmatInfoSource) newsArticle(document *sutil.Element) *schema.NewsArticle {
	selector := src.ArticleSelector
	description := document.ChildOuterHtml(selector.Description[0])
	description = strings.Join(strings.Fields(description), " ")
	return &schema.NewsArticle{
		Description: description,
	}
}
