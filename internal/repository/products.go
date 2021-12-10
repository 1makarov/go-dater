package repository

import (
	"context"
	"github.com/1makarov/go-dater/server/internal/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ascSort  = "asc"
	descSort = "desc"
)

type ProductsRepository struct {
	db *mongo.Collection
}

func newProductsRepo(db *mongo.Collection) *ProductsRepository {
	return &ProductsRepository{db: db}
}

func (r *ProductsRepository) Create(ctx context.Context, v *types.Product) error {
	_, err := r.db.InsertOne(ctx, v)

	return err
}

func (r *ProductsRepository) GetAll(ctx context.Context) (*types.Products, error) {
	result, err := r.db.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var products []types.Product
	if err = result.All(ctx, &products); err != nil {
		return nil, err
	}

	return &types.Products{Slice: products}, nil
}

func (r *ProductsRepository) GetByParameters(ctx context.Context, p types.GetByParametersInput) (products []types.Product, err error) {
	query := bson.M{}

	switch p.Sort {
	case ascSort:
		query[p.Entity] = 1
	case descSort:
		query[p.Entity] = -1
	}

	opts := getPaginationOpts(p.Offset, p.Limit).SetSort(query)

	result, err := r.db.Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}

	return products, result.All(ctx, &products)
}

func (r *ProductsRepository) Update(ctx context.Context, v *types.Product) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"Name": v.Name}, bson.M{"$set": bson.M{"Price": v.Price}})

	return err
}
