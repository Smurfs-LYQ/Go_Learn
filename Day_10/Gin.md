#### <center>Gin框架</center>


1. > Gin框架介绍

    - Go世界里最流行的Web框架，[Github](https://github.com/gin-gonic/gin)上有 `34K+` star。基于[httprouter](https://github.com/julienschmidt/httprouter)开发的Web框架。[中文文档](https://gin-gonic.com/zh-cn/docs/)齐全，简单易用的轻量级框架。

2. > Gin框架安装与使用

    - 安装

        下载并安装 `Gin` :

        ```go
        go get -u github.com/gin-gonic/gin
        ```

    - 第一个Gin示例

        ```go
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
        ```

        使用浏览器打开 `127.0.0.1:8080/index` 就能看到一串JSON字符串

3. > RESTful API

    - REST与技术无关，代表的是一种软件架构风格，REST是Representational State Transfer的简称，中文翻译为 "表征状态转移"或"表现层状态转化"。

    推荐阅读 [阮一峰 理解RESTful架构](http://www.ruanyifeng.com/blog/2011/09/restful.html)

    简单来说，REST的含义就是客户端与Web服务器之间进行交互的时候，使用HTTP协议的4个请求方法代表不同的动作。

    - `GET` 用来获取资源 (查)
    - `POST` 用来新建资源 (增)
    - `PUT` 用来更新资源 (改)
    - `DELETE` 用来删除资源 (删)

    只要API程序遵循了REST风格，那就可以称其为RESTful API。目前在前后端分离的架构中，前后端基本都是通过RESTful API来进行交互。

    例如，我现在要编写一个管理数据的系统，我可以查询对一本书进行查询、创建、更新和删除等操作，在编写程序的时候就要设计客户端浏览器与Web服务端交互端方式和路径。按照经验通常会设计成如下模式:

    |   请求方法   |   URL   |   含义   |
    | ---- | ---- | ---- |
    |  GET  |  /book  |  查询书籍信息  |
    |  POST  |  /create_book  |  创建书籍信息  |
    |  POST  |  /update_book  |  更新书籍信息  |
    |  POST  |  /delete_book  |  删除书籍信息  |

    同样的需求我们按照RESTful API设计如下:

    |   请求方法   |   URL   |   含义   |
    | ---- | ---- | ---- |
    |  GET  |  /book  |  查询书籍信息  |
    |  POST |  /book  |  创建书籍信息  |
    |  PUT  |  /book  |  更新书籍信息  |
    | DELETE |  /book  |  删除书籍信息  |

4. > 