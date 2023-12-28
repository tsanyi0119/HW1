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

// 新增商品
func CreateProduct(c *gin.Context) {
	db := database.ConnectDB()
	defer database.CloseDB(db)

	var inputProduct Product

	if err := c.ShouldBindJSON(&inputProduct); err != nil {
		c.JSON(400, gin.H{
			"message": "請提供有效的商品數據",
		})
		return
	}

	if err := db.Create(&inputProduct).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "新增商品失敗: " + err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "商品新增成功",
		"data":    inputProduct,
	})
}

// 刪除商品
func DeleteProduct(c *gin.Context) {
	db := database.ConnectDB()
	defer database.CloseDB(db)

	// 從 URL 中獲取商品 ID
	productID := c.Param("id")

	// 檢查商品是否存在
	var existingProduct Product
	if err := db.First(&existingProduct, productID).Error; err != nil {
		c.JSON(404, gin.H{
			"message": "商品不存在",
		})
		return
	}

	// 刪除商品
	if err := db.Delete(&existingProduct).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "刪除商品失敗: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "商品刪除成功",
	})
}

// 修改商品價格
func UpdateProductPrice(c *gin.Context) {
	db := database.ConnectDB()
	defer database.CloseDB(db)

	// 從 URL 中獲取商品 ID
	productID := c.Param("id")

	// 檢查商品是否存在
	var existingProduct Product
	if err := db.First(&existingProduct, productID).Error; err != nil {
		c.JSON(404, gin.H{
			"message": "商品不存在",
		})
		return
	}

	// 解析請求中的 JSON 數據到 existingProduct
	if err := c.ShouldBindJSON(&existingProduct); err != nil {
		c.JSON(400, gin.H{
			"message": "請提供有效的商品數據",
		})
		return
	}

	// 更新商品價格
	if err := db.Save(&existingProduct).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "更新商品價格失敗: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "商品價格更新成功",
		"data":    existingProduct,
	})
}