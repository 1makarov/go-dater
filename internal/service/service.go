package service

import (
	"github.com/1makarov/go-dater/internal/repository"
	"github.com/1makarov/go-dater/pkg/transport/http"
)

type Service struct {
	*ProductService
}

func New(repo *repository.Repository, provider *http.Provider) *Service {
	return &Service{
		newProductService(repo.Products, provider),
	}
}
