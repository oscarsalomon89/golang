package itemhdl

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	uc "github.com/osalomon/market-api/internal/application/item"
	"github.com/osalomon/market-api/internal/domain/item"
)

type itemHandler struct {
	useCase uc.ItemUseCase
}

func NewHandler(useCase uc.ItemUseCase) *itemHandler {
	return &itemHandler{useCase: useCase}
}

type ItemRequest struct {
	Name        string  `json:"name" validate:"required,gt=3"`
	Price       float64 `json:"price" validate:"required,numeric,gt=0"`
	Description string  `json:"description" validate:"required"`
	Stock       int     `json:"stock" validate:"required,numeric,gte=0"`
}

func (i ItemRequest) ToItemDomain() item.Item {
	return item.Item{
		Name:        i.Name,
		Price:       i.Price,
		Description: i.Description,
		Stock:       i.Stock,
	}
}

func (h *itemHandler) CreateItem(c *gin.Context) {
	var body ItemRequest
	validate := validator.New()

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(body)

	itemDomain := body.ToItemDomain()

	err := h.useCase.CreateItem(&itemDomain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Item criado com sucesso", "item": itemDomain})
}

func (h *itemHandler) GetItemByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	itemDomain, err := h.useCase.GetItemByID(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"item": itemDomain})
}

func (h *itemHandler) GetItems(c *gin.Context) {
	items, err := h.useCase.GetItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": items})
}

func (h *itemHandler) DeleteItem(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.useCase.DeleteItem(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item deletado com sucesso"})
}
