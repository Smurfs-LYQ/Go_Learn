package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "主页",
	})
}

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// GET: 请求方式: /index: 请求地址
	// 当客户端以GET方式请求/index路径时，会执行后面的匿名函数
	r.GET("./hello", func(c *gin.Context) {
		// c.JSON: 返回JSON格式的数据
		c.JSON(200, gin.H{
			"msg": "hello world",
		})
	})

	r.GET("/", index)

	// 启动HTTP服务，默认以0.0.0.0:8080启动服务
	r.Run()
	// r.Run(":9090") // 以9090端口运行
	// r.Run("127.0.0.1:9090") // 设置IP:端口
}
