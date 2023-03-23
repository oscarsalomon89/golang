package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/teamcubation/neocamp-meli/clean-architecture/internal/domain"
)

type BookController interface {
	GetBooks(ctx *gin.Context)
	GetBook(ctx *gin.Context)
	AddBook(ctx *gin.Context)
	UpdateBook(ctx *gin.Context)
	DeleteBook(ctx *gin.Context)
}

type bookController struct {
	bookUsecase domain.BookUsecase
}

func NewBookController(bookUC domain.BookUsecase) BookController {
	return &bookController{
		bookUsecase: bookUC,
	}
}

func (c *bookController) GetBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  c.bookUsecase.GetBooks(),
	})
}

type ResponseInfo struct {
	Error bool   `json:"error"`
	Data  string `json:"data"`
}

func (c *bookController) GetBook(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  err.Error(),
		})
		return
	}

	book := c.bookUsecase.GetBook(id)
	if book != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": false,
			"data":  book,
		})
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"error": true,
		"data":  "book not found",
	})
}

func (c *bookController) AddBook(ctx *gin.Context) {
	request := ctx.Request

	var book domain.Book
	err := json.NewDecoder(request.Body).Decode(&book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  err.Error(),
		})
		return
	}

	result, err := c.bookUsecase.AddBook(book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  result,
	})
}

func (c *bookController) UpdateBook(ctx *gin.Context) {
	r := ctx.Request
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  err.Error(),
		})
		return
	}

	var book domain.Book
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  err.Error(),
		})
		return
	}

	result, err := c.bookUsecase.UpdateBook(id, book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"data":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  result,
	})
}

func (c *bookController) DeleteBook(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  err.Error(),
		})
		return
	}

	_ = c.bookUsecase.DeleteBook(id)

	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  "deleted",
	})
}
