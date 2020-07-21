package models

import "github.com/jinzhu/gorm"

// User 用户信息
type User struct {
	gorm.Model
	Name		string
	Phone		string
	Type		int
}

// GetUser 获取用户信息
func GetUser(phone string) (user *User, err error) {
	user = &User{
		Phone: phone,
	}
	DB = DB.Where(user).Find(user)
	err = DB.Error

	return
}

// AddUser 添加用户信息
func AddUser(name, phone string, utype int) (exist bool, err error) {
	user := &User{
		Phone: phone,
	}

	if DB.First(user).RecordNotFound(){
		user.Name = name
		user.Type = utype
		DB = DB.Create(user)
		exist = false
		return
	}

	exist = true
	return
}