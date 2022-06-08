package handler

import (
	"net/http"
	"strconv"

	port "github.com/dignelidxdx/HackthonGo/wrapper/application/port/in"
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
	GetOne() gin.HandlerFunc
}

func (e *employee) GetAll() gin.HandlerFunc {
	return func(context *gin.Context) {

		employees, err := e.employeeService.GetAll(context)
		if err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
		} else if len(employees) == 0 {
			context.JSON(http.StatusBadRequest, err.Error())
		} else {
			context.JSON(http.StatusOK, employees)
			return
		}
	}
}

func (e *employee) GetOne() gin.HandlerFunc {
	return func(context *gin.Context) {

		idParam, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}

		employee, err := e.employeeService.GetOne(context, idParam)

		if err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
		} else {
			context.JSON(http.StatusOK, employee)
			return
		}
	}
}
