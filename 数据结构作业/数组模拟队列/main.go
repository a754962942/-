//约瑟夫问题

//创建一个数组模拟队列，每隔一段时间[随机的]，给数组添加一个数
//启动两个协程，没隔一个时间取一个数据
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type arr struct {
	maxSize int
	array   [5]int
	front   int //指向队列最前
	rear    int //指向队列最后
}

func (a *arr) AddArr(val int) error {
	if a.IsFull() {
		err := errors.New("队列已满")
		return err
	} else {
		a.rear++
		a.array[a.rear] = val
		fmt.Println("添加成功")
	}
	return nil
}
func (a *arr) ListArr() {
	for i := a.front + 1; i <= a.rear; i++ {
		fmt.Printf("第%d:%d\n", i, a.array[i])

	}
}
func (a *arr) GetArr(i int) {
	if a.IsEmpty() {
		errors.New("队列为空，取出失败")
		return
	}
	for {
		if a.front >= a.maxSize-1 {
			wg.Done()
			break
		}
		var sync sync.Mutex
		val := a.array[a.front+1]
		sync.Lock()
		fmt.Printf("协程：%d,第%d:%d\n", i, a.front+1, val)
		sync.Unlock()
		sync.Lock()
		a.front++
		sync.Unlock()

	}
}
func (a *arr) IsFull() bool {
	if a.rear == a.maxSize-1 {
		return true
	}
	return false
}
func (a *arr) IsEmpty() bool {
	if a.rear == a.front {
		return true
	}
	return false
}

var wg sync.WaitGroup

func main() {

	queue := arr{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}
	// var key string
	for {
		RandNum := rand.Intn(10)
		Randtime := rand.Intn(3) * int(time.Second)
		time.Sleep(time.Duration(Randtime))
		err := queue.AddArr(RandNum)
		if err != nil {
			fmt.Println(err)
			queue.ListArr()
			break
		}
	}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go queue.GetArr(i)
	}
	wg.Wait()

}
