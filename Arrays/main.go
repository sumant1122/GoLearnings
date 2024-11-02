package main

import "fmt"

func main() {
	var myArray [2]int
	myArray[0] = 124
	myArray[1] = 112
	fmt.Println(myArray)

	//another way
	anotherArray := [5]int{123, 146, 434, 676, 345}
	fmt.Println(anotherArray)

	//length
	fmt.Println("length of myArray:", len(myArray))

	//loop
	for index, element := range anotherArray {
		fmt.Printf("element at index %d is %d\n", index, element)
	}

}
