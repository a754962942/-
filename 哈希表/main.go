package main

import (
	"fmt"
	"os"
)

//定义EMP
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

func (e *Emp) ShowMe() {
	fmt.Printf("链表:%d,Id:%d,Name:%s\n->", e.Id%7, e.Id, e.Name)
}

//定义EmpLink
//EmpLink不带表头，即第一个结点就存放雇员
type EmpLink struct {
	Head *Emp
}

//定义HashTable
type HashTable struct {
	LinkArr [7]EmpLink
}

//方法特定
//1.添加员工方法，保证添加时编号从小到大
func (e *EmpLink) Insert(emp *Emp) {
	cur := e.Head      //这是一个辅助指针
	var pre *Emp = nil //这是一个辅助指针pre 在cur前边
	//如果当前Emplink就是一个空链表
	if cur == nil {

		e.Head = emp
		return
	}
	//如果不是一个空链表，给emp找到对应的位置并插入
	//让cur和emp比较，让pre保持在cur前面
	for {
		if cur != nil {
			//比较
			if cur.Id >= emp.Id {
				//找到位置
				break
			}
			//保证同步
			pre = cur
			cur = cur.Next
		} else {
			//如果找不到位置
			break
		}
	}
	pre.Next = emp
	emp.Next = cur

}

//给HashTable编写insert雇员的方法
func (h *HashTable) Insert(emp *Emp) {
	//使用散列函数，确定将该雇员添加到那个链表
	linkNo := h.HashFun(emp.Id)
	//使用对应的链表添加
	h.LinkArr[linkNo].Insert(emp)
}

//编写一个散列方法
func (h *HashTable) HashFun(id int) int {
	//得到一个值，这个值就是我们要添加的链表下标
	return id % 7
}

//显示当前链表的信息
func (e *EmpLink) ShowLink(no int) {
	if e.Head == nil {
		fmt.Printf("链表%d为空\n", no)
		return
	}
	//遍历当前列表，并显示数据
	cur := e.Head //辅助指针
	for {
		if cur != nil {
			fmt.Printf("链表:%d,Id:%d,Name:%s\n", no, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println()
}

//根据Id查找对应的雇员，如果没有就返回nil
func (e *EmpLink) FindById(id int) *Emp {
	cur := e.Head
	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}

//编写方法，现实的时hashtable的所有雇员
func (h *HashTable) ShowAll() {
	for i := 0; i < len(h.LinkArr); i++ {
		h.LinkArr[i].ShowLink(i)
	}
}

//编写一个方法，完成查找
func (h *HashTable) FindById(id int) *Emp {
	//使用散列函数，确定该雇员在那个表内
	linkNo := h.HashFun(id)
	return h.LinkArr[linkNo].FindById(id)

}
func main() {
	key := ""
	id := 0
	name := ""
	var hashtable HashTable
	for {
		fmt.Println("=========雇员系统菜单栏==========")
		fmt.Println("input 表示添加雇员")
		fmt.Println("show 表示添加雇员")
		fmt.Println("find 表示添加雇员")
		fmt.Println("exit 表示添加雇员")
		fmt.Scan(&key)
		switch key {
		case "input":
			fmt.Print("请输入雇员Id:")
			fmt.Scan(&id)
			fmt.Println("")
			fmt.Print("请输入雇员姓名：")
			fmt.Scan(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashtable.Insert(emp)
		case "find":
			fmt.Println("请输入Id")
			fmt.Scan(&id)

			emp := hashtable.FindById(id)
			if emp == nil {
				fmt.Printf("Id=%d的雇员不存在\n", id)
			} else {
				//编写一个方法，编写雇员信息
				emp.ShowMe()
			}
		case "show":
			hashtable.ShowAll()
		case "exit":
			os.Exit(0)
		}
	}
}
