package storage

import (
<<<<<<< HEAD
	"github.com/Levantate-Labs/sainterview-backend/auth-service/config"
	"github.com/Levantate-Labs/sainterview-backend/auth-service/models"
=======
	"github.com/AlfrinP/point_calculator/config"
	"github.com/AlfrinP/point_calculator/models"
>>>>>>> 52a2cfba8417f30f47f3a85feb3c92850e82f352
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var defaultDB *gorm.DB

func ConnectDB(config *config.Config) {
	db, err := gorm.Open(postgres.Open(config.DBUrl), &gorm.Config{})
	if err != nil {
		panic("DB Connection failed")
	}

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	db.Logger = logger.Default.LogMode(logger.Info)

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic("DB Migrations Failed")
	}

	defaultDB = db
}

func GetDB() *gorm.DB {
	return defaultDB
}
