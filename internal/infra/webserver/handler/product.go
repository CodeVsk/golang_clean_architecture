package handler

import (
	"net/http"

	"github.com/codevsk/golang_clean_architecture/internal/contract"
	"github.com/codevsk/golang_clean_architecture/internal/dto"
	"github.com/codevsk/golang_clean_architecture/internal/usecase"
	"github.com/codevsk/golang_clean_architecture/pkg/web/request"
	"github.com/gawbsouza/boot-help/response"
)

type ProductHandler struct {
	ProductRepository contract.ProductRepository
}

func NewProductHandler(productRepository contract.ProductRepository) *ProductHandler {
	return &ProductHandler{
		ProductRepository: productRepository,
	}
}

func (ph *ProductHandler) Get(w http.ResponseWriter, r *http.Request) {
	getProduct := usecase.NewGetProductUsecase(ph.ProductRepository)
	output, err := getProduct.Execute()
	if err != nil {
		response.To(w).InternalErr(err.Error())
		return
	}

	response.To(w).Content(output).SendJSON()
}

func (ph *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateProductInputDto

	if err := request.JSON(r, &input); err != nil {
		response.To(w).Status(err.Code).Content(err.Message)
		return
	}

	createProduct := usecase.NewCreateProductUsecase(ph.ProductRepository)
	_, err := createProduct.Execute(input)
	if err != nil {
		response.To(w).InternalErr(err.Error())
		return
	}

	response.To(w).Status(http.StatusNoContent).Send()
}
