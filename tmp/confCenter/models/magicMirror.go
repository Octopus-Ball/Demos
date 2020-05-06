package models

// Gender MagicMirror.Gender的字段映射
var Gender = []string{
	"男",
	"女",
	"通用",
}

// MagicMirror 魔镜
type MagicMirror struct {
	ID            uint   `gorm:"primary_key"`
	MagicMirrorID uint   // 编号
	Type          uint   // 类型
	Gender        uint   //	性别专属
	MaterialID    uint   //	素材编号
	Describe      string //	描述
	Coin          uint   // 价格_金币
	Diamond       uint   //	价格_钻石
}
