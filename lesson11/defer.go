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

func openFile(path string) string {
	file, err := os.Open(path)
	// 对已知的问题的处理
	if err != nil {
		if pathError, ok := err.(*os.PathError); ok {
			fmt.Println(pathError.Err)
		} else {
			// 未知问题特殊处理
			fmt.Println("Unkown error", err)
		}
	}
	return file.Name()
}

func main() {
	tryDefer()
	writeFile("fib.txt")
}
