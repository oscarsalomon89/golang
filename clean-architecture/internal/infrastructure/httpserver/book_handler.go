package router

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teamcubation/neocamp-meli/clean-architecture/internal/domain/model"
	"github.com/teamcubation/neocamp-meli/clean-architecture/internal/usecase"
)

type BookHandler interface {
	GetBooks(ctx *gin.Context)
	SaveBook(ctx *gin.Context)
}

type bookHandler struct {
	bookUsecase usecase.BookUsecase
}

func NewBookHandler(bookUsecase usecase.BookUsecase) BookHandler {
	return &bookHandler{
		bookUsecase: bookUsecase,
	}
}

func (h *bookHandler) GetBooks(ctx *gin.Context) {
	books := h.bookUsecase.GetAllBooks()

	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  books,
	})
}

func (h *bookHandler) SaveBook(ctx *gin.Context) {
	var book model.Book
	if err := json.NewDecoder(ctx.Request.Body).Decode(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": false,
			"data":  err.Error(),
		})
		return
	}

	result := h.bookUsecase.SaveBook(book)

	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  result,
	})
}
