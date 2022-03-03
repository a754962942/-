//使用数组来模拟一个栈的使用

package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Stack struct {
	MaxTop int //栈最大可以存放的个数
	Top    int //表示栈顶，因为栈底固定，直接使用Top
	arr    [5]int
}

//push入栈
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

//pop出栈
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

//计算
func Calc(num1, num2, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 4:
		res = num2 / num1
	default:
		fmt.Println("运算符有误")
	}
	return res
}

//优先级
func Peiority(val int) int {
	res := 0
	if val == 42 || val == 47 {
		res = 1
	} else if val == 43 || val == 45 {
		res = 0
	}
	return res
}

//判断是不是运算符
func Isoper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}
func main() {
	//numStack
	numStack := Stack{
		MaxTop: 20,
		Top:    -1,
	}
	//operStack
	operStack := Stack{
		MaxTop: 20,
		Top:    -1,
	}
	exp := "3+2*6-2"
	//定义一个index，帮助扫描exp
	index := 0
	//配合运算的变量
	num1 := 0
	num2 := 0
	oper := 0
	result := 0
	for {
		ch := exp[index : index+1] //字符串，只含有一个字符
		temp := int([]byte(ch)[0]) //字符对应的ASCII码
		if Isoper(temp) {          //说明是符号
			//如果operStack 是一个空栈，直接入栈
			if operStack.Top == -1 {
				operStack.Push(temp)
			} else {
				//如果operStack不是空栈
				//判断运算符优先级
				//从numStackPop出两个数进行运算然后再push回NumStack
				if Peiority(operStack.arr[operStack.Top]) >= Peiority(temp) {
					// 判断当前拿到的运算符优先级
					// 是否高于栈内已经存在的运算符的优先级
					// 如果高于，那么计算
					// 如果不高于，那么继续压栈
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = operStack.Pop()

					//将计算结果重新压入栈内
					numStack.Push(result)
					// 当前的符号压入符号栈
					operStack.Push(temp)
				} else {
					operStack.Push(temp)
				}
			}

		} else {
			val, _ := strconv.ParseInt(string(temp), 10, 64)

			numStack.Push(int(val))
		}
		//继续扫描
		//先判断index是否已经扫描到计算表达式的最后
		if index > len(exp)-1 {
			break
		}
		index++
	}
	//如果扫描表达式完毕，依次从符号栈取出符号，然后从数栈取出两个数，
	//运算后的结果，入数栈，直到符号栈为空
	for {
		if operStack.Top == -1 {
			break
		}
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()
		result = Calc(num1, num2, oper)
		//结果重新入栈
		numStack.Push(result)
	}
	res, _ := numStack.Pop()
	fmt.Printf("表达式为%s = %d", exp, res)
}
