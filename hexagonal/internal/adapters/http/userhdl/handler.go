package userhdl

import (
	"encoding/json"
	"net/http"

	"github.com/osalomon/market-api/internal/application/user"
)

type Handler struct {
	usecase user.UserUseCase
}

func NewHandler(usecase user.UserUseCase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name string `json:"name"`
	}
	json.NewDecoder(r.Body).Decode(&body)
	err := h.usecase.CreateUser(body.Name)
	if err != nil {
		http.Error(w, "Erro ao criar usuário", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
