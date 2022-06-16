package main

import "fmt"

func main() {

	// for initialization; condition; modification { }

	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}

	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))

	// for loop to find character at each position
	for ix := 0; ix < len(str); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str[ix])
	}
	str2 := "Chinese: 日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))

	// for loop to find character at each position
	for ix := 0; ix < len(str2); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str2[ix])
	}

	// for range
	for pos, char := range str {
		fmt.Printf("Character on position %d is: %c \n", pos, char)
	}

LABEL1: // adding a label for location
	for i := 0; i <= 5; i++ { // outer loop
		for j := 0; j <= 5; j++ { // inner loop
			if j == 4 {
				continue LABEL1 // jump to the label
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
}
