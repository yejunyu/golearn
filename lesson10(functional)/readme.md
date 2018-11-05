#### Table of contents
- [闭包](#闭包)
    - [python 闭包](#python-闭包)
    - [java 闭包](#java-闭包)
    - [go 闭包](#go-闭包)
- [函数式编程入门](#函数式编程入门)
- [goimports](#goimports)
- [总结](#总结)

### 闭包
通过一个累加器来看闭包的概念

#### python 闭包
```python

def fun1():
    sum = 0
    def fun2(v):
        nonlocal sum
        sum += v
        return sum
    return fun2

a = fun1()
for i in range(10):
    print(a(i))
```
fun1返回的不是一个值,而是一个函数 fun2,a = fun2,所以 a(i)会打印 sum 的值,为什么 sum 一直在加呢,函数里的值为什么可以带到函数体外呢,这就是闭包的神奇之处,闭包是离散数学的一个概念,可以多看看网上的讲解加深印象
其实可以把闭包看做一个类, sum 就是类里的属性, fun2就是类的方法
所以 fun2可以使用 sum(自由变量)

#### java 闭包
```java
static Function<Integer, Integer> adder() {
        final Holder<Integer> sum = new Holder<>(0);
        return (Integer value) -> {
            sum.value += value;
            return sum.value;
        };
    }

public static void main(String[] args) {
    Function a = adder();
    for (int i = 0; i < 10; i++) {
        System.out.println(a.apply(i));
    }
}
```
java 里函数不能像变量一样传递,但也能模拟闭包这里的 adder 其实是一个 Function 对象
上面`python`代码里 `sum`前有个`nonlocal`修饰,表明`sum`不是一个局部变量,这里直接用了 `final`修饰

> 闭包就是能够读取其他函数内部变量的函数。例如在javascript中，只有函数内部的子函数才能读取局部变量，所以闭包可以理解成“定义在一个函数内部的函数“。在本质上，闭包是将函数内部和函数外部连接起来的桥梁

#### go 闭包
```go
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}
```
但是正统的函数式编程不是这样,函数其实是一个系统,我们只关心,入参(x)和返回值(y)是什么,其实里面是怎么实现的我们并不关心,现代的很多业务代码,其实在函数体内做了很多事情,创造了很多变量和对象,这其实被称为函数的'副作用'
还是看累加器
```go
// 正统的函数式编程
// 只有常量和函数
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(i int) (int, iAdder) {
		return base + i, adder2(base + i)
	}
}

func main(){
    a := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Println(s)
	}
}
```

### 函数式编程入门
斐波那契数列
```go
func fib() func() int {
	a, b := 1, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

f := fib()
for i := 0; i < 10; i++ {
	fmt.Println(f())
}
// 1 2 3 ... 55 89
```

这是用 print 打印的 fib 数,之前说道了 read 和 write 这两个基本接口.
现在让 fib这个函数实现一个 read 接口,然后任何能接收 reader 的方法都能输出这个 fib 数了
```go
// 定义一个函数的结构体,用函数实现接口,函数和普通变量一样
type intGen func() int

func fib1() intGen {
	a, b := 1, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
```
把鼠标放到intGen 上,然后右键
![](http://oqb4aabpb.bkt.clouddn.com/18-8-11/45524464.jpg)
![](http://oqb4aabpb.bkt.clouddn.com/18-8-11/68743049.jpg)
![](http://oqb4aabpb.bkt.clouddn.com/18-8-11/73783184.jpg)

```go
func (g intGen) Read(p []byte) (n int, err error) {
	// 下一个 fib 数
	next := g()
	//fib 数读不完,需要有一个结束条件
	if next > 1000 {
		return 0, io.EOF
	}
	// 底层找一个已经实现的
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}
```
注释已经写得很清楚了,让 intGen 这个函数结构体实现 reader 接口,等会就可以写一个 接收 reader 参数的print 函数,把intGen函数当做参数传进去了
```go
/**
打印的方法
让 fib 实现 Reader 接口,就可以用 print 方法打印了
*/
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

f1 := fib1()
printFileContents(f1)

// 1 2 3 5 8 13 21 34 ... 610 987
```

### goimports
一个好用的工具
![](http://oqb4aabpb.bkt.clouddn.com/18-8-11/94132184.jpg)
能够自动整理imports
把没用到的去除,用到的,但系统没有的,自动 go get
但是正常是下不下来的,因为需要下载`golang.org/x/tools/cmd/goimports`,而 `golang.org` 在国内是被墙的

1. `go get -v github.com/gpmgo/gopm`,github 在国内没被墙,先下载 `gopm` 这个工具
2. 配置`$ GOPATH:bin`
3. `gopm get -v -g -u golang.org/x/tools/cmd/goimports`用 `gopm`下载谷歌的工具包
4. `go install golang.org/x/tools/cmd/goimports` 把 `goimports`安装到$ GOPATH 下


### 总结
- 函数是一等公民,可以当做参数传给另一个函数
- 本来一个函数只能打印,但是改写一下之后,让他能传一个函数进来,就可以干任何事了,扩展性非常强(根据传进来的函数来工作)
- 推荐一本函数式编程的书 `scip`





