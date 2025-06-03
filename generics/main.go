package main

import "fmt"

func main() {
	result := Max(5, 10.5)
	fmt.Println("result:", result)
}

// func Max(a, b any) any {
// 	switch a := a.(type) {
// 	case int:
// 		if a > b.(int) {
// 			return a
// 		}
// 		return b
// 	case float64:
// 		if a > b.(float64) {
// 			return a
// 		}
// 		return b
// 	default:
// 		panic("Tipo não suportado")
// 	}
// }

func Max[T int | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}
