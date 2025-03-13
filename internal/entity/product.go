package entity

import "github.com/google/uuid"

type Product struct {
	Id    string
	Name  string
	Slug  string
	Price float64
}

func NewProduct(name string, slug string, price float64) Product {
	id, _ := uuid.NewUUID()

	return Product{
		Id:    id.String(),
		Name:  name,
		Slug:  slug,
		Price: price,
	}
}
