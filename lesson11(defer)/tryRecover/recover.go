package main

import (
	"errors"
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(r)
		}
	}()
	panic(errors.New("this is a new error"))
}

func main() {
	tryRecover()
}
