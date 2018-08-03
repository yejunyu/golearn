package main

import "fmt"

func main() {
	m := map[string]string{
		"name":     "神奇大叶子",
		"age":      "26",
		"language": "golang",
	}
	m2 := make(map[string]string) // m2 == empty map
	var m3 map[string]string      // m3 == nil
	fmt.Println(m, m2, m3)

	fmt.Println(m2 == nil, m3 == nil)

	for key, value := range m {
		fmt.Println(key + " => " + value)
	}

	fmt.Println("Getting values")
	name, ok := m["name"]
	fmt.Println(name, ok)

	if name1, ok := m["name1"]; ok {
		fmt.Println(name1)
	} else {
		fmt.Println("Key not exist")
	}

	fmt.Println("Deleting values")
	delete(m, "age")
	age, ok := m["age"]
	fmt.Println(age, ok)

}
