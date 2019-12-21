#### <center>Day_08</center>

1. Go语言获取请求提交的数据(form)
2. Go语言实现动态数据页面的示例
3. template简单语法示例
4. template

#### <center>笔记</center>
1. > Sonar
  
    - Sonar是大部分公司会用的代码审查工具
   
 2. > Web本质

   - C/S 架构:

     描述: 客户端/服务端

     优势: 可定制化高，用户体验好

     劣势: 开发成本高(需要适配多种不同的平台)，添加新功能需要客户端升级

   - B/S 架构:

     描述: 浏览器/服务端，基于浏览器的架构(Web开发)

     优势: 开发成本低

     劣势: 没办法做很多复杂的功能

3. > GET和POST

   - GET: 把请求的数据拼接到URL后面
   - POST: 把请求的数据放到请求体中

4. > HTTP服务端

   - 设置访问路径与其对应的函数

     ```go
     func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
     // HandleFunc注册一个处理器函数handler和对应的模式pattern（注册到DefaultServeMux）。ServeMux的文档解释了模式的匹配机制。
     ```

   - 监听一个指定的TCP地址

     ```go
     func ListenAndServe(addr string, handler Handler) error
     // ListenAndServe监听TCP地址addr，并且会使用handler参数调用Serve函数处理接收到的连接。handler参数一般会设为nil，此时会使用DefaultServeMux。
     ```

   - 例如

     ```go
     func Hello(w http.ResponseWriter, r *http.Request) {
     	// w: 代表了跟响应相关的所有内容
     	// r: 代表了跟请求相关的所有内容
     	fmt.Println("hello world")
       w.Write([]byte("hello golang"))
     }
     
     func main() {
     	/*
     		当访问127.0.0.1:8080/index的时候执行hello函数，在终端打印"hello world"，在页面上打印"hello golang"
     	*/
     	http.HandleFunc("/index", hello)
     	http.ListenAndServe("127.0.0.1:8080", nil)
     }
     ```

5. > HTTP服务端处理Form表单

   - 页面Form表单提交数据

     ```html
     <!--提交到127.0.0.1:8080/login中-->
     <form action="/login" method="POST">
       <!--作用: 在点击用户名的时候跳转到id等于username的input中-->
       <label for="username">用户名: </label>
       <input type="text" id="username" name="username" placeholder="请输入用户名">
       <br/>
       
       <label for="password">密码: </label>
       <input type="password" id="password" name="password" placeholder="请输入密码">
       <br/>
       
       <button type="submit">提交</button>
     </form>
     ```

   - HTTP服务端

     ```go
     // 加载index页面
     func index(w http.ResponseWriter, r *http.Request) {
       res, err := ioutil.ReadFile("./index.html")
       if err != nil {
         panic(fmt.Sprintln("页面加载失败, err:", err))
       }
       w.Write(res)
     }
     
     // 处理Form表单发送来的数据
     func login(w http.ResponseWriter, r *http.Request) {
     	// 解析Form表单的数据
       r.ParseForm()
       
       // 打印Form表单的请求方法(GET/POST)
       fmt.Println(r.Method)
       
       // 在终端上打印Form表单所有的数据
       fmt.Println(r.Form)
       
       // 获取用户名和密码
       username := r.Form.Get("username")
       password := r.Form.Get("password")
       // 在页面上打印用户名和密码
       w.Write([]byte(username))
       w.Write([]byte(password))
     }
     
     func main() {
     	http.HandleFunc("/", index)
     	http.HandleFunc("/login", login)
     	http.ListenAndServe("127.0.0.1:8080", nil)
     }
     ```

6. > Go语言的模板引擎

   - Go语言内置了文本模板引擎 `text/template` 和用于HTML文档的 `html/template`。它们的作用机制可以简单归纳如下：
     - 模板文件通常定义为 `.tmpl` 和 `.tpl` 为后缀 (也可以使用其他的后缀)，必须使用UTF8编码。
     - 模板文件中使用 `{{` 和 `}}` 包裹和标识需要传入的数据。
     - 传给模板这样的数据就可以通过点号 `.` 来访问，如果数据是复杂类型的数据，可以通过 `{{.FieldName}}` 来访问它的字段。
     - 除 `{{` 和 `}}` 包裹的内容外，其它内容均不做修改原样输出。

7. > 标准库template

   - `html/template` 包实现了数据驱动的模板，用于生成可防止代码注入的安全的HTML内容。它提供了和 `test/template` 包相同的接口，Go语言中输出HTML的场景应该使用 `html/template` 这个包。
   
   - **模板引擎的使用**
   
     ​	Go语言模板引擎的使用可以分为三部分:
   
       1. 定义模板文件
   
          - 编写页面文件，详见 **6. Go语言的模板引擎** 就在上面...
   
       2. 解析模板文件
   
          - 上面定义好模板文件之后，可以使用下面的常用方法去解析模板文件，得到模板对象:
   
            ```go
            func (t *Template) Parse(src string) (*Template, error)
            func ParseFiles(filenames ...string) (*Template, error)
            func ParseGlob(pattern string) (*Template, error)
            ```
   
          - 当然，也可以使用 `func New(name string) *Template` 函数创建一个名为 `name` 的模板，然后对其调用上面的方法去解析模板字符串或模板文件。
   
       3. 模板渲染
   
          - 模板渲染简单来说就是使用数据去填充模板，本质上就是高级的字符串替换。
   
            ```go
            func (t *Template) Execute(wr io.Writer, data interface{}) error
            func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
            ```
   
     - 基本示例见: [Github](https://github.com/Smurfs-LYQ/Go_Learn/tree/master/Day_08/03_template_demo)
   
   - 
   
   - 