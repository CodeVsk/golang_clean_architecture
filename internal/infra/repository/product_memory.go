package repository

import "github.com/codevsk/golang_clean_architecture/internal/entity"

type ProductMemoryRepository struct {
	products  map[int]entity.Product
	lastIndex int
}

func NewProductMemoryRepository() *ProductMemoryRepository {
	return &ProductMemoryRepository{
		products:  make(map[int]entity.Product),
		lastIndex: 0,
	}
}

func (pr *ProductMemoryRepository) Get() (e []entity.Product) {
	productsLength := len(pr.products)

	e = make([]entity.Product, productsLength)

	for i := 0; i < productsLength; i++ {
		e[i] = pr.products[i]
	}

	return
}

func (pr *ProductMemoryRepository) Create(model entity.Product) {
	pr.products[pr.lastIndex] = model
	pr.lastIndex++
}
