package main

import (
	"fmt"
	"golearn/lesson8/queue"
	"golearn/lesson8/tree"
)

type MyTreeNode struct {
	node *tree.Node
}

// 后序遍历
func (myNode *MyTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := MyTreeNode{myNode.node.Left}
	right := MyTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}           // root->3
	root.Left = &tree.Node{}             // left->0
	root.Right = &tree.Node{5, nil, nil} // right->5
	root.Right.Left = new(tree.Node)     // right->left->0
	root.Left.Right = tree.CreateNode(2) // left->right->2
	root.Right.Left.SetValue(4)
	// 0->2->3->4->5
	root.Traverse()
	fmt.Println("后序遍历")
	myRoot := MyTreeNode{&root}
	myRoot.postOrder()
	count := 0
	root.TraverseFunc(func(node *tree.Node) {
		count++
	})
	fmt.Println("count: ", count)

	s := queue.Queue{1, 3, 5, 7, 9}
	s.Push(11)
	fmt.Println(s)
	for !s.IsEmpty() {
		s.Pop()
		fmt.Println(s)
	}
}
