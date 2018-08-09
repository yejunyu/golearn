package main

import (
	"bufio"
	"fmt"
	"golearn/lesson11/fib"
	"os"
	"strconv"
)

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	defer write.Flush()

	f := fib.Fib()
	for i := 0; i < 20; i++ {
		fmt.Fprint(write, strconv.Itoa(f())+"\n")
	}
}

func main() {
	tryDefer()
	writeFile(".fib.txt")
}
