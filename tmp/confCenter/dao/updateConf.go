package dao

import (
	"confCenter/models"
	"confCenter/common"
	"fmt"
)

// GetUpdateInfo 根据建筑ID和建筑等级获取升级所需信息
func GetUpdateInfo(buildingID, leve uint) ([]models.UpdateBuildingConf, error) {
	DB := common.GetDB()

	updateInfo := make([]models.UpdateBuildingConf, 1)
	status := DB.Find(&updateInfo, "building_id = ? and leve = ?", buildingID, leve)
	if status.Error != nil {
		if status.RecordNotFound() { // 判断错误是否是因为没有匹配到数据
			fmt.Printf("未查询到相关数据\n")
		} else {
			fmt.Printf("查询语句出错, err: %v\n", status.Error)
		}
	}

	return updateInfo, status.Error
}
