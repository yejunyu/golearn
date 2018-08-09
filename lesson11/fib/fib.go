package fib

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func PrintFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 定义一个函数的结构体,用函数实现接口,函数和普通变量一样
type intGen func() int

func Fib() intGen {
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
