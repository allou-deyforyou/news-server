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

	"news/internal/store"
	"news/internal/store/schema"
	"news/internal/util"

	"github.com/PuerkitoBio/goquery"
)

const AfrikMagName = "AfrikMag"

type AfrikMagSource struct {
	*store.NewsArticleSource
	*http.Client
}

func NewAfrikMagSource(source *store.NewsArticleSource) *AfrikMagSource {
	return &AfrikMagSource{
		Client:     http.DefaultClient,
		NewsArticleSource: source,
	}
}

/// NewsLatest
///
///
func (src *AfrikMagSource) LatestPost(ctx context.Context) []*schema.NewsArticlePost {
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

func (src *AfrikMagSource) latestPost(document *util.Element) []*schema.NewsArticlePost {
	selector := src.LatestPostSelector
	filmList := make([]*schema.NewsArticlePost, 0)
	document.ForEach(selector.List[0],
		func(i int, element *util.Element) {
			// category := element.ChildText(selector.Category[0])
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			link := element.ChildAttribute(selector.Link[0], selector.Link[1])
			title := element.ChildText(selector.Title[0])
			date := element.ChildText(selector.Date[0])
			if len(image) == 0 {
				image = element.ChildAttribute(selector.Image[2], selector.Image[3])
			}

			image = util.ParseURL(src.URL, image)
			date, _ = util.ParseTime(date)

			image = strings.ReplaceAll(image, fmt.Sprintf("-220x150%v", path.Ext(image)), path.Ext(image))
			filmList = append(filmList, &schema.NewsArticlePost{
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

func (src *AfrikMagSource) CategoryPost(ctx context.Context, category string, page int) []*schema.NewsArticlePost {
	category, err := util.ParseCategorySource(src.NewsArticleSource, category)
	if err != nil {
		return nil
	}
	response, err := util.RodPostRequest(
		fmt.Sprintf("%s%s", src.URL, *src.CategoryPostURL),
		url.Values{
			"query":    []string{fmt.Sprintf("{'cat':%v,'lazy_load_term_meta':true,'posts_per_page':16,'order':'DESC'}", category)},
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
	return src.categoryPost(util.NewElement(document.Selection))
}

func (src *AfrikMagSource) categoryPost(document *util.Element) []*schema.NewsArticlePost {
	selector := src.CategoryPostSelector
	filmList := make([]*schema.NewsArticlePost, 0)
	document.ForEach(selector.List[0],
		func(i int, element *util.Element) {
			// category := element.ChildText(selector.Category[0])
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			link := element.ChildAttribute(selector.Link[0], selector.Link[1])
			title := element.ChildText(selector.Title[0])
			date := element.ChildText(selector.Date[0])

			image = util.ParseURL(src.URL, image)
			date, _ = util.ParseTime(date)

			image = strings.ReplaceAll(image, fmt.Sprintf("-220x150%v", path.Ext(image)), path.Ext(image))
			filmList = append(filmList, &schema.NewsArticlePost{
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

func (src *AfrikMagSource) NewsArticle(ctx context.Context, link string) *schema.NewsArticlePost {
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

func (src *AfrikMagSource) newsArticle(document *util.Element) *schema.NewsArticlePost {
	selector := src.ArticleSelector
	description := strings.Join(document.ChildrenOuterHtmls(selector.Description[0]), "<br>")
	description = strings.Join(strings.Fields(description), " ")
	return &schema.NewsArticlePost{
		Description: description,
	}
}
