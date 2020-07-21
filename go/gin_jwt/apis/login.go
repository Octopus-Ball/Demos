package apis

import (
	"log"
	"net/http"
	"time"

	"gin_jwt/models"
	"gin_jwt/tools/myjwt"

	"github.com/gin-gonic/gin"
)

type LoginRequestData struct {
	Phone string `json:"phone"`
	Pwd   string `json:"pwd"`
}

// LoginResult 登录结果结构
type LoginResult struct {
	Token string `json:"token"`
	models.User
}

func Login(c *gin.Context) {
	var data = new(LoginRequestData)
	if err := c.BindJSON(data); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "参数解析错误",
		})
		return
	}

	user, err := models.GetUser(data.Phone)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -2,
			"msg":    "无效用户",
		})
		return
	}

	generateToken(c, user)
}

// 生成令牌
func generateToken(c *gin.Context, user *models.User) {
	j := &myjwt.JWT{
		SigningKey: []byte("newtrekWang"),
	}
	claims := myjwt.CustomClaims{
		ID:    user.ID,
		Name:  user.Name,
		Phone: user.Phone,
	}
	claims.NotBefore = int64(time.Now().Unix() - 1000)
	claims.ExpiresAt = int64(time.Now().Unix() + 3600)
	claims.Issuer = "newtrekWang"

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}

	log.Println(token)

	data := LoginResult{
		User:  *user,
		Token: token,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登录成功！",
		"data":   data,
	})
	return
}
