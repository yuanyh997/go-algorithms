package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type Node struct {
	val  int
	next *Node
}
type List struct {
	head *Node
}

func NewList() *List {
	return new(List)
}

func (this *List) Append(v int) {
	node := &Node{val: v}
	if this.IsEmpty() {
		this.head = node
		return
	}
	cur := this.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = node
}

func (this *List) Prepend(v int) {
	head := &Node{
		val:  v,
		next: this.head,
	}
	this.head = head
}

func (this *List) Remove(v int) bool {
	if this.IsEmpty() {
		return false
	}
	if this.head.val == v {
		this.head = this.head.next
		return true
	}
	cur := this.head
	for cur.next != nil {
		if cur.next.val == v {
			cur.next = cur.next.next
			return true
		}
		cur = cur.next
	}
	return false
}

func (this *List) IsEmpty() bool {
	return this.head == nil
}

func (this *List) String() string {
	if this.IsEmpty() {
		return ""
	}

	buf := bytes.NewBuffer(nil)
	cur := this.head
	for cur != nil {
		buf.WriteString(strconv.Itoa(cur.val) + "->")
		cur = cur.next
	}
	buf.WriteString("Nil")
	return buf.String()
}

func main() {
	l := NewList()
	l.Append(1)
	l.Append(4)
	//l.Prepend(2)
	//l.Remove(2)
	l.Append(3)
	fmt.Println(l)
}
