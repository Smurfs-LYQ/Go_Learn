package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func user_cookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		name, err := c.Cookie("username")
		if err != nil {
			c.Redirect(http.StatusMovedPermanently, "/login")
		}
		c.Set("username", name)
		c.Next()
	}
}

func loginHandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		var u User
		err := c.ShouldBind(&u)
		if err != nil {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"msg": "用户名和密码不能为空",
			})
		}
		if u.Username == "smurfs" && u.Password == "123" {
			c.SetCookie("username", u.Username, 20, "/", "127.0.0.1", false, true)
			c.Redirect(http.StatusFound, "/user/index")
		} else {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"msg": "用户名或密码错误",
			})
		}
	} else {
		c.HTML(http.StatusOK, "login.html", nil)
	}
}

func userIndexHandler(c *gin.Context) {
	username, ok := c.Get("username")
	if !ok {
		c.Redirect(http.StatusMovedPermanently, "/login")
	}
	c.HTML(http.StatusOK, "user.html", gin.H{
		"username": username.(string),
	})
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.Any("/login", loginHandler)
	userGroup := r.Group("/user", user_cookie())
	{
		userGroup.GET("/index", userIndexHandler)
	}

	r.Run()
}
