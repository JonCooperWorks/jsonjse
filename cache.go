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
		err = c.Database.AddDailyPrices(date, symbols)
		if err != nil {
			return symbols, err
		}
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
		err = c.Database.AddNewsArticles(date, news)
		if err != nil {
			return news, err
		}
	}
	return news, nil
}
