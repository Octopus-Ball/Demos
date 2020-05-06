package dao

import (
	"confCenter/models"
	"confCenter/common"
	"fmt"
)

// GetBuildingInfoByUser 根据用户ID获取其建筑信息
func GetBuildingInfoByUser(userID uint) ([]models.UserMapInfo, error) {
	DB := common.GetDB()

	buildingInfo := make([]models.UserMapInfo, 1)
	status := DB.Find(&buildingInfo, "user_id = ?", userID)
	if status.Error != nil {
		if status.RecordNotFound() { // 判断错误是否是因为没有匹配到数据
			fmt.Printf("未查询到相关数据\n")
		} else {
			fmt.Printf("查询语句出错, err: %v\n", status.Error)
		}
	}

	return buildingInfo, status.Error
}
