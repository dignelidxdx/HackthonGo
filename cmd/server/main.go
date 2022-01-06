package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	router := gin.Default()

	// es necesario recalcular con los datos que dispone entre sales, invoices y products
	router.GET("/customers/total/:condition", getTotalCustomersByCondition)
	// para pasar datos a la base de datos
	router.POST("/datas/backups", doBackUpAllData)
	// para pasar datos a la base de datos
	router.GET("/invoices", getAllInvoices)
	router.GET("/products", getAllProducts)
	router.GET("/customers", getAllCustomers)
	router.GET("/sales", getAllSales)
	// INSERT
	router.POST("/invoices", createInvoice)
	router.POST("/products", createProduct)
	router.POST("/customers", createCustomer)
	router.POST("/sales", createSale)
	// UPDATE
	router.PUT("/invoices", updateInvoiceById)
	router.PUT("/products", updateProductById)
	router.PUT("/customers", updateCustomerById)
	router.PUT("/sales", updateSaleById)

	router.Run(":9090")
}

// TOTAL

func getTotalCustomersByCondition(context *gin.Context) {

	context.JSON(200, gin.H{"total": "100"})
}

func doBackUpAllData(context *gin.Context) {

	context.JSON(200, gin.H{"total": "100"})
}

// GET ALL

func getAllCustomers(context *gin.Context) {

	context.JSON(200, gin.H{"total": "100"})
}

func getAllSales(context *gin.Context) {

	context.JSON(200, gin.H{"total": "100"})
}

func getAllProducts(context *gin.Context) {

	context.JSON(200, gin.H{"total": "100"})
}

func getAllInvoices(context *gin.Context) {

	context.JSON(200, gin.H{"total": "100"})
}

// CREATE

func createCustomer(context *gin.Context) {

	context.JSON(200, gin.H{"total": "100"})
}

func createSale(context *gin.Context) {

	context.JSON(200, gin.H{"total": "100"})
}

func createProduct(context *gin.Context) {

	context.JSON(200, gin.H{"total": "100"})
}

func createInvoice(context *gin.Context) {

	context.JSON(200, gin.H{"total": "100"})
}

// UPDATE

func updateCustomerById(context *gin.Context) {

	context.JSON(200, gin.H{"total": "100"})
}

func updateSaleById(context *gin.Context) {

	context.JSON(200, gin.H{"total": "100"})
}

func updateProductById(context *gin.Context) {

	context.JSON(200, gin.H{"total": "100"})
}

func updateInvoiceById(context *gin.Context) {

	context.JSON(200, gin.H{"total": "100"})
}
