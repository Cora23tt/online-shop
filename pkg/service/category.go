package service

import (
	"github.com/cora23tt/onlinedilerv3"
	"github.com/cora23tt/onlinedilerv3/pkg/repository"
)

type CategoryService struct {
	repo repository.Categories
}

func NewCategoryService(repo repository.Categories) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAll(lang string) ([]onlinedilerv3.Category, error) {
	return s.repo.GetAll(lang)
}

func (s *CategoryService) Get(lang string, id int) (onlinedilerv3.Category, error) {
	return s.repo.Get(lang, id)
}

func (s *CategoryService) Create(input onlinedilerv3.Category) (int, error) {
	return s.repo.Create(input)
}

func (s *CategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *CategoryService) Update(id int, input onlinedilerv3.Category) error {
	return s.repo.Update(id, input)
}

func (s *CategoryService) Search(lang string, categoryName string) ([]onlinedilerv3.Category, error) {
	return s.repo.Search(lang, categoryName)
}
