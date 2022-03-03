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

//Push
func (c *CircleQueue) Push(val int) error {
	if c.IsFull() {
		return errors.New("队列已满")
	}
	c.array[c.tail] = val
	c.tail = (c.tail + 1) % c.maxSize
	return nil
}

//Pop
func (c *CircleQueue) Pop() (int, error) {
	if c.isEmpty() {
		return -1, errors.New("队列为空")

	}
	val := c.array[c.head]
	c.head = (c.head + 1) % c.maxSize
	return val, nil
}

//IsFull
func (c *CircleQueue) IsFull() bool {
	return (c.tail+1)%c.maxSize == c.head
}

//IsEmpty
func (c *CircleQueue) isEmpty() bool {
	return c.tail == c.head
}

//Listqueue
func (c *CircleQueue) ListQueue() {
	size := c.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}
	//辅助变量
	temp := c.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]:%d\t", temp, c.array[temp])
		temp = (temp + 1) % c.maxSize
	}
	fmt.Println()
}

// Size
func (c *CircleQueue) Size() int {
	return (c.tail + c.maxSize - c.head) % c.maxSize
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
