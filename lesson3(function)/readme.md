#### Table of contents
- [函数](#函数)
- [多返回值函数](#多返回值函数)
- [函数式编程](#函数式编程)
- [可变参数](#可变参数)
- [值传递 or 引用传递](#值传递-or-引用传递)
- [总结](#总结)


### 函数
来写一个加减乘除吧
```go
func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unknown operation " + op)
	}
}
```
函数三要素,`入参` `函数名` `返回值`
a,b 都是 int 我就不用 a int ,b int 这么写了,变量放前面,类型放后面
所以函数返回值也是一样,返回值放后面

### 多返回值函数
```go
/**
多返回值
 */
func div(a, b int) (q, r int) {
	return a / b, a % b
}
```
返回值那也可以只写个 int, 但是有时候为了告诉别人你这个返回值是什么意思,可以命名一下,比如这里的 q,r 代表商和余数
那刚刚的加减乘除就可以改写一下
```go
func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		q, _ := div(a, b) // _表示占位符, 啊 go 真的很严格,不允许有无用的变量
		return q
	default:
		panic("unknown operation " + op)
	}
}
```
上节也说到了`_`代表占位符,其实很多语言比如 `php`,`python` 也是用`_`当占位符,`go` 里不允许有无用变量,所以这里不需要余数这个返回值可以用`_`代替
下面说个好玩牛逼的东西
###函数式编程
还是刚刚的功能
```go
func evalfunc(op func(int, int) int, a, b int) int {
	return op(a, b)
}
```
`go`里函数是一等公民,函数也可以像变量一样传递
假如我想算 $a^b$
```go
/**
函数式编程
 */
func evalfunc(op func(int, int) int, a, b int) int {
	// 通过反射获取函数的指针
	p := reflect.ValueOf(op).Pointer()
	// 通过指针获取函数名
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+
		"(%d %d)\n", opName, a, b)

	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func main() {
	fmt.Println(evalfunc(pow, 3, 4))
}

// Calling function main.pow with args (3 4)
// 81
```
就是这么酷炫,直接把函数传到另一个函数体里
再看一个例子
```go
// 匿名函数
fmt.Println(evalfunc(
	func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))
// Calling function main.main.func1 with args (3 4)
// 81
```
接着上段代码,我也可以写成匿名函数的样子,就不用特意出来定义一个函数了,比如`python`里的 `lambda` 表达式
> `函数式编程`现在已经很流行了,推荐学一学用一用,买不了吃亏也买不了上当的

### 可变参数
`go`里没有`java`那复杂的`函数重载`也没有 python 的`参数默认值`,因为`go`崇尚简洁
但是`可变参`还是有的
```go
/**
可变参数
 */
func sum(num ...int) int {
	s := 0
	for _, value := range num {
		s += value
	}
	return s
}
fmt.Println(sum(1, 2, 3)) // 6
```
`range` 后面再讲,也是用的非常多的一种用法,`...`代表不定参,好像 `js`也有这种写法,其实现在很多语言都是互相借鉴,所以还是非常推荐大家学`go`的,语言简洁,性能非常高,又没有历史包袱

### 值传递 or 引用传递
复习一下指针的知识吧
```go
func passByVal(a int){
	a++
}
func passByRef(a *int)  {
	*a++
}
func main() {
	a := 3
	passByVal(a)
	fmt.Println(a)
	passByRef(&a)
	fmt.Println(a)
}
```
这段代码会打印什么呢
没错
就是3,4
在看一个例子
```go
func swap(a, b *int) {
	*a, *b = *b, *a
}
func main(){
    b := 3
	c := 4
	swap(&b, &c)
	fmt.Println(b, c) // 4,3
}
```
其他的语言学习的时候可能要分值传递,引用传递,新手在这上面很容易出错.
记住了`go`
**只有值传递,不论传值还是传指针他都是拷贝一份到函数体里**
指针和函数是重点,现在只是回顾一下,后面还会讲到

### 总结
总结一下吧

- `go`函数返回值写在最后面,你会发现很多新的语言都是这么设计的
- 函数可以有多个返回值
- 函数可以作为参数传给另一个函数
- 没有默认参数和可选参数
- `go`都是值传递没有引用传递


