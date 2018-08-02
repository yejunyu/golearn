package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

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

/**
多返回值
*/
func div(a, b int) (q, r int) {
	return a / b, a % b
}

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

func main() {
	fmt.Println(evalfunc(pow, 3, 4))

	// 匿名函数
	fmt.Println(evalfunc(
		func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))

	fmt.Println(sum(1, 2, 3))
}
