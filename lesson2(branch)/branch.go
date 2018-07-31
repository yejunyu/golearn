package main

import (
	"fmt"
	"io/ioutil"
)

func readFile() {
	const filename = "./lesson2(branch)/abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 1000:
		g = "A"
	default:
		panic(fmt.Sprintf("Wrong score: %s", score))
	}
	return g
}

func main() {
	readFile()
}
