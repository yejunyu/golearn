#### Table of contents
- [包管理](#包管理)
- [扩展](#扩展)
- [总结](#总结)


### 包管理
我用过的语言,比如`java`,`python`都是按包来区分命名空间的
`go`里也是一样
```go
import (
	"golearn/lesson8/tree"
	"fmt"
	"golearn/lesson8/queue"
)
```
这里`fmt`就是`go`里的标准库的包
前面说到过,`go`里常量不推荐大写,为什么,因为,`go`里首字母大写代表`public`
我们来看一下上一章的二叉树,如果要把中序遍历扩展一下成后序遍历,怎么做呢
首先得改造一下,首字母都大写
```go
package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Println(node.Value)
}

func CreateNode(value int) *Node {
	// go 里不需要关心变量是在堆上创建还是栈上创建的
	return &Node{Value: value}
}

func (node *Node) SetValue(value int) {
	node.Value = value
}

func (node *Node) Traverse() {
	if node == nil {
		return
	}
	// 其他语言需要判断 null,go 不用
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
```
这是之前的代码,现在把方法名和变量名全都首字母大写

> idea 里装上UpperLowerCapitalize插件后
Alt+P // to uppercase
Alt+L // to lowercase
Alt+C // 首字母大写

把 `tree.go`这个文件移到`tree`这个文件夹下,因为一个文件夹下只能有一个包名, `MyTreeNode`是 `main`包

> 这里有个小插曲,我本来这个包叫 lesson8-package,发现引入后怎么都是报错,才发现,包名不能带-号

之前的中序遍历是这样的
```go
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
}

```
来体验一下`go`的面向接口编程
```go
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
```
沿用之前的结构体,定义一个新的结构体`MyTreeNode`
然后给`MyTreeNode`增加一个方法
`func (结构体指针) 方法名([参数])[返回值]`
后序遍历就是`左右中`,如上面的代码
来验证一下
```go
fmt.Println("后序遍历")
myRoot := MyTreeNode{&root}
myRoot.postOrder()
// 2 0 4 5 3
```

### 扩展
我想把之前的数组的 push 和 pop 封装一下
和二叉树一样,单独放在一个包里

```go
package queue

type Queue []int

func (q *Queue) Push(v int) {
	*q = append(*q, v)

}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	if len(*q) == 0 {
		return true
	} else {
		return false
	}
}

```

```go
package main

import (
	"golearn/lesson8/tree"
	"fmt"
	"golearn/lesson8/queue"
)

func  main(){
    s := queue.Queue{1, 3, 5, 7, 9}
    s.Push(11)
    fmt.Println(s)
    for !s.IsEmpty() {
    	s.Pop()
    	fmt.Println(s)
    }

}

```

### 总结
- 要想让其他包能访问你的包,变量和方法首字母大写,代表`public`
