package books_1

type book struct {
	Name    string
	Writer  string
	Price   float32
	Publish bool
}

var book_list = make([]*book, 0, 10)
