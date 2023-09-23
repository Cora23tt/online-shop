package service

import (
	"github.com/cora23tt/onlinedilerv3"
	"github.com/cora23tt/onlinedilerv3/pkg/repository"
)

type OrderService struct {
	repo repository.Orders
}

func NewOrderService(repo repository.Orders) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) GetAll() ([]onlinedilerv3.Order, error) {
	return s.repo.GetAll()
}

func (s *OrderService) GetByID(id int) (onlinedilerv3.Order, error) {
	return s.repo.GetByID(id)
}

func (s *OrderService) Create(input onlinedilerv3.Order) (int, error) {
	return s.repo.Create(input)
}

func (s *OrderService) Update(id int, input onlinedilerv3.Order) error {
	return s.repo.Update(id, input)
}

func (s *OrderService) Delete(id int) error {
	return s.repo.Delete(id)
}
