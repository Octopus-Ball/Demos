package dao

import (
	"confCenter/models"
	"confCenter/common"
	"fmt"
)

// GetMapInfo 获取所有地图数据
func GetMapInfo() ([]models.MapConf, error) {
	DB := common.GetDB()

	maps := make([]models.MapConf, 1)
	status := DB.Find(&maps)
	if status.Error != nil {
		if status.RecordNotFound() { // 判断错误是否是因为没有匹配到数据
			fmt.Printf("未查询到相关数据\n")
		} else {
			fmt.Printf("查询语句出错, err: %v\n", status.Error)
		}
	}

	return maps, status.Error
}
