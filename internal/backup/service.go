package internal

import (
	"fmt"
	"strconv"

	customer "github.com/dignelidxdx/HackthonGo/internal/customers"
	invoice "github.com/dignelidxdx/HackthonGo/internal/invoices"
	"github.com/dignelidxdx/HackthonGo/internal/models"
	product "github.com/dignelidxdx/HackthonGo/internal/products"
	sale "github.com/dignelidxdx/HackthonGo/internal/sales"
	"github.com/dignelidxdx/HackthonGo/pkg/util"
)

type BackUpService interface {
	IsSaved(nameFile string) (bool, error)
	ToSave(id int, nameFile string) (bool, error)
	SaveFile(nameFile string) ([][]string, error)
	SaveElementToDB(key string) error
}

type backUpService struct {
	repository     BackUpRepository
	customerS      customer.CustomerService
	invoiceS       invoice.InvoiceService
	productService product.ProductService
	saleService    sale.SaleService
}

func NewBackUpService(
	repository BackUpRepository,
	customerService customer.CustomerService,
	invoiceService invoice.InvoiceService,
	productService product.ProductService,
	saleService sale.SaleService) BackUpService {
	return &backUpService{
		repository:     repository,
		customerS:      customerService,
		invoiceS:       invoiceService,
		productService: productService,
		saleService:    saleService}
}

func (s *backUpService) IsSaved(nameFile string) (bool, error) {

	return s.repository.isLocked(nameFile)

}

func (s *backUpService) ToSave(id int, nameFile string) (bool, error) {

	return s.repository.SaveToLock(nameFile, id)

}

func (s *backUpService) SaveFile(nameFile string) ([][]string, error) {

	err := util.ConvertToCsv(nameFile)
	if err != nil {
		return nil, err
	}
	lines, err := util.ReadCsv(nameFile)
	if err != nil {
		panic(err)
	}
	return lines, nil

}

func (s *backUpService) SaveElementToDB(key string) error {

	fmt.Println(key)
	isSaved, err := s.IsSaved(key)

	if err != nil {
		return err
	}
	if isSaved {
		return fmt.Errorf("ya se guardo el txt de %v", key)
	}
	err = util.ConvertToCsv(key)
	if err != nil {
		return err
	}
	lines, err := util.ReadCsv(key)
	if err != nil {
		panic(err)
	}

	var indx int

	switch key {
	case "customers":
		customers := []models.Customer{}

		// Loop through lines & turn into object
		for _, line := range lines {
			id, _ := strconv.Atoi(line[0])
			data := models.Customer{
				ID:        id,
				LastName:  line[1],
				FirstName: line[2],
				Condition: line[3],
			}
			customers = append(customers, data)

		}
		fmt.Println("antes del customerService customer")
		err = s.customerS.SaveFile(customers)
		fmt.Println("despues del customerService")
		if err != nil {
			return err
		}
		indx = 1
	case "sales":
		sales := []models.Sale{}

		// Loop through lines & turn into object
		for _, line := range lines {
			quantity, err := strconv.ParseFloat(line[3], 64)
			if err != nil {
				return err
			}
			invoiceID, _ := strconv.Atoi(line[1])
			productID, _ := strconv.Atoi(line[2])
			invoice, err := s.invoiceS.GetOneByID(invoiceID)
			if err != nil {
				return err
			}
			product, err := s.productService.GetOneByID(productID)
			if err != nil {
				return err
			}
			id, _ := strconv.Atoi(line[0])
			data := models.Sale{
				ID:       id,
				Invoice:  invoice,
				Product:  product,
				Quantity: quantity,
			}
			sales = append(sales, data)

		}
		err = s.saleService.SaveFile(sales)
		if err != nil {
			return err
		}
		indx = 2
	case "products":
		products := []models.Product{}
		// Loop through lines & turn into object
		for _, line := range lines {
			id, _ := strconv.Atoi(line[0])
			price, err := strconv.ParseFloat(line[2], 64)
			if err != nil {
				return err
			}
			data := models.Product{
				ID:          id,
				Description: line[1],
				Price:       price,
			}
			products = append(products, data)

		}
		err = s.productService.SaveFile(products)
		if err != nil {
			return err
		}
		indx = 3
	case "invoices":
		invoices := []models.Invoice{}
		// Loop through lines & turn into object
		for _, line := range lines {
			dateString := line[1]
			//convert string to time.Time type

			id, _ := strconv.Atoi(line[0])
			customer, err := s.customerS.GetOneByID(id)
			if err != nil {
				return err
			}

			customerID, _ := strconv.Atoi(line[2])
			customer, err = s.customerS.GetOneByID(customerID)
			if err != nil {
				return err
			}

			data := models.Invoice{
				Datetime: dateString,
				Customer: customer,
				ID:       id,
			}
			fmt.Println("invoices:", data)
			invoices = append(invoices, data)

		}
		err = s.invoiceS.SaveFile(invoices)
		if err != nil {
			return err
		}
		indx = 4
	default:
		return fmt.Errorf("la key es incorrecta")
	}

	_, err = s.ToSave(indx, key)
	if err != nil {
		return err
	}
	return nil
}
