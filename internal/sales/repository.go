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
}

type saleRepository struct {
}

func NewSaleRepository() SaleRepository {
	return &saleRepository{}
}

func (r *saleRepository) Save(sale models.Sale) (models.Sale, error) {
	db := db.StorageDB

	stmt, err := db.Prepare("INSERT INTO sales(description,price) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(sale.Quantity, sale.Product)
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

		stmt, err := db.Prepare("INSERT INTO sales(last_name,first_name,`condition`) VALUES(?,?,?)")
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(sale.Quantity, sale.Product)
		if err != nil {
			return err
		}
	}
	return nil
}
