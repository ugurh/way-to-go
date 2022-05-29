package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	s := "Hello " + "World"
	fmt.Println(s)
	fmt.Println("Lengt:", len(s))

	var str string = "This is an example of a string"
	fmt.Printf("\n%t\n\n", strings.HasPrefix(str, "Th")) // Finding prefix

	fmt.Printf("Does the string \"%s\" have suffix %s? ", str, "ting")
	fmt.Printf("\n%t\n\n", strings.HasSuffix(str, "ting")) // Finding suffix

	str = "Hi, I'm Marc, Hi."
	fmt.Printf("The position of the first instance of\"Marc\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Marc")) // Finding first occurence
	fmt.Printf("The position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Hi")) // Finding first occurence
	fmt.Printf("The position of the last instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str, "Hi")) // Finding last occurence
	fmt.Printf("The position of the first instance of\"Burger\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Burger")) // Finding first occurence

	str = "Hello, how is it going, Hugo?"
	var manyG = "gggggggggg"
	fmt.Printf("Number of H's in %s is: ", str)
	fmt.Printf("%d\n", strings.Count(str, "H")) // count occurences
	fmt.Printf("Number of double g's in %s is: ", manyG)
	fmt.Printf("%d\n", strings.Count(manyG, "gg")) // count occurences

	var origS string = "Hi there! "
	var newS string
	newS = strings.Repeat(origS, 3) // Repeating origS 3 times
	fmt.Printf("The new repeated string is: %s\n", newS)

	var orig = "Hey, how are you George?"
	fmt.Printf("The original string is: %s\n", orig)
	fmt.Printf("The lowercase string is: %s\n", strings.ToLower(orig))
	fmt.Printf("The uppercase string is: %s\n", strings.ToUpper(orig))

	intToStr := strconv.Itoa(5)
	fmt.Println(intToStr)

	strToInt, _ := strconv.Atoi("5")
	fmt.Println(strToInt)

}
