package apis

import (
	"gin_jwt/tools/myjwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	claims := c.MustGet("claims").(*myjwt.CustomClaims)
	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "token有效",
			"data":   claims,
		})
	}
}

// https://www.jianshu.com/p/1f9915818992
