package main

import (
	"github.com/dignelidxdx/HackthonGo/config"
	"github.com/dignelidxdx/HackthonGo/pkg/db"
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

	dataSource, port := config.BuildDataSource()
	//Employee
	db, _ := db.Gorm(dataSource)
	repo, _ := persistence.NewEmployeeRepository()
	repoGorm := persistence.NewGormRepository(db)
	circuitBreaker := owner.NewCircuitBreaker()
	newClient := owner.NewClient("", circuitBreaker)
	service := service.NewEmployeeService(repo, repoGorm, newClient, circuitBreaker)
	controller := handlerE.NewEmployee(service)

	router.GET("/employees", controller.GetAll())
	router.GET("/employees/:id", controller.GetOne())
	router.POST("/employees", controller.CreateOne())
	router.DELETE("/employees/:id", controller.DeleteOne())

	router.Run(port)
}
