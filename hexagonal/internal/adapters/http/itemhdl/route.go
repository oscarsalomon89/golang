package itemhdl

import "github.com/gin-gonic/gin"

const (
	itemPath = "/items"
)

type ItemHandlerRouter struct {
	itemhdl *itemHandler
}

func NewRouter(itemhdl *itemHandler) *ItemHandlerRouter {
	return &ItemHandlerRouter{
		itemhdl: itemhdl,
	}
}

func (r *ItemHandlerRouter) AddRoutesV1(v1 *gin.RouterGroup) {
	v1.POST(itemPath, r.itemhdl.CreateItem)
	v1.GET(itemPath, r.itemhdl.GetItems)
	v1.GET(itemPath+"/:id", r.itemhdl.GetItemByID)
	v1.DELETE(itemPath+"/:id", r.itemhdl.DeleteItem)
}
