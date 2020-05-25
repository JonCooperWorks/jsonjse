package jsonjse

type JSECache struct {
	JSE      *JSE
	Database *Database
}

func (c *JSECache) DailyPrices(date string) ([]Symbol, error) {
	symbols, err := c.Database.GetPricesForDate(date)
	if err != nil {
		symbols, err = c.JSE.GetTodaysPrices()
		if err != nil {
			return symbols, err
		}

		// If we got data from the JSE, continue even if updating the cache fails: we'll try again next time.
		_ = c.Database.AddDailyPrices(date, symbols)
	}
	return symbols, nil
}

func (c *JSECache) DailyNews(date string) ([]NewsArticle, error) {
	news, err := c.Database.GetArticlesForDate(date)
	if err != nil {
		news, err = c.JSE.GetTodaysNews()
		if err != nil {
			return news, err
		}

		// If we got news from the JSE, continue even if updating the cache fails: we'll try again next time.
		_ = c.Database.AddNewsArticles(date, news)
	}
	return news, nil
}
