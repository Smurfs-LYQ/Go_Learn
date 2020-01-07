package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// StatCost 定义一个统计请求耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// 按照顺序调用下一个handler函数
		c.Next()
		// 计算耗时
		cost := time.Since(start)
		log.Println(cost)
	}
}

// 设置一个值，所有调用过这个中间件的都可以通过Get/MustGet方法调用到这个值
func T1() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("name", "Smurfs")
		c.Next()
	}
}

func userIndexHandler(c *gin.Context) {
	// name, _ := c.Get("name")
	name := c.MustGet("name")
	time.Sleep(time.Second)
	c.JSON(http.StatusOK, gin.H{"msg": fmt.Sprintf("%s的用户主页\n", name.(string))})
}

func userLoginHandler(c *gin.Context) {
	name, _ := c.Get("name")

	c.JSON(http.StatusOK, gin.H{"msg": fmt.Sprintf("%s用户登录\n", name)})
}

func main() {
	r := gin.Default()

	// 设置所有路由都走这个中间件
	r.Use(StatCost(), T1())

	// 设置一个路由组走指定的中间件
	// userGroup := r.Group("/user", StatCost(), T1())
	userGroup := r.Group("/user")
	{
		/*
			// 设置单独某个路由的中间件
			// 访问路径 中间件函数 处理函数 调用顺序是按照编写的先后顺序来的
			userGroup.GET("/index", StatCost(), T1(), userIndexHandler)
		*/
		userGroup.GET("/index", userIndexHandler)
		userGroup.GET("/login", userLoginHandler)
	}

	r.Run()
}
