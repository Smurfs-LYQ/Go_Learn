package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB

	r.GET("/", indexHandler)

	r.POST("/file", func(c *gin.Context) {
		// 单个文件
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			return
		}

		log.Println(file.Filename)
		// 设定保存上传文件的路径
		dst := fmt.Sprintf("./photo/%s", file.Filename)
		// 上传文件到指定的目录
		err = c.SaveUploadedFile(file, dst)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
			})
		} else {
			msg := fmt.Sprintf("'%s' not uploaded!", file.Filename)
			log.Println(msg, err)
			c.JSON(http.StatusOK, gin.H{
				"message": msg,
			})
		}
	})

	r.POST("/file", func(c *gin.Context) {
		// 多文件上传
		// Multipart form
		form, err := c.MultipartForm()
		if err != nil {
			log.Println("调取MultipartForm失败，err:", err)
		}
		files := form.File["file"]

		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("./photo/%d_%s", index, file.Filename)
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	})

	r.Run()
}
