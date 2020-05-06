package models

// PositionConf 可装饰位置配置表
type PositionConf struct {
	ID           uint   `gorm:"primary_key"`
	PositionName string //	部位描述
	ConfType     uint   // 可配置类型
}
