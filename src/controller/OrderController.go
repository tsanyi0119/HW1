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

// 新增訂單
func CreateOrder(c *gin.Context) {
	db := database.ConnectDB()
	defer database.CloseDB(db)

	var inputOrder Order

	// 解析請求中的 JSON 數據到 inputOrder
	if err := c.ShouldBindJSON(&inputOrder); err != nil {
		c.JSON(400, gin.H{
			"message": "請提供有效的訂單數據",
		})
		return
	}

	// 將訂單信息插入到數據庫
	if err := db.Create(&inputOrder).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "新增訂單失敗: " + err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "訂單新增成功",
		"data":    inputOrder,
	})
}
