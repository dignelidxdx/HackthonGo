package main

import (
	circuitBreaker "github.com/dignelidxdx/HackthonGo/pkg/lib/circuitBreaker"
	handlerE "github.com/dignelidxdx/HackthonGo/wrapper/adapter/in/handler"
	"github.com/dignelidxdx/HackthonGo/wrapper/adapter/out/owner"
	"github.com/dignelidxdx/HackthonGo/wrapper/adapter/out/persistence"
	"github.com/dignelidxdx/HackthonGo/wrapper/application/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	router := gin.Default()

	//Employee

	repo, _ := persistence.NewEmployeeRepository()
	newClient := owner.NewClient("")
	circuitBreaker := circuitBreaker.NewCircuitBreaker()
	service := service.NewEmployeeService(repo, newClient)
	controller := handlerE.NewEmployee(service)

	router.GET("/employees", controller.GetAll())

	router.Run(":8080")
}
