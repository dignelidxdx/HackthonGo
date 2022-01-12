package internal

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/dignelidxdx/HackthonGo/internal/models"
	"github.com/dignelidxdx/HackthonGo/pkg/db"
	"github.com/dignelidxdx/HackthonGo/pkg/web"
)

type ConditionEnum int

const (
	Locked   ConditionEnum = 0
	Inactive               = 1
	Active                 = 2
)

type CustomerRepository interface {
	SaveFile(customers []models.Customer) error
	Save(customer models.Customer) (models.Customer, error)
	GetTotalSecludedByCondition() []models.Customer
	GetCustomersOrderByLastName() []models.Customer
	GetOneByID(id int) (models.Customer, error)
	//Update(Customer models.Customer) (models.Customer, error)
	GetCustomerByMostCheapProduct() ([]web.ResponseCustomerCheapest, error)
}

type customerRepository struct {
}

func NewCustomerRepository() CustomerRepository {
	return &customerRepository{}
}

func (r *customerRepository) SaveFile(customers []models.Customer) error {
	db := db.StorageDB

	fmt.Println("repository customer")
	for _, c := range customers {

		stmt, err := db.Prepare("INSERT INTO customers(id, last_name,first_name,`condition`) VALUES(?,?,?,?)")
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(c.ID, c.LastName, c.FirstName, c.Condition)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *customerRepository) Save(customer models.Customer) (models.Customer, error) {
	db := db.StorageDB

	stmt, err := db.Prepare("INSERT INTO customers(last_name,first_name,`condition`) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(customer.LastName, customer.LastName, customer.Condition)
	if err != nil {
		return models.Customer{}, err
	}
	idCreado, _ := result.LastInsertId()
	customer.ID = int(idCreado)

	return customer, nil
}

func (r *customerRepository) GetTotalSecludedByCondition() []models.Customer {
	return []models.Customer{}
}

func (r *customerRepository) GetCustomersOrderByLastName() []models.Customer {
	return []models.Customer{}
}

func (r *customerRepository) GetOneByID(id int) (models.Customer, error) {

	db := db.StorageDB
	var customerRead models.Customer
	rows, err := db.Query("SELECT id, last_name, first_name, `condition` FROM customers WHERE id = ?", id)

	if err != nil {
		log.Fatal(err)
		return customerRead, err
	}

	for rows.Next() {
		err = rows.Scan(&customerRead.ID, &customerRead.LastName, &customerRead.FirstName, &customerRead.Condition)
		if err != nil {
			log.Fatal(err)
			return customerRead, err
		}
	}

	return customerRead, nil
}

func (r *customerRepository) GetCustomerByMostCheapProduct() ([]web.ResponseCustomerCheapest, error) {
	var customers []web.ResponseCustomerCheapest
	db := db.StorageDB
	var customerRead web.ResponseCustomerCheapest
	rows, err := db.Query("SELECT last_name, first_name FROM customers AS c INNER JOIN invoices AS i ON c.id = i.idcustomer INNER JOIN sales AS s ON s.idinvoice = i.id AND s.idproduct = (SELECT id FROM products ORDER BY price ASC LIMIT 1) ORDER BY last_name LIMIT 5")

	if err != nil {
		log.Fatal(err)
		return []web.ResponseCustomerCheapest{}, err
	}

	for rows.Next() {
		err = rows.Scan(&customerRead.LastName, &customerRead.FirstName)
		if err != nil {
			log.Fatal(err)
			return []web.ResponseCustomerCheapest{}, err
		}
		customers = append(customers, customerRead)
	}
	return customers, nil
}
