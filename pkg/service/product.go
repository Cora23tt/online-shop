package service

import (
	"github.com/cora23tt/onlinedilerv3"
	"github.com/cora23tt/onlinedilerv3/pkg/repository"
)

type ProductService struct {
	repo repository.Products
}

func NewProductService(repo repository.Products) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Search(lang, productName string) ([]onlinedilerv3.ProductComplex, error) {
	return s.repo.Search(lang, productName)
}

func (s *ProductService) SearchWithLimit(limit, offset int, lang, productName string) ([]onlinedilerv3.ProductComplex, error) {
	return s.repo.SearchWithLimit(limit, offset, lang, productName)
}

func (s *ProductService) ByCategory(language string, categoryID int) ([]onlinedilerv3.ProductTranslationConsignment, error) {
	return s.repo.ByCategory(language, categoryID)
}

func (s *ProductService) TopRated(lang string) ([]onlinedilerv3.ProductComplex, error) {
	return s.repo.TopRated(lang)
}

func (s *ProductService) TopRatedWithLimit(limit, offset int, lang string) ([]onlinedilerv3.ProductComplex, error) {
	return s.repo.TopRatedWithLimit(limit, offset, lang)
}

func (s *ProductService) GetAll(lang string) ([]onlinedilerv3.ProductTranslationConsignment, error) {
	return s.repo.GetAll(lang)
}

func (s *ProductService) Create(product onlinedilerv3.Product, translations []onlinedilerv3.ProductTranslation) (int, error) {
	return s.repo.Create(product, translations)
}

func (s *ProductService) GetByID(lang string, id int) (onlinedilerv3.ProductTranslationConsignment, error) {
	return s.repo.GetByID(lang, id)
}

func (s *ProductService) Update(id int, input onlinedilerv3.ProductComplect) error {
	return s.repo.Update(id, input)
}

func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}
