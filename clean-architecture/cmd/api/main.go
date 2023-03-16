package main

import (
	"github.com/teamcubation/neocamp-meli/clean-architecture/internal/domain/model"
	repo "github.com/teamcubation/neocamp-meli/clean-architecture/internal/infrastructure/database/memory"
	router "github.com/teamcubation/neocamp-meli/clean-architecture/internal/infrastructure/httpserver"
	"github.com/teamcubation/neocamp-meli/clean-architecture/internal/usecase"
)

const port = "9000"

var db []model.Book = []model.Book{
	{
		ID:     1,
		Title:  "Dune",
		Price:  1965,
		Author: "Frank Herbert",
	},
	{
		ID:     2,
		Title:  "Cita con Rama",
		Price:  1974,
		Author: "Arthur C. Clarke",
	},
	{
		ID:     3,
		Title:  "Un guijarro en el cielo",
		Price:  500,
		Author: "Isaac Asimov",
	},
}

func main() {
	repository := repo.NewBookRepository(db)
	usecase := usecase.NewBookUsecase(repository)
	handler := router.NewBookHandler(usecase)
	server := router.NewHTTPServer(handler)

	server.RegisterRouter()

	if err := server.Run(port); err != nil {
		panic(err.Error())
	}
}
