package repository

import (
	"context"
	"github.com/1makarov/go-dater/internal/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
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

	return products, result.All(ctx, &products)
}

type ByParameters struct {
	Offset int64
	Limit  int64
	Entity string
	Sort   string
}

func (r *ProductsRepository) GetByParameters(ctx context.Context, p ByParameters) (products []types.Product, err error) {
	p.Entity = strings.ToLower(p.Entity)

	opts := getPaginationOpts(p.Offset, p.Limit)
	query := bson.M{}

	switch p.Sort {
	case "asc":
		query[p.Entity] = 1
	case "desc":
		query[p.Entity] = -1
	}

	opts.SetSort(query)

	result, err := r.db.Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}

	return products, result.All(ctx, &products)
}

func (r *ProductsRepository) Update(ctx context.Context, v *types.Product) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"name": v.Name}, bson.M{"$set": bson.M{"price": v.Price}})

	return err
}
