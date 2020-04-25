// 连接数据库，返回数据库连接

package init

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// DB 数据库连接
var DB *gorm.DB

// InitDB 连接数据库
func InitDB() {
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.database"),
		viper.GetString("db.charset"),
	)
	fmt.Println(args)
	db, err := gorm.Open(viper.GetString("db.driverName"), args)
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}
	DB = db
}

func init() {
	InitDB()
}
