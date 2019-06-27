package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	maxSize int     //队列大小
	array   [10]int //队列
	front   int     //队首
	rear    int     //队尾
}

//添加数据到队列
func (this *Queue) Enqueue(val int) (err error) {
	//判断队列是否满
	if this.rear == this.maxSize-1 {
		return errors.New("queue full")
	}
	this.rear++
	this.array[this.rear] = val
	return
}

//出队
func (this *Queue) Dequeue() (val int, err error) {
	//判断队列是否为空
	if this.front == this.rear {
		return -1, errors.New("queue empty")
	}
	this.front++
	return this.array[this.front], err
}

//打印队列
func (this *Queue) ShowQueue() (err error) {
	fmt.Println("Show Queue: ")
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, this.array[i])
	}
	fmt.Println()
	return
}

func main() {
	queue := &Queue{
		maxSize: 10,
		front:   -1,
		rear:    -1,
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
