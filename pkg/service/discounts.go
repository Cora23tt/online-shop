package service

import (
	"github.com/cora23tt/onlinedilerv3"
	"github.com/cora23tt/onlinedilerv3/pkg/repository"
)

type DiscountService struct {
	repo repository.Discounts
}

func NewDiscountService(repo repository.Discounts) *DiscountService {
	return &DiscountService{repo: repo}
}

func (s *DiscountService) GetAll() ([]onlinedilerv3.Discount, error) {
	return s.repo.GetAll()
}

func (s *DiscountService) GetByID(id int) (onlinedilerv3.Discount, error) {
	return s.repo.GetByID(id)
}

func (s *DiscountService) Create(input onlinedilerv3.DiscountInput) (int, error) {
	return s.repo.Create(input)
}

func (s *DiscountService) Update(id int, input onlinedilerv3.DiscountInput) error {
	return s.repo.Update(id, input)
}

func (s *DiscountService) Delete(id int) error {
	return s.repo.Delete(id)
}
