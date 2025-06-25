package main

import (
	"log"
	"net/http"

	repository "github.com/osalomon/market-api/internal/adapters/db"
	"github.com/osalomon/market-api/internal/adapters/http/itemhdl"
	"github.com/osalomon/market-api/internal/adapters/http/userhdl"
	itemApp "github.com/osalomon/market-api/internal/application/item"
	userApp "github.com/osalomon/market-api/internal/application/user"
	"github.com/osalomon/market-api/internal/domain/item"
	"github.com/osalomon/market-api/internal/domain/user"
)

type fakeUserRepo struct{}

func (f *fakeUserRepo) Save(u user.User) error { return nil }

type fakeItemRepo struct{}

func (f *fakeItemRepo) Save(i item.Item) error { return nil }

func main() {
	// user
	userRepo := &fakeUserRepo{}
	userUseCase := userApp.NewUserUseCase(userRepo)
	userHandler := userhdl.NewHandler(userUseCase)

	// item
	itemRepo := repository.NewItemRepository(nil)
	itemUseCase := itemApp.NewItemUseCase(itemRepo)
	itemHandler := itemhdl.NewHandler(itemUseCase)

	http.HandleFunc("/users", userHandler.CreateUser)
	http.HandleFunc("/items", itemHandler.CreateItem)

	log.Println("Servidor rodando em :8080")
	http.ListenAndServe(":8080", nil)
}
