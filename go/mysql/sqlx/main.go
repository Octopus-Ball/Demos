// sqlx包提供了对sql包的扩展， 它无缝封装sql包
//在原sql包方法未改动的前提下，提供了更方便的方法
// 新功能包括：
//1.将查询结果扫描到结构体
//2.支持命名查询
//3.针对不同驱动程序重新绑定查询

package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // import _ 包 只是执行包的init函数，而不把整个包导入
	"github.com/jmoiron/sqlx"
)

// DB 数据库连接池
var DB *sqlx.DB

// Person 与数据库表字段对应的struct
type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Age       int
	Email     string
}

// 初始化数据库连接池
func initDB(user, passworld, addr, port, dbname string) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, passworld, addr, port, dbname)
	// Connect：打开一个数据库连接池并链接
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("Connect DB failed, err:%v\n", err)
		return
	}
	DB.SetMaxOpenConns(10) // 设置连接池里最大连接数
	DB.SetMaxIdleConns(5)  // 设置连接池里的最大空闲连接数
	return
}

// 建表
func createTable() {
	schema := `
		CREATE TABLE person (
			first_name text,
			last_name text,
			age int,
			email text
		);
	`
	rst := DB.MustExec(schema)
	rst.LastInsertId()
}

// 插入、更新、删除、建表操作(DB.MustExec)
func update() {
	sqlStr := `insert into person(first_name, last_name, age, email) value (?, ?, ?, ?);`
	rst := DB.MustExec(sqlStr, "c", "d", 2, ".com")
	id, err := rst.LastInsertId() // 获取插入数据的ID
	if err != nil {
		fmt.Printf("insert failed err:%v", err)
	}
	fmt.Printf("insert sccess id is:%d\n", id)
}

// 查询单条数据(Get)
func queryOne() {
	sqlStr := `select * from person where first_name = ?;`
	var person Person
	err := DB.Get(&person, sqlStr, "c")
	if err != nil {
		fmt.Printf("Get failed err:%v", err)
	}
	fmt.Printf("%#v\n", person)
}

// 查询多条数据(Select)
func queryMore() {
	var persons []Person
	sqlStr := `select * from person where age > ?;`
	err := DB.Select(&persons, sqlStr, 3)
	if err != nil {
		fmt.Printf("Select failed err:%v\n", err)
	}
	fmt.Printf("%#v\n", persons)
}

// 事务
func transaction() {
	sqlStr1 := `update person set age = age - 2 where first_name = ?;`
	sqlStr2 := `update person set age = age + 2 where first_name = ?;`
	tx := DB.MustBegin()
	tx.MustExec(sqlStr1, "j")
	tx.MustExec(sqlStr2, "a")
	tx.Commit()
}

func main() {
	initDB("root", "123456", "zy.server", "3306", "go_test")
	// createTable()
	// update()
	// queryOne()
	// queryMore()
	// transaction()
}
