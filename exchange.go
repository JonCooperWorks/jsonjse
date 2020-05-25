package jsonjse

import (
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocarina/gocsv"
)

const (
	jseMarketDataURL = "https://www.jamstockex.com/market-data/download-data/price-history/"
	jseNewsURL       = "https://www.jamstockex.com/"
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
	articles := doc.Find("#primary h5 a")
	articles.Each(func(i int, article *goquery.Selection) {
		rawURL, ok := article.Attr("href")
		if !ok {
			return
		}

		articleURL, err := url.Parse(rawURL)
		if err != nil {
			return
		}
		newsArticle := NewsArticle{
			Title: article.Text(),
			URL:   articleURL.String(),
		}
		newsArticles = append(newsArticles, newsArticle)
	})

	return newsArticles, nil
}
