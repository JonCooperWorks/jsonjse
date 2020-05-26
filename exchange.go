package jsonjse

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocarina/gocsv"
	"net/http"
	"net/url"
)

const (
	jseMarketDataURL = "https://www.jamstockex.com/market-data/download-data/price-history/"
	jseNewsURL       = "https://www.jamstockex.com/news/"
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
		newsArticles[i].Datetime = dateSelection.Text()
	})

	articleDescriptions.Each(func(i int, description *goquery.Selection) {
		newsArticles[i].Summary = description.Text()
	})

	return newsArticles, nil
}
