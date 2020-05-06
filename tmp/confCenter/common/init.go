package common

import (
	// "confCenter/common"
	"flag"
	"fmt"
)

func init() {
	var confFile string
	// 初始化配置
	flag.StringVar(&confFile, "c", "config.json", "conf file")
	flag.Parse()
	if err := InitConfig(confFile); err != nil {
		panic(err)
	}

	// 初始化数据库
	mysqlCof := GetMysqlConfig()
	fmt.Println(mysqlCof.User, mysqlCof.Password, mysqlCof.Addr, mysqlCof.Db)
	db := ininDb(mysqlCof.User, mysqlCof.Password, mysqlCof.Addr, mysqlCof.Db)
	createTable(db)
}
