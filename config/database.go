package config

import (
	"fmt"
	"os"

	"github.com/UsefulForMe/go-ecommerce/logger"
	"github.com/UsefulForMe/go-ecommerce/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func getDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh", Cfg.PostresHost, Cfg.PostresUser, Cfg.PostresPass, Cfg.PostresDB, Cfg.PostresPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("Error when connect to database " + err.Error())
		os.Exit(1)
	}

	if err := db.Exec(`create extension if not exists "uuid-ossp"`).Error; err != nil {
		logger.Error("Failed to create 'uuid-ossp' extension, but got error " + err.Error())
	}

	logger.Info("Connect to database successfully")

	err = db.AutoMigrate(
		&models.User{},
		&models.Payment{},
		&models.UserAddress{},
		&models.Category{},
		&models.Seller{},
		&models.Product{},
		&models.Stock{},
		&models.StockReport{},
		&models.StockItem{},
		&models.Order{},
		&models.OrderItem{},
		&models.Track{})

	if err != nil {
		logger.Error("Error when auto migrate database " + err.Error())
		os.Exit(1)
	}
	logger.Info("Migrated database")

	return db
}

func InitDatabase() {
	DB = getDatabase()
}
