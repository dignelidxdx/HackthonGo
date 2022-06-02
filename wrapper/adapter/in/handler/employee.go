package handler

import (
	"net/http"

	port "github.com/dignelidxdx/HackthonGo/wrapper/application/port/in"
	"github.com/dignelidxdx/HackthonGo/wrapper/domain"
	"github.com/gin-gonic/gin"
)

type employee struct {
	// Acá agregas la interfaz (en /port/in) como propiedad para llamar al service
	employeeService port.GetAllEmployeeUseCase
}

// Constructor
func NewEmployee(employeeService port.GetAllEmployeeUseCase) EmployeeController {
	return &employee{
		employeeService: employeeService,
	}
}

// Interfaz del controller
type EmployeeController interface {
	GetAll() gin.HandlerFunc
}

func (e *employee) GetAll() gin.HandlerFunc {
	return func(context *gin.Context) {

		employees, err := e.employeeService.GetAll(context)
		if err != nil {
			context.JSON(http.StatusBadRequest, domain.Employee{})
		} else if len(employees) == 0 {
			context.JSON(http.StatusBadRequest, domain.Employee{})
		} else {
			context.JSON(http.StatusOK, employees)
			return
		}
	}
}
