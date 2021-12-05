package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Input struct {
	Host     string
	User     string
	Password string
	Port     string
}

func (i *Input) url() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/", i.User, i.Password, i.Host, i.Port)
}

func Open(ctx context.Context, input Input) (*mongo.Client, error) {
	urlConn := input.url()
	opts := options.Client().ApplyURI(urlConn)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	return client, client.Ping(ctx, nil)
}
