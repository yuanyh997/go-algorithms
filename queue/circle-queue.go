package main

import (
	"errors"
	"fmt"
	"os"
)

//需要在队尾空出一个位置，即实际容量为maxSize-1
type CircleQueue struct {
	maxSize int    //队列大小
	array   [5]int //队列
	head    int    //队首，指向下一个出队的位置
	tail    int    //队尾，指向下一个入队的位置
}

//添加数据到队列
func (this *CircleQueue) Enqueue(val int) (err error) {
	if this.IsFull() {
		return errors.New("queue full")
	}
	this.array[this.tail] = val
	this.tail = (this.tail + 1) % this.maxSize
	return
}

//出队
func (this *CircleQueue) Dequeue() (val int, err error) {
	if this.IsEmpty() {
		return -1, errors.New("queue empty")
	}
	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return
}

//打印队列
func (this *CircleQueue) ShowQueue() (err error) {
	fmt.Println("Show Queue: ")
	size := this.Size()
	fmt.Println(this.tail, size)
	if size == 0 {
		fmt.Println("队列为空")
	}
	tempHead := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("array[%d]=%d\t", tempHead, this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxSize
	}
	fmt.Println()
	return
}

func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}

func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}

func (this *CircleQueue) Size() (size int) {
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

func main() {
	queue := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}

	for i := 0; i < 4; i++ {
		queue.Enqueue(i)
		queue.ShowQueue()
	}
	var key string
	var val int
	for {
		fmt.Println("1. 输入 add  表示添加数据到队列")
		fmt.Println("2. 输入 get  表示从队列获取数据")
		fmt.Println("3. 输入 show 表示显示队列")
		fmt.Println("4. 输入 exit 表示显示队列")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队的数字")
			fmt.Scanln(&val)
			err := queue.Enqueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列 ok")
			}
		case "get":
			val, err := queue.Dequeue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出的数字为：", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
