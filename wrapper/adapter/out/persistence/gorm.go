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
func NewGormRepository(db *gorm.DB) out.EmployeeGorm {
	return &gormRepository{db: db}
}

// Es la implementacion de la inferfaz en application/port/out/GetAllEmployeePort.go
func (r *gormRepository) GetOneEmployee(id int) (*domain.Employee, error) {

	var result domain.Employee
	r.db.Table("customers").First(&result, id)
	return &result, nil
}

func (r *gormRepository) GetAllEmployees() ([]domain.Employee, error) {
	var employees []domain.Employee
	r.db.Table("customers").Find(&employees)
	return employees, nil
}
func (r *gormRepository) CreateEmployee(employee *domain.Employee) (int, error) {

	r.db.Table("customers").Create(&employee)
	return employee.ID, nil
}
func (r *gormRepository) DeleteEmployee(id int) (int, error) {

	r.db.Table("customers").Delete(&domain.Employee{ID: id})
	return id, nil
}
func (r *gormRepository) DeleteEmployeeByFisrtNameAndLastName(id int, firstName, lastName string) (int, error) {

	result := r.db.Table("customers").Where("first_name = ?", firstName, "last_name = ?", lastName).Delete(&domain.Employee{ID: id})
	if result.Error != nil {
		return 0, result.Error
	}
	return id, nil
}
func (r *gormRepository) UpdateEmployee(employee *domain.Employee) (*domain.Employee, error) {

	result := r.db.Table("customers").Model(&employee).Updates(employee)
	if result.Error != nil {
		return &domain.Employee{}, result.Error
	}
	return employee, nil
}
