## <center>流程控制</center>
Go语言中最常用的流程控制有 `if` 和 `for` ，而 `switch` 和 `goto` 主要是为了简化代码、降低重复代码而生的结构，属于扩展类的流程控制。

> ### `if else` (分支结构)
1. `if` 条件判断基本写法
    ```
    age := 21
    
    if age < 18 {
        fmt.Println("这玩意是童工")
    } else if age >= 18 && age <= 60 {
        fmt.Println("兄弟你马上就要退休了")
    } else {
        fmt.Println("终于退休了")
    }
    ```
2. `if` 条件判断特殊写法
    ```
    // 可以在if表达式之前添加一个执行语句，再根据变量值进行判断
    if age := 21;age < 18 {
        fmt.Println("这玩意是童工")
    } else if age >= 18 && age <= 60 {
        fmt.Println("兄弟你马上就要退休了")
    } else {
        fmt.Println("终于退休了")
    }
    ```

> ### `for` (循环结构)
1. `for` 循环的基本格式如下
    ```
    for 初始语句;条件表达式;结束语句 {
        循环体语句
    }
    ```
2. `for` 循环的初始语句可以被忽略，但是初始语句后的分号必须写
    ```
    i := 0
    for ; i < 10; i++ {
        fmt.Println(i)
    }
    ```
3. `for` 循环的初始语句和结束语句都可以省略
    ```
    // 这类写法类似于其他编程中的 while ，在 while 后添加一个条件表达式，满足条件表达式时持续循环，否则结束循环
    i := 0
    for i < 10 {
        fmt.Printf("%d ", i)
        i++
    }
    ```
4. `for` 循环的无限循环
    ```
    for {
        fmt.Println("你好")
    }
    ```
5. `for` 循环可以通过 `continue` 跳出当前循环继续下一次循环，也可以通过 `break` 、`goto`、`return`、`panic` 语句强制退出循环 
    - `continue` 语句可以结束当前循环，开始一下次的循环迭代过程，仅限在 `for` 循环内使用。在 `continue` 语句后添加标签时，表示开始标签对应的循环。
        ```
        for i := 0; i < 5; i++ {
            // 设置退出Label标签
        exitTag:
            for o := 0; o <= i; o++ {
                if o == 3 && i == 3 {
                    fmt.Printf("这里是跳过不输入的地点")
                    // continue到退出标签
                    continue exitTag
                }
                fmt.Printf("%d-%d ", o, i)
            }
            fmt.Println()
        }
        ```
    - `break` 语句可以结束 `for`、`switch` 和 `select` 的代码块。`break` 语句还可以在语句后面添加标签，表示退出某个**标签对应的代码块**，标签要求必须定义在对应的 `for`、`switch` 和 `select` 的代码块上。
        ```
        for i := 0; i < 5; i++ {
        // 设置退出Label标签
        exitTag:
            for o := 0; o < i; o++ {
                if i == 3 {
                    // break到退出标签
                    break exitTag
                }
                fmt.Printf("%d-%d ", o, i)
            }
            fmt.Println()
        }
        ```
    - `goto` 跳转到指定的 `label` 标签。`goto` 语句通过标签进行代码间的无条件跳转。`goto` 语句可以在快速跳出循环、避免重复跳出上有一定的帮助。
        ```
            for i := 0; i < 10; i++ {
                if i > 6 {
                    // goto到退出标签
                    goto exitTag
                }
                fmt.Printf("%d ", i)
            }
            
            // 设置退出labal标签
        exitTag:
            fmt.Printf("\n%s\n", "到6就行了")
        ```

> ### `for range` (键值循环)
1. Go语言中可以使用 `for range` 遍历数组、切片、字符串、map已经通道(channel)。通过 `for range` 遍历的返回值有一下规律: 
    - 数组、切片、字符串返回索引和值
    - map返回键和值
    - 通道(channel)返回通道内的值

> ### `switch case`
1. 使用 `switch` 语句可方便的对大量的值进行条件判断
    ```
    finger := 3
    switch finger {
    case 1:
        fmt.Println("大拇指")
    case 2:
        fmt.Println("食指")
    case 3:
        fmt.Println("中指")
    case 4:
        fmt.Println("无名指")
    case 5:
        fmt.Println("小拇指")
    default:
        fmt.Println("鸡爪？")
    }
    ```
2. 一个分支可以有多个值，多个case值之间使用英文逗号分隔
    ```
    switch i := 9; i {
    case 1, 3, 5, 7, 9:
        fmt.Println(i, "是基数")
    case 2, 4, 6, 8:
        fmt.Println(i, "是偶数")
    default:
        fmt.Println("超出测试范围了")
    }
    ```
3. 分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量
    ```
    age := 21
    switch {
    case age < 18:
        fmt.Println("虽然你未成年，但是你也不能为所欲为")
    case age >= 18 && age <= 60:
        fmt.Println("没权没势的，老实做人吧")
    case age > 60:
        fmt.Println("咱是有素质的人啊，碰瓷那种事咱不能干")
    default:
        fmt.Println("你已超出三界之外，不在无形之中")
    }
    ```
4. `fallthrough` 语法可以无条件执行满足条件的case的下一个case，是为了兼容C语言中的case设计的
    ```
    switch s := "a"; s {
    case "a":
        fmt.Println("a, 如果打印出了a，那么下面应该也会打印出b")
        fallthrough
    case "b":
        fmt.Println("b")
    case "c":
        fmt.Println("c")
    default:
        fmt.Println("...")
    }
    ```
