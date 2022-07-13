package trash

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"news/internal/store"
	"news/internal/store/schema"
	"news/internal/util"

	"github.com/PuerkitoBio/goquery"
)

const AbidjanNetName = "Abidjan.Net"

type AbidjanNetSource struct {
	*store.NewsArticleSource
	*http.Client
}

func NewAbidjanNetSource(source *store.NewsArticleSource) *AbidjanNetSource {
	return &AbidjanNetSource{
		Client:     http.DefaultClient,
		NewsArticleSource: source,
	}
}

/// LatestPost
///
///
func (src *AbidjanNetSource) LatestPost(ctx context.Context) []*schema.NewsArticlePost {
	response, err := util.RodNavigate(fmt.Sprintf("%s%s", src.URL, *src.LatestPostURL))
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

func (src *AbidjanNetSource) latestPost(document *util.Element) []*schema.NewsArticlePost {
	selector := src.LatestPostSelector
	filmList := make([]*schema.NewsArticlePost, 0)

	elementCallBack := func(element *util.Element) {
		// category := element.ChildText(selector.Category[0])
		image := element.ChildAttribute(selector.Image[0], selector.Image[1])
		title := element.ChildText(selector.Title[0])
		link := element.Attribute(selector.Link[0])
		date := element.ChildText(selector.Date[0])
		if len(image) == 0 {
			image = element.ChildAttribute(selector.Image[0], selector.Image[2])
			if len(image) == 0 {
				style := element.ChildAttribute(selector.Image[3], "style")
				exp := regexp.MustCompile(`(http(s|)://.*')`)
				image = strings.Trim(exp.FindString(style), "'")
			}
		}
		value := strings.Split(date, "-")
		date = strings.TrimSpace(value[len(value)-1])

		image = util.ParseURL(src.URL, image)
		link = util.ParseURL(src.URL, link)
		date, _ = util.ParseTime(date)

		if strings.Contains(image, "defaut-cover-photo.svg") {
			image = ""
		}

		if !strings.Contains(strings.Join(value, ""), "Fraternité Matin") && len(image) != 0 {
			filmList = append(filmList, &schema.NewsArticlePost{
				Source: src.Name,
				Logo:   src.Logo,
				Image:  image,
				Title:  title,
				Link:   link,
				Date:   date,
			})
		}
	}

	elementCallBack(util.NewElement(document.Selection.Find(selector.List[1])))

	document.ForEach(selector.List[0],
		func(i int, element *util.Element) {
			elementCallBack(element)
		})
	return filmList
}

/// CategoryPost
///
///
func (src *AbidjanNetSource) CategoryPost(ctx context.Context, category string, page int) []*schema.NewsArticlePost {
	category, err := util.ParseCategorySource(src.NewsArticleSource, category)
	if err != nil {
		log.Println(err)
		return nil
	}
	response, err := util.RodNavigate(fmt.Sprintf("%s%s", src.URL, fmt.Sprintf(*src.CategoryPostURL, category, page)))
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

func (src *AbidjanNetSource) categoryPost(document *util.Element) []*schema.NewsArticlePost {
	selector := src.LatestPostSelector
	filmList := make([]*schema.NewsArticlePost, 0)
	document.ForEach(selector.List[0],
		func(i int, element *util.Element) {
			// category := element.ChildText(selector.Category[0])
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			title := element.ChildText(selector.Title[0])
			link := element.Attribute(selector.Link[0])
			date := element.ChildText(selector.Date[0])

			value := strings.Split(date, "-")
			date = strings.TrimSpace(value[len(value)-1])

			image = util.ParseURL(src.URL, image)
			link = util.ParseURL(src.URL, link)
			date, _ = util.ParseTime(date)

			if strings.Contains(image, "defaut-cover-photo.svg") {
				image = ""
			}

			if !strings.Contains(strings.Join(value, ""), "Fraternité Matin") && len(image) != 0 {
				filmList = append(filmList, &schema.NewsArticlePost{
					Source: src.Name,
					Logo:   src.Logo,
					Image:  image,
					Title:  title,
					Link:   link,
					Date:   date,
				})
			}
		})
	return filmList
}

/// NewsArticle
///
///

func (src *AbidjanNetSource) NewsArticle(ctx context.Context, link string) *schema.NewsArticlePost {
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

func (src *AbidjanNetSource) newsArticle(document *util.Element) *schema.NewsArticlePost {
	selector := src.ArticleSelector
	description := document.ChildOuterHtml(selector.Description[0])
	description = strings.Join(strings.Fields(description), " ")
	return &schema.NewsArticlePost{
		Description: description,
	}
}
