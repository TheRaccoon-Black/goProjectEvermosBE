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
 	categoryRepo := repository.NewCategoryRepository(db)
 	productRepo := repository.NewProductRepository(db)
	trxRepo := repository.NewTransactionRepository(db)
	

	authUsecase := usecase.NewAuthUsecase(userRepo, tokoRepo)
 	userUsecase := usecase.NewUserUsecase(userRepo)
 	alamatUsecase := usecase.NewAlamatUsecase(alamatRepo)
 	tokoUsecase := usecase.NewTokoUsecase(tokoRepo)
 	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)
 	productUsecase := usecase.NewProductUsecase(productRepo, tokoRepo)
	trxUsecase := usecase.NewTransactionUsecase(trxRepo, alamatRepo, productRepo)
	provCityUsecase := usecase.NewProvCityUsecase() 

	authHandler := delivery.NewAuthHandler(authUsecase)
 	userHandler := delivery.NewUserHandler(userUsecase)
 	alamatHandler := delivery.NewAlamatHandler(alamatUsecase)
 	tokoHandler := delivery.NewTokoHandler(tokoUsecase)
 	categoryHandler := delivery.NewCategoryHandler(categoryUsecase)
 	productHandler := delivery.NewProductHandler(productUsecase)
	trxHandler := delivery.NewTransactionHandler(trxUsecase)
	provCityHandler := delivery.NewProvCityHandler(provCityUsecase)

	app := fiber.New()
	app.Static("/public", "./public")
	api := app.Group("/api/v1")


	// Rute Auth (Publik)
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	provCity := api.Group("/provcity")
	provCity.Get("/listprovincies", provCityHandler.GetAllProvinces)
	provCity.Get("/listcities/:prov_id", provCityHandler.GetCitiesByProvinceID)
	provCity.Get("/detailprovince/:prov_id", provCityHandler.GetProvinceByID)
	provCity.Get("/detailcity/:city_id", provCityHandler.GetCityByID)

	// Rute User (Terlindungi)
	user := api.Group("/user", delivery.AuthMiddleware()) 
	user.Get("/", userHandler.GetProfile)
	user.Put("/", userHandler.UpdateProfile)
	user.Post("/alamat", alamatHandler.CreateAlamat)
	user.Get("/alamat", alamatHandler.GetAllAlamat)
	user.Get("/alamat/:id", alamatHandler.GetAlamatByID)
	user.Put("/alamat/:id", alamatHandler.UpdateAlamat)
	user.Delete("/alamat/:id", alamatHandler.DeleteAlamat)

	// Rute Toko (Campuran Publik & Terlindungi)
	toko := api.Group("/toko")
	// Rute Terlindungi (Spesifik) harus di atas
	toko.Get("/my", delivery.AuthMiddleware(), tokoHandler.GetMyToko)
	toko.Put("/:id_toko", delivery.AuthMiddleware(), tokoHandler.UpdateToko)
	// Rute Publik (Umum) di bawah
	toko.Get("/", tokoHandler.GetAllToko)
	toko.Get("/:id_toko", tokoHandler.GetTokoByID)

	// Rute Produk (Campuran Publik & Terlindungi)
	product := api.Group("/product")
	// Rute Terlindungi
	product.Post("/", delivery.AuthMiddleware(), productHandler.CreateProduct)
	product.Put("/:id", delivery.AuthMiddleware(), productHandler.UpdateProduct)
	product.Delete("/:id", delivery.AuthMiddleware(), productHandler.DeleteProduct)
	// Rute Publik
	product.Get("/", productHandler.GetAllProducts)
	product.Get("/:id", productHandler.GetProductByID)

	// Rute Kategori (Hanya Admin)
	category := api.Group("/category", delivery.AuthMiddleware(), delivery.AdminMiddleware()) // Dua middleware
	category.Post("/", categoryHandler.CreateCategory)
	category.Get("/", categoryHandler.GetAllCategories)
	category.Get("/:id", categoryHandler.GetCategoryByID)
	category.Put("/:id", categoryHandler.UpdateCategory)
	category.Delete("/:id", categoryHandler.DeleteCategory)


	trx := api.Group("/trx", delivery.AuthMiddleware())
	trx.Post("/", trxHandler.CreateTransaction)
	trx.Get("/", trxHandler.GetAllTransactions)
	trx.Get("/:id", trxHandler.GetTransactionByID)


	log.Println("Server berjalan di port :8000")
	log.Fatal(app.Listen(":8000"))
}