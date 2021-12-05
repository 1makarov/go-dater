package dater

import "github.com/1makarov/go-dater/client/pkg/go-dater-sdk/proto"

const (
	DESC = Sort(proto.Sort_desc)
	ASC  = Sort(proto.Sort_asc)

	Name  = Entity(proto.Entity_Name)
	Price = Entity(proto.Entity_Price)
)

type Sort proto.Sort
type Entity proto.Entity

type ParamInput struct {
	Sort   Sort
	Entity Entity
	Offset int64
	Limit  int64
}

type Product struct {
	Name  string
	Price int64
}
