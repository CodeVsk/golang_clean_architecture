package usecase

import (
	"github.com/codevsk/golang_clean_architecture/internal/contract"
	"github.com/codevsk/golang_clean_architecture/internal/converter"
	"github.com/codevsk/golang_clean_architecture/internal/dto"
	"github.com/codevsk/golang_clean_architecture/internal/entity"
)

type CreateProductUsecase struct {
	ProductRepository contract.ProductRepository
}

func NewCreateProductUsecase(productRepository contract.ProductRepository) *CreateProductUsecase {
	return &CreateProductUsecase{
		ProductRepository: productRepository,
	}
}

func (c *CreateProductUsecase) Execute(input dto.CreateProductInputDto) (output any, err error) {
	productEntity := entity.NewProduct(input.Name, input.Slug, input.Price)

	c.ProductRepository.Create(productEntity)

	output = converter.ToDto(productEntity)

	return
}
