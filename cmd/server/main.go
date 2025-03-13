package main

import (
	"github.com/codevsk/golang_clean_architecture/internal/infra/repository"
	"github.com/codevsk/golang_clean_architecture/internal/infra/webserver"
	"github.com/codevsk/golang_clean_architecture/internal/infra/webserver/handler"
)

func main() {
	ws := webserver.NewWebServer(":3030")

	/* Products */

	pr := repository.NewProductMemoryRepository()
	ph := handler.NewProductHandler(pr)

	ws.RegisterHandler("/products", "GET", ph.Get)
	ws.RegisterHandler("/products", "POST", ph.Create)

	/* -------- */

	ws.Run()
}
