package itemhdl

import (
	"encoding/json"
	"net/http"

	"github.com/osalomon/market-api/internal/application/item"
)

type itemhdl struct {
	useCase item.ItemUseCase
}

func NewHandler(useCase item.ItemUseCase) *itemhdl {
	return &itemhdl{useCase: useCase}
}

func (h *itemhdl) CreateItem(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name string `json:"name"`
	}

	json.NewDecoder(r.Body).Decode(&body)
	err := h.useCase.CreateItem(body.Name)
	if err != nil {
		http.Error(w, "Erro ao criar item", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
