package repository

import "go.mongodb.org/mongo-driver/mongo"

const collectionProducts = "products"

type Repository struct {
	Products *ProductsRepository
}

func New(db *mongo.Database) *Repository {
	return &Repository{
		Products: newProductsRepo(db.Collection(collectionProducts)),
	}
}
