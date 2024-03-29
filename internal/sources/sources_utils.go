package sources

import (
	"errors"
	"fmt"
	"io"
	"net/url"
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
	u := launcher.New().Leakless(false).NoSandbox(true).Bin(path).MustLaunch()
	browser = rod.New().ControlURL(u).MustConnect()
}

func RodNavigate(url string, load ...bool) (io.Reader, error) {
	page := browser.MustPage(url)
	defer page.Close()

	page.Timeout(1 * time.Millisecond).WaitLoad()
	page.WaitRepaint()

	return strings.NewReader(page.MustHTML()), nil
}

func RodGetRequest(url string) (io.Reader, error) {
	page := browser.MustPage(url)
	defer page.Close()

	page.Timeout(1 * time.Millisecond).WaitLoad()
	value := page.MustEval(`
	(url) => {
		let xhr = new XMLHttpRequest();
		xhr.open('GET', url, false);
		xhr.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');
		try {
			xhr.send();
			} catch (e) {
			return e;
		}
		return xhr.response;
	}
	`, url).Str()
	return strings.NewReader(value), nil
}

func RodPostRequest(url string, data string) (io.Reader, error) {
	page := browser.MustPage(url)
	defer page.Close()

	page.Timeout(1 * time.Millisecond).WaitLoad()
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
	return strings.NewReader(value), nil
}

func ParseURL(baseURL, rawURL string) string {
	bu, _ := url.Parse(baseURL)
	u, err := bu.Parse(rawURL)
	if err != nil {
		fmt.Println(err)
		return rawURL
	}
	return u.String()
}

func ParseCategorySource(categories map[string]string, category string) (string, error) {
	for key, value := range categories {
		if key == category {
			return value, nil
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
		`\d{1,2}-\d{1,2}-\d{4}`:       "2-1-2006",
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
func ParseTime(value string) (time.Time, error) {
	value, _, err := transform.String(transform.Chain(norm.NFD, runes.Remove(runes.Predicate(unicode.IsMark))), value)
	if err != nil {
		return time.Time{}, err
	}
	for expression, layout := range layouts {
		if date, er := time.Parse(layout, dateInEnglish(regexp.MustCompile(expression).FindString(strings.Join(strings.Fields(value), " ")))); er != nil {
			err = er
		} else {
			if date.Year() == 0 {
				now := time.Now()
				date = date.AddDate(now.Year(), int(now.Month()), now.Day())
			}
			return date, nil
		}
	}
	return time.Time{}, err
}


var videoRegexpList = []*regexp.Regexp{
	regexp.MustCompile(`(?:v|embed|shorts|watch\?v)(?:=|/)([^"&?/=%]{11})`),
	regexp.MustCompile(`(?:=|/)([^"&?/=%]{11})`),
	regexp.MustCompile(`([^"&?/=%]{11})`),
}

// ExtractVideoID extracts the videoID from the given string
func ExtractVideoID(videoID string) (string) {
	if strings.Contains(videoID, "youtu") || strings.ContainsAny(videoID, "\"?&/<%=") {
		for _, re := range videoRegexpList {
			if isMatch := re.MatchString(videoID); isMatch {
				subs := re.FindStringSubmatch(videoID)
				videoID = subs[1]
			}
		}
	}

	if strings.ContainsAny(videoID, "?&/<%=") {
		return "";
	}
	if len(videoID) < 10 {
		return ""
	}

	return videoID
}