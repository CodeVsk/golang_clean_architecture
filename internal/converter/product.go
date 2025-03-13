package converter

import (
	"github.com/codevsk/golang_clean_architecture/internal/dto"
	"github.com/codevsk/golang_clean_architecture/internal/entity"
)

func ToEntity(input *dto.CreateProductInputDto) entity.Product {
	return entity.Product{
		Name:  input.Name,
		Price: input.Price,
		Slug:  input.Slug,
	}
}

func ToDto(input entity.Product) dto.ProductOutputDto {
	return dto.ProductOutputDto{
		Id:    input.Id,
		Name:  input.Name,
		Price: input.Price,
	}
}

func ToListDto(input []entity.Product) (output []dto.ProductOutputDto) {
	output = make([]dto.ProductOutputDto, len(input))

	for i := 0; i < len(input); i++ {
		output[i] = dto.ProductOutputDto{
			Id:    input[i].Id,
			Name:  input[i].Name,
			Price: input[i].Price,
		}
	}

	return
}
