package main

import (
	"fmt"
	"golearn/lesson9(duckTyping)/mock"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	url := "abcd"
	return r.Get(url)
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is a fake mock"}
	fmt.Println(download(r))
	fmt.Println(r)
	inspect(r)
	// Type assertion
	if mockRetriever, ok := r.(mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}
}

// type assertion
func inspect(r Retriever) {
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents: ", v.Contents)
	case *mock.Retriever:
		fmt.Println("point: ", v.Contents)
	}
}
