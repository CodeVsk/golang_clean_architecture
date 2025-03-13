package contract

import "github.com/codevsk/golang_clean_architecture/internal/entity"

type ProductRepository interface {
	Get() (e []entity.Product)
	Create(model entity.Product)
}
