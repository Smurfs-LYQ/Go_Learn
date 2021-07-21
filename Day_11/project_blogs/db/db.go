package db

import (
	"Go_Learn/Day_11/project_blogs/models"

	_ "github.com/go-sql-driver/mysql" // 导入MySQL数据库驱动
	"github.com/jmoiron/sqlx"          // 导入第三方标志库
)

var DB *sqlx.DB

// InitDB 初始化连接数据库
func InitDB() (err error) {
	dsn := "root:smurfs@tcp(127.0.0.1:3306)/go_test_blogs"

	// 连接数据库，执行了open和ping的操作
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}

	// 设置数据库最大连接数
	DB.SetMaxOpenConns(10)
	// 设置数据库最大空闲连接数
	DB.SetMaxIdleConns(5)

	return
}

// Index_articles 首页文章列表
func Index_articles() (article_list []models.Article, err error) {
	sql := "select id, title, summary from article where status = ?"

	// 数据库查询
	err = DB.Select(&article_list, sql, 1)
	if err != nil {
		return
	}

	return
}

// Index_article_type 首页文章类型
func Index_article_type() (article_type []models.Article_type, err error) {
	sql := "select id, name from type"

	// 数据库查询
	err = DB.Select(&article_type, sql)
	if err != nil {
		return
	}

	return
}