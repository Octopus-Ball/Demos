package models

// UpdateBuildingConf 建筑升级配置表
type UpdateBuildingConf struct {
	ID                   uint   `gorm:"primary_key"`
	UpdateBuildingConfID uint   // 编号
	BuildingID           uint   // 建筑编号
	Leve                 uint   // 级别
	Coin                 uint   // 金币价格
	Diamond              uint   //	钻石价格
	Star                 uint   //	星星
	ImgURL               string //	图片素材
	NeedType             uint   //	需求条件类型
	NeedSum              uint   // 需求条件数量
	NeedCont             uint   //	需求条件内容
}
