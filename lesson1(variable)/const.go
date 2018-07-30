package main

import (
	"fmt"
	"math"
)

func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c) // abc.txt 5
}

func enums() {
	const (
		cpp = iota
		_
		python
		golang
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
	)
	fmt.Println(cpp, python, golang) // 0 2 3
	fmt.Println(kb, mb, gb, tb)      // 1024 1048576 1073741824 1099511627776
}

func main() {
	consts()
	enums()
}
