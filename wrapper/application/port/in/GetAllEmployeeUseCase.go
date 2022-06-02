package in

import (
	"context"

	"github.com/dignelidxdx/HackthonGo/wrapper/domain"
)

type GetAllEmployeeUseCase interface {
	GetAll(ctx context.Context) ([]domain.Employee, error)
}
