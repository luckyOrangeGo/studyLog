package main

import (
	"fmt"
)

//定义Node结构体类型
type Node struct {
	//节点值为int型
	Value int

	//左子树和右子树都为Node结构体，使用指针调用自己的类型
	Left, Right *Node
}

//使用自定义工厂函数创建Node 方法--4
func CreateNode(v int) *Node {
	// v int 是局部变量，返回的是局部变量的地址
	return &Node{Value: v}
}

//建立打印函数
//	(node Node)是接收者，调用时使用 root.Print()
//	如果是print(node Node)，调用时使用print()
//	此函数传递的是值，如果修改值在只能作用在该函数内
func (node *Node) Print() {
	fmt.Print(node.Value, "\n")
}

//建立setValue 注意不是set函数内的值，需要传递指针
//只有使用指针才可以改变结构内容
func (node *Node) SetValue(v int) {
	//nil指针也可以调用方法
	if node == nil {
		fmt.Println("忽略给 nil node 设置值")
		return
	}
	node.Value = v
	fmt.Print("setThisNode is ")
	node.Print()
}

func main() {

	//创建Node 方法--1
	var root Node

	//给Node:value赋值
	root = Node{Value: 1000}

	//给Node左，右子树赋值
	//	因为left, right是指针，所以要取地址
	root.Left = &Node{}
	root.Right = &Node{400, nil, nil}
	//打印，注意分支是指针，要用*取值
	fmt.Println(root, *root.Left, *root.Right, "方法--1")

	//创建Node 方法--2
	//	使用Go内建函数new()
	//	不管是指针还是实例都使用 "."
	root.Right.Left = new(Node)

	root.Right.Left.SetValue(100)
	fmt.Print("root.Right.Left=")
	root.Right.Left.Print()

	//打印，注意root.Right.left是指针，要用*取值
	fmt.Println(*root.Right, *root.Right.Left, "方法--2")

	//创建Slice定义的nodes 方法--3
	nodes := []Node{
		{Value: 3}, // 注意有逗号,
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes, "方法--3")

	root.Left.Right = CreateNode(100)
	fmt.Println(*root.Left.Right, "方法--4")

	//取 root值 传递给print
	fmt.Print("root=")
	root.Print()

	//取 root指针 传递给print
	pRoot := &root
	fmt.Print("root=")
	pRoot.Print()

	//编译器自动找到函数需要的是值还是指针，把它交给函数
	pRoot.SetValue(2000)

	fmt.Println("set pRootNew")
	var pRootNew *Node
	//尝试给nil node赋值
	pRootNew.SetValue(10000)
	pRootNew = &root
	pRootNew.SetValue(20000)

	fmt.Println("中序遍历")
	root.TraverseMiddle()
	fmt.Println("")
}

//中序遍历
func (node *Node) TraverseMiddle() {
	//允许空指针进入函数，但遇到时返回
	if node == nil {
		return
	}
	node.Left.TraverseMiddle()
	node.Print()
	node.Right.TraverseMiddle()
}
