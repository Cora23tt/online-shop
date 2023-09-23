package service

import (
	"github.com/cora23tt/onlinedilerv3"
	"github.com/cora23tt/onlinedilerv3/pkg/repository"
)

type ConsignmentService struct {
	repo repository.Consignments
}

func NewConsignmentService(repo repository.Consignments) *ConsignmentService {
	return &ConsignmentService{repo: repo}
}

func (s *ConsignmentService) GetAll() ([]onlinedilerv3.Consignment, error) {
	return s.repo.GetAll()
}

func (s *ConsignmentService) GetByID(id int) (onlinedilerv3.Consignment, error) {
	return s.repo.GetByID(id)
}

func (s *ConsignmentService) Create(input onlinedilerv3.Consignment) (int, error) {
	return s.repo.Create(input)
}

func (s *ConsignmentService) Update(id int, input onlinedilerv3.Consignment) error {
	return s.repo.Update(id, input)
}

func (s *ConsignmentService) Delete(id int) error {
	return s.repo.Delete(id)
}
