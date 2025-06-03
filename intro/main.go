package main

import (
	"fmt"
)

// func main() {
// 	var ola string = "hello"
// 	var exist bool
// 	number := 10

// 	number = 5

// 	fmt.Println(ola)
// 	fmt.Println(exist)
// 	fmt.Println(number)
// }

// func main() {
// 	number := 10

// 	//true            //true
// 	if number > 0 && !(number == 15) {
// 		fmt.Println("ok")
// 	} else {
// 		fmt.Println("error")
// 	}
// }

// func main() {
// 	number := 10

// 	switch {
// 	case number > 0:
// 		fmt.Println("o número é positivo")
// 	case number < 0:
// 		fmt.Println("o número é negativo")
// 	default:
// 		fmt.Println("o número é zero")
// 	}
// }

// func main() {
// 	// 	// ARRAY: tamanho fixo, declarado com [n]
// 	//var array [3]int = [3]int{1, 2, 3}
// 	array := [3]int{1, 2, 3}
// 	array[1] = 5

// 	fmt.Println("Array:", array)
// 	fmt.Printf("Tamanho do array: %d\n", len(array))

// 	// SLICE: tamanho dinâmico, declarado com []
// 	slice := []int{1, 2, 3, 4, 5}
// 	fmt.Println("Slice:", slice)
// 	fmt.Println("Tamanho do slice:", len(slice))

// 	// Podemos adicionar elementos no slice (não no array)
// 	slice = append(slice, 14)
// 	fmt.Println("Slice depois do append:", slice)
// 	//1 2 3 4 5 14

// 	//slice = append(slice[:index], slice[index+1:]...)
// 	// Usamos append para "deletar" o elemento na posição `index`.
// 	// slice[:index] pega todos os elementos antes do índice.
// 	// slice[index+1:] pega todos os elementos depois do índice.
// 	// O append junta esses dois pedaços, pulando o elemento na posição `index`.

// 	// 1 2 3 4 5 14
// 	slice = slices.Delete(slice, 1, 3)

// 	fmt.Println("Slice depois de remover:", slice)
// }

func main() {
	hello("oscar", "salomon", 35)

	substracao, err, message := resta(14, 6)
	if err {
		fmt.Println("Error substracao")
	} else {
		fmt.Println("Resultado substracao:", substracao)
	}
	fmt.Println(message)

	// substracao, _, _ := resta(4, 6)
	// fmt.Println("Resultado substracao:", substracao)

	resultado := soma(10, 5)
	fmt.Println("O resultado da soma é:", resultado)
}

func soma(a int, b int) int {
	return a + b
}

func resta(a, b int) (int, bool, string) {
	if a < b {
		return 0, true, "error"
	}

	return a - b, false, "ok"
}

func hello(nome, sobrenome string, idade int) {
	fmt.Println("Hello", nome, sobrenome, idade)
}
