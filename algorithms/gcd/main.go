package main

import (
	"fmt"
)

func gcdSimple(a, b int) int {
	min := a
	if b < a {
		min = b
	}

	gcd := 1
	for i := 1; i <= min; i++ {
		if a%i == 0 && b%i == 0 {
			gcd = i
		}
	}

	return gcd
}

func main() {
	var num1, num2 int

	fmt.Print("Enter the first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter the second number: ")
	fmt.Scanln(&num2)

	if num1 <= 0 || num2 <= 0 {
		fmt.Println("Please enter positive integers only.")
		return
	}

	if num1 == num2 {
		fmt.Printf("The Greatest Common Divisor (GCD) of %d and %d is: %d\n", num1, num2, num1)
		return
	}

	result := gcdSimple(num1, num2)
	fmt.Printf("The Greatest Common Divisor (GCD) of %d and %d is: %d\n", num1, num2, result)
}

// // Função que calcula o MDC usando o algoritmo de Euclides
// func gcd(a, b int) int {
// 	for b != 0 {
// 		a, b = b, a%b
// 	}
// 	return a
// }

// func main() {
// 	var num1, num2 int

// 	fmt.Print("Enter the first number: ")
// 	fmt.Scanln(&num1)

// 	fmt.Print("Enter the second number: ")
// 	fmt.Scanln(&num2)

// 	if num1 <= 0 || num2 <= 0 {
// 		fmt.Println("Please enter positive integers only.")
// 		return
// 	}

// 	result := gcd(num1, num2)
// 	fmt.Printf("The Greatest Common Divisor (GCD) of %d and %d is: %d\n", num1, num2, result)
// }
