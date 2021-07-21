package main

import (
	"Go_Learn/Day_11/project_blogs/controller"
	"Go_Learn/Day_11/project_blogs/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库链接
	err := db.InitDB()
	if err != nil {
		log.Println("数据库链接失败, err:", err)
	}

	r := gin.Default()

	// gin.SetMode(gin.ReleaseMode)

	// 加载静态文件
	r.Static("/static", "./static")

	// 加载模板文件
	r.LoadHTMLGlob("./views/*")

	// 路由
	r.GET("/", controller.IndexHandler)

	// 如果访问不存在的路由地址, 自动跳转到主页
	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/")
	})

	r.Run()
}
