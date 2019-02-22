package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func CreateNode(value int) *Node {
	// go 里不需要关心变量是在堆上创建还是栈上创建的
	return &Node{Value: value}
}

func (node *Node) SetValue(value int) {
	node.Value = value
}

func (node Node) Print() {
	fmt.Println(node.Value)
}

func (node *Node) Traverse() {
	//if node == nil {
	//	return
	//}
	//// 其他语言需要判断 null,go 不用
	//node.Left.Traverse()
	//node.Print()
	//node.Right.Traverse()

	// 函数式编程
	node.TraverseFunc(func(node *Node) {
		node.Print()
	})
	fmt.Println()
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}
