package models

// Shop 服装店
type Shop struct {
	ID          uint   `gorm:"primary_key"`
	ShopID      uint   // 编号
	Type        uint   // 类型
	ChildType   uint   // 细分类型
	Gender      uint   //	性别专属
	Material1ID uint   //	素材编号
	Material2ID uint   //	素材编号
	Describe    string //	描述
	Coin        uint   // 价格_金币
	Diamond     uint   //	价格_钻石
}
