package dao

import (
	"confCenter/common"
	"confCenter/models"
	"fmt"
)

// GetBuildingInfoByBuildingID 根据建筑ID获取建筑信息
func GetBuildingInfoByBuildingID(BuildingID uint) ([]models.BuildingConf, error) {
	DB := common.GetDB()

	buildingInfo := make([]models.BuildingConf, 1)
	status := DB.Find(&buildingInfo, "building_id = ?", BuildingID)
	if status.Error != nil {
		if status.RecordNotFound() { // 判断错误是否是因为没有匹配到数据
			fmt.Printf("未查询到相关数据\n")
		} else {
			fmt.Printf("查询语句出错, err: %v\n", status.Error)
		}
	}

	return buildingInfo, status.Error
}
