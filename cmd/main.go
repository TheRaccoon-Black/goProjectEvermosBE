package main

import (
	"fmt"
	"log"

	"goProjectEvermos/config"
	"goProjectEvermos/internal/domain" 
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
    "goProjectEvermos/internal/repository" 

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
    &domain.User{},
    &domain.Toko{},
    &domain.AlamatKirim{},
    &domain.Category{},
    &domain.Product{},
    &domain.ProductPhoto{},
    &domain.Transaction{},
    &domain.TransactionDetail{},
    &domain.LogProduct{},
)

	if err != nil {
		log.Fatal("Gagal melakukan migrasi database:", err)
	}

	log.Println("Migrasi database berhasil!")
}