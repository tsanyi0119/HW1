package main

import (
	"fmt"
	"go_gin_example/controller"
	"go_gin_example/database"
	"go_gin_example/envconfig"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	// 創建資料表
	// createTable()
	// 刪除資料表
	// dropFakeData()
	// 創建假資料
	// createFakeData()

	server := gin.Default()

	// 查詢商品清單
	server.GET("/products", controller.GetProducts)
	// 查詢客戶清單
	server.GET("/customers", controller.GetCustomers)
	// 查詢客戶的訂單清單
	server.GET("/customers/:customer_id/orders", controller.GetCustomerOrders)
	// 查詢訂單的明細
	server.GET("/orders/:order_id/items", controller.GetOrderDetails)

	if err := server.Run(":" + envconfig.GetEnv("PORT")); err != nil {
		log.Fatalln(err.Error())
		return
	}

}

// 創建資料表
func createTable() {

	var db = database.ConnectDB()
	defer database.CloseDB(db)

	// 建立資料表
	db.AutoMigrate(&controller.Product{}, &controller.Category{}, &controller.Customer{}, &controller.Order{}, &controller.Item{}) // 使用控制器中的結構

	// 關閉資料庫連線
	sqlDB, err := db.DB()
	if err != nil {
		panic("無法關閉資料庫連線")
	}
	sqlDB.Close()
}

// 創建假資料
func createFakeData() {

	var db = database.ConnectDB()

	// 創建10筆商品資料
	for i := 1; i <= 10; i++ {
		product := controller.Product{
			Name:       fmt.Sprintf("Product %d", i),
			Price:      float64(i * 10),
			CategoryID: uint(i % 3), // 假設有3個商品類別
		}
		db.Create(&product)
	}

	// 創建10筆類別資料
	for i := 1; i <= 10; i++ {
		category := controller.Category{
			Name: fmt.Sprintf("Category %d", i),
		}
		db.Create(&category)
	}

	// 創建10筆客戶資料
	for i := 1; i <= 10; i++ {
		customer := controller.Customer{
			Name: fmt.Sprintf("Customer %d", i),
		}
		db.Create(&customer)
	}

	// 創建10筆訂單資料
	for i := 1; i <= 10; i++ {
		order := controller.Order{
			CustomerID: uint(i),
			IsPaid:     i%2 == 0, // 偶數訂單為已付款
		}
		db.Create(&order)
	}

	// 創建10筆訂單明細資料
	for i := 1; i <= 10; i++ {
		item := controller.Item{
			OrderID:   uint(i),
			ProductID: uint(i % 5), // 假設有5個商品
			IsShipped: i%2 == 0,    // 偶數明細為已出貨
		}
		db.Create(&item)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("無法關閉資料庫連線")
	}
	sqlDB.Close()
}

// 刪除資料表
func dropFakeData() {

	var db = database.ConnectDB()

	migrator := db.Migrator()

	// 刪除資料表 (Product, Category, Customer, Order, Item)
	migrator.DropTable(&controller.Product{}, &controller.Category{}, &controller.Customer{}, &controller.Order{}, &controller.Item{})

	sqlDB, err := db.DB()
	if err != nil {
		panic("無法關閉資料庫連線")
	}
	sqlDB.Close()
}
