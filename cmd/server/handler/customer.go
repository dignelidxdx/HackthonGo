package handler

import (
	"fmt"

	customer "github.com/dignelidxdx/HackthonGo/internal/customers"
	"github.com/dignelidxdx/HackthonGo/internal/models"
	"github.com/dignelidxdx/HackthonGo/pkg/web"
	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	service customer.CustomerService
}

func NewCustomer(s customer.CustomerService) *CustomerHandler {
	return &CustomerHandler{service: s}
}

func (customer *CustomerHandler) SaveCustomer() gin.HandlerFunc {

	return func(context *gin.Context) {

		var request models.Customer
		err := context.ShouldBindJSON(&request)
		if err != nil {
			context.JSON(400, web.NewResponse(400, "", fmt.Sprintf("There was a error %v", err)))
		} else {

			customer, err := customer.service.SaveCustomer(request)
			if err != nil {
				context.JSON(400, web.NewResponse(400, "", fmt.Sprintf("There was a error %v", err)))
			} else {
				context.JSON(200, web.NewResponse(200, customer.FirstName, ""))
			}
		}

	}
}
