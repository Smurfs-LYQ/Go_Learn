package article

import "time"

// Article 文章 redis表名 article:文章ID
type Article struct {
	Title  string `json:"title"`// 标题
	Url    string `json:"link"`// 连接
	Poster string `json:"poster"`// 作者
	Time   string `json:"time"`// 发布时间
	Votes  string `json:"votes"`// 票数
}

// Time 文章发布时间 redis表名 time:
type Time struct {
	Title string     // 文章标题
	ID    string     // 文章id
	Time  *time.Time // 文章发布时间
}

// Source 文章的票数 redis表名 source:
type Source struct {
	Title string // 文章标题
	ID    int    // 文章id
	Votes int    // 票数
}

// Voted 投票了的用户 redis表名 voted:文章ID
type Voted struct {
	ID int // 文章的ID
}
