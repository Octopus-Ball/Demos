package web

import "github.com/gin-gonic/gin"

// InitRouter 初始化路由
func InitRouter(g *gin.Engine) *gin.Engine {
	g.POST("/add", AddController)
	g.POST("/del", DelController)
	g.GET("/show", ShowController)
	return g
}