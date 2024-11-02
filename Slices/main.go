package main

import "fmt"

func main() {
	s := make([]int, 3)
	s[0] = 12
	s[1] = 14

	customers := []string{"ABC", "DEF", "GHI", "JKL"} //not providing any size
	customersSlice := customers[0:2]
	fmt.Println(customersSlice)
	anotherCustomerSlice := customers[2:4]
	fmt.Println(anotherCustomerSlice)

	//modify
	customers[0] = "XYZ"
	fmt.Println(customersSlice)
}
