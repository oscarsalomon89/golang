package router

import (
	"log"

	"github.com/gin-gonic/gin"
)

type HTTPServer interface {
	RegisterRouter()
	Run(port string) error
}

type httpServer struct {
	router      *gin.Engine
	bookHandler BookHandler
}

func NewHTTPServer(bookHandler BookHandler) HTTPServer {
	ginrouter := gin.Default()

	return &httpServer{
		router:      ginrouter,
		bookHandler: bookHandler,
	}
}

func (srv *httpServer) RegisterRouter() {
	basePath := "/api/v1/books"
	r := srv.router.Group(basePath)

	r.GET("/", srv.bookHandler.GetBooks)
	r.POST("/", srv.bookHandler.SaveBook)
}

func (srv *httpServer) Run(port string) error {
	log.Println("Server listening on port", port)

	return srv.router.Run(":" + port)
}
