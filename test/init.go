package test

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB //mysql

// NewDB 初始化数据库(用于测试)
func NewDB(dsn string) {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalln("[DB ERROR] : ", err)
	}
	err = db.AutoMigrate()
	if err != nil {
		log.Fatalln("[DB ERROR] : ", err)
	}
	DB = db
}
