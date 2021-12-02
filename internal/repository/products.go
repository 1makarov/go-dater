package repository

import (
	"context"
	"github.com/1makarov/go-dater/internal/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductsRepository struct {
	db *mongo.Collection
}

func newProductsRepo(db *mongo.Collection) *ProductsRepository {
	return &ProductsRepository{
		db: db,
	}
}

func (r *ProductsRepository) Create(ctx context.Context, v *types.Product) error {
	_, err := r.db.InsertOne(ctx, v)

	return err
}

func (r *ProductsRepository) GetAll(ctx context.Context) (products []types.Product, err error) {
	result, err := r.db.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	err = result.All(ctx, &products)

	return
}

func (r *ProductsRepository) Update(ctx context.Context, v *types.Product) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"name": v.Name}, bson.M{"$set": bson.M{"price": v.Price}})

	return err
}
