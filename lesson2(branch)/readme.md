[toc]
### if
先看一个最常见的 `if` `else`分支
```go
func readFile() {
	const filename = "./lesson2(branch)/abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}
```
发现什么了吗, `go`里的`if`后面不用加括号,直接写语句就行,等下学的 `for` 循环也是不用加括号,因为`go`是个崇尚简洁的语言,下面再来看这个例子还能这么写
```go
func readFile() {
	const filename = "./lesson2(branch)/abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}
```
 可以把读操作和判断 error 写在同一行,`if` `else`没啥讲的,其他语言你会用,这也一样
### switch
直接看代码
```go
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 1000:
		g = "A"
	default:
		panic(fmt.Sprintf("Wrong score: %s", score))
	}
	return g
}
```
与 java 些许不同的是,这些 case 里面不用写 break, 编译器会自动帮你退出,很贴心有木有,`panic`是个关键字,抛出异常,相当于`new throw`.分支就讲到这,下面讲循环
### for
直接看代码
```go
/**
传统循环
*/
func sum(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum
}
```
`go`的 `for`不光省略了括号,还可以省略初始条件,终止条件
```go
/**
省略初始条件
*/
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

/**
只有终止条件的循环
*/
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

/**
死循环
*/
func forever() {
	for {
		fmt.Println("我是死循环")
		time.Sleep(1 * time.Second)
	}
}
```
`go`里死循环就是 `for` 后面什么都不加,就是 `while(true)`的意思
###总结
- `for`,`if`后面的条件没有括号
- `if`条件里也可以定义变量
- 没有`while`
- `switch`不需要`break`,也可以直接`switch`多个条件,条件写`case`里
