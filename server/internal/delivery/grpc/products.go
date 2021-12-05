package grpcHandler

import (
	"context"
	"github.com/1makarov/go-dater/server/internal/proto"
	"github.com/1makarov/go-dater/server/internal/types"
)

func (h *Handler) Fetch(ctx context.Context, v *proto.FetchRequest) (*proto.Empty, error) {
	return &proto.Empty{}, h.services.Products.Fetch(ctx, v.Url)
}

func (h *Handler) List(ctx context.Context, v *proto.ListRequest) (*proto.ListResponse, error) {
	input := types.GetByParametersInput{
		Offset: v.Paging.Offset,
		Limit:  v.Paging.Limit,
		Entity: v.Sorting.Entity.String(),
		Sort:   v.Sorting.Sort.String(),
	}

	products, err := h.services.Products.List(ctx, input)
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
