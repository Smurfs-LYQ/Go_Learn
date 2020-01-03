package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	/*
		返回HTML格式的数据
		参数：
			1. 状态码
			2. 模板文件
			3. 数据
	*/
	c.HTML(http.StatusOK, "index.html", gin.H{
		"msg":  "这是主页",
		"logo": "/static/images/logo.jpg",
	})
}

func loginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"msg": "<h1>登录页面</h1>",
		"url": "<a href='https://baidu.com'>百度一下</a>",
	})
}

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()

	// 自定义模板函数 参数: template.FuncMap类型的map
	r.SetFuncMap(template.FuncMap{
		// safe函数的作用是 将传入的字符串转换为html格式并返回
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	/*
		设置静态文件的目录
		参数
			1. 代码里使用的路径
			2. 实际保存静态文件的路径
	*/
	r.Static("/static", "./static")

	// 加载模板文件
	r.LoadHTMLGlob("templates/**/*")

	// 当客户端以GET方式请求 / 路径时，返回json格式的数据 200 和 {"msg":"hello"}
	r.GET("/index", indexHandler)
	r.GET("/login", loginHandler)

	// 启动服务
	r.Run()
}
