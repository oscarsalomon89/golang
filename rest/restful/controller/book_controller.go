package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/teamcubation/neocamp-meli/rest/restful/entity"
)

var db []entity.Book
var id int = 4

type ResponseInfo struct {
	Error  bool `json:"error"`
	Result any  `json:"result"`
}

func GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, ResponseInfo{
		Error:  false,
		Result: db,
	})
}

func GetBook(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: "o parametro nao e um numero",
		})
		return
	}

	//chamar a repository para procurar na tabela getBook(id) book

	for _, v := range db {
		if v.ID == id {
			c.JSON(http.StatusOK, ResponseInfo{
				Error:  false,
				Result: v,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, ResponseInfo{
		Error:  true,
		Result: "book not found",
	})
}

func GetBookByName(c *gin.Context) {
	title := c.Param("name")

	for _, v := range db {
		if v.Title == title {
			c.JSON(http.StatusOK, ResponseInfo{
				Error:  false,
				Result: v,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, ResponseInfo{
		Error:  true,
		Result: "book not found",
	})
}

func AddBook(c *gin.Context) {
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

	book.ID = id
	db = append(db, book)

	id++

	c.JSON(http.StatusCreated, ResponseInfo{
		Error:  false,
		Result: "criado com sucesso",
	})
}

func UpdateBook(c *gin.Context) {
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

	for i, v := range db {
		if v.ID == id {
			db[i] = book
		}
	}

	c.JSON(http.StatusOK, ResponseInfo{
		Error:  false,
		Result: db,
	})
}

func DeleteBook(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	for i, v := range db {
		if v.ID == id {
			db = append(db[:i], db[i+1:]...)
		}
	}

	c.JSON(http.StatusOK, ResponseInfo{
		Error:  false,
		Result: "deletado com sucesso",
	})
}
