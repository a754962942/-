//使用数组来模拟一个栈的使用

package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	MaxTop int //栈最大可以存放的个数
	Top    int //表示栈顶，因为栈底固定，直接使用Top
	arr    [5]int
}

func (s *Stack) Push(val int) error {
	//判断栈内是否已满
	if s.Top == s.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	//先自增
	s.Top++
	s.arr[s.Top] = val
	fmt.Println("添加成功")
	return nil
}

//遍历栈，需要从栈顶开始遍历
func (s *Stack) List() {
	//判断栈是否为空
	if s.Top == -1 {
		fmt.Println("stack empty")
		return
	}
	fmt.Println("栈内情况如下")
	for i := s.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]:%d\n", i, s.arr[i])
	}
}
func (s *Stack) Pop() (val int, err error) {
	//判断栈是否为空
	if s.Top == -1 {
		// fmt.Println("stack empty")
		return -1, errors.New("Stack Empty")
	}
	//先取值，再top--
	val = s.arr[s.Top]
	s.Top--
	return val, nil

}
func main() {
	stack := Stack{
		MaxTop: 5,
		Top:    -1,
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	stack.List()

	for {
		val, err := stack.Pop()
		if err != nil {
			break
		}
		fmt.Println(val)
	}
}
