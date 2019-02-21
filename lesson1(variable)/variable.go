package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

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

// 函数外面就不能用简写定义了
var (
	a = 3
	b = "str"
)

func euler() {
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	// Exp表示以e为底
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	// (0+1.2246467991473515e-16i)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b))) // go 只有强制类型转换
	fmt.Println(c)                         // 5

}

func main() {
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	euler()
	triangle()
}
