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
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const YecloName = "Yeclo"

type YecloSource struct {
	*store.NewsSource
	*http.Client
}

func NewYecloSource(source *store.NewsSource) *YecloSource {
	return &YecloSource{
		Client:     http.DefaultClient,
		NewsSource: source,
	}
}

/// LatestPost
///
///
func (src *YecloSource) LatestPost(ctx context.Context) []*schema.NewsPost {
	response, err := sutil.RodNavigate(fmt.Sprintf("%s%s", src.URL, *src.LatestPostURL))
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

func (src *YecloSource) latestPost(document *sutil.Element) []*schema.NewsPost {
	selector := src.LatestPostSelector
	result := make([]*schema.NewsPost, 0)
	document.ForEach(selector.List[0], func(i int, element *sutil.Element) {
		// category := element.ChildText(selector.Category[0])
		image := element.ChildAttribute(selector.Image[0], selector.Image[1])
		link := element.ChildAttribute(selector.Link[0], selector.Link[1])
		date := element.ChildAttribute(selector.Date[0], selector.Date[1])
		title := element.ChildText(selector.Title[0])

		if len(image) != 0 {
			if len(regexp.MustCompile(`-\d{3}x\d{3}.`).FindString(image)) != 0 {
				rawPath := strings.Split(image, "-")
				image = strings.ReplaceAll(image, fmt.Sprintf("-%v", rawPath[len(rawPath)-1]), path.Ext(image))
			}

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

/// CategoryPost
////////////////
func (src *YecloSource) CategoryPost(ctx context.Context, category string, page int) []*schema.NewsPost {
	category, err := sutil.ParseCategorySource(src.NewsSource, category)
	if err != nil {
		log.Println(err)
		return nil
	}
	response, err := sutil.RodNavigate(fmt.Sprintf("%s%s", src.URL, fmt.Sprintf(*src.CategoryPostURL, category, page)))
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

func (src *YecloSource) categoryPost(document *sutil.Element) []*schema.NewsPost {
	selector := src.CategoryPostSelector
	result := make([]*schema.NewsPost, 0)
	document.ForEach(selector.List[0], func(i int, element *sutil.Element) {
		// category := element.ChildText(selector.Category[0])
		image := element.ChildAttribute(selector.Image[0], selector.Image[1])
		link := element.ChildAttribute(selector.Link[0], selector.Link[1])
		date := element.ChildAttribute(selector.Date[0], selector.Date[1])
		title := element.ChildText(selector.Title[0])

		if len(image) != 0 {
			if len(regexp.MustCompile(`-\d{3}x\d{3}.`).FindString(image)) != 0 {
				rawPath := strings.Split(image, "-")
				image = strings.ReplaceAll(image, fmt.Sprintf("-%v", rawPath[len(rawPath)-1]), path.Ext(image))
			}

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
func (src *YecloSource) NewsArticle(ctx context.Context, link string) *schema.NewsArticle {
	response, err := sutil.RodNavigate(link)
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

func (src *YecloSource) newsArticle(document *sutil.Element) *schema.NewsArticle {
	selector := src.ArticleSelector
	contents := document.ChildrenOuterHtmls(selector.Description[0])
	description := strings.Join(contents, "")
	return &schema.NewsArticle{
		Description: description,
	}
}
