package user

import (
	"Go_Learn/Day_11/05_Session_redis/session"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	UserName string `form:"username"`
	Password string `form:"password"`
}

// Login 登录
func Login(c *gin.Context) {
	if c.Request.Method != "POST" {
		c.HTML(http.StatusOK, "login.html", nil)
	} else {
		var user user

		// 接收信息
		err := c.ShouldBind(&user)
		if err != nil {
			log.Println(err)
			c.HTML(http.StatusFound, "login.html", gin.H{
				"msg": "用户名或密码不能为空",
			})
			return
		}

		log.Println(user)

		if user.UserName == "smurfs" && user.Password == "123" {
			// 初始化Session
			sess := session.InitSession()
			// 添加Session信息
			sess.Add("username", user.UserName)
			sess.Add("password", user.Password)
			// 保存Session信息
			id, err := sess.Save()
			if err != nil {
				log.Println(err)
				c.HTML(http.StatusFound, "login.html", gin.H{
					"msg": "登录失败",
				})
			}
			// 将唯一ID保存到Cookie中
			c.SetCookie("UUID", id, session.Session_time, "/", "127.0.0.1", false, true)

			// 跳转到用户界面
			c.Redirect(http.StatusFound, "/user")
		} else {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"msg": "用户名或密码错误",
			})
		}

	}
}

// User 用户页面
func User(c *gin.Context) {
	// 从cookie中获取UUID
	user_uuid, err := c.Cookie("UUID")
	if err != nil {
		log.Println("UUID cookie not found, err:", err)
		c.Redirect(http.StatusFound, "/login")
		return
	}

	// 判断Redis中是否保存了对应UUID的信息
	sess, err := session.Sel_Session(user_uuid)
	if err != nil {
		log.Println("redis session not found, err:", err)
		c.Redirect(http.StatusFound, "/login")
		return
	}

	// 获取用户名
	username, err := sess.Sel("username")
	if err != nil {
		log.Println("获取用户名失败, err:", err)
		c.Redirect(http.StatusFound, "/login")
		return
	}

	c.HTML(http.StatusOK, "user.html", gin.H{
		"username": username,
	})
}
