// gorm是go语言编写的orm框架
// 优点：提高开发效率
// 缺点：牺牲执行性能、牺牲灵活性

package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// 定义模型
// **********************************************************
// 简介：
// 在代码中定义Models与数据库中的数据表进行映射
// 在GORM中，模型通常是正常定义的结构体、基本go类型或它们的指针
// 同时也支持sql.Scanner、driver.Valuer、interface
// gorm.Model
// GORM内置了gorm.Model结构体
// 其中包含ID、CreatedAt、UpdateAt、DeletedAt四个字段
// 可将它们嵌入到自己的模型中，也可以完全自定义模型
// (当删除一个字段时，会给DeletedAt赋为当前时间，而不是正真删除)
// 主键
// 若有嵌入gorm.Model,则其中的ID字段为主键
// 或者自建的名为ID的主键，会被当作组件
// 若想自定义某字段为主键可用tag指明 `gorm:"primary_key"`
// 数据库表的列名：
// 若模型字段名为单个单词，则列名为首字母小写后的单词
// 若模型字段名为驼峰式的多个单词，列名为下划线连接的多个单词
// 也可使用tag指定列明 `gorm:"colum:列名"`

type User struct {
	gorm.Model
	Name string
	Age  int64 `gorm:"default:'5'"` // 设置默认值为5
	/*
		Birthday		*time.Time
		Email			string	`gorm:"type:varchar(100);unique_index"`
		Role			string	`gorm:"size:255"`	// 设置字段大小为255
		MemberNumber	*string	`gorm:"unique;not null"`	// 会员号唯一且不为空
		Num				int	`gorm:"AUTO_INCREMENT"`	// 设置Num为自增类型
		Address			string	`gorm:"index:addr"`	// 创建名为addr的索引
		IgnoreMe		int	`gorm:"-"`	// 忽略本字段(不会出现在数据库里)
	*/
}

// 默认生成的数据库表的名字为模型名的复数形式,驼峰命名将被该为下划线
// 也可为该模型绑定TableName方法来设定数据库表的名字
func (User) TableName() string {
	return "my_users"
}

// **********************************************************

// 连接数据库
// **********************************************************
var DB *gorm.DB

func initDB(user, passworld, addr, port, dbname string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, passworld, addr, port, dbname)
	var err error
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
}

// **********************************************************

// 创建表
// **********************************************************
func creatTable() {
	DB.SingularTable(true)  // 禁用默认的让表名为模型名的复数形式(需要在AutoMigrate之前)
	DB.AutoMigrate(&User{}) // 创建表并自动迁移

	// u1 := UserInfo{ID: 1, Name: "aa", Gender: "man", Hobby: "run"}
	// DB.Create(u1)	// 插入数据
}

// **********************************************************

// 插入数据
// **********************************************************
func insert(name string, age int64) {
	u := User{Name: name, Age: age} // 实例化一个User结构体
	if DB.NewRecord(&u) {           // 判断是否为新数据(即数据库是否已经存有该数据)，若还无该数据，则返回True
		DB.Create(&u) // 在数据库里存储一条记录
		fmt.Println("存入数据")
	}

}

// 一般查询
// **********************************************************
func query() {
	// 根据主键查询第一条记录
	u1 := new(User)
	DB.Where("deleted_at = ?", "nil").First(u1)
	fmt.Printf("第一条记录为: name:%s, age:%d\n", u1.Name, u1.Age)
	// 根据主键查询最后一条记录
	u2 := new(User)
	DB.Last(u2)
	fmt.Printf("最后一条记录为: name:%s, age:%d\n", u2.Name, u2.Age)
	// 随机获取一条记录
	u3 := new(User)
	DB.Take(u3)
	fmt.Printf("随机获取一条记录为: name:%s, age:%d\n", u3.Name, u3.Age)
	// 查询所有记录
	users := make([]User, 2)
	DB.Find(&users)
	fmt.Println("列举所有记录")
	for _, u := range users {
		fmt.Printf("name:%s, age:%d\n", u.Name, u.Age)
	}
}

// **********************************************************

func main() {
	initDB("root", "123456", "zy.server", "3306", "go_test")
	defer DB.Close()
	creatTable()
	// insert("aa", 1)
	// insert("bb", 2)
	// insert("cc", 3)
	// insert("dd", 4)
	query()
}
