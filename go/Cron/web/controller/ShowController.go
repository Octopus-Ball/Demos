package web

import (
	"Cron/init"
	"Cron/model"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// ShowController 注册定时任务
func ShowController(g *gin.Context) {
	db := init.DB
	j := new(model.Job)
	// 参数解析绑定
	if err := g.ShouldBind(j); err != nil {
		log.Printf("参数解析错误:%#\n", err)
	}
	fmt.Println("暂时省略参数校验")
	jobs := make([]model.Job, 10)
	db.Find(&jobs)
	g.JSON(200, gin.H{
		"code": 200,
		"data": jobs,
	})
}
