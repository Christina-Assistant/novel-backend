package middleware

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/novel_db?charset=utf8mb4&parseTime=True&loc=Local" // 替换为你的数据库信息
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}
	DB = database
}
