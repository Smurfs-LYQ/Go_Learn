package main

import (
	"Go_Learn/Day_11/04_Session/session"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

var SessionMgr *session.SessionMgr

func user_session() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 先获取cookie 判断有没有用户的id
		uuid, err := c.Cookie("uuid")
		if err != nil {
			log.Println(err)
			c.Redirect(http.StatusMovedPermanently, "/login")
			return
		}

		// 去session中判断，session中有没有对应的user session_data
		session_data, err := SessionMgr.GetSession(uuid)
		if err != nil {
			log.Println(err)
			c.Redirect(http.StatusMovedPermanently, "/login")
		}

		// 去session_data中判断是否有username
		val, err := session_data.GetSessionData("username")
		if err != nil {
			log.Println(err)
			c.Redirect(http.StatusMovedPermanently, "/login")
		}

		// 设置用户名
		c.Set("username", val)

		c.Next()
	}
}

func indexHandler(c *gin.Context) {
	username, ok := c.Get("username")
	if !ok {
		log.Println("username获取失败")
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
			log.Println(err)
			c.HTML(http.StatusOK, "login.html", gin.H{
				"msg": "用户名或密码不能为空",
			})
		}

		if u.Username == "smurfs" && u.Password == "123" {
			// 创建Cookie
			uuid := SessionMgr.AddSession()
			SessionMgr.Session[uuid].AddSessionData("username", u.Username)
			c.SetCookie("uuid", uuid, 3, "/", "127.0.0.1", false, true)

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

func main() {
	SessionMgr = session.NewSessionMgr()

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.Any("/login", loginHandler)
	r.GET("/user/index", user_session(), indexHandler)

	r.Run()
}
