package main

import (
	"fmt"
)

type treeNode struct {
	value       int
	left, right *treeNode
}

// 没有构造函数,但是可以自定义工厂函数
func createNode(value int) *treeNode {
	// go 里不需要关心变量是在堆上创建还是栈上创建的
	return &treeNode{value: value}
}

// 函数名前面代表是这个结构体有的方法
// go 是面向接口编程,这其实相当于是实现接口
// treeNode 这个结构体有 print 这个方法
func (node treeNode) print() {
	fmt.Println(node.value)
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	// 其他语言需要判断 null,go 不用
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func (node *treeNode) setValue(value int) {
	node.value = value
}

func main() {
	var root treeNode
	fmt.Println(root)

	// 结构的创建
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	fmt.Println(root)
	root2 := []treeNode{
		{value: 3},
		{},
		{6, nil, root.left},
	}
	root.left.right = createNode(2)
	fmt.Println(root2)

	root.print()
	root.setValue(100)
	proot := root
	proot.print()
	proot.setValue(200)
	proot.print()

	// root 100->0->2 right 100->5->0
	fmt.Println("traverse~")
	root.traverse()
}
