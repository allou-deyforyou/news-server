package sources

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"news/internal/storage"
	"news/internal/storage/custom"

	"github.com/PuerkitoBio/goquery"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const RTIInfoName = "RTI Info"

type RTIInfoSource struct {
	*storage.Source
	*http.Client
}

func NewRTIInfoSource(source *storage.Source) *RTIInfoSource {
	return &RTIInfoSource{
		Client: http.DefaultClient,
		Source: source,
	}
}

/// ArticleFeaturedPostList
///
func (src *RTIInfoSource) ArticleFeaturedPostList(ctx context.Context) []*custom.ArticlePost {
	response, err := RodGetRequest(fmt.Sprintf("%s%s", src.URL, src.ArticleFeaturedPostURL))
	if err != nil {
		log.Println(err)
		return nil
	}
	data := make(map[string]interface{})
	json.NewDecoder(response).Decode(&data)
	return src.articleFeaturedPostList(data)
}

func (src *RTIInfoSource) articleFeaturedPostList(data map[string]interface{}) []*custom.ArticlePost {
	selector := src.ArticleFeaturedPostSelector
	result := make([]*custom.ArticlePost, 0)
	for _, value := range data[selector.List[0]].([]interface{}) {
		data := value.(map[string]interface{})
		image := fmt.Sprintf(selector.Image[1], data[selector.Image[0]].(string))
		link := fmt.Sprintf(selector.Link[1], data[selector.Link[0]].(float64))
		title := data[selector.Title[0]].(string)
		date := data[selector.Date[0]].(string)

		link = src.URL + link
		dateTime, _ := ParseTime(date)

		result = append(result, &custom.ArticlePost{
			Date:   timestamppb.New(dateTime),
			Source: src.Name,
			Logo:   src.Logo,
			Title:  title,
			Image:  image,
			Link:   link,
		})
	}
	return result
}

/// ArticleCategoryPostList
////////////////
func (src *RTIInfoSource) ArticleCategoryPostList(ctx context.Context, category string, page int) []*custom.ArticlePost {
	response, err := RodGetRequest(fmt.Sprintf("%s%s", src.URL, fmt.Sprintf(src.ArticleCategoryPostURL, src.ArticleCategories[category])))
	if err != nil {
		log.Println(err)
		return nil
	}
	data := make(map[string]interface{})
	json.NewDecoder(response).Decode(&data)
	return src.articleCategoryPostList(data)
}

func (src *RTIInfoSource) articleCategoryPostList(data map[string]interface{}) []*custom.ArticlePost {
	selector := src.ArticleCategoryPostSelector
	result := make([]*custom.ArticlePost, 0)
	for _, value := range data[selector.List[0]].(map[string]interface{})[selector.List[1]].([]interface{}) {
		data := value.(map[string]interface{})
		link := fmt.Sprintf(selector.Link[1], data[selector.Link[0]].(float64))
		title := data[selector.Title[0]].(string)
		image := fmt.Sprintf(selector.Image[1], data[selector.Image[0]].(string))
		date := data[selector.Date[0]].(string)

		link = src.URL + link
		dateTime, _ := ParseTime(date)

		result = append(result, &custom.ArticlePost{
			Date:   timestamppb.New(dateTime),
			Source: src.Name,
			Logo:   src.Logo,
			Title:  title,
			Image:  image,
			Link:   link,
		})
	}
	return result
}

/// ArticleContent
///////////////
func (src *RTIInfoSource) ArticleContent(ctx context.Context, link string) *custom.ArticlePost {
	response, err := RodGetRequest(link)
	if err != nil {
		log.Println(err)
		return nil
	}
	data := make(map[string]interface{})
	json.NewDecoder(response).Decode(&data)
	return src.articleContent(data)
}

func (src *RTIInfoSource) articleContent(data map[string]interface{}) *custom.ArticlePost {
	selector := src.ArticleContentSelector
	value := data[selector.Content[0]].(map[string]interface{})[selector.Content[1]].(string)
	document, _ := goquery.NewDocumentFromReader(strings.NewReader(value))
	element := NewElement(document.Selection.Find("*").RemoveClass().RemoveAttr("style"))
	element.ForEach("p", func(_ int, e *Element) {
		innerHtml := strings.TrimSpace(e.InnerHtml())
		if innerHtml == "<br/>" || innerHtml == "<span> </span>" || len(innerHtml) == 0 {
			e.Selection.Remove()
		}
	})
	content := element.OuterHtml()
	return &custom.ArticlePost{Content: content}
}

/// MediaLivePostList
//////////////
func (src *RTIInfoSource) MediaLivePostList(ctx context.Context) []*custom.MediaPost {
	return nil
}

/// MediaCategoryPostList
//////////////
func (src *RTIInfoSource) MediaCategoryPostList(ctx context.Context, category string, page int) []*custom.MediaPost {
	response, err := RodGetRequest(fmt.Sprintf("%s%s", src.URL, fmt.Sprintf(src.MediaCategoryPostURL, src.MediaCategories[category])))
	if err != nil {
		log.Println(err)
		return nil
	}
	data := make(map[string]interface{})
	json.NewDecoder(response).Decode(&data)
	return src.mediaCategoryPostList(data)
}

func (src *RTIInfoSource) mediaCategoryPostList(data map[string]interface{}) []*custom.MediaPost {
	selector := src.MediaCategoryPostSelector
	result := make([]*custom.MediaPost, 0)
	for _, value := range data[selector.List[0]].(map[string]interface{})[selector.List[1]].([]interface{}) {
		data := value.(map[string]interface{})
		link := fmt.Sprintf(selector.Link[1], data[selector.Link[0]].(float64))
		title := data[selector.Title[0]].(string)
		image := fmt.Sprintf(selector.Image[1], data[selector.Image[0]].(string))
		date := data[selector.Date[0]].(string)

		link = src.URL + link
		dateTime, _ := ParseTime(date)

		result = append(result, &custom.MediaPost{
			Date:   timestamppb.New(dateTime),
			Source: src.Name,
			Logo:   src.Logo,
			Title:  title,
			Image:  image,
			Link:   link,
		})
	}
	return result
}

/// MediaContent
//////////////
func (src *RTIInfoSource) MediaContent(ctx context.Context, link string) *custom.MediaPost {
	response, err := RodGetRequest(link)
	if err != nil {
		log.Println(err)
		return nil
	}
	data := make(map[string]interface{})
	json.NewDecoder(response).Decode(&data)
	return src.mediaContent(data)
}

func (src *RTIInfoSource) mediaContent(data map[string]interface{}) *custom.MediaPost {
	selector := src.MediaContentSelector
	data = data[selector.Content[0]].(map[string]interface{})
	for _, value := range data[selector.Content[1]].([]interface{}) {
		data := value.(map[string]interface{})
		for _, value := range data[selector.Content[2]].([]interface{}) {
			data := value.(map[string]interface{})
			file := data[selector.Content[3]].(string)
			return &custom.MediaPost{
				Type:    custom.MediaPost_YOUTUBE,
				Content: ExtractVideoID(file),
			}
		}
	}
	return nil
}
