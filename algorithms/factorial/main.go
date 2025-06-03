package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	var n int
	fmt.Print("Enter a number N: ")
	fmt.Scanln(&n)

	if n < 0 {
		fmt.Println("Factorial is not defined for negative numbers.")
		return
	}

	result := factorial(n)
	fmt.Printf("The factorial of %d is: %d\n", n, result)
}
