package source

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"

	"news/internal/store"
	"news/internal/store/schema"
	"news/internal/util"

	"github.com/PuerkitoBio/goquery"
)

const RFIName = "RFI"

type RFISource struct {
	*store.NewsSource
	*http.Client
}

func NewRFISource(source *store.NewsSource) *RFISource {
	return &RFISource{
		Client:     http.DefaultClient,
		NewsSource: source,
	}
}

/// LatestPost
///
///
func (src *RFISource) LatestPost(ctx context.Context) []*schema.NewsPost {
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

func (src *RFISource) latestPost(document *util.Element) []*schema.NewsPost {
	selector := src.LatestPostSelector
	result := make([]*schema.NewsPost, 0)
	document.ForEach(selector.List[0], func(i int, element *util.Element) {
		// category := element.ChildText(selector.Category[0])
		image := element.ChildAttribute(selector.Image[0], selector.Image[1])
		link := element.ChildAttribute(selector.Link[0], selector.Link[1])
		title := element.ChildText(selector.Title[0])

		link = util.ParseURL(src.URL, link)
		if strings.HasPrefix(link, src.URL) {
			if len(title) != 0 {
				rawImage := strings.Split(image, ",")
				image = strings.Fields(rawImage[len(rawImage)-1])[0]

				date := strings.Split(path.Base(link), "-")[0]
				date = fmt.Sprintf("%v-%v-%v", string(date[:4]), string(date[4:6]), string(date[6:8]))

				image = util.ParseURL(src.URL, image)
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
		}

	})
	return result
}

/// NewsCategory
////////////////
func (src *RFISource) CategoryPost(ctx context.Context, category string, page int) []*schema.NewsPost {
	if page != 1 {
		return nil
	}
	category, err := util.ParseCategorySource(src.NewsSource, category)
	if err != nil {
		log.Println(err)
		return nil
	}
	response, err := util.RodNavigate(fmt.Sprintf("%s%s", src.URL, fmt.Sprintf(*src.CategoryPostURL, category)))
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

func (src *RFISource) categoryPost(document *util.Element) []*schema.NewsPost {
	selector := src.CategoryPostSelector
	result := make([]*schema.NewsPost, 0)
	document.ForEach(selector.List[0], func(i int, element *util.Element) {
		// category := element.ChildText(selector.Category[0])
		image := element.ChildAttribute(selector.Image[0], selector.Image[1])
		link := element.ChildAttribute(selector.Link[0], selector.Link[1])
		title := element.ChildText(selector.Title[0])

		if strings.HasPrefix(link, src.URL) {
			rawImage := strings.Split(image, ",")
			image = strings.Fields(rawImage[len(rawImage)-1])[0]

			date := strings.Split(path.Base(link), "-")[0]
			date = fmt.Sprintf("%v-%v-%v", string(date[:4]), string(date[4:6]), string(date[6:8]))

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
func (src *RFISource) NewsArticle(ctx context.Context, link string) *schema.NewsArticle {
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

func (src *RFISource) newsArticle(document *util.Element) *schema.NewsArticle {
	selector := src.ArticleSelector
	contents := document.ChildrenOuterHtmls(selector.Description[0])
	description := strings.Join(contents, "")
	return &schema.NewsArticle{
		Description: description,
	}
}
