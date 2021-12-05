package get

import (
	"github.com/1makarov/go-dater/server/internal/types"
	"github.com/1makarov/go-dater/server/pkg/transport/http"
	"strconv"
	"strings"
)

type Client struct {
	http *http.Provider
}

func New() *Client {
	return &Client{
		http: http.New(),
	}
}

func (c *Client) GetProductsFromURL(url, sep string) (products []types.Product, err error) {
	body := new(string)
	if err = c.http.Get(url, body); err != nil {
		return nil, err
	}

	for _, p := range strings.Split(*body, "\n") {
		data := strings.Split(p, sep)

		if len(data) != 2 {
			continue
		}

		price, err := strconv.Atoi(data[1])
		if err != nil {
			continue
		}

		product := types.Product{
			Name:  data[0],
			Price: int64(price),
		}

		products = append(products, product)
	}

	return
}
