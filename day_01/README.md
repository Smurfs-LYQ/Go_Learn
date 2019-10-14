#### <center>Day_01</center>

1. 数据类型
2. 流程控制


#### <center>笔记</center>
1. > 本地查询Go自带的文档方法 1
	```
	godoc -http=:8080
	```
	- 该命令相当于在本地起了一个http服务，等号后面跟的是端口号，打开浏览器访问本机的8080端口就可以查看Go的本地文档了	
2. > 本地查询Go自带的文档方法 2
	```
	// 查询指定包中有哪些方法
	go doc strings

	// 查询指定方法的用法
	go doc strings.HasPrefix
	```
	- go doc 后面跟的是要查询的包名或这指定的方法名
3. 
