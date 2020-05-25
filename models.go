package jsonjse

import "go.mongodb.org/mongo-driver/bson/primitive"

type DailyPrices struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Date    string             `bson:"date"`
	Symbols []Symbol           `bson:"symbols"`
}

type DailyNews struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Date         string             `bson:"date"`
	NewsArticles []NewsArticle      `bson:"articles"`
}
