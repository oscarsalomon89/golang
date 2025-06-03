package main

import (
	"fmt"
)

func fibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}

	seq := make([]int, n)
	seq[0] = 0

	if n > 1 {
		seq[1] = 1
		for i := 2; i < n; i++ {
			seq[i] = seq[i-1] + seq[i-2]
		}
	}

	return seq
}

func main() {
	var n int
	fmt.Print("Digite o número de termos da sequência de Fibonacci: ")
	fmt.Scan(&n)

	resultado := fibonacci(n)
	fmt.Println("Sequência de Fibonacci:", resultado)
}
