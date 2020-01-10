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

    同样的需求按照RESTful API设计如下:

    |   请求方法   |   URL   |   含义   |
    | ---- | ---- | ---- |
    |  GET  |  /book  |  查询书籍信息  |
    |  POST |  /book  |  创建书籍信息  |
    |  PUT  |  /book  |  更新书籍信息  |
    | DELETE |  /book  |  删除书籍信息  |

    Gin框架支持开发RESTful API的开发。

    ```go
    func main() {
        r := gin.Default()
        r.GET("/book", func(c *gin.Context){
            c.JSON(200, gin.H{
                "msg": "GET",
            })
        })

        r.POST("/book", func(c *gin.Context){
            c.JSON(200, gin.H{
                "msg": "POST",
            })
        })

        r.PUT("/book", func(c *gin.Context){
            c.JSON(200, gin.H{
                "msg": "PUT",
            })
        })

        r.DELETE("/book", func(c *gin.Context{
            c.JSON(200, gin.H{
                "msg": "DELETE",
            })
        }))
    }
    ```

    开发RESTful API的时候通常使用 Postman 来作为客户端的测试工具

4. > Gin渲染-HTML渲染

    - 首先定义一个存放模板文件的 `template` 文件夹，然后在其内部按照业务分别定义一个 `posts` 文件夹和一个 `users` 文件夹。

        `posts/index.html` 文件的内容如下:

        ```html
        {{define "posts/index.html"}}
        <!DOCTYPE html>
        <html lang="en">

        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <meta http-equiv="X-UA-Compatible" content="ie=edge">
            <title>posts/index</title>
        </head>
        <body>
            {{.title}}
        </body>
        </html>
        {{end}}
        ```

        `users/index.html` 文件的内容如下:

        ```html
        {{defind "users/index.html"}}
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <meta http-equiv="X-UA-Compatible" content="ie=edge">
            <title>users/index</title>
        </head>
        <body>
            {{.title}}
        </body>
        </html>
        {{end}}
        ```

        Gin框架中使用 `LoadHTMLGlob()` 或者 `LoadHTMLFiles()` 方法进行HTML模板渲染。

        ```go
        func main() {
            r := gin.Default()

            // 可以一次性渲染指定文件夹中的所有文件 下面的意思就是: templates文件夹下面的 所有文件夹下面的 所有文件
            r.LoadHTMLGlob("templates/**/*")
            // 必须把想要渲染的文件都写在里面
            // r.LoadHTMLFiles("templates/posts/index.html", "template/users/index.html")

            r.GET("/posts/index", func(c *gin.Context) {
                c.HTML(http.StatusOK, "posts/index.html", gin.H{
                    "title": "posts/index",
                })
            })

            r.GET("/users/index", func(c *gin.Context) {
                c.HTML(http.statusOK, "users/index.html", gin.H{
                    "title": "users/index",
                })
            })

            r.Run(":9090")
        }
        ```

5. > Gin渲染-自定义模板函数

    - 定义一个不转义相应内容的 `safe` 的模板函数如下:

        ```go
        func main() {
            r := gin.Default()

            // 设置模板函数
            r.SetFuncMap(template.FuncMap{
                "safe": func(str string) template.HTML{
                    return template.HTML(str)
                },
            })

            r.LoadHTMLFiles("./index.html")

            r.GET("/index", func(c *gin.Context) {
                c.HTML(http.StatusOK, "index.html", "<a href='https://baidu.com'>百度一下</a>")
            })

            r.Run()
        }
        ```

        在 `index.html` 中使用定义好的 `safe` 模板函数:

        ```html
        <!DOCTYPE html>
        <html lang="zh-CN">
        <head>
            <title>title</title>
        </head>
        <body>
        <div>{{ . | safe }}</div>
        </body>
        </html>
        ```

6. > Gin渲染-静态文件处理

    - 当渲染的HTMl文件中引用了静态文件(例如: css、js、image)时，只需要按照一下方式在渲染页面前调用 `gin.Static` 方法即可。

        ```go
        func main() {
            r := gin.Default()

            /*
            参数
                1. 代码里使用的路径
                2. 实际保存静态文件的路径
            */
            r.Static("/static", "./static")
            r.LoadHTMLGlob("templates/**/*")
            // ...
            r.Run()
        }
        ```

7. > Gin渲染-使用模板继承

    - Gin框架默认都是使用单模板，如果需要使用 `block template` 功能，可以通过 `"github.com/gin-contrib/multitemplate"` 库实现，具体示例如下：

        首先，假设项目目录下的templates文件夹下有以下模板文件，其中 `home.html` 和 `index.html` 继承了 `base.html`：

        ```txt
        templates
        ├── includes
        │   ├── home.html
        │   └── index.html
        ├── layouts
        │   └── base.html
        └── scripts.html
        ```

        然后定义一个 `loadTemplates` 函数如下:

        ```go
        func loadTemplates(templatesDir string) multitemplate.Renderer {
            r := multitemplate.NewRenderer()
            layouts, err := filepath.Glob(templatesDir + "/layouts/*.tmpl")
            if err != nil {
                panic(err.Error())
            }
            includes, err := filepath.Glob(templatesDir + "/includes/*.tmpl")
            if err != nil {
                panic(err.Error())
            }
            // 为layouts/和includes/目录生成 templates map
            for _, include := range includes {
                layoutCopy := make([]string, len(layouts))
                copy(layoutCopy, layouts)
                files := append(layoutCopy, include)
                r.AddFromFiles(filepath.Base(include), files...)
            }
            return r
        }
        ```

        在 `main` 函数中:

        ```go
        func indexFunc(c *gin.Context){
            c.HTML(http.StatusOK, "index.tmpl", nil)
        }

        func homeFunc(c *gin.Context){
            c.HTML(http.StatusOK, "home.tmpl", nil)
        }

        func main(){
            r := gin.Default()
            r.HTMLRender = loadTemplates("./templates")
            r.GET("/index", indexFunc)
            r.GET("/home", homeFunc)
            r.Run()
        }
        ```

8. > Gin渲染-补充文件路径处理

    - 关于模板文件和静态文件的路径，需要根据公司/项目的要求进行设置。可以使用下面的函数获取当前执行程序的路径。

        ```go
        func getCurrentPath() string {
            if ex, err := os.Executable(); err != nil {
                return filepath.Dir(ex)
            }
            return "./"
        }
        ```

9.  > Gin渲染-JSON渲染

    ```go
    func main() {
        r := gin.Default()

        r.GET("/someJSON", func(c *gin.Context) {
            // 方式一：自己拼接JSON
            c.JSON(http.StatusOK, gin.H{"msg": "hello world"})
        })

        r.GET("/moreJSON", func(c *gin.Context) {
            // 方式二：使用结构体
            type msg struct {
                Name    string
                Message string
                Age     int
            }

            msg.Name = "Smurfs"
            msg.Message = "Hello World"
            msg.Age = 18
            c.JSON(http.StatusOK, msg)
        })

        r.Run()
    }
    ```

10. > Gin渲染-XML渲染

    注意需要使用具名的结构体类型

    ```go
    func main() {
        r := gin.Default()

        r.GET("/someXML", func(c *gin.Context) {
            // 方式一：自己拼接JSON
            c.XML(http.StatusOK, gin.H{"msg": "hello world"})
        })

        r.GET("/moreXML", func(c *gin.Context) {
            // 方式二：使用结构体
            type msg struct {
                Name    string
                Message string
                Age     int
            }

            msg.Name = "Smurfs"
            msg.Message = "Hello World"
            msg.Age = 18
            c.XML(http.StatusOK, msg)
        })

        r.Run()
    }
    ```

11. > Gin渲染-YAML渲染

    多用于配置文件，需要特殊的协议才可以收到数据

    ```go
    r.GET("/someYAML", func(c *gin.Context) {
        c.YAML(http.StatusOK, gin.H{"msg": "ok", "status": http.StatusOK})
    })
    ```

12. > Gin渲染-protobuf渲染

    做微服务二进制的

    ```go
    r.GET("/someProtoBuf", func(c *gin.Context) {
        reps := []int64{int64(1), int64(2)}
        label := "test"
        // protobuf 的具体定义写在 testdata/protoexample 文件中。
        data := &protoexample.Test{
            Label: &label,
            Reps:  reps,
        }
        // 请注意，数据在响应中变为二进制数据
        // 将输出被 protoexample.Test protobuf 序列化了的数据
        c.ProtoBuf(http.StatusOK, data)
    })
    ```

13. > 获取参数-获取querystring参数

    `querystring` 指的是URL中 `?` 后面携带的参数，例如: `/user/search?username=Smurfs&address=北京`。获取请求的querystring参数的方法如下:

    ```go
    func main() {
        r := gin.Default()

        r.GET("/user/search", func(c *gin.Context) {
            // 接收参数，有默认值
            username := c.DefaultQuery("username", "Smurfs")
            // 接收参数，无默认值
            // username := c.Query("username")
            address := c.Query("address")
            // 输出json结果给调用方
            c.JSON(http.StatusOK, gin.H{
                "message": "ok",
                "username": username,
                "address": address,
            })
        })

        r.Run()
    }
    ```

14. > 获取参数-获取form参数

    请求的数据通过form表单来提交，例如向 `/user/search` 发送一个POST请求，获取请求数据的方式如下:

    ```go
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
    ```

15. > 获取参数-获取path参数

    请求的参数通过URL路径传递，例如: `/user/search/smurfs/北京`。获取请求URL路径中的参数的方式如下。

    ```go
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
    ```

16. > 获取参数-参数绑定

    为了能够更方便的获取请求相关参数，提高开发效率，我们可以基于请求的 `Context-Type` 识别请求数据类型并利用反射机制自动提取请求中 `QueryString`、`form表单`、`JSON`、`XML` 等参数到结构体中。下面的示例代码演示了 `.ShouldBind()` 强大的功能，它能够基于请求自动提取 `JSON`、`form表单` 和 `QueryString` 类型的数据，并把值绑定到指定的结构体对象。

    ```go
    // Binding from JSON
    type Login struct {
        User     string `form:"user" json:"user" binding:"required"`
        Password string `form:"password" json:"password" binding:"required"`
    }

    func main() {
        router := gin.Default()

        // 绑定JSON的示例 ({"user": "q1mi", "password": "123456"})
        router.POST("/loginJSON", func(c *gin.Context) {
            var login Login

            if err := c.ShouldBind(&login); err == nil {
                fmt.Printf("login info:%#v\n", login)
                c.JSON(http.StatusOK, gin.H{
                    "user":     login.User,
                    "password": login.Password,
                })
            } else {
                c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            }
        })

        // 绑定form表单示例 (user=q1mi&password=123456)
        router.POST("/loginForm", func(c *gin.Context) {
            var login Login
            // ShouldBind()会根据请求的Content-Type自行选择绑定器
            if err := c.ShouldBind(&login); err == nil {
                c.JSON(http.StatusOK, gin.H{
                    "user":     login.User,
                    "password": login.Password,
                })
            } else {
                c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            }
        })

        // 绑定QueryString示例 (/loginQuery?user=q1mi&password=123456)
        router.GET("/loginForm", func(c *gin.Context) {
            var login Login
            // ShouldBind()会根据请求的Content-Type自行选择绑定器
            if err := c.ShouldBind(&login); err == nil {
                c.JSON(http.StatusOK, gin.H{
                    "user":     login.User,
                    "password": login.Password,
                })
            } else {
                c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            }
        })

        // Listen and serve on 0.0.0.0:8080
        router.Run(":8080")
    }
    ```

    `ShouldBind` 会按照下面的顺序解析请求中的数据完成绑定:

    - 如果是 `GET` 请求，只使用 `Form` 绑定引擎 (query)。
    - 如果是 `POST` 请求，首先检查 `content-type` 是否为 `JSON` 或 `XML`，然后在使用 `FORM` (form-data)。

    `ShouldBind` 默认情况下不传应该有的参数，对应的参数为空，不会报错，但是在声明的结构体中参数的后面加上tag `binding:"required"` 之后，如果再不传应该有的参数，那么会导致报错

17. > 文件上传—文件上传

    - 文件上传前段页面代码:

        [demo](https://github.com/Smurfs-LYQ/Go_Learn/blob/master/Day_10/09_Gin_file_upload/templates/index.html)

    - 后端gin框架代码

        [demo](https://github.com/Smurfs-LYQ/Go_Learn/blob/master/Day_10/09_Gin_file_upload/main.go)

18. > 中间件

    - Gin框架允许开发者在处理请求的过程中，加入用户自己的钩子(Hook)函数。这个钩子函数就叫中间件，中间件适合处理一下公共的业务逻辑，比如登录检验、日志打印、耗时统计等。

    - Gin中的中间件必须是一个 `gin.HandlerFunc` 类型。例如我们像下面的代码一样定义一个中间件。

        ```go
        // StatCost 是一个统计耗时请求耗时的中间件
        func StatCost() gin.HandlerFunc {
            return func(c *gin.Context) {
                start := time.Now()
                c.Set("name", "smurfs")
                // 执行其他中间件
                c.Next()
                // 计算耗时
                cost := time.Since(start)
                log.Println(cost)
            }
        }
        ```

    - 然后注册中间件的时候，可以在全局注册

        ```go
        func main() {
            // 新建一个没有任何默认中间件的路由
            r := gin.New()
            // 注册一个全局中间件
            r.Use(StatCost())

            r.GET("/test", func(c *gin.Context) {
                name := c.MustGet("name").(string)
                log.Println(name)
                c.JSON(http.StatusOK, gin.H{
                    "message": "Hello world!",
                })
            })
            r.Run()
        }
        ```

    - 也可以给某个路由单独注册中间件

        ```go
        // 给/test2路由单独注册中间件(可注册多个)
        r.GET("/test2", StatCost(), func(c *gin.Context) {
            name := c.MustGet("name").(string)
            log.Println(name)
            c.JSON(http.StatusOK, gin.H{
                "message": "Hello world!",
            })
        })
        ```

19. > 重定向-HTTP重定向

    - HTTP重定向很容易。内部、外部重定向均支持

        ```go
        r.GET("/test", func(c *gin.Context) {
            c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
        })
        ```

20. > 重定向-路由重定向

    - 路由重定向，使用 `HandleContext`:

        ```go
        r.GET("/test", func(c *gin.Context) {
            // 指定重定向的URL
            c.Request.URL.Path = "/test2"
            r.HandleContext(c)
        })
        r.GET("/test2", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{"hello": "world"})
        })
        ```

21. > Gin路由-普通路由

    ```go
    r.GET("/index", func(c *gin.Context) {...})
    r.GET("/login", func(c *gin.Context) {...})
    r.POST("/index", func(c *gin.Context) {...})
    ```

    此处，还有一个可以匹配所有请求方法的 `Any` 方法如下:

    ```go
    r.Any("/test", func(c *gin.Context) {...})
    ```

    为没有配置处理函数的路由添加处理程序，默认情况下它返回404代码，下面的代码为没有匹配到路由的请求都返回 `views/404.html` 页面。

    ```go
    r.NoRoute(func(c *gin.Context) {
        c.HTML(http.StatusNotFound, "views/404.html", nil)
    })
    ```

22. > Gin路由-路由组

    我们可以将拥有共同URL前缀的路由划分为一个路由组。习惯性使用 `{}` 包裹同组的路由，这只是为了看着清晰，你用不用 `{}` 包裹功能上没什么区别。

    ```go
    func main() {
        r := gin.Default()
        userGroup := r.Group("/user")
        {
            userGroup.GET("/index", func(c *gin.Context) {...})
            userGroup.GET("/login", func(c *gin.Context) {...})
            userGroup.POST("/login", func(c *gin.Context) {...})
        }
        shopGroup := r.Group("/shop")
        {
            userGroup.GET("/index", func(c *gin.Context) {...})
            userGroup.GET("/cart", func(c *gin.Context) {...})
            userGroup.POST("/checkout", func(c *gin.Context) {...})
        }
        r.Run()
    }
    ```

    ```txt
    访问:
        user组
        127.0.0.1:8080/user/index GET方式请求
        127.0.0.1:8080/user/login GET方式请求
        127.0.0.1:8080/user/login POST方式请求

        shop组
        27.0.0.1:8080/user/index GET方式请求
        127.0.0.1:8080/user/cart GET方式请求
        127.0.0.1:8080/user/checkout POST方式请求
    ```

    通常我们将路由分组用在划分业务逻辑或划分API版本时。

23. > Gin路由-路由原理

    - Gin框架中的路由使用的是 [httprouter](https://github.com/julienschmidt/httprouter)这个库
    - 其基本原理就是构造一个路由地址的前缀树。