package persistence

import (
	"github.com/dignelidxdx/HackthonGo/wrapper/application/port/out"
	"github.com/dignelidxdx/HackthonGo/wrapper/domain"
	"gorm.io/gorm"
)

type gormRepository struct {
	db     *gorm.DB
	config gorm.Config
}

// Constructor
func NewGormRepository(db *gorm.DB) (out.GetOneEmployee, out.CreateEmployee) {
	return &gormRepository{db: db}, &gormRepository{db: db}
}

// Es la implementacion de la inferfaz en application/port/out/GetAllEmployeePort.go
func (r *gormRepository) GetOneEmployee(id int) (*domain.Employee, error) {

	var result domain.Employee
	r.db.Table("customers").First(&result, id)
	return &result, nil
}

// Es la implementacion de la inferfaz en application/port/out/CreateEmployeePort.go
func (r *gormRepository) CreateEmployee(e domain.Employee) (domain.Employee, error) {

	return domain.Employee{}, nil
}
