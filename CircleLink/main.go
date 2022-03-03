package main

import "fmt"

type CatNode struct {
	no   int      //猫猫的编号
	name string   //猫猫的名字
	next *CatNode //猫猫下一个指向
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) *CatNode {
	//判断是不是添加的第一只猫猫
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head //构成一个环形
		fmt.Println(newCatNode.no, "加入到环形的链表")
		return head
	}
	//如果不是添加的第一支猫猫
	//那么定义一个临时变量，找到环形最后的结点
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	temp.next = newCatNode
	newCatNode.next = head

	return newCatNode
}

//输出环形链表
func ListCat(head *CatNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("环形链表为空")
		return
	}
	for {
		if temp.next == head {
			fmt.Printf("猫的信息为[id=%d,name=%s]->指向[%#v]\n", temp.next.no, temp.next.name, temp.next.next)
			break
		}
		fmt.Printf("猫的信息为[id=%d,name=%s]->指向[%#v]\n", temp.next.no, temp.next.name, temp.next.next)
		temp = temp.next
	}
}

//删除一只猫
func DeleteCat(head *CatNode, id int) *CatNode {
	//思路
	//先让temp指向head
	//定义一个helper指向环形链表最后
	//让temp和要删除的id进行比较，如果相同，则同helper完成删除
	//必须考虑如果删除的就是头结点
	temp := head
	helper := head
	//空链表
	if temp.next == nil {
		fmt.Println("这是一个空的环形链表,不能删除")
		return head
	}
	if temp.next == head { //说明环形链表只有一个结点
		temp.next = nil
		return head
	}

	//将helper定位到环形链表最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}
	//如果包含两个以上的结点
	flag := true
	for {
		if temp.next == head { //如果到这里，说明比较到最后一个[最有一个还没比较]
			break
		}
		if temp.no == id {
			if temp == head {
				head = head.next
			}
			//可以直接删除
			helper.next = temp.next
			fmt.Printf("猫猫%d\n", id)
			flag = false
			break
		}
		temp = temp.next     //移动[比较]
		helper = helper.next //移动[一旦找到要删除的结点，helper发挥作用]
	}
	if flag {
		if temp.no == id {
			helper.next = temp.next
			fmt.Printf("猫猫%d\n", id)
		} else {
			fmt.Println("没有这个猫猫")
		}
	}

	return head
}
func main() {
	//这里初始化一个环形链表的头结点
	head := &CatNode{}
	//创建一只猫
	cat1 := &CatNode{
		no:   1,
		name: "Tom",
	}
	cat2 := &CatNode{
		no:   2,
		name: "Jack",
	}
	cat3 := &CatNode{
		no:   3,
		name: "Rose",
	}
	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	ListCat(head)
	head = DeleteCat(head, 20)
	ListCat(head)
}
