package main

import (
	"log"
	"os/user"
	"time"
)

type Cart struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	lockedAt  time.Time
	user.User
	CurrencyCode string
	isLocked     bool
}

func (c *Cart) TotalPrice() (*money.Money, error) {
	//...
	return nil, nil
}

func (c *Cart) Lock() error {
	//...
	return nil
}

func (c *Cart) delete() error { //not public as method name is lowercase
	// to implement
	return nil
}

func main() {
	newCart := cart.Cart{}

	totalPrice, err := newCart.TotalPrice()
	if err != nil {
		log.Printf("impossible to compute price of the cart: %s", err)
		return
	}
	log.Println("Total Price", totalPrice.Display())

	err = newCart.Lock()
	if err != nil {
		log.Printf("impossible to lock the cart: %s", err)
		return
	}

}
