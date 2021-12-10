package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const collectionProducts = "products"

type Repository struct {
	Products *ProductsRepository
}

func New(db *mongo.Database) *Repository {
	return &Repository{
		Products: newProductsRepo(db.Collection(collectionProducts)),
	}
}

func getPaginationOpts(offset, limit int64) *options.FindOptions {
	return &options.FindOptions{
		Skip:  &(offset),
		Limit: &(limit),
	}
}
