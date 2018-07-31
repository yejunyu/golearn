package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

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

func main() {
	fmt.Println(sum(5))
	fmt.Printf(convertToBin(5))
	printFile("./lesson2(branch)/abc.txt")
	forever()
}
