#### <center>Day_11</center>

1. Logrus-第三方日志库
2. Gin框架操作Cookie
3. Gin框架-用户登录 写入cookie记录登录状态
4. Gin框架-用户登录 缓存版本session
5. Gin框架-用户登录 Redis版本session

#### <center>笔记</center>

1. > Go操作Cookie

    - Cookie

        标准库 `net/http` 中定义了Cookie，它代表一个出现在HTTP响应头中Set-Cookie的值里或者HTTP请求头中Cookie的值的 `HTTP cookie`。

        ```go
        type Cookie struct {
            Name       string       // 名
            Value      string       // 值
            Path       string       // 路径
            Domain     string       // 域名
            Expires    time.Time    // 超时时间(有缺陷，不推荐使用) 绝对时间
            RawExpires string
            // MaxAge=0表示未设置Max-Age属性
            // MaxAge<0表示立刻删除该cookie，等价于"Max-Age: 0"
            // MaxAge>0表示存在Max-Age属性，单位是秒
            MaxAge   int            // 生命时长 相对时间
            Secure   bool           // 是否启用安全测量
            HttpOnly bool           // 只允许HTTP进行访问，防止别人通过JS来获取Cookie
            Raw      string         // 
            Unparsed []string       // 未解析的“属性-值”对的原始文本
        }
        ```

    - 设置Cookie

        `net/http` 中提供了如下 `SetCookie` 函数，它在w的头域中添加Set-Cookie头，该HTTP头的值为cookie。

        ```go
        func SetCookie(w ResposeWriter, cookie *Cookie)
        ```

    - 获取Cookie

        `Request` 对象拥有两个获取Cookie的方法和一个添加Cookie的方法:

        获取Cookie的两种方法:

        ```go
        // 解析并返回该请求的Cookie头设置的所有cookie
        func (r *Request) Cookies() []*Cookie

        // 返回请求中名为name的cookie，如果未找到该cookie会返回nil，ErrNoCookie
        func (r *Request) Cookie(name string) (*Cookie, error)
        ```

        添加Cookie的方法:

        ```go
        // AddCookie向请求中添加一个cookie
        func (r *Request) AddCookie(c *Cookie)
        ```

    - Gin框架操作Cookie

    ```go
    func main() {
        r := gin.Default()

        r.GET("/", func(c *gin.Context) {
            cookie, err := c.Cookie("gin_cookie") // 获取cookie
            if err != nil {
                cookie = "NotSet"
                // 设置cookie
                // 参数：cookie名，cookie值，生命时长，路径，域名，cookie属性，是否启用安全策略，防止别人通过JS来获取cookie
                c.SetCookie("gin_cookie", "test", 3600, "/", "127.0.0.1", http.SameSiteDefaultMode, false, true)
            }

            fmt.Printf("Cookie value: %s\n", cookie)
        })

        r.Run()
    }
    ```

    - Cookie的缺点

        - 数据量最大4K

        - 保存在客户端(浏览器)，不安全

2. > Session

    Cookie虽然在一定程度上解决了“保持状态”的需求，但是由于Cookie本身最大支持4096字节，以及Cookie本身保存在客户端，可能被拦截或窃取，因此就需要有一种新的东西，它能支持更多的字节，并且他保存在服务器，有较高的安全性。这就是 `Session`。

    Session的存在必须依赖于Cookie，Cookie中保存了没个用户Session的唯一标识。