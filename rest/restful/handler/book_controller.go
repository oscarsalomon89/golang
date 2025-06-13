package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/teamcubation/neocamp-meli/rest/restful/entity"
	"github.com/teamcubation/neocamp-meli/rest/restful/repository"
)

type ResponseInfo struct {
	Error  bool `json:"error"`
	Result any  `json:"result"`
}

type BookHandler struct {
	Repo repository.BookRepository
}

func (h *BookHandler) GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, ResponseInfo{
		Error:  false,
		Result: h.Repo.GetBooks(),
	})
}

func (h *BookHandler) GetBook(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: "o parametro nao e um numero",
		})
		return
	}

	book := h.Repo.GetBook(id)
	if book.ID == 0 {
		c.JSON(http.StatusNotFound, ResponseInfo{
			Error:  true,
			Result: "book not found",
		})
		return
	}

	c.JSON(http.StatusOK, ResponseInfo{
		Error:  false,
		Result: book,
	})
}

func (h *BookHandler) GetBookByName(c *gin.Context) {
	title := c.Param("name")

	book := h.Repo.GetBookByName(title)
	if book.ID == 0 {
		c.JSON(http.StatusNotFound, ResponseInfo{
			Error:  true,
			Result: "book not found",
		})
		return
	}

	c.JSON(http.StatusOK, ResponseInfo{
		Error:  false,
		Result: book,
	})
}

func (h *BookHandler) AddBook(c *gin.Context) {
	var book entity.Book
	err := json.NewDecoder(c.Request.Body).Decode(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	if book.Author == "" || book.Title == "" || book.Price == 0 {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: "todos os campos sao obrigatorios",
		})
		return
	}

	h.Repo.AddBook(book)

	c.JSON(http.StatusCreated, ResponseInfo{
		Error:  false,
		Result: "criado com sucesso",
	})
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	var book entity.Book
	err = json.NewDecoder(c.Request.Body).Decode(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	book.ID = id
	h.Repo.UpdateBook(book)

	c.JSON(http.StatusOK, ResponseInfo{
		Error:  false,
		Result: book,
	})
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	h.Repo.DeleteBook(id)

	c.JSON(http.StatusOK, ResponseInfo{
		Error:  false,
		Result: "deletado com sucesso",
	})
}
