package out

import (
	"context"

	"github.com/dignelidxdx/HackthonGo/wrapper/domain"
)

type GetAllEmployees interface {
	GetAllEmployees() ([]domain.Employee, error)
}

type EmployeeGorm interface {
	GetOneEmployee(id int) (*domain.Employee, error)
	GetAllEmployees() ([]domain.Employee, error)
	CreateEmployee(*domain.Employee) (int, error)
	DeleteEmployee(id int) (int, error)
	DeleteEmployeeByFisrtNameAndLastName(id int, FirstName, lastName string) (int, error)
	UpdateEmployee(*domain.Employee) (*domain.Employee, error)
}

type GetEmployee interface {
	GetEmployees(ctx context.Context, options *domain.Employee) (*domain.Employee, error)
}
