package services

import (
	"context"
	"github.com/1makarov/go-dater/server/internal/getter"
	"github.com/1makarov/go-dater/server/internal/repository"
	"github.com/1makarov/go-dater/server/internal/types"
	"net/url"
)

type ProductsService struct {
	repo   *repository.ProductsRepository
	getter *get.Client
}

func newProductService(repo *repository.ProductsRepository, getter *get.Client) *ProductsService {
	return &ProductsService{
		repo:   repo,
		getter: getter,
	}
}

func (s *ProductsService) Fetch(ctx context.Context, inputURL string) error {
	if _, err := url.ParseRequestURI(inputURL); err != nil {
		return err
	}

	products, err := s.getter.GetProductsFromURL(inputURL, ";")
	if err != nil {
		return err
	}

	dbProducts, err := s.repo.GetAll(ctx)
	if err != nil {
		return err
	}

	mapProducts := dbProducts.Map()
	for _, product := range products {
		productDB, ok := mapProducts[product.Name]

		if !ok {
			if err = s.repo.Create(ctx, &product); err != nil {
				return err
			}
			continue
		}

		if product.Price != productDB.Price {
			if err = s.repo.Update(ctx, &product); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *ProductsService) List(ctx context.Context, input types.GetByParametersInput) ([]types.Product, error) {
	return s.repo.GetByParameters(ctx, input)
}
