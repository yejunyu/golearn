package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:])
	s1 := arr[2:]
	s2 := arr[:]
	fmt.Println(s2)
	// Slice 本身没有数据,只是对底层 array 的一个 view
	fmt.Println("After updateSlice")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("reSlice")
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	// cap的概念
	arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 = arr[2:6]
	s2 = s1[3:6]
	fmt.Println("s1 = ", s1)
	fmt.Println("s2 = ", s2)
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n",
		s1, len(s1), cap(s1))
	fmt.Printf("s2=%v,len(s2)=%d,cap(s2)=%d\n",
		s2, len(s2), cap(s2))

}
