package jsonjse

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const (
	databaseName         = "jse"
	pricesCollectionName = "prices"
	newsCollectionName   = "news"
)

type Database struct {
	MongoDB *mongo.Client
}

func (d *Database) AddDailyPrices(date string, symbols []Symbol) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pricesCollection := d.MongoDB.Database(databaseName).Collection(pricesCollectionName)
	dailyPrices := DailyPrices{
		Date:    date,
		Symbols: symbols,
	}
	_, err := pricesCollection.InsertOne(ctx, dailyPrices)
	return err
}

func (d *Database) AddNewsArticles(date string, newsArticles []NewsArticle) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	newsCollection := d.MongoDB.Database(databaseName).Collection(newsCollectionName)
	dailyNews := DailyNews{
		Date:         date,
		NewsArticles: newsArticles,
	}
	_, err := newsCollection.InsertOne(ctx, dailyNews)
	return err
}

func (d *Database) GetPricesForDate(date string) ([]Symbol, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var datePrices DailyPrices
	pricesCollection := d.MongoDB.Database(databaseName).Collection(pricesCollectionName)
	if err := pricesCollection.FindOne(ctx, bson.M{"date": date}).Decode(&datePrices); err != nil {
		return []Symbol{}, err
	}
	return datePrices.Symbols, nil
}

func (d *Database) GetArticlesForDate(date string) ([]NewsArticle, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var dateNews DailyNews
	newsCollection := d.MongoDB.Database(databaseName).Collection(newsCollectionName)
	if err := newsCollection.FindOne(ctx, bson.M{"date": date}).Decode(&dateNews); err != nil {
		return []NewsArticle{}, err
	}
	return dateNews.NewsArticles, nil
}
