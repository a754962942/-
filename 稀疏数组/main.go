package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type nodeval struct {
	row int
	col int
	val interface{}
}

//文件读取转成原始数据
func ReadData(filename string) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer file.Close()
	bfnr := bufio.NewReader(file)
	var chess [][]int
	index := 0
	for {
		byteslice, err := bfnr.ReadBytes('\n')
		if err != nil {
			break
		}
		index++
		temp := strings.Split(string(byteslice), " ")
		row, _ := strconv.Atoi(string(temp[0]))
		col, _ := strconv.Atoi(string(temp[1]))
		val, _ := strconv.Atoi(string(temp[2]))
		if index == 1 {
			for i := 0; i < row; i++ {
				var arr []int
				for j := 0; j < col; j++ {
					arr = append(arr, val)
				}
				chess = append(chess, arr)
			}
		}
		if index != 1 {
			chess[row][col] = val
		}
	}
	fmt.Println("")
	fmt.Println("原始数组：")
	fmt.Println("")
	for _, value := range chess {
		for _, value2 := range value {
			fmt.Printf("%d\t", value2)
		}
		fmt.Println("")
	}
}

func main() {

	var chessmap [11][11]int
	chessmap[1][2] = 1
	chessmap[2][3] = 2

	// 看看原始数据
	for _, v := range chessmap {
		for _, v1 := range v {
			fmt.Printf("%d\t", v1)
		}
		fmt.Println()
	}

	// 转成稀疏数据
	var sparseArr []nodeval
	// 数据规模
	sparseArr = append(sparseArr, nodeval{
		row: 11,
		col: 11,
		val: 0,
	})
	//稀疏数组
	for row, val := range chessmap {
		for col, val1 := range val {
			if val1 != 0 {
				sparseArr = append(sparseArr, nodeval{
					row: row,
					col: col,
					val: val1,
				})
			}
		}
	}
	// 稀疏数组存盘
	filepath := "./chess.data"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, node := range sparseArr {
		str := fmt.Sprintf("%d %d %d \n", node.row, node.col, node.val)
		writer.WriteString(str)
	}
	writer.Flush()

	// 稀疏数据从磁盘读取转换成原始数据
	ReadData(filepath)
}
