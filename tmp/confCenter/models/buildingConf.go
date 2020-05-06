package models

// BuildingType BuildingConf.Type的字段映射
var BuildingType = []string{
	"功能建筑",
	"非功能建筑",
	"地形装饰物",
}

// BuildingChildType BuildingConf.ChildType的字段映射
var BuildingChildType = []string{
	"城堡",
	"可建造功能建筑",
	"农场类",
	"魔法类",
	"日常类",
	"动物类",
	"树木",
	"花朵",
	"设施",
}

// BuildingConf 建筑配置表
type BuildingConf struct {
	ID         uint   `gorm:"primary_key"`
	BuildingID uint   `gorm:"not null;unique"` // 编号、建筑ID
	Name       string // 中文名
	Type       uint   // 主类型
	ChildType  uint   // 细分类型
}
