package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

func (c *Config) url() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/", c.User, c.Password, c.Host, c.Port)
}

func Open(ctx context.Context, cfg Config) (*mongo.Client, error) {
	url := cfg.url()
	opts := options.Client().ApplyURI(url)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	return client, client.Ping(ctx, nil)
}
