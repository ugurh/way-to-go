package main

import (
	"fmt"
	"io"
	"log"
)

/*
	Defer allow us to guarantee that
	- Closing file stream
	- Unlocking a locked resource
	- Printing a footer in a report
	- Closing a database connection
*/
func main() {

	trace()
	defer untrace()
	Function2()
	fmt.Println("Ended...")
	Defer()

	Logging(4)
}

func Function1() {
	fmt.Println("Function1 executed...")
}

func Function2() {
	fmt.Println("Function2 executed...")
	defer Function1()
	fmt.Println("Function2 ended...")
}

func Defer() {
	for i := 0; i < 6; i++ {
		defer fmt.Println("i: ", i)
	}
}

func trace() {
	fmt.Println("Entering....")
}

func untrace() {
	fmt.Println("Leaving....")
}

func Logging(number int) (n int, err error) {
	defer func() {
		log.Printf("Logging(%q) = %d, %v\n", number, n, err)
	}()
	return 0, io.EOF
}
