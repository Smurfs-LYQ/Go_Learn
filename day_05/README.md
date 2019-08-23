#### <center>Day05</center>

1. struct 结构体
2. 链表定义
3. 二叉树
4. 结构体
    * struct和数据类型的别名
    * 工厂模式
5. struct中的tag
6. struct中的匿名字段
7. struct中的方法
8. struct中的继承
9. struct中的组合
10. 接口示例：如果一个变量实现了String()这个方法，那么fmt.Println()默认会调用这个变量的String()进行输出
11. 接口

#### <center>笔记</center>
1. > struct
    * Go语言中没有像其他语言中的class类，但是struct和class的功能差不多，可以把struct当做class类来使用
2. > 函数和struct中方法的区别
    * 函数调用：函数名(参数列表)
    * 方法调用：结构体变量名.函数名(参数列表)
3. > struct中的方法访问控制
    * 和struct中的字段元素一样，通过大小写控制，如果首字母小写，外界则无法访问
4. > struct中组合和继承的区别
    * ```有名嵌套```为组合
    * ```匿名嵌套```为继承
5. > 强调
    * ```make``` 用来创建map、slice、channel等引用类型
    * ```new``` &nbsp; 用来创建值类型
