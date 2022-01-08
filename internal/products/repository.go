package internal

import (
	"database/sql"
	"log"

	"github.com/dignelidxdx/HackthonGo/internal/models"
	"github.com/dignelidxdx/HackthonGo/pkg/db"
	"github.com/dignelidxdx/HackthonGo/pkg/store"
)

type ProductRepository interface {
	Save(product models.Product) (models.Product, error)
	GetAll() ([]models.Product, error)
	Update(Product models.Product) (models.Product, error)
	SaveFile(products []models.Product) error
}

type productRepository struct {
	db store.Store
}

func NewProductRepository(db store.Store) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) Save(product models.Product) (models.Product, error) {
	db := db.StorageDB

	stmt, err := db.Prepare("INSERT INTO products(description,price) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(product.Description, product.Price)
	if err != nil {
		return models.Product{}, err
	}
	idCreado, _ := result.LastInsertId()
	product.ID = int(idCreado)

	return product, nil
}

func (r *productRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	db := db.StorageDB
	var productsRead models.Product
	rows, err := db.Query("SELECT id, description, price FROM products")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&productsRead.ID, &productsRead.Description, &productsRead.Price)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		products = append(products, productsRead)
	}
	return products, nil
}

func (r *productRepository) Update(Product models.Product) (models.Product, error) {
	return models.Product{}, nil
}

func (r *productRepository) SaveFile(products []models.Product) error {
	db := db.StorageDB

	for _, product := range products {

		stmt, err := db.Prepare("INSERT INTO products(last_name,first_name,`condition`) VALUES(?,?,?)")
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(product.Description, product.Price)
		if err != nil {
			return err
		}
	}
	return nil
}
