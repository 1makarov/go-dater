package types

type Product struct {
	Name  string `json:"name" bson:"name"`
	Price int64  `bson:"price" bson:"price"`
}
