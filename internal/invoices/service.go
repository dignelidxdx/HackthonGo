package internal

import "github.com/dignelidxdx/HackthonGo/internal/models"

type InvoiceService interface {
	SaveInvoice(models.Invoice) (models.Invoice, error)
	SaveFile([]models.Invoice) error
	GetOneByID(id int) (models.Invoice, error)
}

type invoiceService struct {
	repository InvoiceRepository
}

func NewInvoiceService(repository InvoiceRepository) InvoiceService {
	return &invoiceService{repository: repository}
}

func (s *invoiceService) SaveInvoice(invoice models.Invoice) (models.Invoice, error) {

	invoice, err := s.repository.Save(invoice)
	if err != nil {
		return models.Invoice{}, err
	}
	return invoice, nil
}

func (s *invoiceService) SaveFile(invoices []models.Invoice) error {

	err := s.repository.SaveFile(invoices)
	if err != nil {
		return err
	}
	return nil
}

func (s *invoiceService) GetOneByID(id int) (models.Invoice, error) {
	return s.repository.GetOneByID(id)
}
