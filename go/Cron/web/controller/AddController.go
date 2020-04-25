package web

import (
	"Cron/init"
	"Cron/model"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// AddController 注册定时任务
func AddController(g *gin.Context) {
	db := init.DB
	j := new(model.Job)
	// 参数解析绑定
	if err := g.ShouldBind(j); err != nil {
		log.Printf("参数解析错误:%#\n", err)
	}
	fmt.Println("暂时省略参数校验")
	db.Create(j)
	g.JSON(200, gin.H{
		"code": 200,
		"data": "任务注册成功",
	})
}
