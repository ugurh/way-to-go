package main

import "fmt"

func main() {

	language := "Go"
	fmt.Printf("A string: %s, it's location in memory: %p\\n", language, &language)

	var langP *string // Pointer variable
	langP = &language // Storing address of language in pointer variable
	fmt.Printf("\nThe value at memory location %p is %s\n", langP, *langP)

	fmt.Println("---------------------------------------")
	s := "good bye"
	var p *string = &s
	*p = "ciao" // changing the value at &s

	fmt.Printf("Here is the pointer p: %p\n", p)  // prints address
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	fmt.Printf("Here is the string s: %s\n", s)   // prints same string

	var x *int = nil // Making a nil pointer
	*x = 0

	fmt.Println(&x)

}
