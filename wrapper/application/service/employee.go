package service

import (
	"context"

	"github.com/dignelidxdx/HackthonGo/wrapper/application/port/in"
	"github.com/dignelidxdx/HackthonGo/wrapper/application/port/out"
	"github.com/dignelidxdx/HackthonGo/wrapper/domain"
)

type service struct {
	// Aca agregas la interfaz (en /port/out) como propiedad para llamar al repository
	repository out.GetAllEmployees
	client     out.GetEmployee
}

// Constructor
func NewEmployeeService(repository out.GetAllEmployees, client out.GetEmployee) in.GetAllEmployeeUseCase {
	return &service{repository: repository, client: client}
}

// Es la implementacion de la inferfaz en application/port/in/GetAllEmployeeUseCase.go
func (service *service) GetAll(ctx context.Context) ([]domain.Employee, error) {

	service.repository.GetAllEmployees()
	res, err := service.client.GetEmployees(ctx, &domain.Employee{})

	if err != nil {
		return nil, err
	}
	res2 := append([]domain.Employee{}, *res)
	return res2, nil
}
