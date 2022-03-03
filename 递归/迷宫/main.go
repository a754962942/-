package main

import "fmt"

//编写一个函数，完成老鼠找路
//mymap *[8][7]int：地图，保证是同一个地图
//i，j表示对地图的那个点进行测试
func FindWay(mymap *[8][7]int, i, j int) bool {
	//分析出什么情况下，就找到出路
	if mymap[6][5] == 2 {
		return true
	} else {
		//说明继续找
		if mymap[i][j] == 0 {
			//如果该点可以探测

			//假设这个点是可以通的，但是需要探测.上下左右
			mymap[i][j] = 2
			if FindWay(mymap, i-1, j) {
				//向上
				return true

			} else if FindWay(mymap, i+1, j) {
				//向下
				return true
			} else if FindWay(mymap, i, i-1) {
				//向左
				return true
			} else if FindWay(mymap, i, i+1) {
				//向右走
				return true
			} else {
				//死路
				mymap[i][j] = 3
				return false
			}
		} else { //说明这个点不能探测，为墙
			return false
		}
	}
}
func main() {
	//先创建一个二维数组，模拟迷宫
	//规则
	//1.如果元素值为1，就是墙
	//2.如果元素值为2，是一个通路
	//3.如果元素值为3，是走过的点，但是没走通
	//4.如果元素值为0，那就是没有走过的点
	var myMap [8][7]int
	//先设置地图的墙
	for i := 0; i < 8; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}
	for i := 0; i < 7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
	}
	//显示地图
	// for i := 0; i < 8; i++ {
	// 	for j := 0; j < 7; j++ {
	// 		fmt.Print(myMap[i][j], "\t")
	// 	}
	// 	fmt.Println("")
	// }
	FindWay(&myMap, 1, 1)
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], "\t")
		}
		fmt.Println("")
	}
}
