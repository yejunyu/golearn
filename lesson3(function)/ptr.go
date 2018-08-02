package main

import "fmt"

func passByVal(a int) {
	a++
}

func passByRef(a *int) {
	*a++
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	a := 3
	passByVal(a)
	fmt.Println(a)
	passByRef(&a)
	fmt.Println(a)
	b := 3
	c := 4
	swap(&b, &c)
	fmt.Println(b, c)
}
