package internal

import "github.com/dignelidxdx/HackthonGo/internal/models"

type ProductService interface {
	SaveProduct(models.Product) (models.Product, error)
	SaveFile([]models.Product) error
	GetOneByID(id int) (models.Product, error)
}

type productService struct {
	repository ProductRepository
}

func NewProductService(repository ProductRepository) ProductService {
	return &productService{repository: repository}
}

func (s *productService) SaveProduct(product models.Product) (models.Product, error) {

	product, err := s.repository.Save(product)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (s *productService) SaveFile(products []models.Product) error {

	err := s.repository.SaveFile(products)
	if err != nil {
		return err
	}
	return nil
}

func (s *productService) GetOneByID(id int) (models.Product, error) {
	return s.repository.GetOneByID(id)
}

func (s *productService) MustProductSold() error {
	return nil
}
