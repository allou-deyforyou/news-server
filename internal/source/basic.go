package source

import (
	"errors"
	"fmt"
	"io"
	"net/url"
	"news/internal/store"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
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

func rodPostRequest(url string, data string) (io.Reader, error) {
	page := browser.MustPage(url)
	value := page.MustEval(`
	(url, data) => {
		let xhr = new XMLHttpRequest();
		xhr.open('POST', url, false);
		xhr.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');
		try {
			xhr.send(data);
			} catch (e) {
			return e;
		}
		return xhr.response;
	}
	`, url, data).Str()
	defer page.Close()
	return strings.NewReader(value), nil
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

var (
	months = map[string]string{
		"janvier":   "January",
		"fevrier":   "February",
		"mars":      "March",
		"avril":     "April",
		"mai":       "May",
		"juin":      "June",
		"juillet":   "July",
		"aout":      "August",
		"septembre": "September",
		"octobre":   "October",
		"novembre":  "November",
		"decembre":  "December",
	}

	// Regexp for date time
	layouts = map[string]string{
		`\d{1,2} sema`:                "1 sema",
		`\d{1,2}/\d{1,2}/\d{2}`:       "2/1/06", // get date in format 02/01/06
		`\d{1,2} ans`:                 "1 ans",
		`hier a \d{1,2}`:              "hier a 02",
		`\d{1,2} jour`:                "1 jour",
		`\d{1,2} mois`:                "1 mois",
		`\d{4}-\d{1,2}-\d{1,2}`:       "2006-1-2",            // get date in format 2006/01/02
		`\d{1,2}/\d{1,2}/\d{4}`:       "2/1/2006",            // get date in format 02/01/2006
		`\d{1,2} \w* \d{4}`:           "2 January 2006",      // get date in format 02 January 2006
		`il y a \d{1,2} heure`:        "il y a 15 heure",     // get date in format il y a 15 heure
		`il y a \d{1,2} minute`:       "il y a 4 minute",     // get date in format il y a 15 minute
		`il y a \d{1,2} seconde`:      "il y a 5 seconde",    // get date in format il y a 15 seconde
		`.* \d{1,3}, \d{4} at`:        "January 2, 2006 at",  // get date in format January 2, 2006 at
		`aujourd’hui a \d{1,2}h\d{2}`: "aujourd’hui a 15h04", // get date in format aujourd’hui à 15h04
	}
)

//dateInEnglish Format date from french ton english
func dateInEnglish(date string) string {
	for k, v := range months {
		date = strings.Replace(date, k, v, -1)
	}
	return date
}

//parseTime parse time from string to golang time
func parseTime(value string) (string, error) {
	value, _, err := transform.String(transform.Chain(norm.NFD, runes.Remove(runes.Predicate(unicode.IsMark))), value)
	if err != nil {
		return "", err
	}
	for expression, layout := range layouts {
		if date, er := time.Parse(layout, dateInEnglish(regexp.MustCompile(expression).FindString(strings.Join(strings.Fields(value), " ")))); er != nil {
			err = er
		} else {
			return date.Format(time.RFC3339), nil
		}
	}
	return "", err
}
