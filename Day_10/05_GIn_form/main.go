package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/user/search", func(c *gin.Context) {
		// DefaultPostForm取不到值时会返回指定的默认值
		// username := c.DefaultPostForm("username", "Smurfs")
		username := c.PostForm("username")
		address := c.PostForm("address")

		// 输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"address":  address,
		})
	})

	r.Run()
}
