package service

import (
	"github.com/cora23tt/onlinedilerv3"
	"github.com/cora23tt/onlinedilerv3/pkg/repository"
)

type OrderItemsService struct {
	repo repository.OrderItems
}

func NewOrderItemsService(repo repository.OrderItems) *OrderItemsService {
	return &OrderItemsService{repo: repo}
}

func (s *OrderItemsService) GetItems(orderID int) ([]onlinedilerv3.OrderItem, error) {
	return s.repo.GetItems(orderID)
}

func (s *OrderItemsService) Add(orderID int, input onlinedilerv3.OrderItem) ([]onlinedilerv3.OrderItem, error) {
	return s.repo.Add(orderID, input)
}

func (s *OrderItemsService) Update(orderID int, input onlinedilerv3.OrderItem) error {
	return s.repo.Update(orderID, input)
}

func (s *OrderItemsService) Delete(orderID int, itemID int) error {
	return s.repo.Delete(orderID, itemID)
}
