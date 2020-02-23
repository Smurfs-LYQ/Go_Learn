package test

import "fmt"

type LinkNode struct {
	data interface{}
	next *LinkNode
}

type Link struct {
	head *LinkNode
	tail *LinkNode
}

func (this *Link) InsertHead(data interface{}) {
	node := &LinkNode{
		data: data,
		next: nil,
	}

	if this.head == nil && this.tail == nil {
		this.head = node
		this.tail = node
		return
	}

	node.next = this.head
	this.head = node
}

func (this *Link) InsertTail(data interface{}) {
	node := &LinkNode{
		data: data,
		next: nil,
	}

	if this.head == nil && this.tail == nil {
		this.head = node
		this.tail = node
		return
	}

	this.tail.next = node
	this.tail = node
}

func (p *Link) Trans() {
	q := p.head
	for q != nil {
		fmt.Println(q.data)
		q = q.next
	}
}

func Test() {
	var intLink Link
	for i := 0; i < 10; i++ {
		intLink.InsertHead(i)
		intLink.InsertTail(i)
	}

	intLink.Trans()
}
