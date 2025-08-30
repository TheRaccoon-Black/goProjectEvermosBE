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
	alamatRepo := repository.NewAlamatRepository(db)

	authUsecase := usecase.NewAuthUsecase(userRepo, tokoRepo)
	userUsecase := usecase.NewUserUsecase(userRepo)
	alamatUsecase := usecase.NewAlamatUsecase(alamatRepo)
	tokoUsecase := usecase.NewTokoUsecase(tokoRepo)

	authHandler := delivery.NewAuthHandler(authUsecase) 
	userHandler := delivery.NewUserHandler(userUsecase)
	alamatHandler := delivery.NewAlamatHandler(alamatUsecase)
	tokoHandler := delivery.NewTokoHandler(tokoUsecase)
	
	
	app := fiber.New() 

	app.Static("/public", "./public")
	
	api := app.Group("/api/v1")
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	api.Get("/toko", tokoHandler.GetAllToko)
	api.Get("/toko/:id_toko", tokoHandler.GetTokoByID)


	user := api.Group("/user")
	user.Use(delivery.AuthMiddleware()) 
	user.Get("/", userHandler.GetProfile) 
	user.Put("/", userHandler.UpdateProfile)
	user.Post("/alamat", alamatHandler.CreateAlamat)
	user.Get("/alamat", alamatHandler.GetAllAlamat)      
	user.Get("/alamat/:id", alamatHandler.GetAlamatByID)
	user.Put("/alamat/:id", alamatHandler.UpdateAlamat)     
	user.Delete("/alamat/:id", alamatHandler.DeleteAlamat)

	toko := api.Group("/toko")
	toko.Use(delivery.AuthMiddleware()) 
	toko.Get("/my", tokoHandler.GetMyToko)
	toko.Put("/:id_toko", tokoHandler.UpdateToko) 

	log.Println("Server berjalan di port :8000")
	log.Fatal(app.Listen(":8000")) 
}