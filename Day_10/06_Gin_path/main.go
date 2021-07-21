package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/user/search/:username/:address", func(c *gin.Context) {
		// 提取路径参数
		username := c.Param("username")
		address := c.Param("address")
		// 输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"address":  address,
		})
	})

	r.Run()
}
