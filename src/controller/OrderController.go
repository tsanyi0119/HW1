package controller

import (
	"github.com/gin-gonic/gin"
	"go_gin_example/database"
)

type Order struct {
	ID         uint `gorm:"primaryKey" json:"id"`
	CustomerID uint `json:"customer_id"`
	IsPaid     bool `json:"is_paid"`
}

// 查詢客戶的訂單清單
func GetCustomerOrders(c *gin.Context) {
	db := database.ConnectDB()
	defer database.CloseDB(db)

	customerID := c.Param("customer_id")

	var orders []Order

	if err := db.Where("customer_id = ?", customerID).Find(&orders).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "查詢客戶的訂單清單失敗: " + err.Error(),
		})
		return
	}

	c.JSON(200, orders)
}
