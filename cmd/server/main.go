package main

import (
	"github.com/dignelidxdx/HackthonGo/cmd/server/handler"
	backup "github.com/dignelidxdx/HackthonGo/internal/backup"
	customer "github.com/dignelidxdx/HackthonGo/internal/customers"
	invoice "github.com/dignelidxdx/HackthonGo/internal/invoices"
	product "github.com/dignelidxdx/HackthonGo/internal/products"
	sale "github.com/dignelidxdx/HackthonGo/internal/sales"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	router := gin.Default()

	// CUSTOMER
	repoC := customer.NewCustomerRepository()
	repoP := product.NewProductRepository()
	repoS := sale.NewSaleRepository()
	repoI := invoice.NewInvoiceRepository()

	serviceP := product.NewProductService(repoP)
	serviceS := sale.NewSaleService(repoS)
	serviceI := invoice.NewInvoiceService(repoI, serviceS, serviceP)
	serviceC := customer.NewCustomerService(repoC, serviceI)

	controllerC := handler.NewCustomer(serviceC)
	controllerI := handler.NewInvoice(serviceI)

	// BACKUP
	repoBa := backup.NewBackUpRepository()
	serviceBa := backup.NewBackUpService(repoBa, serviceC, serviceI, serviceP, serviceS)
	controllerBa := handler.NewBackUp(serviceBa)

	// es necesario recalcular con los datos que dispone entre sales, invoices y products
	//router.GET("/customers/total", controller.SaveFile())
	// producto top 5 mas vendidos

	// para pasar datos a la base de datos
	router.POST("/datas/backups", controllerBa.SaveFiles())
	// update total
	router.PATCH("/invoices/total", controllerI.UpdateAllTotal())
	// Get total by Condition
	router.GET("/customers/totales", controllerC.GetTotalesByCondition())
	// Get customer by cheapest product
	router.GET("/customers/cheapest-product", controllerC.GetCustomerByMostCheapProduct())

	// INSERT
	router.POST("/invoices", createInvoice)
	router.POST("/products", createProduct)
	router.POST("/customers", controllerC.SaveCustomer())
	router.POST("/sales", createSale)
	// UPDATE
	router.PUT("/invoices", updateInvoiceById)
	router.PUT("/products", updateProductById)
	router.PUT("/customers", updateCustomerById)
	router.PUT("/sales", updateSaleById)

	router.Run(":9090")
}

// TOTAL

func getTotalCustomers(context *gin.Context) {

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
