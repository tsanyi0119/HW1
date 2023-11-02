package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"
	"fmt"
	"go_gin_example/envconfig"
)

type Product struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryID uint    `json:"category_id"`
}

// 查詢商品清單
func GetProducts(c *gin.Context) {
	db := connectDB()
	defer closeDB(db)

	var products []Product

	if err := db.Find(&products).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "查詢商品清單失敗: " + err.Error(),
		})
		return
	}

	c.JSON(200, products)
}

func closeDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to close database")
	}
	sqlDB.Close()
}

func connectDB() *gorm.DB {
	var dsn string = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Taipei",
		envconfig.GetEnv("DB_HOST"), envconfig.GetEnv("DB_USER"), envconfig.GetEnv("DB_PASSWORD"), envconfig.GetEnv("DB_NAME"), envconfig.GetEnv("DB_PORT"), envconfig.GetEnv("DB_WITH_SSL"))

	var db *gorm.DB
	var err error
	if envconfig.GetEnv("DB_Source") == "gcp" {
		db, err = gorm.Open(postgres.New(postgres.Config{
			DriverName: "cloudsqlpostgres",
			DSN:        dsn,
		}), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	} else {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		panic("failed to connect database")
	}
	return db
}
