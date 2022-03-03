package main

import "fmt"

//定义结点
type Hero struct {
	no       int
	name     string
	nickname string
	pre      *Hero //这个表示指向前一个结点
	next     *Hero //这个表示指向下一个结点
}

//增加一个结点
//方法一
func InsertHero(head *Hero, newhero *Hero) {
	//思路
	//1.先找到该链表的最后这个结点
	//创建辅助结点
	temp := head
	for {
		if temp.next == nil { //表示找到最后
			break
		}
		temp = temp.next //让temp不断指向下一个结点
	}
	//3.将newhero加入到链表最后
	temp.next = newhero
	newhero.pre = temp
}

//增加一个结点
//方法二，按照no排序
func InsertHero2(head *Hero, newhero *Hero) {
	//思路
	//1.先找到该链表的最后这个结点
	//创建辅助结点
	temp := head
	flag := true
	for {
		if temp.next == nil { //找到最后一个结点
			break
		} else if temp.next.no > newhero.no {
			break

		} else if temp.next.no == newhero.no {
			flag = false
			break
		}
		temp = temp.next
	}
	if flag == false {
		fmt.Println("该No已存在")
		return
	} else {
		newhero.pre = temp
		newhero.next = temp.next
		if temp.next != nil {
			temp.next.pre = newhero
		}
		temp.next = newhero
	}
}

//查询链表
//这里仍然使用单向列表的查询方式
func ListHero(head *Hero) {
	//创建一个辅助结点
	temp := head
	if temp.next == nil {
		fmt.Println("该链表为空")
		return
	}
	//2.遍历这个链表
	for {
		//输出头结点next结点信息
		fmt.Printf("[%d,%s,%s]->", temp.next.no, temp.next.name, temp.next.nickname)
		//辅助变量该为下一节点
		temp = temp.next
		//如果temp.next 等于nil，则退出
		if temp.next == nil {
			break
		}
	}
	fmt.Println("")
}

//倒着显示双向链表
func ListHero2(head *Hero) {
	//创建一个辅助结点
	temp := head
	if temp.next == nil {
		fmt.Println("该链表为空")
		return
	}
	for {
		if temp.next == nil { //找到链表最后一位
			break
		}
		temp = temp.next
	}
	for {
		fmt.Printf("[%d,%s,%s]->", temp.no, temp.name, temp.nickname)
		temp = temp.pre
		if temp.pre == nil {
			break
		}
	}

	//2.遍历这个链表
	// for {
	// 	//输出头结点next结点信息
	// 	fmt.Printf("[%d,%s,%s]->", temp.next.no, temp.next.name, temp.next.nickname)
	// 	//辅助变量该为下一节点
	// 	temp = temp.next
	// 	//如果temp.next 等于nil，则退出
	// 	if temp.next == nil {
	// 		break
	// 	}
	// }
	fmt.Println("")
}

//删除结点
func DeleteHero(head *Hero, id int) {

	temp := head
	flag := false
	for {
		if temp.next == nil { //找到最后一个结点
			break
		} else if temp.next.no > id {
			break

		} else if temp.next.no == id {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		}
	} else {
		fmt.Println("输入的Id不存在")
	}
}
func main() {
	//1.先创建一个头结点
	head := &Hero{}
	//2.创建一个新的英雄
	hero1 := &Hero{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}
	hero2 := &Hero{
		no:       2,
		name:     "李逵",
		nickname: "黑旋风",
	}
	hero3 := &Hero{
		no:       3,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}
	hero4 := &Hero{
		no:       4,
		name:     "林冲",
		nickname: "教头",
	}
	InsertHero2(head, hero1)
	InsertHero2(head, hero4)
	InsertHero2(head, hero3)
	ListHero(head)
	InsertHero2(head, hero2)
	ListHero2(head)
	DeleteHero(head, 3)
	ListHero(head)

}
