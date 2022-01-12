package internal

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/dignelidxdx/HackthonGo/internal/models"
	"github.com/dignelidxdx/HackthonGo/pkg/db"
)

type InvoiceRepository interface {
	GetAll() ([]models.Invoice, error)
	SaveFile(invoices []models.Invoice) error
	Save(invoice models.Invoice) (models.Invoice, error)
	GetOneByID(id int) (models.Invoice, error)
	UpdateTotal(id int, total float64) error
}

type invoiceRepository struct {
}

func NewInvoiceRepository() InvoiceRepository {
	return &invoiceRepository{}
}

func (r *invoiceRepository) GetAll() ([]models.Invoice, error) {
	var invoices []models.Invoice
	db := db.StorageDB
	var invoiceRead models.Invoice
	rows, err := db.Query("SELECT id, `datetime`, idcustomer, total FROM invoices")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&invoiceRead.ID, &invoiceRead.Datetime, &invoiceRead.Customer.ID, &invoiceRead.Total)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		invoices = append(invoices, invoiceRead)
	}
	return invoices, nil
}

func (r *invoiceRepository) SaveFile(invoices []models.Invoice) error {
	db := db.StorageDB

	for _, invoice := range invoices {

		// TODO: cambiar datos del insert
		stmt, err := db.Prepare("INSERT INTO invoices (id, `datetime`,idcustomer,total) VALUES(?, ?, ?, ?)")
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(invoice.ID, invoice.Datetime, invoice.Customer.ID, invoice.Total)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *invoiceRepository) Save(invoice models.Invoice) (models.Invoice, error) {
	db := db.StorageDB

	stmt, err := db.Prepare("INSERT INTO invoices (`datetime`,idcustomer,total) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(invoice.Datetime, invoice.Customer.ID, invoice.Total)
	if err != nil {
		return models.Invoice{}, err
	}
	idCreated, _ := result.LastInsertId()
	invoice.ID = int(idCreated)

	return invoice, nil
}

func (r *invoiceRepository) GetOneByID(id int) (models.Invoice, error) {

	db := db.StorageDB
	var invoiceRead models.Invoice
	rows, err := db.Query("SELECT id, `datetime`, idcustomer, coalesce(total, 0.00) FROM invoices WHERE id = ?", id)

	if err != nil {
		log.Fatal(err)
		return invoiceRead, err
	}

	for rows.Next() {

		err = rows.Scan(&invoiceRead.ID, &invoiceRead.Datetime, &invoiceRead.Customer.ID, &invoiceRead.Total)
		if err != nil {
			log.Fatal(err)
			return invoiceRead, err
		}
	}

	return invoiceRead, nil
}

func (r *invoiceRepository) UpdateTotal(id int, total float64) error {

	db := db.StorageDB

	stmt, err := db.Prepare("UPDATE invoices SET total = ? WHERE id = ?")
	if err != nil {
		log.Fatal("err", err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(total, id)
	if err != nil {
		fmt.Println("error:", err)
		return err
	}
	filasActualizadas, _ := result.RowsAffected()

	if filasActualizadas == 0 {
		fmt.Println("error:", err)
		return errors.New("No se encontro invoice por id")
	}

	return nil
}
