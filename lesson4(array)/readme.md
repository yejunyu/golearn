#### Table of contents
- [数组的声明](#数组的声明)
- [数组的遍历](#数组的遍历)
- [值传递 or 引用传递](#值传递-or-引用传递)
- [切片slice](#切片slice)
- [append](#append)
    - [扩容的规律](#扩容的规律)
    - [删除数组元素](#删除数组元素)
- [总结](#总结)


### 数组的声明
```go
func array() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	// 自行推断数组长度
	arr3 := [...]int{1, 3, 5, 7, 9}
	// 二维数组
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3, grid)
	// [0 0 0 0 0] [1 3 5] [1 3 5 7 9] [[0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0]]
}
```
变量的声明已经讲过啦,不熟悉的可以看第二章

### 数组的遍历
```go
func arrRange() {
	arr1 := [...]int{2, 4, 6, 8, 10}
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	for _, value := range arr1 {
		fmt.Println(value)
	}

```
两种方法,一种传统的下标遍历一种上一章讲到的`range`
推荐遍历用 range, 很方便
在idea里,传统的 `for` 循环智能提示是 `for`,`forr`就可以自动补全 `range`了
### 值传递 or 引用传递
上一章讲到这个概念了,这一章继续举个🌰
```go
func printArray(arr [5]int) {
	arr[0] = 100
	for _, v := range arr {
		println(v)
	}
}

func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := [3]int{1, 2, 3}
	printArray(arr1)
	printArray(arr2)
}
```
这个结果是什么呢
```
# command-line-arguments
lesson4(array)/array.go:45:12: cannot use arr2 (type [3]int) as type [5]int in argument to printArray
```
说的是类型不匹配,我需要一个`[5]int`的,你给我传的是`[3]int`的,当然报错啦
```go
func printArray(arr [5]int) {
	arr[0] = 100
	for _, v := range arr {
		println(v)
	}
}
func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	printArray(arr1)
	fmt.Println(arr1)
}
```
打印的结果是什么?
当然还是[1,2,3,4,5]啦,因为你是把 arr1拷贝了一份,函数体内的变化并不影响外部的变量
如果想改变 arr[0]的值你得这样
```go
func printArray2(arr *[5]int) {
	arr[0] = 100
	for _, v := range arr {
		println(v)
	}
}

func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	printArray(arr1)
	printArray2(&arr1)
	fmt.Println(arr1)
}
```
传一个指向 arr1的指针,懂了吗,在`go`里只有值传递,而且 `array` 这种类型也是值,而在其他语言里比如`python`,`javascript` `array` 都是引用类型
这几段代码可能让你觉得`go`里的数组很难用,其实不然,下面我们讲下数组的切片

### 切片slice
```go
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[2:6] = ", arr[:6])
	fmt.Println("arr[2:6] = ", arr[2:])
	fmt.Println("arr[2:6] = ", arr[:])
}
```
有编程基础的应该都能知道,结果是:
```
arr[2:6] =  [2 3 4 5]
arr[2:6] =  [0 1 2 3 4 5]
arr[2:6] =  [2 3 4 5 6 7]
arr[2:6] =  [0 1 2 3 4 5 6 7]
```
再来看看切片
```go
func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:]
	s2 := arr[:]
	fmt.Println(s2) // [0 1 2 3 4 5 6 7]
	fmt.Println("After updateSlice")
	updateSlice(s1)
	fmt.Println(s1) // [100 3 4 5 6 7]
	fmt.Println(arr) // [0 1 100 3 4 5 6 7]
}
```
**Slice 本身没有数据,只是对底层 array 的一个 view**
巩固一下
```go
// s2 = [0 1 100 3 4 5 6 7]
fmt.Println("reSlice")
s2 = s2[:5]
fmt.Println(s2) // [0 1 100 3 4]
s2 = s2[2:]
fmt.Println(s2) // [100 3 4]
```
这些都理解,下面再看个例子
```go
arr = [...]int{0,1,2,3,4,5,6,7}
s1 = arr[2:6]
s2 = s1[3:6]
fmt.Println("s1 = ",s1)
fmt.Println("s2 = ",s2)
```
这是不是有问题呢? s1是 arr 的第3个到第6个元素,总共4个
s2是不是下标越界了呢
运行一下
```
s1 =  [2 3 4 5]
s2 =  [5 6 7]
```
再来看看
`fmt.Println(s1[4])`
**panic: runtime error: index out of range**
很奇怪对不对,奇怪就对了
记住这个概念
**Slice 本身没有数据,只是对底层 array 的一个 view**
来看个图
![](http://oqb4aabpb.bkt.clouddn.com/FpTOPmxH6q-IdV2mlHLaiQWTEBHo)
- s1为[2 3 4 5],s2为[5 6 7]
- slice 可以向后扩展,不可以向前扩展
- s[i]不可以超越len(s),向后扩展不可以超越cap(s)
这就是`cap` `capacity(容量)`的概念
```go
fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n",
		s1, len(s1), cap(s1))
fmt.Printf("s2=%v,len(s2)=%d,cap(s2)=%d\n",
		s2, len(s2), cap(s2))

// s1=[2 3 4 5],len(s1)=4,cap(s1)=6
// s2=[5 6 7],len(s2)=3,cap(s2)=3
```
`go`的数组可以指定起始,结束和 `capacity`,而`python`是指定起始,结束和步长,不要混为一谈
`s1[1:2:3]`指定 `cap` 为3
`slice`起初我也觉得有点绕,多练习就好了
### append
```go
arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
s1 := arr[2:6]
s2 := s1[3:5]
s3 := append(s2, 10)
s4 := append(s3, 11)
s5 := append(s4, 12)
fmt.Println(s1, s2, s3, s4, s5)

// [2 3 4 5] [5 6] [5 6 10] [5 6 10 11] [5 6 10 11 12]
```
append和其他语言一样,可以突破`capacity`的限制,依次在数组后面添加元素
#### 扩容的规律
```go
func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
}

var s []int
for i := 0; i < 100; i++ {
	printSlice(s)
	s = append(s, 2*i+1)
}

```
贴几个结果感受下
```
len=14, cap=16
len=15, cap=16
len=16, cap=16
...
len=31, cap=32
len=32, cap=32
...
len=63, cap=64
len=64, cap=64
len=65, cap=128
len=66, cap=128
```
和`java`不同,`java`默认负载因子是`0.75`,`go`看的出来是`capacity`满了才扩容,每次扩容两倍,所以和`java`一样,数组最好知道容量,上来就建好

还有一种创建数组的方法`make`
```
s2 = make([]int, 16)
s3 = make([]int, 10, 32)
printSlice(s2)
printSlice(s3)

// len=16, cap=16
// len=10, cap=32
```
后面会经常用到,创建数组和`channel`都很常用
`copy(s2, s1)`
```go
func copy(dst, src []Type) int
```
```go
fmt.Println(s2)

// [2 3 4 5 0 0 0 0 0 0 0 0 0 0 0 0]
```

#### 删除数组元素
```go
fmt.Println("Deleting elements from slice")
s2 = append(s2[:3], s2[4:]...)
printSlice(s2)

fmt.Println("Poping from front")
front := s2[0]
s2 = s2[1:]
fmt.Println(front)
printSlice(s2)

fmt.Println("Poping from tail")
tail := s2[len(s2)-1]
s2 = s2[:len(s2)-1]
fmt.Println(tail)
printSlice(s2)

/**
Deleting elements from slice
len=15, cap=16
Poping from front
2
len=14, cap=15
Poping from tail
0
len=13, cap=15
*/
```

### 总结
- `go`里只有值传递
- 不要直接用 array, 要用切片来操作
- 注意`cap` 这个隐含的值,你如果不想别人能访问到切片以外的数据,可以加上 `cap` 比如`s[1:2:3]`
- 切片只是对底层 `array` 的 `view`


