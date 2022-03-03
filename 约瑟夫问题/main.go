package main

import "fmt"

//小孩的结构体
type Boy struct {
	No   int  //小孩编号
	Next *Boy //指向下一个小孩的指针
}

//编写一个函数，构成单向的环形链表
//num 表示小孩个数
//*Boy :返回该环形链表的第一个小孩的指针
func Addboy(num int) *Boy {
	//首先先创建一个空结点
	first := &Boy{}
	curBoy := &Boy{}
	if num < 1 {
		fmt.Println("num输入有误")
		return first
	}
	//循环构成环形链表
	for i := 1; i <= num; i++ {
		boy := &Boy{
			No: i,
		}
		//分析构成循环链表，需要一个辅助指针帮忙
		//1.因为第一个小孩比较特殊
		if i == 1 { //第一个小孩
			first = boy //头指针固定
			curBoy = boy
			curBoy.Next = first
		} else {
			curBoy.Next = boy   //先将当前环形链表next指向新添加的boy
			curBoy = boy        //将curBoy赋值为新添加的Boy
			curBoy.Next = first //将curBoy即新添加的Boy.next指向头指针
		}
	}
	return first
}
func ShowBoy(first *Boy) {
	//如果链表为空，就退出
	if first.Next == nil {
		fmt.Println("链表为空")
		return
	}
	//设置辅助指针
	curBoy := first
	//开始循环遍历
	for {
		fmt.Printf("Boy%d ->", curBoy.No)
		//退出条件为curBoy.next指向first
		if curBoy.Next == first {
			break
		}
		curBoy = curBoy.Next
	}
	fmt.Println("")
}

/*
设编号为1，2，...，n的n个人围坐一圈，约定编号为k(1<=k<=n)的人从1开始报数
数到m的人出列，它的下一位又从1开始报数，数到m的又出列，以此类推
直到所有人都出列位置，由此产生出一个出队编号的序列
*/
//分析思路
//1.编写一个函数PlayGame(first *boy,stratNo int,countNum int)
//2.最后我们使用一个算法，按照要求，在环形链表中留下最后一个人
func PlayGame(first *Boy, startNo, countNum int) {
	//空链表单独设置
	if first.Next == nil {
		fmt.Println("空链表")
		return
	}
	//留一个，判断startNo<=小孩总数
	//定义辅助指针，帮助删除小孩
	tail := first
	//让tail执行环形链表的最后一个小孩
	//因为tail在删除小孩时需要用到
	for {
		if tail.Next == first { //说明tail到了最后的小孩
			break
		}
		tail = tail.Next
	}
	//让first移动到stratNo[删除小孩以first为准]
	for i := 1; i <= startNo-1; i++ {
		first = first.Next
		tail = tail.Next
	}
	//开始数CountNum,然后就删除first
	for {
		//开始数countNum-1次
		for i := 0; i < countNum-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		//删除first指向的小孩
		fmt.Printf("出圈的为%d号\n", first.No)
		first = first.Next
		tail.Next = first
		if first == tail {
			fmt.Printf("最后剩下%d号\n", first.No)
			break
		}
	}

}
func main() {
	first := Addboy(5)
	ShowBoy(first)
	PlayGame(first, 2, 3)
}