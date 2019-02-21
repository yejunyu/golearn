#### Table of contents
- [go变量的基本类型](#go变量的基本类型)
- [变量的声明](#变量的声明)
- [常量](#常量)
- [总结](#总结)
> 任何编程语言都有变量,下面来学一学 `go` 的变量与其他语言有什么异同

### go变量的基本类型

- bool,string
- (u)int,(u)int8,(u)int16,(u)int32,(u)int64,uintptr(指针)
- byte,rune(go 的 char,4字节32位)
- float32,float64,complex64,complex128(实部虚部都是64位)

一个比较大的区别, `go` 作为一种静态语言,没有 long,double 这种常见的类型,而是只有 int 和 float.
这是因为 `go` 语言崇尚简洁,你觉得 int 不够用可以用 int32 或者 int64,float64也是一样的道理
complex 很特别,别的语言都没有,看来 go 野心不小, python 的科学计算它也想掺一脚☺
下面来复习一下高中数学:
最美数学公式,欧拉公式
e<sup>iπ</sup>+1=0
用 go 语言怎么表示呢
```go
func euler() {
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	// Exp表示以e为底
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	// (0+1.2246467991473515e-16i)
}
有兴趣的同学可以自己验证一下,因为计算机浮点数精度的问题,这里肯定是不等于0的
```
### 变量的声明
`go`语言有很多种声明变量的方法,下面一起来看一看
和 `java` 不同, `go` 的设计哲学就是简洁,变量的声明就体现了他的这个理念
```go
/**
不赋初值
*/
func variableZeroValue() {
	var a int
	var s string
	// quotation 可以把引号打出来
	fmt.Printf("%d %q\n", a, s) // 0 ""
}

/**
赋初值
*/
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s) // 3 4 abc
}
```
用一个 `var` 关键字就声明了一个变量,`go`和`scala`一样是把变量名放前面,类型放后面,起初不太习惯,用多了也就习惯了,而且想一想,确实变量名比类型重要,确实应该放前面

---------
`go`还有一个很好的地方--类型推断(`java`10也有了)
```go
/**
类型推断
*/
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s) // 3 4 true def
}

/**
简写变量定义
*/
func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s) // 3 5 true def
}
```
由上面代码可知,声明一个变量我也可以不写出变量的类型,而是由编译器根据变量的值来自行推断变量的类型,这点和动态类型很相似,再也不用写冗长的类型声明了
`go`还有一种更简洁的写法就是用`:`来代替 var, 小伙伴们使用 `go` 的时候可以试试简写的声明变量的方法,真的很方便

----
刚才都是函数里的,也可以在函数外声明变量,不过函数外面就不能用简写的声明方式了
```go
var (
	a = 3
	b = "str"
)
```
-------
类型转换
```
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b))) // go 只有强制类型转换
	fmt.Println(c)                         // 5

}

```
注意,这里 `math.Sqrt`里需要传 `float` 类型,而 `go` 是没有隐式转换的,必须显式的转换
### 常量
直接看代码吧
```go
func consts(){
	const filename = "abc.txt"
	const a, b  = 3,4
	var c int
	c = int(math.Sqrt(a*a+b*b))
	fmt.Println(filename,c) // abc.txt 5
}
```
和声明变量一样,只不过关键字改成 `const`, 注意通常其他语言常量都是全大写表示,`go`里不建议这么做,因为 `go`里首字母大写代表是 `public`
再看一个常量的例子,这里用了 `go`里的一个关键字`iota`(自增值)
```go
func enums()  {
	const (
		cpp = iota
		_
		python
		golang
	)
	const (
		b = 1 << (10*iota)
		kb
		mb
		gb
		tb
	)
	fmt.Println(cpp,python,golang) // 0 2 3
	fmt.Println(kb,mb,gb,tb) // 1024 1048576 1073741824 1099511627776
}
```
`iota`默认是0,下面的变量或常量都会依次`+1`,这里`_`代表占位符,以后会经常用到
### 总结
- `go`没有 long和 double,但是可以指定 int 和 float 的长度,注意 int 不指定是根据系统来的,32位系统就是 int32,64位系统就是 int64,float 必须指定32或者64
- `go`可以用 `var` 关键字或者直接 `:=`来声明一个变量
- 在函数体外的时候不能用简写,只能用`var`关键字来声明变量
- `go` 语言必须显式的强制类型转换
- `go`里首字母大写的变量或者函数都代表 `public`


