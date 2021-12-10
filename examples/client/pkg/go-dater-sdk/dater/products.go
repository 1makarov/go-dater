package dater

import (
	"context"
	"github.com/1makarov/go-dater/client/pkg/go-dater-sdk/proto"
)

func (c *Client) Fetch(ctx context.Context, url string) error {
	input := &proto.FetchRequest{Url: url}

	if _, err := c.daterClient.Fetch(ctx, input); err != nil {
		return err
	}

	return nil
}

func (c *Client) List(ctx context.Context, i ParamInput) (products []Product, err error) {
	input := &proto.ListRequest{
		Sorting: &proto.ListSortingStruct{
			Sort:   proto.Sort(i.Sort),
			Entity: proto.Entity(i.Entity),
		},
		Paging: &proto.ListPagingStruct{
			Offset: i.Offset,
			Limit:  i.Limit,
		},
	}

	resp, err := c.daterClient.List(ctx, input)
	if err != nil {
		return nil, err
	}

	for _, product := range resp.Products {
		products = append(products, Product{
			Name:  product.Name,
			Price: product.Price,
		})
	}

	return
}
