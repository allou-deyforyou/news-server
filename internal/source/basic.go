package source

import (
	"errors"
	"fmt"
	"io"
	"net/url"
	"news/internal/store"
	"strings"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

var browser *rod.Browser

func init() {
	path, _ := launcher.LookPath()
	u := launcher.New().Bin(path).NoSandbox(true).MustLaunch()
	browser = rod.New().ControlURL(u).MustConnect()
}

func rodGetRequest(url string, wait string) (io.Reader, error) {
	page := browser.MustPage(url)
	defer page.Close()
	return strings.NewReader(page.MustElement(wait).MustHTML()), nil
}

func parseURL(baseURL, rawURL string) string {
	bu, _ := url.Parse(baseURL)
	u, err := bu.Parse(rawURL)
	if err != nil {
		fmt.Println(err)
		return rawURL
	}
	return u.String()
}

func parseCategorySource(source *store.NewsSource, name string) (string, error) {
	for _, category := range source.Categories {
		if strings.HasPrefix(category, name) {
			return strings.TrimPrefix(category, name+":"), nil
		}
	}
	return "", errors.New("no found")
}
