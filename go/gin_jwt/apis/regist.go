package apis

import (
	"net/http"

	"gin_jwt/models"

	"github.com/gin-gonic/gin"
)

type RegistRequestData struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Pwd   string `json:"pwd"`
}

func Register(c *gin.Context) {
	var data = new(RegistRequestData)
	if err := c.BindJSON(data); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "参数解析错误",
		})
		return
	}

	exist, err := models.AddUser(data.Name, data.Phone, 0)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -2,
			"msg":    err,
		})
		return
	}

	if exist {
		c.JSON(http.StatusOK, gin.H{
			"status": -3,
			"msg":    "已存在该用户",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "注册成功",
	})
}
