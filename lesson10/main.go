package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

// 正统的函数式编程
// 只有常量和函数
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(i int) (int, iAdder) {
		return base + i, adder2(base + i)
	}
}

func fib() func() int {
	a, b := 1, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

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

// 定义一个函数的结构体,用函数实现接口,函数和普通变量一样
type intGen func() int

func fib1() intGen {
	a, b := 1, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

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

func main() {
	adder := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(adder(i))
	}
	fmt.Println("======")
	a := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Println(s)
	}
	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	f1 := fib1()
	printFileContents(f1)
}
