package main

import "fmt"

func main() {

	var result = func() {
		fmt.Println("Function without name")
	}
	result()

	var sum = func(a int, b int) {
		sum := a + b
		fmt.Println("Sum is ", sum)
	}
	sum(2, 4)

	a := displayNumber()
	fmt.Println(a())
}

func displayNumber() func() int { //returning a function : This is also called Closure. A closure is a function that has access to its own scope and the scope of its surrounding functions.

	number := 10
	return func() int {
		number++
		return number
	}
}
