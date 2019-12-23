#### <center>Day_08</center>

1. Go语言获取请求提交的数据(form)
2. Go语言实现动态数据页面的示例
3. template简单语法示例
4. template 变量和条件判断
5. template with和range
6. template 预定义函数
7. template 自定义函数
8. template 嵌套
9. 链式操作
10. template block

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
   
8. > template模板语法

   - {{.}}

     模板语法都包含在 `{{` 和 `}}` 中间，其中 `{{.}}` 中的点表示当前对象。[示例代码](https://github.com/Smurfs-LYQ/Go_Learn/blob/master/Day_08/04_template/index.html)

     当传入一个结构体对象时，可以根据 `.` 来访问结构体的对应字段。同理，当传入的变量是map时，也可以在模板文件中通过 `.` 根据key来取值。[示例代码](https://github.com/Smurfs-LYQ/Go_Learn/blob/master/Day_08/04_template/main.go)

   - 注释

     ```go
     {{/* 注释内容 */}}
     注释，执行时会忽略。可以多行。注释不能嵌套，而且必须紧贴分界符始止
     ```

   - pipeline

     "pipeline" 是管道的意思。

     `pipeline` 是指产生数据的操作。比如 `{{.}}` 、`{{.Name}}` 等。Go的模板语法中支持使用管道符号 `|` 链接多个命令，用法和unix下的管道类似: `|` 前面的命令会将运算结果(或返回值)传递给后一个命令的最后一个位置。

     **注意:** 并不是只有使用了 `|` 才是pipeline。Go的模板语法中，`pipeline` 的概念是传递数据，只要能产生数据的，都是 `pipeline`。相当于管道

   - 变量

     可以在模板中声明变量，用来保存传入模板的数据或其它语句生成的结果。示例代码如下:

     ```go
     %obj := {{.}}
     ```

     其中 `$obj` 是变量的名字，在后续的代码中就可以使用该变量了。

   - 移除空格

     `{{-` 语法可以去除模板内容左侧的所有空白符号，`-}}` 可以去除模板内容右侧的所有空白符号。示例代码如下:

     ```go
     {{- .Name -}}
     ```

     **注意:** `-` 要紧挨 `{{` 和 `}}`，同时与模板值之间需要使用空格分隔。

   - 条件判断

     Go模板语法中的条件判断有以下几种:

     ```go
     {{if pipeline}} T1 {{end}}
     {{if pipeline}} T1 {{else}} T0 {{end}}
     {{if pipeline}} T1 {{else if pipeline}} T0 {{end}}
     ```

   - range

     Go的模板语法中使用 `range` 关键字进行遍历，有以下两种写法，其中 `pipeline` 的值必须是数组、切片、字典或管道。

     ```go
     {{range pipeline}} T! {{end}}
     如果pipeline的值其长度为0，不会有任何输出
     
     {{range pipeline}} T1 {{else}} T0 {{end}}
     如果pipeline的值其长度为0，则不会执行T0
     ```

   - with

     ```go
     {{with pipeline}} T1 {{end}}
     如果pipeline为empty不产生输出，否则将dot设为pipeline的值并执行T1。不修改外面的dot。
     
     {{with pipeline}} T1 {{else}} T0 {{end}}
     如果pipeline为empty，不改变dot并执行T0，否则dot设为pipeline的值并执行T1。
     ```

   - 预定义函数

     执行模板时，函数从两个函数字典中查找: 首先是模板函数字典，然后是全局函数字典。一般不在模板内定义函数，而是使用Funcs方法添加函数到模板里。

     预定义的全局函数如下:

     ```txt
     and
     	函数返回它的第一个empty参数或者最后一个参数;
     	就是说"and x y"等价于"if x then y else x"; 所有参数都会执行。
     or
     	返回第一个非empty参数或者最后一个参数;
     	亦即"or x y"等价于"if x then x else y"; 所有参数都会执行。
     not
     	返回它的单个参数的布尔值的设定。
     len
     	返回它的参数的整数类型长度。
     index
     	执行结果为第一个参数以剩下的参数为索引/键指向的值;
     	如"index x 1 2 3"返回x[1][2][3]的值; 每个被索引的主体必须是数组、切片或字典。
     print
     	即fmt.Sprint
     printf
     	即fmt.Sprintf
     println
     	即fmt.Sprintln
     html
     	返回与其参数的文本表示形式等效的转义HTML。
     	这个函数在html/template中不可用。
     urlquery
     	以适合嵌入到网址查询中的形式返回其参数的文本表示的转义值。
     	这个函数在html/template中不可用。
     js
     	返回与其参数的文本表示形式等效的转义JavaScript。
     call
     	执行结果是调用第一个参数的返回值，该参数必须是函数类型，其余参数作为调用该函数的参数;
     	如"call .X.Y 1 2"等价于go语言里的dot.X.T(1, 2);
     	其中Y是函数类型的字段或者字典的值，或者其他类似情况;
     	call的第一个参数的执行结果必须是函数类型的值(和预定义函数如print明显不同);
     	该函数类型值必须有1到2个返回值，如果有2个则后一个必须是error接口类型;
     	如果有2个返回值的方法返回的error非nil，模板执行会终端并返回给调用模板执行者该错误;
     ```

   - 比较函数

     布尔函数会将任何类型的零值视为假，其余视为真。

     下面定义是为函数的二元比较运算的集合:

     ```go
     eq	如果arg1 == arg2则返回真
     ne	如果arg1 != arg2则返回真
     lt	如果arg1 < arg2则返回真
     le	如果arg1 <= arg2则返回真
     gt	如果arg1 > arg2则返回真
     ge	如果arg1 >= arg2则返回真
     ```

     为了简化多参数相等检测，eq (只有eq) 可以接受两个或多个参数，它会将第一个参数和其余参数依次比较，返回下列的结果:

     ```go
     {{eq arg1 arg2 arg3}}
     ```

     比较函数只适用于基本类型 (或重定义的基本类型，如"type Celsius float32")。但是，整数和浮点数不能互相比较。

   - 自定义函数

     [示例代码](https://github.com/Smurfs-LYQ/Go_Learn/tree/master/Day_08/07_template_custom)
   
     **注意:**  自定义函数需要在页面解析之前加入到模板中。
   
   - 嵌套template
   
     template中可以还可以嵌套其他的template。这个template可以是单独的文件，也可以是通过 `define` 定义的template。
   
     [示例代码](https://github.com/Smurfs-LYQ/Go_Learn/tree/master/Day_08/08_template_nest)
   
     **注意:** 在解析模板时，被嵌套的模板一定要在后面解析。
   
9. > 链式操作

   - 每一次执行完方法之后返回操作的对象本身

     ```go
     type Moto struct {
     	name string
     }
     
     func (m Moto) start() Moto {
     	fmt.Printf("%s 点火\n", m.name)
     	return m
     }
     
     func (m Moto) stop() Moto {
     	fmt.Printf("%s 熄火\n", m.name)
     	return m
     }
     
     func main() {
     	Ducati := Moto{"杜卡迪V4R"}
     	Ducati.start().stop()
     }
     ```

10. > block

    ```go
    {{block "name" pipeline}} T1 {{end}}
    ```

    `block` 是定义模板 `{{define "name"}} T1 {{end}}` 和执行 `{{template "name" pipeline}}` 缩写，经典的用法是定义一组根模板，然后通过在其中重新定义块模板进行自定义。

    示例:

    定义一个根模板 `template/base.html`

    ```html
    <!DOCTYPE html>
    <html lang="zh-CN">
    <head>
        <title>Go Templates</title>
    </head>
    <body>
    <div class="container-fluid">
        {{block "content" . }}{{end}}
    </div>
    </body>
    </html>
    ```

    然后定义一个 `template/index.html`，"继承" `base.html`：

    ```html
    {{template "base.html"}}
    
    {{define "content"}}
    	<div>Hello World</div>	
    {{end}}
    ```

    然后使用 `template.ParseGlob` 按照正则匹配规则解析模板文件，然后通过 `ExecuteTemplate` 渲染指定的模板:

    ```go
    func index(w http.ResponseWriter, r *http.Request) {
    	// 按照正则匹配规则解析模板文件
      tmpl, err := template.ParseGlob("templates/*.html")
      if err != nil {
        fmt.Println("页面加载失败, err:", err)
        return
      }
      err = tmpl.ExecuteTemplate(w, "index.html", nil)
      if err != nil {
        fmt.Println("渲染模板失败, err:", err)
        return
      }
    }
    ```

    如果模板名称冲突了，例如不同业务线下都定义了一个 `index.html` 模板，我们可以通过下面两种方法来解决:

    	1. 在模板文件开头使用 `{{define 模板名}}` 语句显式的为模板命名。

     	2. 可以把模板文件存放在 `templates` 文件夹下面的不同目录中，然后使用 `template.ParseGlob("templates/**/*.html")` 解析模板。

 11. > 修改默认的标识符

     Go标准库的模板引擎使用的花括号 `{{` 和 `}}` 作为标识，而许多前段框架（如 `Vue` 和 `AngularJS`）也使用`{{` 和 `}}`作为标识符，所以当我们同时使用Go语言模板引擎和以上前端框架时就会出现冲突，这个时候我们需要修改标识符，修改前端的或者修改Go语言的。

     这里演示如何修改Go语言模板引擎默认的标识符:

     ```go
     // 创建一个模板对象，修改其模板标识符为 "{[" 和 "]}" ，并且渲染模板文件"./index.html"
     template.New("test").Delims("{[", "]}").ParseFiles("./index.html")
     ```

12. > text/template 与 html/template 的区别

    `html/template` 针对的是需要返回HTML内容的场景，在模板渲染过程中会对一些有风险的内容进行转义，以此来防范跨站脚本攻击。

    例如: 

    ```html
    <!DOCTYPE html>
    <html lang="zh-CN">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta http-equiv="X-UA-Compatible" content="ie=edge">
        <title>Hello</title>
    </head>
    <body>
        {{.}}
    </body>
    </html>
    ```

    这个时候传入一段JS代码并使用 `html/template` 去渲染该文件，会在页面上显示出转义后的JS内容。`<script>alert('嘿嘿嘿')</script>` 这就是 `html/template` 为我们做的事。

    但是在某些场景下，我们如果相信用户输入的内容，不想转义的话，可以自行编写一个safe函数，手动返回一个 `template.HTML` 类型的内容。示例如下: 

    ```go
    func index(w http.ResponseWriter, r *http.Request) {
    	tmpl, err := template.New("index.html").Funcs(template.FuncMap{
    		"safe": func(s string) template.HTML {
    			return template.HTML(s)
    		},
    	}).ParseFiles("./index.html")
    	if err != nil {
    		fmt.Println("创建模板失败, err:", err)
    	}
    	jsStr := "<script>alert('嘿嘿嘿')</script>"
    	err = tmpl.Execute(w, jsStr)
    	if err != nil {
    		fmt.Println(err)
    	}
    }
    ```

    这里只需要在模板文件不需要转义的内容后面使用我们定义好的safe函数就可以了。

    ```HTML
    {{ . | safe }}
    ```
