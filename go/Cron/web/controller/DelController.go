package web

import (
	"Cron/init"
	"Cron/model"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// DelController 删除定时任务
func DelController(g *gin.Context) {
	db := init.DB
	j := new(model.Job)
	// 参数解析绑定
	if err := g.ShouldBind(j); err != nil {
		log.Printf("参数解析错误:%#\n", err)
	}
	fmt.Println("暂时省略参数校验")
	db.Delete(j)
	g.JSON(200, gin.H{
		"code": 200,
		"data": "任务删除成功",
	})
}