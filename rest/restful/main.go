package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teamcubation/neocamp-meli/rest/restful/controller"
)

func main() {
	r := gin.Default()

	r.GET("", index)
	r.GET("/books", controller.GetBooks)
	r.POST("/books", controller.AddBook)
	r.GET("/books/:id", controller.GetBook)
	r.GET("/books/find/:name", controller.GetBookByName)
	r.PUT("/books/:id", controller.UpdateBook)
	r.DELETE("/books/:id", controller.DeleteBook)

	r.Run(":8080")
}

func index(c *gin.Context) {
	c.JSON(http.StatusOK, "Bienvenido a mi increible API!")
}
