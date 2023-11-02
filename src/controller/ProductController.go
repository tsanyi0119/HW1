package controller

import (
	"github.com/gin-gonic/gin"
	"go_gin_example/database"
)

type Product struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryID uint    `json:"category_id"`
}

// 查詢商品清單
func GetProducts(c *gin.Context) {
	db := database.ConnectDB()
	defer database.CloseDB(db)

	var products []Product

	if err := db.Find(&products).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "查詢商品清單失敗: " + err.Error(),
		})
		return
	}
	
	c.JSON(200, products)
}