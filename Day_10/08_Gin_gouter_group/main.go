package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 路由组
func main() {
	r := gin.Default()

	userGroup := r.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "用户主页"})
		})
		userGroup.GET("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "用户登录"})
		})
		userGroup.POST("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "用户登录处理程序"})
		})
	}
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "商品首页"})
		})
		shopGroup.GET("/add", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "添加商品"})
		})
		shopGroup.POST("/add", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "添加商品处理程序"})
		})
	}

	r.Run()
}
