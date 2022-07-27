package trash

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"news/internal/sources"
	"news/internal/storage"
	"news/internal/storage/custom"

	"github.com/PuerkitoBio/goquery"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const AbidjanNetArticleName = "Abidjan.Net"

type AbidjanNetArticleSource struct {
	*storage.Source
	*http.Client
}

func NewAbidjanNetArticleSource(source *storage.Source) *AbidjanNetArticleSource {
	return &AbidjanNetArticleSource{
		Client: http.DefaultClient,
		Source: source,
	}
}

/// LatestPost
///
///
func (src *AbidjanNetArticleSource) LatestPost(ctx context.Context) []*custom.ArticlePost {
	response, err := sources.RodNavigate(fmt.Sprintf("%s%s", src.URL, src.ArticleFeaturedPostURL))
	if err != nil {
		log.Println(err)
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response)
	if err != nil {
		log.Println(err)
		return nil
	}
	return src.latestPost(sources.NewElement(document.Selection))
}

func (src *AbidjanNetArticleSource) latestPost(document *sources.Element) []*custom.ArticlePost {
	selector := src.ArticleFeaturedPostSelector
	result := make([]*custom.ArticlePost, 0)

	elementCallBack := func(element *sources.Element) {
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

		image = sources.ParseURL(src.URL, image)
		if strings.Contains(image, "defaut-cover-photo.svg") {
			image = ""
		}

		link = sources.ParseURL(src.URL, link)
		dateTime, _ := sources.ParseTime(date)

		if !strings.Contains(strings.Join(value, ""), "Fraternité Matin") && len(image) != 0 {
			result = append(result, &custom.ArticlePost{
				Date:   timestamppb.New(dateTime),
				Source: src.Name,
				Logo:   src.Logo,
				Image:  image,
				Title:  title,
				Link:   link,
			})
		}
	}

	elementCallBack(sources.NewElement(document.Selection.Find(selector.List[1])))

	document.ForEach(selector.List[0],
		func(_ int, element *sources.Element) {
			elementCallBack(element)
		})
	return result
}

/// CategoryPost
///
///
func (src *AbidjanNetArticleSource) CategoryPost(ctx context.Context, category string, page int) []*custom.ArticlePost {
	response, err := sources.RodNavigate(fmt.Sprintf("%s%s", src.URL, fmt.Sprintf(src.ArticleCategoryPostURL, src.ArticleCategories[category], page)))
	if err != nil {
		log.Println(err)
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response)
	if err != nil {
		log.Println(err)
		return nil
	}
	return src.categoryPost(sources.NewElement(document.Selection))
}

func (src *AbidjanNetArticleSource) categoryPost(document *sources.Element) []*custom.ArticlePost {
	selector := src.ArticleCategoryPostSelector
	result := make([]*custom.ArticlePost, 0)
	document.ForEach(selector.List[0],
		func(_ int, element *sources.Element) {
			// category := element.ChildText(selector.Category[0])
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			title := element.ChildText(selector.Title[0])
			link := element.Attribute(selector.Link[0])
			date := element.ChildText(selector.Date[0])

			value := strings.Split(date, "-")
			date = strings.TrimSpace(value[len(value)-1])

			image = sources.ParseURL(src.URL, image)
			if strings.Contains(image, "defaut-cover-photo.svg") {
				image = ""
			}

			link = sources.ParseURL(src.URL, link)
			dateTime, _ := sources.ParseTime(date)

			if !strings.Contains(strings.Join(value, ""), "Fraternité Matin") && len(image) != 0 {
				result = append(result, &custom.ArticlePost{
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

/// NewsArticle
///
///

func (src *AbidjanNetArticleSource) NewsArticle(ctx context.Context, link string) *custom.ArticlePost {
	response, err := sources.RodNavigate(link)
	if err != nil {
		log.Println(err)
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response)
	if err != nil {
		log.Println(err)
		return nil
	}
	return src.newsArticle(sources.NewElement(document.Selection))
}

func (src *AbidjanNetArticleSource) newsArticle(document *sources.Element) *custom.ArticlePost {
	selector := src.ArticleContentSelector
	content := strings.Join(strings.Fields(document.ChildOuterHtml(selector.Content[0])), "")
	return &custom.ArticlePost{Content: content}
}
