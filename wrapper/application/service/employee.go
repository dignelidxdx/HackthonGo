package service

import (
	"context"

	"github.com/dignelidxdx/HackthonGo/wrapper/adapter/out/owner"
	"github.com/dignelidxdx/HackthonGo/wrapper/application/port/in"
	"github.com/dignelidxdx/HackthonGo/wrapper/application/port/out"
	"github.com/dignelidxdx/HackthonGo/wrapper/domain"
)

type service struct {
	// Aca agregas la interfaz (en /port/out) como propiedad para llamar al repository
	repository     out.GetAllEmployees
	gormRepo       out.GetOneEmployee
	client         out.GetEmployee
	circuitBreaker *owner.CircuitBreaker
}

// Constructor
func NewEmployeeService(repository out.GetAllEmployees, gorm out.GetOneEmployee, client out.GetEmployee, circuitBreaker *owner.CircuitBreaker) in.GetAllEmployeeUseCase {
	return &service{repository: repository, gormRepo: gorm, client: client, circuitBreaker: circuitBreaker}
}

// Es la implementacion de la inferfaz en application/port/in/GetAllEmployeeUseCase.go
func (service *service) GetAll(ctx context.Context) ([]domain.Employee, error) {

	//employees, err := service.repository.GetAllEmployees()
	res, err := service.client.GetEmployees(ctx, &domain.Employee{})
	if err != nil {
		return nil, err
	}
	res2 := append([]domain.Employee{}, *res)

	return res2, nil
	//return employees, nil
}

func (service *service) GetOne(ctx context.Context, id int) (domain.Employee, error) {

	result, err := service.gormRepo.GetOneEmployee(id)
	if err != nil {
		return domain.Employee{}, err
	}
	return *result, nil
}
