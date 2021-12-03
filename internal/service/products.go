package service

import (
	"context"
	"github.com/1makarov/go-dater/internal/proto"
	"github.com/1makarov/go-dater/internal/repository"
	"github.com/1makarov/go-dater/internal/types"
	"github.com/1makarov/go-dater/pkg/transport/http"
	"net/url"
	"strconv"
	"strings"
)

type ProductService struct {
	repo     *repository.ProductsRepository
	provider *http.Provider
}

func newProductService(repo *repository.ProductsRepository, provider *http.Provider) *ProductService {
	return &ProductService{
		repo:     repo,
		provider: provider,
	}
}

func (s *ProductService) Fetch(ctx context.Context, v *proto.FetchRequest) (*proto.Empty, error) {
	if _, err := url.ParseRequestURI(v.Url); err != nil {
		return nil, err
	}

	csvProducts := new(string)
	if err := s.provider.Get(v.Url, csvProducts); err != nil {
		return nil, err
	}

	dbProducts, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	mapProducts := make(map[string]int64)
	for _, product := range dbProducts {
		mapProducts[product.Name] = product.Price
	}

	var products []types.Product
	for _, product := range strings.Split(*csvProducts, "\n") {
		data := strings.Split(product, ";")

		if len(data) != 2 {
			continue
		}

		price, err := strconv.Atoi(data[1])
		if err != nil {
			continue
		}

		prod := types.Product{
			Name:  data[0],
			Price: int64(price),
		}

		products = append(products, prod)
	}

	for _, product := range products {
		price, ok := mapProducts[product.Name]

		if !ok {
			if err = s.repo.Create(ctx, &product); err != nil {
				return nil, err
			}
			continue
		}

		if product.Price != price {
			if err = s.repo.Update(ctx, &product); err != nil {
				return nil, err
			}
			continue
		}
	}

	return &proto.Empty{}, nil
}

func (s *ProductService) List(ctx context.Context, v *proto.ListRequest) (*proto.ListResponse, error) {
	products, err := s.repo.GetByParameters(ctx, repository.ByParameters{
		Offset: v.Paging.Offset,
		Limit:  v.Paging.Limit,
		Entity: v.Sorting.Entity.String(),
		Sort:   v.Sorting.Sort.String(),
	})
	if err != nil {
		return nil, err
	}

	var protoProducts []*proto.ListProductObject
	for _, product := range products {
		protoProducts = append(protoProducts, &proto.ListProductObject{
			Name:  product.Name,
			Price: product.Price,
		})
	}

	return &proto.ListResponse{Products: protoProducts}, nil
}
