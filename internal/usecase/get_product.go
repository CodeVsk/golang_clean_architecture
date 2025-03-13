package usecase

import (
	"github.com/codevsk/golang_clean_architecture/internal/contract"
	"github.com/codevsk/golang_clean_architecture/internal/converter"
	"github.com/codevsk/golang_clean_architecture/internal/dto"
)

type GetProductUsecase struct {
	ProductRepository contract.ProductRepository
}

func NewGetProductUsecase(productRepository contract.ProductRepository) *GetProductUsecase {
	return &GetProductUsecase{
		ProductRepository: productRepository,
	}
}

func (c *GetProductUsecase) Execute() (output []dto.ProductOutputDto, err error) {
	products := c.ProductRepository.Get()

	output = converter.ToListDto(products)

	return
}
