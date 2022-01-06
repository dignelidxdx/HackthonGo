package main

func main() {

	router := gin.Default()

	router.GET("/invoices/total", getTotalFromInvoice)

	router.Run(":9090")
}

func getTotalFromInvoice(context *gin.Context) {

}
