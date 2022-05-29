package main

import (
	"fmt"
	"math/rand"
)

var number int = 5

func main() {

	var decision bool = true
	println("Number :", number)
	number = 10
	println("Number Updated: ", number)
	println("Decision : ", decision)

	var a int
	var b int32
	a = 15
	// b = a + a  // compiler error
	b = b + 5        // ok. 5 is constant
	b = b + int32(a) // explicit typing
	fmt.Printf("a: %d\nb: %d\n", a, b)

	// Complex Numbers (re + im¡) (real part, imaginary part)
	var c1 complex64 = 5 + 10i
	fmt.Printf("The value is: %v:", c1)

	c2 := complex(5, 20) // create complex number using complex function
	fmt.Printf("\nComplex value: %v", c2)

	// Random Numbers

	// generates a random number
	x := rand.Int()
	fmt.Printf("x is:%d\n", x)

	// generates a random number in [0, n)
	y := rand.Intn(8)
	fmt.Printf("y is:%d\n", y)

	// Character type
	var ch1 byte = 'A'
	var ch2 byte = 65
	var ch3 byte = '\x41'

	fmt.Printf(string(ch1), ch2, ch3)

	fmt.Printf("----------------------------------")

}

/* Bitwise operators
Bitwise AND operator &
Bitwise OR operator |
Bitwise XOR operator ^
Bit CLEAR operator &^
Bitwise COMPLEMENT operator ^
*/
