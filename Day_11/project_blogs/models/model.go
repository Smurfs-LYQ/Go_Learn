package models

// Article 首页文章列表
type Article struct {
	ID int `db:"id"`
	Title string `db:"title"`
	Summary string `db:"summary"`
}

// Article_type 首页文章分类
type Article_type struct {
	ID int `db:"id"`
	Name string `db:"name"`
}