package main

import (
	"bytes"
	"fmt"
	"strconv"
)

//Node定义
type CricleNode struct {
	val  int
	next *CricleNode
}
type CricleList struct {
	head *CricleNode
}

func NewCricleList() *CricleList{
	return new(CricleList)
}

func (this *CricleList) Append(v int) {
	node := &CricleNode{val: v}
	if this.IsEmpty() {
		node.next = node
		this.head = node
		//fmt.Println(this.head.val)
		return
	}
	cur := this.head
	for cur.next != this.head {
		cur = cur.next
	}
	cur.next = node
	node.next = this.head
}

func (this *CricleList) Prepend(v int) {
	head := &CricleNode{
		val:  v,
		next: this.head,
	}
	//修改最后一个节点的next值为新head
	cur := this.head
	for cur.next != this.head {
		cur = cur.next
	}
	cur.next = head
	this.head = head
}

func (this *CricleList) Remove(v int) bool {
	if this.IsEmpty() {
		return false
	}
	if this.head.val == v {
		//找到队尾元素
		tail := this.head
		for tail.next!=this.head{
			tail = tail.next
		}
		//fmt.Println("队尾为",tail.val)
		this.head = this.head.next
		tail.next = this.head
		return true
	}
	//找到元素对应位置
	cur := this.head
	for cur.next != this.head {
		if cur.next.val == v {
			cur.next = cur.next.next
			return true
		}
		cur = cur.next
	}
	return false
}

func (this *CricleList) IsEmpty() bool {
	return this.head == nil
}

func (this *CricleList) String() string {
	if this.IsEmpty() {
		return ""
	}
	buf := bytes.NewBuffer(nil)
	cur := this.head
	buf.WriteString("head->"+strconv.Itoa(cur.val) + "->")
	for cur.next != this.head {
		cur = cur.next
		buf.WriteString(strconv.Itoa(cur.val) + "->")
	}
	buf.WriteString("head")
	return buf.String()
}

func main() {
	l := NewCricleList()
	l.Append(1)
	//fmt.Println(l)
	l.Append(4)
	l.Prepend(2)
	fmt.Println(l)
	l.Remove(4)
	l.Append(5)
	fmt.Println(l)
	l.Prepend(6)
	l.Append(3)
	fmt.Println(l)

}