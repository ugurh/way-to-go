package main

import "fmt"

const version string = "1.3.2"

func main() {

	roomNumber, floorNumber := 154, 3
	fmt.Println("Room No, Floor No", roomNumber, floorNumber)

	var wifiPassword string
	fmt.Println("Wifi Password: ", wifiPassword)
	fmt.Println(version)

	const limit = 12

	var limit1 uint8 = limit
	var limit2 uint16 = limit
	var limit3 float32 = limit

	fmt.Println(limit1, limit2, limit3)

	const name = "No Name"
	var limit4 string
	limit4 = name
	fmt.Println(limit4)

	fmt.Println("==================")

	const isOpen = true
	const MyRune = 'r'
	const occupancyLimit = 12
	const rate = 29.87
	const complexNumber = 1 + 1i
	const hotelName = "Asya Hotel"

	fmt.Println(isOpen, MyRune, occupancyLimit, rate, complexNumber, hotelName)

	const profit = 9223372036854775807
	var profit2 int64 = profit
	fmt.Println(profit2)

	const longitude = 24.806078
	const latitude = -78.243027
	var occupancy = 12
	fmt.Println(hotelName, longitude, latitude)
	fmt.Println(occupancy)
}
