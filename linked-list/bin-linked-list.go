package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type BinNode struct {
	val  int
	pre  *BinNode
	next *BinNode
}

type BinList struct {
	head *BinNode
}

func NewBinList() *BinList {
	return new(BinList)
}

func (this *BinList) Append(v int) {
	node := &BinNode{val:v}
	if this.IsEmpty(){
		this.head = node
		return
	}
	cur := this.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = node
	node.pre = cur
	//fmt.Println("now ",cur.next.pre.val)
}

func (this *BinList) Prepend(v int) {
	head := &BinNode{
		val:  v,
		next: this.head,
	}
	if !this.IsEmpty(){
		fmt.Println(head.val)
		this.head.pre = head
	}
	this.head = head
}

func (this *BinList) Remove(v int) bool {
	if this.IsEmpty() {
		return false
	}
	if this.head.val == v {
		this.head.next.pre = nil
		this.head= this.head.next
		return true
	}
	cur := this.head
	for cur.next != nil {
		if cur.next.val == v {
			cur.next.next.pre = cur
			cur.next = cur.next.next
			return true
		}
		cur = cur.next
	}
	return false
}
func (this *BinList) IsEmpty()bool{
	return this.head == nil
}

func (this *BinList) String() string {
	if this.IsEmpty() {
		return ""
	}
	buf1 := bytes.NewBuffer(nil)
	buf1.WriteString("Nil->")
	cur := this.head
	for cur != nil {
		buf1.WriteString(strconv.Itoa(cur.val) + "->")
		cur = cur.next
	}
	buf1.WriteString("Nil")
	return buf1.String()
}
func (this *BinList) PrintReverse() {
	tail := this.head
	for tail.next!=nil{
		tail = tail.next
	}
	buf1 := bytes.NewBuffer(nil)
	buf1.WriteString("Nil->")
	for tail!=nil{
		//fmt.Println(tail.val)
		buf1.WriteString(strconv.Itoa(tail.val) + "->")
		tail = tail.pre
	}
	buf1.WriteString("Nil")
	fmt.Println(buf1)
}
func main() {
	l := NewBinList()
	l.Append(1)
	l.Append(4)
	//l.Prepend(2)
	//l.Remove(2)
	l.PrintReverse()
	fmt.Println(l)
	l.Append(3)
	l.PrintReverse()
	fmt.Println(l)
}