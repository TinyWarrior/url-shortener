package mongo

import (
	"context"
	"time"
)

type MongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func NewMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURL(mongoURL))
}
