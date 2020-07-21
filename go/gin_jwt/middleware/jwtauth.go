package middleware

import (
	"net/http"

	"gin_jwt/tools/myjwt"

	"github.com/gin-gonic/gin"
)

func JWTMiddleWare(c *gin.Context) {
	j := myjwt.NewJWT()
	token := c.Request.Header.Get("token") // 获取token
	claims, err := j.ParseToken(token)     // 解析token

	switch err {
	case nil:
		c.Set("claims", claims)
	case myjwt.TokenEmpty:
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "未带token",
		})
		c.Abort()
	case myjwt.TokenExpired:
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "授权已过期",
		})
		c.Abort()
	default:
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		c.Abort()
	}

}
