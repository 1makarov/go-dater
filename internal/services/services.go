package services

import (
	"context"
	"github.com/1makarov/go-dater/server/internal/getter"
	"github.com/1makarov/go-dater/server/internal/repository"
	"github.com/1makarov/go-dater/server/internal/types"
)

type Services struct {
	Products Products
}

type Products interface {
	Fetch(ctx context.Context, inputURL string) error
	List(ctx context.Context, input types.GetByParametersInput) ([]types.Product, error)
}

func New(repo *repository.Repository, get *getter.Client) *Services {
	return &Services{
		Products: newProductService(repo.Products, get),
	}
}
