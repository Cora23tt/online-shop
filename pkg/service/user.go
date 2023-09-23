package service

import (
	"github.com/cora23tt/onlinedilerv3"
	"github.com/cora23tt/onlinedilerv3/pkg/repository"
)

type UserService struct {
	repo repository.Users
}

func NewUserService(repo repository.Users) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Search(name string) ([]onlinedilerv3.User, error) {
	return s.repo.Search(name)
}

func (s *UserService) GetAll() ([]onlinedilerv3.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetByID(id int) (onlinedilerv3.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *UserService) Update(id int, input onlinedilerv3.User) error {
	return s.repo.Update(id, input)
}
