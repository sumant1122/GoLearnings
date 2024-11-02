package main

import "fmt"

func main() {

	const hotelName string = "Gopher Hotel"
	const longitude = 24.806078
	const latitude = -78.243027
	var occupancy int = 12
	//occupancy := 12 //<- can also be wriiten like this
	//var occupancy = 12 //<- can also be written like this
	fmt.Println("Hotel:", hotelName, "Lng:", longitude, "Lat:", latitude)
	fmt.Println("============Rooms Available==============")
	fmt.Println(occupancy)

}
