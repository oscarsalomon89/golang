package web

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	ctrl "github.com/teamcubation/neocamp-meli/clean-architecture/internal/infrastructure/web/controller"
)

const port = ":9000"

func NewHTTPServer(bookCtrl ctrl.BookController) error {
	r := gin.Default()

	r.GET("/", index)
	r.GET("/books", bookCtrl.GetBooks)
	r.POST("/books", bookCtrl.AddBook)
	r.GET("/books/:id", bookCtrl.GetBook)
	r.PUT("/books/:id", bookCtrl.UpdateBook)
	r.DELETE("/books/:id", bookCtrl.DeleteBook)

	log.Println("Server listening on port", port)

	return http.ListenAndServe(port, r)
}

/*
w response: respuesta del servidor al cliente
r request: peticion del cliente al servidor
*/
func index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Bienvenido a mi increible API!")
}
