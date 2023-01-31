package clients

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	Database *mongo.Database
}

func NewMongoClient(mongoUri string, mongoDb string) (*MongoClient, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &MongoClient{
		Database: client.Database(mongoDb),
	}, nil
}

func (c *MongoClient) Disconnect(ctx context.Context) error {
	return c.Database.Client().Disconnect(ctx)
}
