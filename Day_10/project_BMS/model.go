package main

// 专门用来定义与数据对应的结构体
type book struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
	Price string `db:"price"`
}
