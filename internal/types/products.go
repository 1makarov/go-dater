package types

type Product struct {
	Name  string `bson:"Name"`
	Price int64  `bson:"Price"`
}

type GetByParametersInput struct {
	Offset int64
	Limit  int64
	Entity string
	Sort   string
}

type Products struct {
	Slice []Product
}

func (p *Products) Map() map[string]Product {
	products := make(map[string]Product, len(p.Slice))

	for _, product := range p.Slice {
		products[product.Name] = product
	}

	return products
}
