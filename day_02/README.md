#### <center>Day02</center>

1. 给一个参数n，列出那些数字组合相加等于n
2. 一个程序包含两个包add和main，其中add包中有两个变量，Name和Age。请问main包中怎么访问Name和Age (也可以使用包别名的方式来做)
3. 每个源文件都可以包含一个init函数(初始化函数)，这个init函数自动被Go运行框架调用。
4. 常量的声明，以及 iota 的用法实例
5. 
#### <center>笔记</center>
1. 常量
   1. 格式: const 常量名 类型 = 值 (类型是可以省略的)
   2. 常量是const修饰，代表永远是只读
   3. const只能修饰boolean、number(int相关类型,浮点类型,complex)和string
2. iota
    * 它的主要作用是从0开始逐1递增，04_const 中有实例