package service

import (
	"github.com/cora23tt/onlinedilerv3"
	"github.com/cora23tt/onlinedilerv3/pkg/repository"
)

type ClientDiscountService struct {
	repo repository.ClientDiscounts
}

func NewClientDiscountService(repo repository.ClientDiscounts) *ClientDiscountService {
	return &ClientDiscountService{repo: repo}
}

func (s *ClientDiscountService) GetAll() ([]onlinedilerv3.ClientDiscount, error) {
	return s.repo.GetAll()
}

func (s *ClientDiscountService) GetByID(id int) (onlinedilerv3.ClientDiscount, error) {
	return s.repo.GetByID(id)
}

func (s *ClientDiscountService) Create(input onlinedilerv3.ClientDiscount) (int, error) {
	return s.repo.Create(input)
}

func (s *ClientDiscountService) Update(id int, input onlinedilerv3.ClientDiscount) error {
	return s.repo.Update(id, input)
}

func (s *ClientDiscountService) Delete(id int) error {
	return s.repo.Delete(id)
}
