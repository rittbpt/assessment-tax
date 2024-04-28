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
	// เตรียมฐานข้อมูล
	connection.PrepareDB()

	// สร้าง instance ของ Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// เชื่อมต่อฐานข้อมูล
	db, err := connection.ConnectDB()
	if err != nil {
		e.Logger.Fatal("Failed to connect to database:", err)
	}
	connection.PrepareDB()

	// สร้าง repository, service และ controller
	taxRepository := repository.NewTaxRepository(db)
	taxService := service.NewTaxService(taxRepository)
	taxController := controller.NewTaxController(taxService)

	// กำหนดเส้นทางของ API
	route.TaxRoutes(e, taxController)

	// เริ่มต้นเซิร์ฟเวอร์
	e.Logger.Fatal(e.Start(":8080"))
}
