package common

import (
	"confCenter/models"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DB 数据库连接池
var DB *gorm.DB

// IninDb 连接数据库
func ininDb(user, passworld, addr, dbname string) (*gorm.DB){
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", user, passworld, addr, dbname)
	var err error
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	log.Printf("数据库%s连接成功\n", addr)
	return DB
}

// createTable 迁移数据库表
func createTable(DB *gorm.DB) {
	DB.SingularTable(true) // 禁用默认的让表名为模型名的复数形式
	DB.AutoMigrate(        // 若结构体变化则自动进行迁移
		&models.BuildingConf{},
		&models.MapConf{},
		&models.UpdateBuildingConf{},
		&models.UserMapInfo{},
		&models.PositionConf{},
		&models.MagicMirror{},
		&models.Shop{},
	)
}

// GetDB 获取数据库连接池
func GetDB() *gorm.DB {
	return DB
}
