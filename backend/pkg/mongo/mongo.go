package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

type MongoInput struct {
	DSN      string
	Database string
}

func NewMongoDatabase(ctx context.Context, input MongoInput) *mongo.Database {
	if db == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		options := options.Client()

		options.ApplyURI(input.DSN)

		client, err := mongo.Connect(ctx, options)

		if err != nil {
			panic(err)
		}
		db = client.Database(input.Database)
	}

	return db
}

func Disconnect(ctx context.Context) {
	_ = db.Client().Disconnect(ctx)
}

func IsAvailable(ctx context.Context) bool {
	err := db.Client().Ping(ctx, nil)

	if err != nil {
		return false
	}

	return true
}
