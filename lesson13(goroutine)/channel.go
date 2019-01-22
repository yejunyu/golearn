package main

import (
	"fmt"
	"time"
)

//func chanDemo() {
//	c := make(chan int)
//	go func() {
//		for {
//			n := <-c
//			fmt.Println(n)
//		}
//	}()
//	c <- 1
//	c <- 2
//	time.Sleep(time.Second)
//}
func worker(id int, c chan int) {
	for {
		fmt.Printf("Worker %d received %c\n", id, <-c)
	}
}

func createWorker(id int) chan int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker %d received %c\n", id, <-c)
		}
	}()
	return c
}

func chanDemo() {
	var channels [10]chan int
	//for i := 0; i < 10; i++ {
	//	channels[i] = make(chan int)
	//	go worker(i, channels[i])
	//}
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Second)
}

func main() {
	chanDemo()
}
