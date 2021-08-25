package main

import "fmt"

func main() {

	PrintHeader()

}

func PrintHeader() { //Merthod first letter is captial so it's public scope else it's private
	fmt.Println("San Francisco, CA")
}
