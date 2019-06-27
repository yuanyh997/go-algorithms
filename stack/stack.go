package main

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

//数组模拟栈
type Stack struct {
	MaxTop int
	Top int
	arr [5]int
}

//入栈
func (this *Stack) Push(val int) (err error){
	//判断栈满
	if this.Top == this.MaxTop{
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	this.Top++
	this.arr[this.Top] = val
	//fmt.Println(this.Top)
	return
}

//出栈
func (this *Stack) Pop() (val int,err error) {
	if this.Top==-1{
		return -1,errors.New("stack empty")
	}
	val = this.arr[this.Top]
	this.Top--
	return
}

//遍历栈
func (this *Stack) List() {
	if this.Top==-1{
		fmt.Println("stack empty")
		return
	}
	buf := bytes.NewBuffer(nil)
	for i:=this.Top;i>=0;i--{
		buf.WriteString(strconv.Itoa(this.arr[i]) + "->")
	}
	buf.WriteString("bottom")
	fmt.Println(buf.String())
}
func main(){
	stack := &Stack{
		MaxTop: 4,
		Top:    -1,
	}
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(4)

	fmt.Println(stack.Pop())
	stack.List()

}