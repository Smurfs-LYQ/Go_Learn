package main

import "fmt"

// 封装单个Link
type LinkNode struct {
	// 这里用interface是因为空接口里面可以保存任意类型的数据
	Data interface{}
	// 下一个Link的地址
	Next *LinkNode
}

// 创建一个链表类
type Link struct {
	// 链表的头
	Head *LinkNode
	// 链表的尾
	Tail *LinkNode
}

// 在链表的头部插入一个Link
func (p *Link) InsertHead(data interface{}) {
	// 首先创建一个Link实例
	node := &LinkNode{Data: data, Next: nil}

	// 将现在链表头的第一个Link的地址赋值给node的Next
	node.Next = p.Head
	// 将链表类的Head赋值为当前node的地址
	p.Head = node
}

// 在链表的尾部插入一个Link
func (p *Link) InsertTail(data interface{}) {
	// 首先创建一个Link实例
	node := &LinkNode{Data: data, Next: nil}

	if p.Tail == nil {
		p.Head.Next = node
		p.Tail = node
	} else {
		// 将Link链表类中的Tail元素(最后一个LinkNode)的Next元素复制为node的地址
		p.Tail.Next = node
		// 将Link链表类中的Tail元素(最后一个LinkNode)复制为node的地址
		p.Tail = node
	}
}

func Put(node *LinkNode) {
	fmt.Println(node.Data)
	if node.Next != nil {
		Put(node.Next)
	}
}

func (link *Link) Trans() {
	p := link.Head
	for p != nil {
		fmt.Println(p.Data)
		p = p.Next
	}
}

func main() {
	list := Link{}
	for i := 0; i < 5; i++ {
		if list.Head == nil {
			list.InsertHead(fmt.Sprintf("链表-%d", i))
		} else {
			list.InsertTail(fmt.Sprintf("链表-%d", i))
		}
	}
	// Put(list.Head)
	list.Trans()
}
