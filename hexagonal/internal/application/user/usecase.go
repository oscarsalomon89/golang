package user

import "github.com/osalomon/market-api/internal/domain/user"

type UserUseCase interface {
	CreateUser(name string) error
}

type userUseCase struct {
	repo user.Repository
}

func NewUserUseCase(repo user.Repository) *userUseCase {
	return &userUseCase{repo: repo}
}

func (uc *userUseCase) CreateUser(name string) error {
	u := user.User{Name: name}
	return uc.repo.Save(u)
}
