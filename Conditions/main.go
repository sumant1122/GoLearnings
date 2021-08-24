package main

import (
	"fmt"
)

func main() {
	a := 5
	b := 10

	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a == b)

	if a > b {
		fmt.Println("a is greater than b")
	}
	fmt.Println("b is greater than a")

	ageJohn := 5
	agePaul := 10
	switch ageSum := ageJohn + agePaul; ageSum {
	case 15:
		fmt.Println("ageJohn + agePaul = 15")
	case 20, 30, 40: //Multiple
		fmt.Println("ageJohn + agePaul = 20 or 30 or 40")
	case 2 * agePaul:
		fmt.Println("ageJohn + agePaul = 2 times agePaul")
	default:
		fmt.Println("Default")
	}

}
