package in

import (
	"context"

	"github.com/dignelidxdx/HackthonGo/wrapper/domain"
)

type GetAllEmployeeUseCase interface {
	GetAll(ctx context.Context) ([]domain.Employee, error)
	GetOne(context.Context, int) (domain.Employee, error)
	CreateOne(context.Context, domain.Employee) (int, error)
	DeleteOne(context.Context, int) (int, error)
}
