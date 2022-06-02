package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// value, ok = sample_function(parameters)
	var orig string = "csac"

	num, err := strconv.Atoi(orig)

	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
		return
	}

	fmt.Printf("The integer is %d\n", num)

	f, err := os.Open("/file")

	if err != nil {
	}

	doSomething(*f)
}

func doSomething(f os.File) {

}
