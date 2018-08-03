package main

import (
	"fmt"
	"strconv"
)

func array() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{1, 3, 5, 7, 9}
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3, grid)
}

func arrRange() {
	arr1 := [...]int{2, 4, 6, 8, 10}
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	for key, value := range arr1 {
		fmt.Println(strconv.Itoa(key) + "=>" + strconv.Itoa(value))
	}
}

func printArray(arr [5]int) {
	for _, v := range arr {
		println(v)
	}
	arr[0] = 100
}

func printArray2(arr *[5]int) {
	arr[0] = 100
	for _, v := range arr {
		println(v)
	}
}

func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	printArray(arr1)
	// printArray2(&arr1)
	fmt.Println(arr1)
}
