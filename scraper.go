package jsonjse

import (
	"net/http"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocarina/gocsv"
)

const (
	jseMarketDataCSVURL = "https://www.jamstockex.com/market-data/download-data/price-history/generate-csv/all-stocks"
	jseMarketDataURL    = "https://www.jamstockex.com/market-data/download-data/price-history/"
)

type Symbol struct {
	Symbol           string  `csv:"Symbol"`
	Date             string  `csv:"Date"`
	FiftyTwoWeekHigh float64 `csv:"52 Week High"`
	FiftyTwoWeekLow  float64 `csv:"52 Week Low"`
	Last             float64 `csv:"Last"`
	Volume           float64 `csv:"Volume (non block)"`
	TodayHigh        float64 `csv:"Today High"`
	TodayLow         float64 `csv:"Today Low"`
	LastTraded       float64 `csv:"Last Traded"`
	ClosePrice       float64 `csv:"Close Price"`
	PreviousYearDiv  float64 `csv:"Previous Year Div"`
	CurrentYearDiv   float64 `csv:"Current Year Div"`
	PriceChange      float64 `csv:"Price Change"`
	ClosingBid       float64 `csv:"Closing Bid"`
	ClosingAsk       float64 `csv:"Closing Ask"`
}

type JSE struct {
	*http.Client
}

func (j *JSE) GetPricesForDate(date time.Time) ([]Symbol, error) {
	var symbols []Symbol
	u, err := j.todayMarketReportURL()
	if err != nil {
		return symbols, err
	}

	resp, err := j.Get(u.String())
	if err != nil {
		return symbols, err
	}
	gocsv.Unmarshal(resp.Body, &symbols)
	return symbols, nil
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
