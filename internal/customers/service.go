package internal

import (
	"fmt"

	invoice "github.com/dignelidxdx/HackthonGo/internal/invoices"
	"github.com/dignelidxdx/HackthonGo/internal/models"
	"github.com/dignelidxdx/HackthonGo/pkg/web"
)

type CustomerService interface {
	SaveCustomer(models.Customer) (models.Customer, error)
	SaveFile(customers []models.Customer) error
	GetOneByID(id int) (models.Customer, error)
	GetTotalesByCondition() ([]web.ResponseTotalesByCondition, error)
	GetCustomerByMostCheapProduct() ([]web.ResponseCustomerCheapest, error)
}

type customerService struct {
	repository     CustomerRepository
	invoiceService invoice.InvoiceService
}

func NewCustomerService(repository CustomerRepository, invoiceService invoice.InvoiceService) CustomerService {
	return &customerService{repository: repository, invoiceService: invoiceService}
}

func (s *customerService) SaveCustomer(customer models.Customer) (models.Customer, error) {

	customer, err := s.repository.Save(customer)
	if err != nil {
		return models.Customer{}, err
	}
	return customer, nil
}

func (s *customerService) SaveFile(customers []models.Customer) error {

	err := s.repository.SaveFile(customers)
	if err != nil {
		return err
	}
	return nil
}

func (s *customerService) GetOneByID(id int) (models.Customer, error) {
	return s.repository.GetOneByID(id)
}

func (s *customerService) GetTotalesByCondition() ([]web.ResponseTotalesByCondition, error) {

	var totalInactivo float64
	var totalBloqueado float64
	var totalActivo float64

	invoices, err := s.invoiceService.GetAll()
	if err != nil {
		return []web.ResponseTotalesByCondition{}, err
	}

	for _, invoice := range invoices {
		// Acá es donde podriamos usar un cache para los id que ya fueron buscados con anterioridad
		customer, err := s.repository.GetOneByID(invoice.Customer.ID)
		if err != nil {
			return []web.ResponseTotalesByCondition{}, err
		}
		switch customer.Condition {
		case "Inactivo":
			totalInactivo += invoice.Total
		case "Bloqueado":
			totalBloqueado += invoice.Total
		case "Activo":
			totalActivo += invoice.Total
		default:
			return []web.ResponseTotalesByCondition{}, fmt.Errorf("la condicion es incorrecta")
		}
	}

	return []web.ResponseTotalesByCondition{{Condition: "Activo", Total: totalActivo}, {Condition: "Bloqueado", Total: totalBloqueado}, {Condition: "Inactivo", Total: totalInactivo}}, nil
}

func (s *customerService) GetCustomerByMostCheapProduct() ([]web.ResponseCustomerCheapest, error) {

	return s.repository.GetCustomerByMostCheapProduct()

}
