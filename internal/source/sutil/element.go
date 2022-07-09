package sutil

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Element struct {
	Selection *goquery.Selection
}

func NewElement(selection *goquery.Selection) *Element {
	return &Element{Selection: selection}
}

func (element *Element) Text() string {
	return element.Selection.Text()
}

func (element *Element) InnerHtml() string {
	content, _ := element.Selection.RemoveClass().Html()
	return strings.Join(strings.Fields(content), " ")
}

func (element *Element) OuterHtml() string {
	content, _ := goquery.OuterHtml(element.Selection.RemoveClass())
	return strings.Join(strings.Fields(content), " ")
}

func (element *Element) ChildInnerHtml(selector string) string {
	content, _ := element.Selection.Find(selector).RemoveClass().Html()
	return strings.Join(strings.Fields(content), " ")
}

func (element *Element) ChildOuterHtml(selector string) string {
	content, _ := element.Selection.Find(selector).RemoveClass().Html()
	return strings.Join(strings.Fields(content), " ")
}

func (element *Element) ChildrenInnerHtmls(selector string) (values []string) {
	element.Selection.Find(selector).Each(func(_ int, selection *goquery.Selection) {
		content, _ := selection.RemoveClass().Html()
		values = append(values, strings.Join(strings.Fields(content), " "))
	})
	return values
}

func (element *Element) ChildrenOuterHtmls(selector string) (values []string) {
	element.Selection.Find(selector).Each(func(_ int, selection *goquery.Selection) {
		content, _ := goquery.OuterHtml(selection.RemoveClass())
		values = append(values, strings.Join(strings.Fields(content), " "))
	})
	return values
}

func (element *Element) ChildText(selector string) string {
	return strings.TrimSpace(element.Selection.Find(selector).Text())
}

func (element *Element) ChildTexts(selector string) (values []string) {
	element.Selection.Find(selector).Each(func(_ int, selection *goquery.Selection) {
		values = append(values, strings.TrimSpace(selection.Text()))
	})
	return values
}

func (element *Element) Attribute(k string) string {
	value, _ := element.Selection.Attr(k)
	return value
}

func (element *Element) ChildAttribute(selector, name string) string {
	attr, _ := element.Selection.Find(selector).Attr(name)
	return strings.TrimSpace(attr)
}

func (element *Element) ChildAttributes(selector, attrName string) (result []string) {
	element.Selection.Find(selector).Each(func(_ int, s *goquery.Selection) {
		if attr, ok := s.Attr(attrName); ok {
			result = append(result, strings.TrimSpace(attr))
		}
	})
	return
}

func (element *Element) ForEach(selector string, callback func(int, *Element)) {
	element.Selection.Find(selector).Each(func(index int, selection *goquery.Selection) {
		for _, node := range selection.Nodes {
			callback(index, NewElement(goquery.NewDocumentFromNode(node).Selection))
		}
	})
}

func (element *Element) ForEachWithBreak(selector string, callback func(int, *Element) bool) {
	element.Selection.Find(selector).EachWithBreak(func(index int, selection *goquery.Selection) bool {
		for _, node := range selection.Nodes {
			if callback(index, NewElement(goquery.NewDocumentFromNode(node).Selection)) {
				return true
			}
		}
		return false
	})
}
