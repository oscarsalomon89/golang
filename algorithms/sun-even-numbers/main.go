package main

import (
	"fmt"
)

func sumEvenNumbersUpTo(n int) int {
	sum := 0
	for i := 2; i <= n; i += 2 {
		sum += i
	}
	return sum
}

func main() {
	var n int
	fmt.Print("Enter a number N: ")
	fmt.Scanln(&n)

	result := sumEvenNumbersUpTo(n)
	fmt.Printf("The sum of even numbers up to %d is: %d\n", n, result)
}
