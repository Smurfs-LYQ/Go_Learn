package main

import (
	"Go_Learn/Day_11/05_Session_redis/db"
	"Go_Learn/Day_11/05_Session_redis/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	err := db.InitRedis()
	if err != nil {
		log.Println("Redis启动失败, err:", err)
		return
	}

	r := gin.Default()

	r.LoadHTMLGlob("./templates")

	r.Any("/login", user.Login)

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "404.html", nil)
	})

	r.Run()
}
