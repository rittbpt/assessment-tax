package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rittbpt/assessment-tax/controller"
	"github.com/rittbpt/assessment-tax/repository"
	"github.com/rittbpt/assessment-tax/service"
	"github.com/rittbpt/assessment-tax/connection"
	"github.com/rittbpt/assessment-tax/route"

)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, err := connection.ConnectDB()
    if err != nil {
        e.Logger.Fatal("Failed to connect to database:", err)
    }
	taxRepository := repository.NewTaxRepository(db)

	taxService := service.NewTaxService(taxRepository)

	taxController := controller.NewTaxController(taxService)

	route.TaxRoutes(e, taxController)

	e.Logger.Fatal(e.Start(":8080"))

}
