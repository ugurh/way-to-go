package main

import (
	"errors"
	"fmt"
	"log"
	"os/user"
	"time"
)

type Item struct {
	ID   string
	Name string
}

type Cart struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	lockedAt  time.Time
	user.User
	Items        []Item
	CurrencyCode string
	isLocked     bool
}

func (c *Cart) TotalPrice() (float32, error) {
	return 0, nil
}

func (c *Cart) Lock() error {
	if c.isLocked {
		return errors.New("cart is already locked")
	}
	c.isLocked = true
	c.lockedAt = time.Now()
	return nil
}

func Display(c *Cart) float64 {
	return 0
}

func main() {
	newCart := Cart{}

	totalPrice, err := newCart.TotalPrice()
	if err != nil {
		log.Printf("impossible to compute price of the cart: %s\n", err)
		return
	}
	log.Println("Total Price", totalPrice)

	err = newCart.Lock()
	if err != nil {
		log.Printf("impossible to lock the cart: %s", err)
		return
	}

	fmt.Println(newCart.isLocked, newCart.lockedAt)

}
