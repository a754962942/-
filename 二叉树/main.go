package main

import "fmt"

type Hero struct {
	No    int
	Name  string
	Left  *Hero
	Right *Hero
}

//前序遍历[先输出root结点，然后再输出左子树，然后再输出右子树]
func PreOrder(node *Hero) {
	if node != nil {
		fmt.Printf("No = %d,Name = %s\n", node.No, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

//中序遍历[先输出Root的左子树，再输root结点，最后输出root的右子树]
func InfixOrder(node *Hero) {
	if node != nil {
		InfixOrder(node.Left)
		fmt.Printf("No = %d,Name = %s\n", node.No, node.Name)
		InfixOrder(node.Right)
	}
}

//后序遍历[先输出Root的左子树，再输出右子树，最后输出root结点]
func BackOrder(node *Hero) {
	if node != nil {
		BackOrder(node.Left)
		BackOrder(node.Right)
		fmt.Printf("No = %d,Name = %s\n", node.No, node.Name)
	}
}

func main() {
	//创建一个二叉树
	root := &Hero{
		No:   1,
		Name: "宋江",
	}
	left1 := &Hero{
		No:   2,
		Name: "吴用",
	}
	right1 := &Hero{
		No:   3,
		Name: "卢俊义",
	}
	right2 := &Hero{
		No:   4,
		Name: "林冲",
	}
	root.Left = left1
	root.Right = right1
	right1.Right = right2
	BackOrder(root)
}
