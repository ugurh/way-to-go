package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		// value, ok = sample_function(parameters)
		var orig string = "csac"

		num, err := strconv.Atoi(orig)

		if err != nil {
			fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
			//return
		}

		fmt.Printf("The integer is %d\n", num)

		f, err := os.Open("/file")

		if err != nil {
		}

		doSomething(*f)
	*/
	lastName := "Harun"
	greeting(lastName)
	fmt.Println("variable lastName is still: ", lastName)

	//Function Call within a Function Call
	fmt.Print(f2(f1(4, 3)))

}

func doSomething(f os.File) {
}

func greeting(name string) {
	fmt.Println("In greeting: Hi!!!!!", name)
	name = "Ugur"
	fmt.Println("In greeting: Hi!!!!!", name)
}

func f1(a, b int) (int, int, int) {
	n1 := a + b
	n2 := a - b
	n3 := a * b
	return n1, n2, n3
}

func f2(a, b, c int) int {
	return a + b + c
}
