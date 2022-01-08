package store

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/dignelidxdx/HackthonGo/internal/models"
	"github.com/gocarina/gocsv"
)

type Store interface {
	WriteCustomers(models.Customer) error
	WriteProducts(models.Product) error
	WriteSales(models.Invoice) error
	WriteInvoice(models.Sale) error
	ReadCustomers([]models.Customer) error
	ReadProducts(models.Product) error
	ReadSales(models.Invoice) error
	ReadInvoice(models.Sale) error
}

const (
	FileTypeC Type = "customer"
	FileTypeS Type = "sale"
	FileTypeP Type = "product"
	FileTypeI Type = "invoice"
)

type Type string

func New(store Type, fileName string) Store {

	switch store {
	case FileTypeC:
	case FileTypeS:
	case FileTypeP:
	case FileTypeI:
		fmt.Println("store:", store)
		return &FileStore{FileName: fileName}
	}
	return nil
}

type FileStore struct {
	FileName string
}

func (fs *FileStore) WriteCustomers(customer models.Customer) error {
	file, err := json.MarshalIndent(customer, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FileName, file, 0644)
}
func (fs *FileStore) WriteProducts(product models.Product) error {
	file, err := json.MarshalIndent(product, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FileName, file, 0644)
}
func (fs *FileStore) WriteSales(invoice models.Invoice) error {
	file, err := json.MarshalIndent(invoice, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FileName, file, 0644)
}
func (fs *FileStore) WriteInvoice(sale models.Sale) error {
	file, err := json.MarshalIndent(sale, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FileName, file, 0644)
}
func (fs *FileStore) ReadCustomers(customers []models.Customer) error {

	file, err := os.ReadFile(fs.FileName)

	if err != nil {

		return err
	}

	data := string(file)

	newData := strings.ReplaceAll(data, "#$%#", ";")

	os.WriteFile("../../datos/csv/customers2.csv", []byte(newData), 0644)

	in, err := os.Open("../../datos/csv/customers2.csv")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	clients := []*models.Customer{}

	if err := gocsv.UnmarshalFile(in, &clients); err != nil {
		panic(err)
	}
	for _, client := range clients {
		customers = append(customers, *client)
		fmt.Println("Hello, ", client.FirstName, client.LastName)
	}

	if err != nil {
		return err
	}
	return nil
}
func (fs *FileStore) ReadProducts(product models.Product) error {
	file, err := os.ReadFile(fs.FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &product)
}
func (fs *FileStore) ReadSales(sale models.Invoice) error {
	file, err := os.ReadFile(fs.FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &sale)
}
func (fs *FileStore) ReadInvoice(invoice models.Sale) error {
	file, err := os.ReadFile(fs.FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &invoice)
}
