package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/invoices/total", getTotalFromInvoice)

	router.Run(":9090")
}

func getTotalFromInvoice(context *gin.Context) {

	context.JSON(200, gin.H{"total": "100"})
}
