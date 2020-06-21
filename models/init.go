package models

import (
	"log"

	// mysql 数据库驱动
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DB 数据库
var DB *gorm.DB

// Setup 初始化数据库
func Setup() {
	var err error

	DB, err = gorm.Open("mysql", "root:keyboard@(keyboard-man.site)/iv-go?charset=utf8&parseTime=True&loc=Local&timeout=300ms")
	if err != nil {
		log.Printf("Mysql connect error %v", err)
	}

	if DB.Error != nil {
		log.Printf("Database error %v", DB.Error)
	}

	AutoMigrateAll()
}

// AutoMigrateAll 自动迁移数据库
func AutoMigrateAll() {
	DB.AutoMigrate(&User{})
}
