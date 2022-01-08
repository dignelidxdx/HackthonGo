package internal

import (
	"database/sql"
	"log"

	"github.com/dignelidxdx/HackthonGo/internal/models"
	"github.com/dignelidxdx/HackthonGo/pkg/db"
)

type InvoiceRepository interface {
	SaveFile(invoices []models.Invoice) error
	Save(invoice models.Invoice) (models.Invoice, error)
}

type invoiceRepository struct {
}

func NewInvoiceRepository() InvoiceRepository {
	return &invoiceRepository{}
}

func (r *invoiceRepository) SaveFile(invoices []models.Invoice) error {
	db := db.StorageDB

	for _, c := range invoices {

		// TODO: cambiar datos del insert
		stmt, err := db.Prepare("INSERT INTO invoices(last_name,first_name,`condition`) VALUES(?,?,?)")
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(c.Total, c.Customer, c.Datetime)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *invoiceRepository) Save(invoice models.Invoice) (models.Invoice, error) {
	db := db.StorageDB

	stmt, err := db.Prepare("INSERT INTO invoices(last_name,first_name,`condition`) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(invoice.Total, invoice.Customer, invoice.Datetime)
	if err != nil {
		return models.Invoice{}, err
	}
	idCreado, _ := result.LastInsertId()
	invoice.ID = int(idCreado)

	return invoice, nil
}
