package internal

import (
	"database/sql"
	"log"

	"github.com/dignelidxdx/HackthonGo/internal/models"
	"github.com/dignelidxdx/HackthonGo/pkg/db"
)

type SaleRepository interface {
	Save(sale models.Sale) (models.Sale, error)
	SaveFile(sales []models.Sale) error
	GetOneByID(id int) (models.Sale, error)
	GetQuantityAndIDProduct() (float64, int, error)
	GetAll() ([]models.Sale, error)
}

type saleRepository struct {
}

func NewSaleRepository() SaleRepository {
	return &saleRepository{}
}

func (r *saleRepository) Save(sale models.Sale) (models.Sale, error) {
	db := db.StorageDB

	stmt, err := db.Prepare("INSERT INTO sales (idinvoice, idproduct, quantity) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(sale.Invoice.ID, sale.Product.ID, sale.Quantity)
	if err != nil {
		return models.Sale{}, err
	}
	idCreado, _ := result.LastInsertId()
	sale.ID = int(idCreado)

	return sale, nil
}

func (r *saleRepository) SaveFile(sales []models.Sale) error {
	db := db.StorageDB

	for _, sale := range sales {

		stmt, err := db.Prepare("INSERT INTO sales (id, idinvoice, idproduct, quantity) VALUES(?,?,?,?)")
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(sale.ID, sale.Invoice.ID, sale.Product.ID, sale.Quantity)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *saleRepository) GetOneByID(id int) (models.Sale, error) {

	db := db.StorageDB
	var SaleRead models.Sale
	rows, err := db.Query("SELECT id, idinvoice, idproduct, quantity FROM sales WHERE id = ?", id)

	if err != nil {
		log.Fatal(err)
		return SaleRead, err
	}

	for rows.Next() {
		err = rows.Scan(&SaleRead.ID, &SaleRead.Invoice.ID, &SaleRead.Product.ID, &SaleRead.Quantity)
		if err != nil {
			log.Fatal(err)
			return SaleRead, err
		}
	}

	return SaleRead, nil
}

func (r *saleRepository) GetQuantityAndIDProduct() (float64, int, error) {

	var sales []models.Sale
	var saleRead models.Sale

	db := db.StorageDB

	rows, err := db.Query("SELECT id, idinvoice, idproduct, quantity FROM sales")

	if err != nil {
		log.Fatal(err)
		return 0.0, 0, err
	}

	for rows.Next() {
		err = rows.Scan(&saleRead.ID, &saleRead.Invoice.ID, &saleRead.Product.ID, &saleRead.Quantity)
		if err != nil {
			log.Fatal(err)
			return 0.0, 0, err
		}
		sales = append(sales, saleRead)
	}
	return saleRead.Quantity, saleRead.Product.ID, nil
}

func (r *saleRepository) GetAll() ([]models.Sale, error) {
	var sales []models.Sale
	db := db.StorageDB
	var saleRead models.Sale
	rows, err := db.Query("SELECT id, idinvoice, idproduct, quantity FROM sales")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&saleRead.ID, &saleRead.Invoice.ID, &saleRead.Product.ID, &saleRead.Quantity)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		sales = append(sales, saleRead)
	}
	return sales, nil
}
