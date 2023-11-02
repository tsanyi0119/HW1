package database
import(
	"fmt"
	"go_gin_example/envconfig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
func ConnectDB() *gorm.DB {
    // 設定 PostgreSQL 連線資訊
    var dsn string = fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Taipei",
        envconfig.GetEnv("DB_HOST"), envconfig.GetEnv("DB_USER"), envconfig.GetEnv("DB_PASSWORD"), envconfig.GetEnv("DB_NAME"), envconfig.GetEnv("DB_PORT"), envconfig.GetEnv("DB_WITH_SSL"))

    // 連線到 PostgreSQL 資料庫
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("無法連線到資料庫")
    }

    return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to close database")
	}
	sqlDB.Close()
}
