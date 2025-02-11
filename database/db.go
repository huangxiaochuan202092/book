package database

import (
	"book-management-api/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	//从环境变量中读取数据库主机地址
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		log.Println("DB_HOST environment variable is not set, using default 127.0.0.1")
		dbHost = "127.0.0.1"
	}
	// 构建数据库连接字符串
	dsn := fmt.Sprintf("root:123456@tcp(%s:3306)/user_db?charset=utf8mb4&parseTime=True&loc=Local", dbHost)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("数据库连接成功")
	}

	DB = db
	DB.AutoMigrate(&models.Book{}) //自动迁移

}
