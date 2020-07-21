package models

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DB 数据库连接池
var DB *gorm.DB

func init(){
	var err error
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
		"root",
		"123456",
		"127.0.0.1",
		"gin_jwt",
	)
	if DB, err = gorm.Open("mysql", dsn); err != nil {
		log.Println("连接数据库失败:", err)
		panic(err)
	}

	DB.DB().SetMaxOpenConns(15)
	DB.DB().SetMaxIdleConns(15)
	DB.LogMode(true)

	DB.AutoMigrate(
		&User{},
	)
}