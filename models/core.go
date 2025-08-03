package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:test123@tcp(127.0.0.1:3306)/ginormdemo?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("获取数据库底层连接失败: %v", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("数据库Ping失败: %v", err)
	}

	log.Println("数据库连接成功")
}
