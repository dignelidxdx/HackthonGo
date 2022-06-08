package persistence

import (
	"database/sql"

	"github.com/dignelidxdx/HackthonGo/pkg/db"
	"github.com/dignelidxdx/HackthonGo/wrapper/application/port/out"
	"github.com/dignelidxdx/HackthonGo/wrapper/config"
	"github.com/dignelidxdx/HackthonGo/wrapper/domain"
)

var (
	QueryGetAll = `SELECT id, last_name, first_name, 'condition' FROM customers`
	QueryCreate = `INSERT INTO customers(last_name, first_name, 'condition') VALUES (?,?,?)`
)

type employeeRepository struct {
	db     *sql.DB
	config config.DBConfiguration
}

// Constructor
func NewEmployeeRepository() (out.GetAllEmployees, out.CreateEmployee) {
	return &employeeRepository{}, &employeeRepository{}
}

// Es la implementacion de la inferfaz en application/port/out/GetAllEmployeePort.go
func (r *employeeRepository) GetAllEmployees() ([]domain.Employee, error) {

	db := db.StorageDB
	rows, err := db.Query(QueryGetAll)

	if err != nil {

		return nil, err
	}

	var employees []domain.Employee

	for rows.Next() {
		e := domain.Employee{}
		_ = rows.Scan(&e.ID, &e.LastName, &e.FirstName, &e.Condition)
		employees = append(employees, e)
	}

	return employees, nil
}

// Es la implementacion de la inferfaz en application/port/out/CreateEmployeePort.go
func (r *employeeRepository) CreateEmployee(e domain.Employee) (domain.Employee, error) {

	db := db.StorageDB
	stmt, err := db.Prepare(QueryCreate)
	if err != nil {
		return domain.Employee{}, err
	}

	res, err := stmt.Exec(&e.LastName, &e.FirstName, &e.Condition)
	if err != nil {
		return domain.Employee{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return domain.Employee{}, err
	}
	e.ID = int(id)

	return e, nil
}
