package main

import (
	"fmt"
	"log"

	"goProjectEvermos/config"
	delivery "goProjectEvermos/internal/delivery/http" 
	"goProjectEvermos/internal/domain"
	"goProjectEvermos/internal/repository"
	"goProjectEvermos/internal/usecase"

	"github.com/gofiber/fiber/v2" 
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}
	log.Println("Koneksi database berhasil!")

	err = db.AutoMigrate(
		&domain.User{}, &domain.Toko{}, &domain.AlamatKirim{}, &domain.Category{},
		&domain.Product{}, &domain.ProductPhoto{}, &domain.Transaction{},
		&domain.TransactionDetail{}, &domain.LogProduct{},
	)
	if err != nil {
		log.Fatal("Gagal melakukan migrasi database:", err)
	}
	log.Println("Migrasi database berhasil!")

	userRepo := repository.NewUserRepository(db)
	tokoRepo := repository.NewTokoRepository(db)

	authUsecase := usecase.NewAuthUsecase(userRepo, tokoRepo)

	authHandler := delivery.NewAuthHandler(authUsecase) 

	app := fiber.New() 

	api := app.Group("/api/v1")
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	log.Println("Server berjalan di port :8000")
	log.Fatal(app.Listen(":8000")) 
}