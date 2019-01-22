#### Table of contents
- [表格驱动测试](#表格驱动测试)
- [命令行测试](#命令行测试)
- [性能调优](#性能调优)
- [总结](#总结)

### 表格驱动测试
传统的`assert`测试形式就不贴代码了,
传统测试的问题是测试数据和代码混在一起了,除非你去看实现代码,要不然你并不知道测试代码是干嘛的
还有一个问题是只要一个出错,整个测试就停止了,得改正 assert, 然后再来一次

```go
tests := []struct {
		a, b, c int
	}{
		{1, 2, 3},
		{0, 2, 2},
		{0, 0, 0,},
		{-1, 1, 0},
		{math.MaxInt32, 1, math.MinInt32},
	}

for _, test := range tests {
	if actual := add(test.a, test.b); actual != test.c {
		
	}
}
```

表格驱动测试的优点:

- 分离的测试数据和测试逻辑
- 明确的出错信息
- 可以部分失败

看一个完整测试的例子
`file: Main.go`
```go
package main

func add(a, b int) int {
	return a + b
}

func main() {

}
```
`file:add_test.go`
**测试文件必须以_test结尾**
```go
package main

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, c int
	}{
		{1, 2, 3},
		{0, 2, 2},
		{0, 0, 0,},
		{-1, 1, 1},
		{math.MaxInt32, 1, math.MinInt32},
	}

	for _, test := range tests {
		if actual := add(test.a, test.b); actual != test.c {
			t.Errorf("add(%d,%d); got %d;but expected %d", test.a, test.b, test.c, actual)
		}
	}
}

```
**要测试什么方法就取名TestFunc(t *testing.T)**


### 命令行测试
测试当前目录
`go test .`

测试代码覆盖率,生成 html 文件(idea也有这个功能)
`go tool cover -html=c.out`

benchmark
`go test -bench .`
```go
func BenchmarkAdd(bench *testing.B) {
	a := math.MaxInt64
	b := 1
	c := math.MinInt64
	for i := 0; i < bench.N; i++ {
		if actual := add(a, b); actual != c {
			bench.Errorf("add(%d,%d); got %d;but expected %d", a, b, c, actual)
		}
	}
}
```
和 Test 类似,要测 Benchmark 就把 Test 换成 Benchmark, 然后类型是 `*testing.B`

Benchmark 要测多次,这个多少次,不用用户操心,用 `bench.N`
就行,来看下结果
![](http://oqb4aabpb.bkt.clouddn.com/18-8-12/90416178.jpg)
循环了20亿次,0.63ns/op 代表每次循环用的时间是0.63ns, 代表了你程序的性能

### 性能调优
换个复杂点的程序吧,之前的二叉树遍历拿过来
```go
func BenchmarkTree(b *testing.B) {
	root := tree.Node{3, nil, nil}
	root.Left = tree.CreateNode(1)
	root.Left.Left = tree.CreateNode(2)
	root.Left.Right = tree.CreateNode(3)
	root.Right = tree.CreateNode(4)
	root.Right.Left = tree.CreateNode(5)
	root.Right.Right = tree.CreateNode(6)
	// 2 1 3 3 5 4 6
	for i := 0; i < b.N; i++ {
		root.TraverseFunc(func(node *tree.Node) {
			node.Print()
		})
	}

}
```
![](http://oqb4aabpb.bkt.clouddn.com/18-8-12/34760110.jpg)

来分析一下
`go test -bench . -cpuprofile cpu.out`
但是是二进制文件,看不懂
用下 go tool
输入
`go tool pprof cpu.out`
是个交互式的工具,有很多功能,这里用最简单的 `web` 工具
![](http://oqb4aabpb.bkt.clouddn.com/18-8-12/82948151.jpg)
需要下载一个包,不是`go`自带的
需要装一个`Graphviz`
> http://www.graphviz.org/download/

![](http://oqb4aabpb.bkt.clouddn.com/18-8-12/60209337.jpg)
会生成一张调用图,框越大代表话费的时间越长,就可以针对性的优化了

### 总结
- 测试文件必须以*_test.go命令
- 测试文件里的测试方法必须以Test*命令
- 功能测试传参为t *testing.T
- 性能测试传参为b *testing.B
- 用pprof来分析性能瓶颈
- go tool中有很多有用的工具


