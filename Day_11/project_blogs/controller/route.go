package controller

import (
	"Go_Learn/Day_11/project_blogs/db"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
)

// IndexHandler 主页路由方法
func IndexHandler(c *gin.Context) {
	// 获取首页文章列表
	article_list, err := db.Index_articles()
	if err != nil {
		log.Println("获取主页文章列表失败, err:", err)
		return
	}

	// 获取首页文章分类
	article_type, err := db.Index_article_type()
	if err != nil {
		log.Println("获取主页文章类型失败, err:", err)
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"article_list": article_list,
		"article_type": article_type,
	})
}
