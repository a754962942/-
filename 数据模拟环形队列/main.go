package main

import (
	"errors"
	"fmt"
	"os"
)

type CircleQueue struct {
	maxSize int
	array   [5]int
	head    int
	tail    int
}

func (c *CircleQueue) Push(val int) (err error) {
	//入队列
	if c.IsFull() {
		return errors.New("queue full")
	}
	//分析出tail在队列尾部，但是不包含最后的元素
	c.array[c.tail] = val
	c.tail = (c.tail+1)%c.maxSize
	return nil
}
func (c *CircleQueue) Pop() (val int, err error) {
	//出队列
	if c.IsEmpty() {
		return -1, errors.New("queue is Empty")
	}
	//head指向队首，并且包含队首元素
	val = c.array[c.head]
	c.head=(c.head+1)%c.maxSize
	return val, nil
}

//判断环形队列是否满了
func (c *CircleQueue) IsFull() bool {
	return (c.tail+1)%c.maxSize == c.head
}

//判断环形队列是否为空
func (c *CircleQueue) IsEmpty() bool {
	return c.tail == c.head
}

//取出环形队列有多少个元素
func (c *CircleQueue) Size() int {
	//关键算法
	return (c.tail + c.maxSize - c.head) % c.maxSize
}

//显示队列
func (c *CircleQueue) ListQueue() {
	//取出当前队列有多少个元素
	size := c.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}
	//设置辅助变量
	tempHead := c.head
	for i := 0; i < size; i++ {
		fmt.Printf("第%d:%d\t", tempHead, c.array[tempHead])
		tempHead = (tempHead + 1) % c.maxSize
	}
}
func main() {
	circleQueue := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}
	var key string
	var val int
	for {
		fmt.Println("1.输入add表示添加数据到队列")
		fmt.Println("2.输入get表示从队列获取数据")
		fmt.Println("3.输入show表示显示队列")
		fmt.Println("4.输入exit表示退出")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			err := circleQueue.Push(val)
			if err == nil {
				fmt.Println("加入队列OK")
			} else {
				fmt.Println(err)
			}
		case "get":
			val, err := circleQueue.Pop()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("你取出的数是：", val)
			}
		case "show":
			circleQueue.ListQueue()
		case "exit":
			os.Exit(0)
		}
	}

}
