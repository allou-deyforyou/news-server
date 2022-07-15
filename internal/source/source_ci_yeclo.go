package source

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"path"
	"regexp"
	"strings"

	"news/internal/store"
	"news/internal/store/schema"
	"news/internal/util"

	"github.com/PuerkitoBio/goquery"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const YecloName = "Yeclo"

type YecloSource struct {
	*store.NewsArticleSource
	*http.Client
}

func NewYecloSource(source *store.NewsArticleSource) *YecloSource {
	return &YecloSource{
		Client:            http.DefaultClient,
		NewsArticleSource: source,
	}
}

/// LatestPost
///
///
func (src *YecloSource) LatestPost(ctx context.Context) []*schema.NewsArticlePost {
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

func (src *YecloSource) latestPost(document *util.Element) []*schema.NewsArticlePost {
	selector := src.LatestPostSelector
	result := make([]*schema.NewsArticlePost, 0)
	document.ForEach(selector.List[0], func(i int, element *util.Element) {
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

/// CategoryPost
////////////////
func (src *YecloSource) CategoryPost(ctx context.Context, category string, page int) []*schema.NewsArticlePost {
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

func (src *YecloSource) categoryPost(document *util.Element) []*schema.NewsArticlePost {
	selector := src.CategoryPostSelector
	result := make([]*schema.NewsArticlePost, 0)
	document.ForEach(selector.List[0], func(i int, element *util.Element) {
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
func (src *YecloSource) NewsArticle(ctx context.Context, link string) *schema.NewsArticlePost {
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

func (src *YecloSource) newsArticle(document *util.Element) *schema.NewsArticlePost {
	selector := src.ArticleSelector
	contents := document.ChildrenOuterHtmls(selector.Description[0])
	description := strings.Join(contents, "")
	return &schema.NewsArticlePost{
		Description: description,
	}
}
