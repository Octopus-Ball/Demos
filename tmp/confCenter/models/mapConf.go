package models

// FloorType MapConf.Type的字段映射
var FloorType = []string{
	"城堡",
	"功能建筑",
	"非功能建筑",
	"地形装饰物",
}

// MapConf 地图配置表
type MapConf struct {
	ID          uint `gorm:"primary_key"`
	Anchor      uint `gorm:"not null;unique"` // 编号、锚点(非空且唯一)
	Type        uint // 地块类型
	UnLockLevel uint // 地块解锁对应城堡等级
}
