package main

import (
	"fmt"
	"math/rand"
	"time"
)

func computePrice(rate float32, nights int) float32 {
	return rate * float32(nights)
}

func computePrice2(rate float32, nights int) (price float32) {
	price = rate * float32(nights)
	return
}

func computePrice3(rate float32, nights int) (price float32) {
	p := rate * float32(nights)
	p = p * 2
	return
}

func vacantRooms() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(100)
}

func printHeader() {
	fmt.Println("Hotel Golang")
	fmt.Println("San Francisco, CA")
}

func main() {
	price1 := computePrice3(1.4, 2)
	price2 := computePrice3(2.4, 3)
	price3 := computePrice3(3.4, 4)

	total := price1 + price2 + price3
	fmt.Printf("TOTAL : %0.2f $\n", total)

	fmt.Println("Vacant Rooms:", vacantRooms())

	printHeader()
}
