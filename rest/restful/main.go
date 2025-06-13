package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teamcubation/neocamp-meli/rest/restful/handler"
	"github.com/teamcubation/neocamp-meli/rest/restful/repository"
)

func main() {
	r := gin.Default()

	bookHandler := &handler.BookHandler{
		Repo: repository.BookRepository{},
	}

	// KVS (DynamoDB)
	// DS (elastic search)
	// MySQL

	//GET /books/2
	r.GET("", index)
	r.GET("/books", bookHandler.GetBooks)
	r.POST("/books", bookHandler.AddBook)
	r.GET("/books/:id", bookHandler.GetBook)
	r.GET("/books/find/:name", bookHandler.GetBookByName)
	r.PUT("/books/:id", bookHandler.UpdateBook)
	r.DELETE("/books/:id", bookHandler.DeleteBook)

	r.Run(":8080")
}

func index(c *gin.Context) {
	c.JSON(http.StatusOK, "Bienvenido a mi increible API!")
}
