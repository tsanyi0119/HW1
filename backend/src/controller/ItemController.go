package controller

import (
	"github.com/gin-gonic/gin"
	"go_gin_example/database"
)

type Item struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	OrderID   uint `json:"order_id"`
	ProductID uint `json:"product_id"`
	IsShipped bool `json:"is_shipped"`
}

// 查詢特定訂單的明細
func GetOrderDetails(c *gin.Context) {
	db := database.ConnectDB()
	defer database.CloseDB(db)

	orderID := c.Param("order_id")

	var items []Item

	if err := db.Where("order_id = ?", orderID).Find(&items).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "查詢訂單的明細失敗: " + err.Error(),
		})
		return
	}

	c.JSON(200, items)
}
