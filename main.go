package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rittbpt/assessment-tax/controller"
	"github.com/rittbpt/assessment-tax/repository"
	"github.com/rittbpt/assessment-tax/service"
	"github.com/rittbpt/assessment-tax/connection"
	"github.com/rittbpt/assessment-tax/route"
	"os"
	"os/signal"
	"time"
	"context"
	"net/http"
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

	go func() {
		if err := e.Start(":" + os.Getenv("PORT")); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatalf("Failed to start server: %s", err)
		}
	}()



	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatalf("Server shutdown error: %s", err)
	}

	e.Logger.Print("Server gracefully stopped")
}
