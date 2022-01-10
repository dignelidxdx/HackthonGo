package internal

import "github.com/dignelidxdx/HackthonGo/internal/models"

type SaleService interface {
	SaveSale(models.Sale) (models.Sale, error)
	SaveFile([]models.Sale) error
	GetOneByID(id int) (models.Sale, error)
}

type saleService struct {
	repository SaleRepository
}

func NewSaleService(repository SaleRepository) SaleService {
	return &saleService{repository: repository}
}

func (s *saleService) SaveSale(sale models.Sale) (models.Sale, error) {

	sale, err := s.repository.Save(sale)
	if err != nil {
		return models.Sale{}, err
	}
	return sale, nil
}

func (s *saleService) SaveFile(sales []models.Sale) error {

	err := s.repository.SaveFile(sales)
	if err != nil {
		return err
	}
	return nil
}

func (s *saleService) GetOneByID(id int) (models.Sale, error) {
	return s.repository.GetOneByID(id)
}
