package request

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
)

func JSON(r *http.Request, ptr any) *Error {
	if r.Header.Get("Content-Type") != "application/json" {
		return NewError(400, "request content type is not application/json")
	}

	if err := json.NewDecoder(r.Body).Decode(ptr); err != nil {
		return NewError(400, "request json invalid")
	}

	validate := validator.New()
	if err := validate.Struct(ptr); err != nil {
		return NewError(422, err.Error())
	}

	return nil
}
