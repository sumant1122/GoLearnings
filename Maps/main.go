package main

import "fmt"

func main() {
	m := make(map[int]string)

	m[1] = "Hello"
	m[2] = "World"

	fmt.Println(m)

	result, ok := m[5]
	if ok {
		fmt.Println(result)
	} else {
		fmt.Println("not found")
	}

	for k, v := range m {
		fmt.Printf("key: %d - Value: %s\n", k, v)
	}
}
