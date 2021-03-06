#### <center>Day_09</center>

1. Go连接Mysql
2. Go使用连接池方式连接MySQL
3. 数据库查询操作
4. 数据库插入操作
5. 数据库更新操作
6. 数据库删除操作
7. 数据库预处理-Query()/QueryRow()
8. 数据库预处理-Exec()
9. 数据库事务
10. 数据库-第三方库 sqlx 连接数据库
11. 数据库-第三方库 sqlx 查询
12. 数据库-第三方库 sqlx 增、删、改
13. 数据库-第三方库 sqlx 事务
14. SQL注入
15. Redis 连接
16. Redis 增删查
17. NSQ

#### <center>笔记</center>

1. > 数据库讲解
  
    - 数据库分类:

      - 关系型数据库:
        - MySQL、postGRESQL、SQL Server、Oracle、SQLite
      - 非关系型数据库:
        - MongoDB、Redis、Memcache、ETCD、TiDB

    - 关系型数据库（SQL）

      - 是指采用了**关系模型**来组织数据的数据库，使用行和列的形势存储数据，以便于用户理解，关系型数据库这一系列的行和列被称为表，一组表组成了一个数据库。

    - 非关系型数据库（NoSQL）

    - 常用引擎

      - MyISAM：
        - 不支持事务，只有表级锁，但是查询速度快
      - InnoDB：
        - 支持事务，支持行级锁
      - MySQL支持插件式引擎

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

3. > CRUD

    - CRUD 概念

      ```txt
      计算处理时的增加(Create)、读取(Retrieve)、更新(Update)和删除(Delete)几个单词的首字母简写
      ```

    - **查询-单行查询**

      单行查询 `db.QueryRow()` 执行一次查询，并期望返回最多一行结果(即Row)。参数args表示query中的占位参数。QueryRow总是返回非nil的值，直到返回值的Scan方法被调用时，才会返回被延迟的错误。(如: 未找到结果)

      ```go
      func (db *DB) QueryRow(query string, args ...interface{}) *Row
      ```

      具体示例代码:

      ```go
      func queryRowDemo() {
        sql := "select id, name, age from uer where id=?"
        var u user
        // 非常重要: 确保QueryRow之后调用Scan方法，否则持有的数据库连接不会被释放
        err := db.QueryRow(sql, 1).Scan(&u.id, &u.name, &u.age)
        if err != nil {
          fmt.Printf("scan failed, err:%v\n", err)
          return
        }
        fmt.Println(u)
      }
      ```

    - **查询-多行查询**

      多行查询 `db.Query()` 执行一次查询，返回多行结果(即Rows)，一般用于执行select命令。参数args表示query中的占位参数。

      ```go
      func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
      ```

      具体示例代码:

      ```go
      func queryMultiRowDemo() {
        sql := "select id, name, age from user where id >= ?"
        rows, err := db.Query(sql, 0)
        if err != nil {
          fmt.Printf("query failed, err:%v\n", err)
          return
        }
        // 非常重要：关闭rows释放持有的数据库链接
        defer rows.Close()

        // 循环读取结果集中的数据
        for rows.Next() {
          var u user
          err := rows.Scan(&u.id, &u.name, &u.age)
          if err != nil {
            fmt.Printf("scan failed, err:%v\n", err)
            return
          }
          fmt.Println(u)
        }
      }
      ```

    - **插入数据**

      插入、更新和删除操作都是用方法。

      ```go
      func (db *DB) Exec(query string, args ...interface{}) (Result, error)
      ```

      Exec执行一次命令(包括查询、删除、更新、插入等)，返回的Result是对已执行的SQL命令的总结。参数args表示query中的占位参数。

      具体插入数据示例代码如下:

      ```go
      func insertRowDemo() {
        sql := "insert into user(name, age) values(?,?)"
        ret, err := db.Exec(sql, "Smurfs", 18)
        if err != nil {
          fmt.Printf("insert failed, err:%v\n", err)
          return
        }
        theID, err := ret.LastInsertId() // 新插入数据的id
        if err != nil {
          fmt.Printf("get lastinsert ID failed, err:%v\n", err)
          return
        }
        fmt.Printf("insert success, the id is %d.\n", theID)
      }
      ```

    - **更新数据**

      具体更新数据示例代码如下:

      ```go
      func updateRowDemo() {
        sql := "update user set age=? where id = ?"
        ret, err := db.Exec(sql, 16, 1)
        if err != nil {
          fmt.Printf("update failed, err:%v\n", err)
          return
        }
        n, err := ret.RowsAffected() // 操作影响的行数
        if err != nil {
          fmt.Printf("get RowsAffected failed, err:%v\n", err)
          return
        }
        fmt.Printf("update success, affected rows:%d\n", n)
      }
      ```

    - **删除数据**

      具体删除数据的示例代码如下:

      ```go
      func deleteRowDemo() {
        sql := "delete from user where id = ?"
        ret, err := db.Exec(sql, 1)
        if err != nil {
          fmt.Printf("delete failed, err:%v\n", err)
          return
        }
        n, err := ret.RowsAffected() // 操作影响的行数
        if err != nil {
          fmt.Printf("get RowsAffected failed, err:%v\n", err)
          return
        }
        fmt.Printf("delete success, affected rows:%d\n", n)
      }
      ```

4. > MySQL预处理

    - **什么是预处理**
      - 普通SQL语句执行过程:
        - 客户端对SQL语句进行占位符替换得到完整的SQL语句。
        - 客户端发送完整SQL语句到MySQL服务端。
        - MySQL服务端执行完整的SQL语句并将结果返回给客户端。
      - 预处理执行过程:
        - 把SQL语句分为两部分，命令部分与数据部分。
        - 先把命令部分发送给MySQL服务端，MySQL服务端进行SQL预处理。
        - 然后把数据部分发送给MySQL服务端，MySQL服务端对SQL语句进行占位符替换。
        - MySQL服务端执行完整的SQL语句并将结果返回给客户端。

    - **为什么要预处理**
      - 优化MySQL服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本。
      - 避免SQL注入问题。

5. > Go实现MySQL预处理

    - `Prepare` 方法会先将sql语句发送给MySQL服务端，返回一个准备好的状态用于之后的查询和命令。返回值可以同时执行多个查询和命令。

      ```go
      func (db *DB) Prepare(query string) (*Stmt, error)
      ```

    - 查询操作的预处理 [示例demo](https://github.com/Smurfs-LYQ/Go_Learn/blob/master/Day_09/07_MySQL_Prepare_Query/main.go)

    - 插入、更新和删除的操作与预处理十分类似 [示例demo](https://github.com/Smurfs-LYQ/Go_Learn/blob/master/Day_09/08_MySQL_Prepare_Exec/main.go)

6. > Go实现MySQL事务

    - **什么是事务？**

      事务: 一个最小的不可再分的工作单元；通常一个事务对应一个完整的业务

      在MySQL中只有使用了 `Innodb` 数据库引擎的数据库或表才支持事务。事务处理可以用来维护数据库的完整性，保证成批的SQL语句要么全部执行，要么全部不执行。

    - **事务的ACID**

      通常事务必须满足4个条件 (ACID) :

        - 原子性(Atomicity, 或称不可分割性)

          ```txt
          一个事务(transaction)中的所有操作，要么全部完成，要么全部不完成，不会结束在中间某个环节。事务在执行过程中发生错误，会被回滚(Rollback)到事务开始前的状态，就像这个事务从来没有执行过一样。
          ```

        - 一致性(Consistency)

          ```txt
          在事务开始之前和事务结束之后，数据库的完整性没有被破坏。这表示写入的资料必须完全符合所有的预设规则，这包含资料的准确度、串联性以及后续数据库可以自发性的完成预定的任务。
          ```

        - 隔离性(Isolation, 又称独立性)

          ```txt
          数据库有允许多个并发事务同时对其数据进行读写和修改的能力，隔离性可以防止多个事务并发执行时由于交叉执行而导致数据的不一致。
          事务隔离分为不同级别，包括读未提交(Read uncommitted)、读提交(read committed)、可重复读(repeatable read)和串行化(Serializable)。
          ```

        - 持久性(Durability)

          ```txt
          事务处理结束后，对数据的修改就是永久的，即便系统故障也不回丢失。
          ```

    - **事务相关方法**

      Go语言使用一下三个方法实现MySQL中的事务操作。

        - 开始事务

          ```go
          func (db *DB) Begin() (*Tx, error)
          ```

        - 提交事务

          ```go
          func (tx *Tx) Commit() error
          ```

        - 回滚事务

          ```go
          func (tx *Tx) Rollback() error
          ```

    - **事务demo** [github](https://github.com/Smurfs-LYQ/Go_Learn/blob/master/Day_09/09_MySQL_Trans/main.go)

7. > sqlx使用

    - 第三方库 `sqlx` 能简化操作，提高开发效率

    - **安装**

      ```go
      go get github.com/jmoiron/sqlx
      ```

    - **基本使用**

    - 连接数据库 [示例demo](https://github.com/Smurfs-LYQ/Go_Learn/blob/master/Day_09/10_MySQL_sqlx_Connect/main.go)

    - 查询 [示例demo](https://github.com/Smurfs-LYQ/Go_Learn/blob/master/Day_09/11_MySQL_sqlx_Get/main.go)

    - 插入、更新和删除 [示例demo](https://github.com/Smurfs-LYQ/Go_Learn/blob/master/Day_09/12_MySQL_sqlx_Exec/main.go)
      - sqlx中的exec方法与原生sql中的exec使用基本一致

    - 事务操作 [示例demo](https://github.com/Smurfs-LYQ/Go_Learn/blob/master/Day_09/13_MySQL_sqlx_Trans/main.go)
      - 对于事务操作，我们可以使用 `sqlx` 中提供的 `db.Beginx()` 和 `db.MustExec()` 方法来简化错误处理过程。

8. > Redis

    - Redis支持的数据结构
      - STRING（字符串）
      - LIST（列表）
      - SET（集合）
      - HASH（哈希）
      - ZSET（有序集合）

    - Redis应用场景
      - 缓存系统，减轻主数据库(MySQL)的压力
      - 计数场景，比如微博、抖音中的关注数和粉丝数
      - 热门排行榜，需要排序的场景特别适合使用ZSET
      - 利用LIST可以实现队列的功能。

    - Redis与Memcached比较
      - Memcached中的值只支持简单的字符串，Redis支持更丰富的5中数据结构类型。
      - Redis的性能比Memcached好很多。
      - Redis支持RDB持久化和AOF持久化。
      - Redis支持master/slave模式。
9. > Go操作Redis

    - 安装

      Go语言中使用第三方库 https://github.com/go-redis/redis 连接Redis连接数据库并进行操作。

      ```go
      go get -u github.com/go-redis/redis
      ```

    - 连接 [demo](https://github.com/Smurfs-LYQ/Go_Learn/blob/master/Day_09/15_Redis_connect/main.go)

    - 设置值 获取值 删除值 [demo](https://github.com/Smurfs-LYQ/Go_Learn/blob/master/Day_09/16_Redis_operation/main.go)

#### <center>注意事项</center>

1. > SQL中的占位符
    - 不同的数据库中，SQL语句使用的占位符语法不尽相同。

      | 数据库      | 占位符语法 |
      | ---------- | ---------- |
      | MySQL      | ?          |
      | postGRESQL | $1, $2等   |
      | SQLite     | ? 和 $1    |
      | Oracle     | :name      |

2. > SQL注入
    - **任何时候都不应该自己拼接SQL语句！**

      示例演示:

      ```go
      // sql注入示例
      func sqlInjectDemo(name string) {
        sqlStr := fmt.Sprintf("select id, name, age from user where name='%s'", name)
        fmt.Printf("SQL:%s\n", sqlStr)

        var users []user
        err := db.Select(&users, sqlStr)
        if err != nil {
          fmt.Printf("exec failed, err:%v\n", err)
          return
        }
        for _, u := range users {
          fmt.Printf("user:%#v\n", u)
        }
      }
      ```

      此时一下输入字符串都可以引发SQL注入问题:

      ```go
      sqlInjectDemo("xxx' or 1=1#")
      sqlInjectDemo("xxx' union select * from user #")
      sqlInjectDemo("xxx' and (select count(*) from user) <10 #")
      ```