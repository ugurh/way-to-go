package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	const firstRoomNumber = 101

	hotelName := "The Gopher Hotel"
	totalRooms, roomsOccupied := 100, rand.Intn(100)
	roomsAvailable := totalRooms - roomsOccupied
	occupancyRate := (float32(roomsOccupied) / float32(totalRooms)) * 100

	var occupancyLevel string
	switch {
	case occupancyRate < 20:
		occupancyLevel = "Low"
	case occupancyRate < 70:
		occupancyLevel = "Medium"
	default:
		occupancyLevel = "High"
	}

	fmt.Println("Hotel:", hotelName)
	fmt.Println("Number of rooms", totalRooms)
	fmt.Println("Rooms available", roomsAvailable)
	fmt.Println("Occupancy Level:", occupancyLevel)
	fmt.Printf("Occupancy Rate: %0.2f %%\n", occupancyRate)

	for i := 0; roomsAvailable > i; i++ {
		roomNumber := firstRoomNumber + i
		size := rand.Intn(6) + 1
		nights := rand.Intn(10) + 1
		fmt.Printf("%d: %d people / %d nights\n", roomNumber, size, nights)
	}

}
