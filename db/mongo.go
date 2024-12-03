package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	client *mongo.Client
	source string
}

func NewMongo(source string) *Mongo {
	return &Mongo{
		source: source,
	}
}

func (m *Mongo) Connect(ctx context.Context) error {
	c, err := mongo.Connect(ctx)
	if err != nil {
		return err
	}

	m.client = c

	return nil
}

func (m *Mongo) Disconnect(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}
