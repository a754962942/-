package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	maxSize int
	array   [5]int //数组——>模拟队列
	front   int    //表示指向队列首(不含第一位)
	rear    int    //表示指向队列尾
}

//添加数据到队列尾
func (this *Queue) AddQueue(val int) error {
	//先判断队列是否已满
	if this.rear == this.maxSize-1 {
		//rear时队列尾部(含最后元素)
		return errors.New("queue full")
	}
	this.rear++ //rear 后移
	this.array[this.rear] = val

	return nil
}

//显示队列，找到队首，然后遍历到队尾
func (this *Queue) ShowQueue() {
	//front表示指向队列首(不包含第一位)
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array%d为:\t%d\n", i, this.array[i])

	}

}

//从队列中取出数据
func (this *Queue) GetQueue() (val int, err error) {
	//先判断队列是否为空
	if this.rear == this.front {
		return -1, nil
	} else {
		this.front++
		val := this.array[this.front]
		return val, err
	}
}
func main() {
	queue := Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
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
			err := queue.AddQueue(val)
			if err == nil {
				fmt.Println("加入队列OK")
			} else {
				fmt.Println(err)
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("你取出的数是：", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
