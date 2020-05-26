package jsonjse

type Symbol struct {
	Symbol           string  `csv:"Symbol" json:"symbol"`
	Date             string  `csv:"Date" json:"date"`
	FiftyTwoWeekHigh float64 `csv:"52 Week High" json:"year_high"`
	FiftyTwoWeekLow  float64 `csv:"52 Week Low" json:"year_low"`
	Last             float64 `csv:"Last" json:"last"`
	Volume           float64 `csv:"Volume (non block)" json:"volume"`
	TodayHigh        float64 `csv:"Today High" json:"today_high"`
	TodayLow         float64 `csv:"Today Low" json:"today_low"`
	LastTraded       float64 `csv:"Last Traded" json:"last_traded"`
	ClosePrice       float64 `csv:"Close Price" json:"close_price"`
	PreviousYearDiv  float64 `csv:"Previous Year Div" json:"previous_year_div"`
	CurrentYearDiv   float64 `csv:"Current Year Div" json:"current_year_div"`
	PriceChange      float64 `csv:"Price Change" json:"price_change"`
	ClosingBid       float64 `csv:"Closing Bid" json:"closing_bid"`
	ClosingAsk       float64 `csv:"Closing Ask" json:"closing_ask"`
}

type NewsArticle struct {
	Title      string `json:"headline"`
	URL        string `json:"url"`
	Summary    string `json:"summary"`
	Source     string `json:"source"`
	Lang       string `json:"lang"`
	HasPaywall bool   `json:"paywall"`
	Datetime   int64  `json:"datetime"`
}
