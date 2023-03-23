package main

import (
	"log"

	"github.com/teamcubation/neocamp-meli/clean-architecture/internal/infrastructure/repository"
	"github.com/teamcubation/neocamp-meli/clean-architecture/internal/infrastructure/web"
	ctrl "github.com/teamcubation/neocamp-meli/clean-architecture/internal/infrastructure/web/controller"
	"github.com/teamcubation/neocamp-meli/clean-architecture/internal/usecase"
)

func main() {
	bookRepository := repository.NewBookRepository()
	bookUsecase := usecase.NewBookUsecase(bookRepository)
	bookCtrl := ctrl.NewBookController(bookUsecase)

	if err := web.NewHTTPServer(bookCtrl); err != nil {
		log.Fatalln(err)
	}
}
