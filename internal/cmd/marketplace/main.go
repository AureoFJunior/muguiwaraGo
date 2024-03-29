package main

import (
	"log"

	"github.com/AureoFJunior/muguiwaraGo/internal/application/services"
	postgres2 "github.com/AureoFJunior/muguiwaraGo/internal/infrastructure/db/postgres"
	"github.com/AureoFJunior/muguiwaraGo/internal/interface/api/rest"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	port := ":8080"

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	//gormDB.AutoMigrate()

	productRepo := postgres2.NewGormProductRepository(gormDB)
	sellerRepo := postgres2.NewGormSellerRepository(gormDB)

	productService := services.NewProductService(productRepo, sellerRepo)
	sellerService := services.NewSellerService(sellerRepo)

	e := echo.New()
	rest.NewProductController(e, productService)
	rest.NewSellerController(e, sellerService)

	if err := e.Start(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
