package dto

type CreateProductInputDto struct {
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required"`
	Slug  string  `json:"slug" validate:"required"`
}

type ProductOutputDto struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
