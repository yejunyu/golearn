package tree

import (
	"golearn/lesson8(package)/tree"
	"testing"
)

func BenchmarkTree(b *testing.B) {
	root := tree.Node{3, nil, nil}
	root.Left = tree.CreateNode(1)
	root.Left.Left = tree.CreateNode(2)
	root.Left.Right = tree.CreateNode(3)
	root.Right = tree.CreateNode(4)
	root.Right.Left = tree.CreateNode(5)
	root.Right.Right = tree.CreateNode(6)
	// 准备数据的时间除外
	b.ResetTimer()
	// 2 1 3 3 5 4 6
	for i := 0; i < b.N; i++ {
		root.TraverseFunc(func(node *tree.Node) {
			node.Print()
		})
	}

}
