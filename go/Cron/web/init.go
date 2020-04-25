package web

import "github.com/gin-gonic/gin"

// G 实例化后的引擎
var G *gin.Engine

// init 初始化web
func init() {
	G = gin.Default()		// 初始化gin引擎
	G = InitRouter(G)		// 初始化路由
}

