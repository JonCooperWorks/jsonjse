package jsonjse

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocarina/gocsv"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

const (
	jseMarketDataURL = "https://www.jamstockex.com/market-data/download-data/price-history/"
	jseNewsURL       = "https://www.jamstockex.com/news/"
)

var (
	dateLineRegexp = regexp.MustCompile(`Posted: (?P<Date>.*) at (?P<Time>.*)`)
)

type JSE struct {
	*http.Client
}

func (j *JSE) GetTodaysPrices() ([]Symbol, error) {
	var symbols []Symbol
	u, err := j.todayMarketReportURL()
	if err != nil {
		return symbols, err
	}

	resp, err := j.Get(u.String())
	if err != nil {
		return symbols, err
	}
	err = gocsv.Unmarshal(resp.Body, &symbols)
	return symbols, err
}

func (j *JSE) todayMarketReportURL() (*url.URL, error) {
	doc, err := goquery.NewDocument(jseMarketDataURL)
	if err != nil {
		return nil, err
	}

	rawURL, found := doc.Find("input[value='Download CSV']").Attr("onclick")
	if !found {
		panic("cannot find download button")
	}

	// We need to remove the text `onclick="window.location = '` from the start of the URL and the trailing '
	return url.Parse(rawURL[19 : len(rawURL)-1])
}

func (j *JSE) GetTodaysNews() ([]NewsArticle, error) {
	doc, err := goquery.NewDocument(jseNewsURL)
	if err != nil {
		return nil, err
	}
	newsArticles := []NewsArticle{}
	articleLinks := doc.Find("#main h1 a")
	articleDates := doc.Find("p.text-muted")
	articleDescriptions := doc.Find("#primary div p")

	articleLinks.Each(func(i int, articleLink *goquery.Selection) {
		rawURL, ok := articleLink.Attr("href")
		if !ok {
			return
		}

		articleURL, err := url.Parse(rawURL)
		if err != nil {
			panic(err)
			return
		}
		newsArticle := NewsArticle{
			Title:      articleLink.Text(),
			URL:        articleURL.String(),
			HasPaywall: false,
			Lang:       "en",
			Source:     "Jamaica Stock Exchange",
		}
		newsArticles = append(newsArticles, newsArticle)
	})

	articleDates.Each(func(i int, dateSelection *goquery.Selection) {
		dateLine := dateSelection.Text()
		match := regexpMap(dateLineRegexp, dateLine)
		dateString := fmt.Sprintf("%v %v", match["Date"], match["Time"])
		date, err := time.Parse("January 2, 2006 3:04 pm", dateString)
		if err != nil {
			log.Println(err)
			return
		}

		jamaica, _ := time.LoadLocation("America/Bogota")
		newsArticles[i].Datetime = date.In(jamaica).UnixNano()
	})

	articleDescriptions.Each(func(i int, description *goquery.Selection) {
		newsArticles[i].Summary = description.Text()
	})

	return newsArticles, nil
}

func regexpMap(r *regexp.Regexp, input string) map[string]string {
	rMap := map[string]string{}
	names := r.SubexpNames()
	res := r.FindStringSubmatch(input)
	for i, _ := range res {
		if i != 0 {
			rMap[names[i]] = res[i]
		}
	}
	return rMap
}