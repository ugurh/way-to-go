package main

import "fmt"

// Celsius Temperature Conversion To Fahrenheit
type Celsius float32
type Fahrenheit float32

func main() {

	// Temperature in ˚F = (Temp in °C * 9/5) + 32

	result := toFahrenheit(100)

	fmt.Printf("%f ˚C is equal to %f ˚F", Celsius(100), result)
}

func toFahrenheit(t Celsius) Fahrenheit {
	return Fahrenheit(t*1.8 + 32)
}
