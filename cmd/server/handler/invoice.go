package handler

import (
	"fmt"

	invoice "github.com/dignelidxdx/HackthonGo/internal/invoices"
	"github.com/dignelidxdx/HackthonGo/pkg/web"
	"github.com/gin-gonic/gin"
)

type InvoiceHandler struct {
	service invoice.InvoiceService
}

func NewInvoice(s invoice.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{service: s}
}

func (invoice *InvoiceHandler) UpdateAllTotal() gin.HandlerFunc {

	return func(context *gin.Context) {

		err := invoice.service.UpdateAllTotal()

		if err != nil {
			context.JSON(400, web.NewResponse(400, "", fmt.Sprintf("There was a error %v", err)))
		} else {
			context.JSON(200, web.NewResponse(200, "Se actualizo correctamente el total de los invoices", ""))
		}

	}
}
