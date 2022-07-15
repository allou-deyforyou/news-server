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
	"google.golang.org/protobuf/types/known/timestamppb"
)

const RFIName = "RFI"

type RFISource struct {
	*store.NewsArticleSource
	*http.Client
}

func NewRFISource(source *store.NewsArticleSource) *RFISource {
	return &RFISource{
		Client:            http.DefaultClient,
		NewsArticleSource: source,
	}
}

/// LatestPost
///
///
func (src *RFISource) LatestPost(ctx context.Context) []*schema.NewsArticlePost {
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

func (src *RFISource) latestPost(document *util.Element) []*schema.NewsArticlePost {
	selector := src.LatestPostSelector
	result := make([]*schema.NewsArticlePost, 0)
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
				dateTime, _ := util.ParseTime(date)

				result = append(result, &schema.NewsArticlePost{
					Date:   timestamppb.New(dateTime),
					Source: src.Name,
					Logo:   src.Logo,
					Image:  image,
					Title:  title,
					Link:   link,
				})
			}
		}

	})
	return result
}

/// NewsCategory
////////////////
func (src *RFISource) CategoryPost(ctx context.Context, category string, page int) []*schema.NewsArticlePost {
	if page != 1 {
		return nil
	}
	category, err := util.ParseCategorySource(src.NewsArticleSource, category)
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

func (src *RFISource) categoryPost(document *util.Element) []*schema.NewsArticlePost {
	selector := src.CategoryPostSelector
	result := make([]*schema.NewsArticlePost, 0)
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
			dateTime, _ := util.ParseTime(date)

			result = append(result, &schema.NewsArticlePost{
				Date:   timestamppb.New(dateTime),
				Source: src.Name,
				Logo:   src.Logo,
				Image:  image,
				Title:  title,
				Link:   link,
			})
		}
	})
	return result
}

/// PostArticle
///////////////
func (src *RFISource) NewsArticle(ctx context.Context, link string) *schema.NewsArticlePost {
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

func (src *RFISource) newsArticle(document *util.Element) *schema.NewsArticlePost {
	selector := src.ArticleSelector
	contents := document.ChildrenOuterHtmls(selector.Description[0])
	description := strings.Join(contents, "")
	return &schema.NewsArticlePost{
		Description: description,
	}
}
