package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	UserName string
	Password string
}

// Login 登录
func Login(c *gin.Context) {
	if c.Request.Method != "POST" {
		c.HTML(http.StatusOK, "login.html", nil)
	} else {
		var user *user

		// 接收信息
		err := c.ShouldBind(user)
		if err != nil {
			fmt.Println(err)
			c.HTML(http.StatusFound, "login.html", gin.H{
				"msg": "用户名或密码不能为空",
			})
		}

	}
}
