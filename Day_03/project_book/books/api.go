package books

type book struct {
	Name    string  // 书名
	Writer  string  // 作者
	Price   float32 // 价格
	Publish bool    // 上架信息
}

// 定义一个切片，长度为0，容量为200
var book_list = make([]*book, 0, 200)
