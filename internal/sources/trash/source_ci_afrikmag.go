package trash

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"

	"news/internal/sources"
	"news/internal/storage"
	"news/internal/storage/custom"

	"github.com/PuerkitoBio/goquery"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const AfrikMagArticleName = "AfrikMag"

type AfrikMagArticleSource struct {
	*storage.Source
	*http.Client
}

func NewAfrikMagArticleSource(source *storage.Source) *AfrikMagArticleSource {
	return &AfrikMagArticleSource{
		Client: http.DefaultClient,
		Source: source,
	}
}

/// NewsLatest
///
///
func (src *AfrikMagArticleSource) LatestPost(ctx context.Context) []*custom.ArticlePost {
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

func (src *AfrikMagArticleSource) latestPost(document *sources.Element) []*custom.ArticlePost {
	selector := src.ArticleFeaturedPostSelector
	result := make([]*custom.ArticlePost, 0)
	document.ForEach(selector.List[0],
		func(_ int, element *sources.Element) {
			// category := element.ChildText(selector.Category[0])
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			link := element.ChildAttribute(selector.Link[0], selector.Link[1])
			title := element.ChildText(selector.Title[0])
			date := element.ChildText(selector.Date[0])
			if len(image) == 0 {
				image = element.ChildAttribute(selector.Image[2], selector.Image[3])
			}

			image = sources.ParseURL(src.URL, image)
			dateTime, _ := sources.ParseTime(date)

			image = strings.ReplaceAll(image, fmt.Sprintf("-220x150%v", path.Ext(image)), path.Ext(image))
			result = append(result, &custom.ArticlePost{
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

func (src *AfrikMagArticleSource) CategoryPost(ctx context.Context, category string, page int) []*custom.ArticlePost {
	response, err := sources.RodPostRequest(
		fmt.Sprintf("%s%s", src.URL, src.ArticleCategoryPostURL),
		url.Values{
			"query":    []string{fmt.Sprintf("{'cat':%v,'lazy_load_term_meta':true,'posts_per_page':16,'order':'DESC'}", src.ArticleCategories[category])},
			"action":   []string{"tie_archives_load_more"},
			"page":     []string{strconv.Itoa(page)},
			"layout":   []string{"default"},
			"settings": []string{"{'uncropped_image':'jannah-image-post','category_meta':false,'post_meta':true,'excerpt':'true','excerpt_length':'20','read_more':'true','read_more_text':false,'media_overlay':false,'title_length':0,'is_full':false,'is_category':true}"},
		}.Encode())
	if err != nil {
		log.Println(err)
		return nil
	}
	data := make(map[string]interface{}, 0)

	b, _ := ioutil.ReadAll(response)
	var value string
	json.Unmarshal(b, &value)
	json.Unmarshal([]byte(value), &data)

	document, err := goquery.NewDocumentFromReader(strings.NewReader((data["code"]).(string)))
	if err != nil {
		log.Println(err)
		return nil
	}
	return src.categoryPost(sources.NewElement(document.Selection))
}

func (src *AfrikMagArticleSource) categoryPost(document *sources.Element) []*custom.ArticlePost {
	selector := src.ArticleCategoryPostSelector
	result := make([]*custom.ArticlePost, 0)
	document.ForEach(selector.List[0],
		func(_ int, element *sources.Element) {
			// category := element.ChildText(selector.Category[0])
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			link := element.ChildAttribute(selector.Link[0], selector.Link[1])
			title := element.ChildText(selector.Title[0])
			date := element.ChildText(selector.Date[0])

			image = sources.ParseURL(src.URL, image)
			dateTime, _ := sources.ParseTime(date)

			image = strings.ReplaceAll(image, fmt.Sprintf("-220x150%v", path.Ext(image)), path.Ext(image))
			result = append(result, &custom.ArticlePost{
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

func (src *AfrikMagArticleSource) NewsArticle(ctx context.Context, link string) *custom.ArticlePost {
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

func (src *AfrikMagArticleSource) newsArticle(document *sources.Element) *custom.ArticlePost {
	selector := src.ArticleContentSelector
	content := strings.Join(strings.Fields(strings.Join(document.ChildrenOuterHtmls(selector.Content[0]), "<br>")), " ")
	return &custom.ArticlePost{Content: content}
}
