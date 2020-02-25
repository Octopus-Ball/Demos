// 使用sql模块用go语言操作mysql

package main

import (
	// sql包提供了保证SQL或类SQL数据库的泛用接口(使用sql包时必须注入数据库驱动)
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // import _ 包 只是执行包的init函数，而不把整个包导入
)

// DB 数据库连接池
var DB *sql.DB

// USER 对应user表的字段
type USER struct {
	id   int
	name string
	age  int
}

// 初始化数据库连接池
func initDb(addr, port, user, passworld, dbname string) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, passworld, addr, port, dbname)
	// db是数据库操作句柄代表一个数据库连接池，可安全的被多个go程同时使用
	DB, err = sql.Open("mysql", dsn) // Open方法会校验dsn格式正确与否
	if err != nil {
		return
	}
	err = DB.Ping() // Ping方法会校验用户名和密码
	if err != nil {
		return
	}
	fmt.Println("数据库连接成功")
	DB.SetMaxOpenConns(10) // 设置连接池里最大连接数
	DB.SetMaxIdleConns(5)  // 设置连接池里的最大空闲连接数
	return
}

// 查询单条记录(DB.QueryRow)
func queryOne() {
	var user USER
	sqlStr := `select id, name, age from user where id=?;`
	rowObj := DB.QueryRow(sqlStr, 5)            // 查询行
	rowObj.Scan(&user.id, &user.name, user.age) // 解析行，并将连接归还到连接池(查询完一定要调用Scan来归还连接)
	fmt.Printf("u:%#v\n", user)
}

// 查询多条记录(DB.Query)
func queryMore() {
	sqlStr := `select id, name, age from user where id > ?;`
	rows, err := DB.Query(sqlStr, 3)
	if err != nil {
		fmt.Printf("Query failed, err:%v\n", err)
		return
	}
	defer rows.Close() // rows用完要关闭，以释放数据库连接
	for rows.Next() {
		var user USER
		err := rows.Scan(&user.id, &user.name, &user.age)
		if err != nil {
			fmt.Printf("Scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("user:%#v\n", user)
	}
}

// 插入数据(DB.Exec)
func insert() {
	sqlStr := `insert into user(name, age) value("孙", 10)`
	rst, err := DB.Exec(sqlStr)
	if err != nil {
		fmt.Printf("insert failed, err:%v", err)
		return
	}
	id, err := rst.LastInsertId() // 获取插入数据的ID
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}
	fmt.Println("id:", id)
}

// 更新或删除数据(DB.Exec)
func updateRow() {
	sqlStr1 := `update user set age = ? where id = ?`
	// sqlStr2 := `delete from user where id = ?`
	rst, err := DB.Exec(sqlStr1, 9000, 2)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := rst.RowsAffected() // 获取更新行数
	if err != nil {
		fmt.Printf("get id failed, err:%v\n", err)
	}
	fmt.Printf("更新了%d行数据", n)
}

// mysql预处理
func prepare() {
	// 普通SQL语句执行流程
	//1.客户端对SQL语句占位符替换得到完整SQL语句
	//2.客户端发送完整SQL语句到mysql服务端
	//3.mysql服务端执行完整sql语句并返回结果给客户端
	// 预处理执行流程
	//1.把SQL语句分为命令部分和数据部分
	//2.先把命令部分发送给服务端，mysql服务端进行SQL预处理
	//3.把数据部分发送给mysql服务端，服务端对SQL语句进行占位符替换
	//4.服务端执行完整SQL语句，并将结果返回给客户端
	// 预处理的优点
	//1.优化服务端重复执行SQL的方法，可提升服务端性能(一次编译多次执行)
	//2.避免SQL注入问题

	// Prepare方法先将sql语句发送给Mysql服务端，返回一个准备好的状态，用于之后的查询和命令
	// 返回值可以同时执行多个查询和命令
	sqlStr := `insert into user(name, age) value(?, ?);`
	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	var m = map[string]int{
		"一": 10,
		"二": 20,
		"三": 30,
	}
	for k, v := range m {
		stmt.Exec(k, v) // 后续只需拿stmt去执行一些操作
	}
}

// 事务
func transaction() {
	// 定义：
	//事务是一个最小的不可再分的工作单元，可用来维护数据库完整性
	//保证成批SQL语句要么全部执行，要么全部不执行
	// 事务的ACID：
	//原子性：要么全执行要么全不执行
	//一致性：事务前后数据库完整性未被破坏
	//隔离性：数据库允许多个并发事务同时对进行读写修改
	//持久性：事务结束后对数据库的修改是永久的，不会丢失
	// 事务相关方法：
	//开始事务：Begin()
	//提交事务：Commit()
	//回滚事务：Rollback()

	// 开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Printf("begin failed,err:%v\n", err)
		return
	}
	// 执行多个SQL操作
	sqlStr1 := `update user set age = age - 2 where id = 1;`
	sqlStr2 := `update user set age = age + 2 where id = 2;`
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		tx.Rollback() // 回滚
		return
	}
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		tx.Rollback() // 回滚
		return
	}
	// 都执行成功则提交本次事务
	err = tx.Commit()
	if err != nil {
		tx.Rollback() // 回滚
	}
}

func main() {
	err := initDb("zy.server", "3306", "root", "123456", "go_test")
	if err != nil {
		fmt.Printf("初始化数据库连接失败,err:%v\n", err)
	}
}
