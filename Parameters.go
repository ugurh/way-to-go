package main

import "fmt"

var num1 int = 10
var num2, num3 int

func main() {
	// Call by value -> function(arg1)
	// Call by reference -> function(&arg1)
	fmt.Printf("Multiply 2 * 5 * 6 = %d\n", MultiplePly3Nums(2, 5, 6))

	num2, num3 = getX2AndX3(num1)
	PrintValues()
	num2, num3 = getX2AndX3V2(num1)
	PrintValues()

	var x1 int
	var x3 float32
	x1, _, x3 = ThreeValues() // Black identifier
	fmt.Printf("x1 = %d, x3 = %f\n", x1, x3)

	n := 0
	reply := &n
	fmt.Println("Before Multiplication:", n)
	fmt.Println("Before Multiplication:", *reply)

	Multiply(10, 5, reply)

	fmt.Println("After Multiplication:", n)
	fmt.Println("After Multiplication:", *reply)

	Print(1, 5, 6, 7, 0, 4)
	sum, prod, diff := SumProductDiff(3, 4)
	fmt.Println("sum, prod, diff = ", sum, prod, diff)

	result := SumNumbers(5, -2, 0, 9)
	fmt.Println("Sum of numbers :", result)
}

func Multiply(n1 int, n2 int, reply *int) {
	*reply = n1 * n2
}

func PrintValues() {
	fmt.Printf("num1 = %d, num2 = %d, num3 = %d\n", num1, num2, num3)
}

func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

func getX2AndX3V2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	return
}

func MultiplePly3Nums(a, b, c int) int {
	return a * b * c
}

func ThreeValues() (int, int, float32) {
	return 2, 3, 7.5
}

func Print(nums ...int) {
	if len(nums) == 0 {
		return
	}

	for _, value := range nums {
		fmt.Printf("The number is: %d\n", value)
	}
}

func SumProductDiff(x, y int) (int, int, int) {
	return x + y, x * y, x - y
}

func SumNumbers(list ...int) int {
	if len(list) == 0 {
		return 0
	}

	sum := 0
	for _, num := range list {
		sum += num
	}
	return sum
}
