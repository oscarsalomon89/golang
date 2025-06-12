package service

import "my-api/repository"

type userService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) userService {
	return userService{repo: repo}
}

func (s userService) FindAll() []string {
	return s.repo.FindAll()
}
