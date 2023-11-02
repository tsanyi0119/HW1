package controller

import (
	"github.com/gin-gonic/gin"
)

type Customer struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

// 查詢客戶清單
func GetCustomers(c *gin.Context) {
	db := connectDB()
	defer closeDB(db)

	var customers []Customer

	if err := db.Find(&customers).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "查詢客戶清單失敗: " + err.Error(),
		})
		return
	}

	c.JSON(200, customers)
}

