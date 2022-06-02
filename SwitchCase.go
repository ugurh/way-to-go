package main

import "fmt"

func main() {

	switch num := 100; {
	case num < 0:
		fmt.Println("Number is negative")
	case num > 0 && num <= 100:
		fmt.Println("Number is between 0 and 100")
	default:
		fmt.Println("Number is greater than 100")
	}
}
