
[toc]

### 面向对象
**`go`只支持封装,不支持继承和多态**
`go`是面向接口的编程,也可以说`go`所有对象都是继承了一个空接口
`java`这类面向对象的语言,三大特点`封装`,`继承`,`多态`,`多态`非常重要,可以说前面两个特点都是为了`多态`,所以学习`go`也会帮助你换个思路理解`面向对象`
`go`暂时没有`泛型`,不过2据说要出泛型
我们都知道`c`语言是典型的面向过程的,但是 `c`有结构体这种结构,其实这就是后面的面向对象的基础

### 结构体定义
来定义一个树型结构
```go
type treeNode struct {
	value       int
	left, right *treeNode
}

var root treeNode
fmt.Println(root)
// {0 <nil> <nil>}
```

```go
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

// 没有构造函数,但是可以自定义工厂函数
func createNode(value int) *treeNode {
	// go 里不需要关心变量是在堆上创建还是栈上创建的
	return &treeNode{value: value}
}
```
一定是要传一个地址出去
### 面向接口编程
假如我想扩展一下这个结构体,想给他增加一个 print 方法
```go
// 函数名前面代表是这个结构体有的方法
// go 是面向接口编程,这其实相当于是实现接口
// treeNode 这个结构体有 print 这个方法
func (node treeNode) print() {
	fmt.Println(node.value)
}

root.print()
```
我这么定义就相当于给 treeNode 这个结构体实现了 print 这个接口
然后就像面向对象那样,点出来就可以有 print 这个方法了

我想改变 treeNode 的 value 值呢
```go
func (node treeNode) setValue(value int) {
	node.value = value
}
```
这样有问题吗,这样肯定是改变不了 treeNode 的 value 值的,**`go`里只有值传递,如果想改变这个对象,就要传指针**
所以应该改为
```
func (node *treeNode) setValue(value int) {
	node.value = value
}
```
一个`*`的差别
```go
root.print()
root.setValue(100)

// 这里两种定义都行,go 的编译器非常聪明
proot := root
// proot := &root

proot.print()
proot.setValue(200)
proot.print()

// 3
// 100
// 200
```
### 遍历二叉树
```go
// root 100->0->2 right 100->5->0
fmt.Println("traverse~")
root.traverse()

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	// 其他语言需要判断 null,go 不用
	node.left.traverse()
	node.print()
	node.right.traverse()
}
```
非常简洁

### 总结
- 要改变内容必须使用指针接收者(传递对象的地址)
- 结构过大也考虑指针接收者(指针只是对象的地址)
- 一致性:如果有指针接收者,最好都是指针接受者




