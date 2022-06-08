package out

import (
	"context"

	"github.com/dignelidxdx/HackthonGo/wrapper/domain"
)

type GetAllEmployees interface {
	GetAllEmployees() ([]domain.Employee, error)
}

type GetOneEmployee interface {
	GetOneEmployee(id int) (*domain.Employee, error)
}

type GetEmployee interface {
	GetEmployees(ctx context.Context, options *domain.Employee) (*domain.Employee, error)
}
