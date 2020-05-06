package models

// LockType UserMapInfo.UnLock的字段映射
var LockType = []string{
	"解锁",
	"锁定",
}

// UserMapInfo 用户个人地图信息表
type UserMapInfo struct {
	ID           uint `gorm:"primary_key"`
	UserID       uint // 用户ID
	FloorID      uint // 地块ID
	Lock         uint // 锁定/解锁
	Building     uint // 地上建筑
	BuildingLeve uint // 地上建筑级别
}
