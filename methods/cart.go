package main

import (
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
	return nil
}

func Display(c *Cart) float64 {
	return 0
}
