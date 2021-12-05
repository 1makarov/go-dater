package main

import (
	"context"
	"fmt"
	"github.com/1makarov/go-dater/client/pkg/go-dater-sdk/dater"
	"github.com/sirupsen/logrus"
)

const (
	addr = ":80"

	productsURL = "http://164.92.251.245:8080/api/v1/products/"
)

func main() {
	client, err := dater.New(addr)
	if err != nil {
		logrus.Errorln(err)
		return
	}

	defer func() {
		if err = client.Close(); err != nil {
			logrus.Errorln(err)
			return
		}
	}()

	ctx := context.Background()

	if err = client.Fetch(ctx, productsURL); err != nil {
		logrus.Errorln(err)
		return
	}

	products, err := client.List(ctx, dater.ParamInput{
		Sort:   dater.DESC, // dater.ASC
		Entity: dater.Name, // dater.Price
		Offset: 0,
		Limit:  50,
	})
	if err != nil {
		logrus.Errorln(err)
		return
	}

	for _, product := range products {
		fmt.Println(product)
	}
}
