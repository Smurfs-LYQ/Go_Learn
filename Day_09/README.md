#### <center>Day_08</center>

1. Go连接Mysql
2. Go使用连接池方式连接MySQL

#### <center>笔记</center>
1. > 数据库讲解
  
    - 数据库分类:
   
      ​	关系型数据库:
   
      ​		MySQL、postGRESQL、SQL Server、Oracle、SQLite
   
      ​	非关系型数据库:
   
      ​		MongoDB、Redis、Memcache、ETCD、TiDB
   
    - 关系型数据库（SQL）
   
      ​	是指采用了**关系模型**来组织数据的数据库，使用行和列的形势存储数据，以便于用户理解，关系型数据库这一系列的行和列被称为表，一组表组成了一个数据库。
   
    - 非关系型数据库（NoSQL）
   
      ​	
   
    - 常用引擎
   
      ​	MyISAM：
   
      ​		不支持事务，只有表级锁，但是查询速度快
   
      ​	InnoDB：
   
      ​		支持事务，支持行级锁
   
      ​	MySQL支持插件式引擎
   
 2. > Go操作MySQL

    - 连接

      ​	Go语言中的 `database/sql` 包提供了保证SQL或类SQL数据库的泛用接口，并不提供具体的数据库驱动。使用 `database/sql` 包时必须注入（至少）一个数据库驱动。例如: [第三方](https://github.com/go-sql-driver/mysql)

    - 下载依赖

      ```go
      go get -u github.com/go-sql-driver/mysql
      // -u: update更新
      ```
    - 使用MySQL驱动

      ```go
  func Open(driverName, dataSourceName string) (*DB, error)
      ```
    
       Open打开一个dirverName指定的数据库，dataSourceName指定数据源，一般至少包括数据库文件名和（可能的）连接信息

       [示例代码](https://github.com/Smurfs-LYQ/Go_Learn/blob/master/Day_09/01_MySQL/main.go)

    - 初始化连接

       Open函数可能只是验证其参数，而不创建与数据库的连接。如果要检查数据源的名称是否合法，应调用返回值的Ping的方法。

       返回的DB可以安全的被多个goroutine同时使用，并会维护自身的闲置连接池。这样一来，Open函数只需调用一次。很少需要关闭DB。

       [示例代码_1](https://github.com/Smurfs-LYQ/Go_Learn/blob/master/Day_09/01_MySQL/main.go)

       [示例代码_2](https://github.com/Smurfs-LYQ/Go_Learn/blob/master/Day_09/02_MySQL/main.go)

       其中 `sql.DB` 是一个数据库（操作）句柄，代表一个具有零到多个底层连接的连接池。它可以安全的被多个go程序同时使用。`database/sql` 包会自动创建和释放连接；它也会维护一个闲置连接的连接池。

    - **SetMaxOpenConns**

      ```go
  func (db *DB) SetMaxOpenConns(n int)
      ```
    
       `SetMaxOpenConns` 设置与数据库建立连接的最大数目。如果n大于0且小于最大闲置连接数，会将最大闲置连接数减小到匹配最大开启连接数的限制。如果n<=0，不会限制最大开启连接数，默认为0（无限制）

    - **SetMaxIdleConns**

      ```go
  func (db *DB) SetMaxIdleConns(n int)
      ```
    
       SetMaxIdleConns设置连接池中最大闲置连接数。如果n大于最大开启连接数，则新的最大闲置连接数会减小到匹配最大开启连接数的限制。如果n<=0，不会保留闲置连接。

    - sd

3. > 