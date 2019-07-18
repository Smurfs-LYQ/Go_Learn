#### <center>Day03</center>

1. strings包的使用
2. strconv包的使用
3. time包的使用
4. 指针类型的介绍
5. 流程控制( if , switch , for )

#### <center>笔记</center>

1. 在Golang中能使用Label的有 goto, break, continue.

   ​	注意点:

   			* `Label`在continue, break中是`可选的`, 但是在`goto`中是`必须的`
   			* 作用范围: 定义`Label`的函数体内
   			* `Label`可以声明在函数体的任何位置, 不管`Label`声明在`调用点`的前面还是后面.