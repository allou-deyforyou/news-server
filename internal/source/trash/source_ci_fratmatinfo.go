package trash

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"news/internal/store"
	"news/internal/store/schema"
	"news/internal/util"

	"github.com/PuerkitoBio/goquery"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const FratmatInfoName = "Fratmat Info"

type FratmatInfoSource struct {
	*store.NewsArticleSource
	*http.Client
}

func NewFratmatInfoSource(source *store.NewsArticleSource) *FratmatInfoSource {
	return &FratmatInfoSource{
		Client:            http.DefaultClient,
		NewsArticleSource: source,
	}
}

/// NewsLatest
//////////////
func (src *FratmatInfoSource) LatestPost(ctx context.Context) []*schema.NewsArticlePost {
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

func (src *FratmatInfoSource) latestPost(document *util.Element) []*schema.NewsArticlePost {
	selector := src.LatestPostSelector
	result := make([]*schema.NewsArticlePost, 0)
	document.ForEach(selector.List[0],
		func(i int, element *util.Element) {
			// category := element.ChildText(selector.Category[0])
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			link := element.ChildAttribute(selector.Link[0], selector.Link[1])
			title := element.ChildText(selector.Title[0])
			date := element.ChildText(selector.Date[0])

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
		})
	return result
}

/// NewsCategory
////////////////
func (src *FratmatInfoSource) CategoryPost(ctx context.Context, category string, page int) []*schema.NewsArticlePost {
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

func (src *FratmatInfoSource) categoryPost(document *util.Element) []*schema.NewsArticlePost {
	selector := src.CategoryPostSelector
	result := make([]*schema.NewsArticlePost, 0)
	document.ForEach(selector.List[0],
		func(i int, element *util.Element) {
			// category := element.ChildText(selector.Category[0])
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			link := element.ChildAttribute(selector.Link[0], selector.Link[1])
			title := element.ChildText(selector.Title[0])
			date := element.ChildText(selector.Date[0])

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
		})
	return result
}

/// PostArticle
///////////////
func (src *FratmatInfoSource) NewsArticle(ctx context.Context, link string) *schema.NewsArticlePost {
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

func (src *FratmatInfoSource) newsArticle(document *util.Element) *schema.NewsArticlePost {
	selector := src.ArticleSelector
	description := document.ChildOuterHtml(selector.Description[0])
	description = strings.Join(strings.Fields(description), " ")
	return &schema.NewsArticlePost{
		Description: description,
	}
}
