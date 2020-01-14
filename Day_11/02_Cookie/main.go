package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// cookie简单示例

type User struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func userHandler(c *gin.Context) {
	// 检查Cookie
	username, err := c.Cookie("username")
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/login")
	}

	c.HTML(http.StatusOK, "user.html", gin.H{
		"username": username,
	})
}

func loginHandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		var u User
		err := c.ShouldBind(&u)
		if err != nil {
			fmt.Println(err)
			c.HTML(http.StatusOK, "login.html", gin.H{
				"err": "用户名和密码不能为空",
			})
			return
		}

		if u.Username == "smurfs" && u.Password == "123" {
			// 设置Cookie
			// 参数：cookie名，cookie值，生命时长，路径，域名，是否启用安全策略，防止别人通过JS来获取cookie
			c.SetCookie("username", "smurfs", 20, "/", "127.0.0.1", false, true)
			c.Redirect(http.StatusFound, "/user")
		} else {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"err": "用户名或密码错误",
			})
		}
	} else {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"err": "请登录",
		})
	}
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/user", userHandler)

	r.Any("/login", loginHandler)

	r.Run()
}
