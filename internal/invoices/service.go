package internal

import (
	"github.com/dignelidxdx/HackthonGo/internal/models"
	product "github.com/dignelidxdx/HackthonGo/internal/products"
	sale "github.com/dignelidxdx/HackthonGo/internal/sales"
	"github.com/dignelidxdx/HackthonGo/pkg/web"
)

type InvoiceService interface {
	SaveInvoice(models.Invoice) (models.Invoice, error)
	SaveFile([]models.Invoice) error
	GetOneByID(id int) (models.Invoice, error)
	UpdateAllTotal() error
	GetAll() ([]models.Invoice, error)
}

type invoiceService struct {
	repository     InvoiceRepository
	saleService    sale.SaleService
	productService product.ProductService
}

func NewInvoiceService(
	repository InvoiceRepository,
	saleService sale.SaleService,
	productService product.ProductService) InvoiceService {
	return &invoiceService{repository: repository, saleService: saleService, productService: productService}
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

func (s *invoiceService) UpdateAllTotal() error {

	sales, err := s.saleService.GetAll()

	if err != nil {
		return err
	}

	for _, sale := range sales {

		product, err := s.productService.GetOneByID(sale.Product.ID)
		if err != nil {
			return err
		}
		invoice, err := s.GetOneByID(sale.Invoice.ID)
		if err != nil {
			return err
		}
		total := invoice.Total + (sale.Quantity * product.Price)

		err = s.repository.UpdateTotal(sale.Invoice.ID, total)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *invoiceService) GetAll() ([]models.Invoice, error) {
	return s.repository.GetAll()
}

func (s *invoiceService) MustProductSold() ([]web.ResponseProductSold, error) {

	sales, err := s.saleService.GetAll()

	if err != nil {
		return []web.ResponseProductSold{}, err
	}
	quantityByProducts := make(map[string]float64)
	totalPriceByProduct := make(map[string]float64)

	totalPriceQuantity := make(map[string]*web.ResultTotalQuantity)
	for _, sale := range sales {

		product, err := s.productService.GetOneByID(sale.Product.ID)
		if err != nil {
			return []web.ResponseProductSold{}, err
		}
		invoice, err := s.GetOneByID(sale.Invoice.ID)
		if err != nil {
			return []web.ResponseProductSold{}, err
		}
		_, isExists := quantityByProducts[product.Description]
		//_, isExists := totalPriceQuantity[product.Description]

		if isExists {
			quantity := quantityByProducts[product.Description]
			quantityByProducts[product.Description] = quantity + sale.Quantity

			total := totalPriceByProduct[product.Description]
			totalPriceByProduct[product.Description] = total + invoice.Total

			resultTotalQuantity := totalPriceQuantity[product.Description]
			totalPriceQuantity[product.Description] = &web.ResultTotalQuantity{Quantity: (resultTotalQuantity.Quantity + sale.Quantity), Total: (resultTotalQuantity.Total + invoice.Total)}

		} else {
			quantityByProducts[product.Description] = sale.Quantity

			totalPriceByProduct[product.Description] = invoice.Total

			totalPriceQuantity[product.Description] = &web.ResultTotalQuantity{Quantity: sale.Quantity, Total: invoice.Total}
		}
	}

	/*type resultToSend []*web.ResultTotalQuantity

	value := make(resultToSend, 0, len(totalPriceQuantity))
	i := 0
	for k, v := range totalPriceQuantity {
		value = append(value, &web.ResultTotalQuantity{v.Quantity, v.Total, k})
		i++
	}
	sort.Sort(*value)*/

	/*var count int
	var response []web.ResponseProductSold
	for _, k := range keys {

		fmt.Println(k, totalPriceQuantity[k])
		response = append(response, web.ResponseProductSold{Description: "s", Total: 0.0})
		if count == 4 {
			break
		}
		count += 1
	} */

	return []web.ResponseProductSold{}, nil
}
