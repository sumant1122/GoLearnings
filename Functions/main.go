package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	PrintHeader()
	price := ComputePrice(10, 4)
	fmt.Println("Price:", price)
	vacant := vacantRooms()
	fmt.Println("Vacant Rooms:", vacant)

}

func ComputePrice(rate float32, nights int) float32 {
	return rate * float32(nights)

}

func PrintHeader() { //Merthod first letter is captial so it's public scope else it's private
	fmt.Println("San Francisco, CA")
}

func vacantRooms() int {
	rand.Seed(time.Now().UTC().UnixMicro())
	return rand.Intn(100)
}
