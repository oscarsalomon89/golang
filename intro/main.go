package main

import (
	"fmt"
)

// func main() {
// 	var ola string = "hello"
// 	var exist bool
// 	number := 10

// 	fmt.Println(ola)
// 	fmt.Println(exist)
// 	fmt.Println(number)
// }

// func main() {
// 	number := 10

// 	if number > 0 {
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
// 	// var array [3]int = [3]int{1, 2, 3}
// 	//array := [3]int{1, 2, 3}

// 	// fmt.Println("Array:", array)
// 	// fmt.Printf("Tamanho do array: %d\n", len(array))

// 	// SLICE: tamanho dinâmico, declarado com []
// 	slice := make([]int, 5)
// 	slice = []int{1, 2, 3, 4, 5}
// 	fmt.Println("Slice:", slice)
// 	fmt.Println("Tamanho do slice:", len(slice))

// 	// Podemos adicionar elementos no slice (não no array)
// 	slice = append(slice, 14)
// 	fmt.Println("Slice depois do append:", slice) //1, 2, 3, 4, 5, 14

// 	slice = slices.Delete(slice, 1, 2)
// 	fmt.Println("Slice depois de remover:", slice) //1, 4, 5, 14
// }

// func main() {
// 	substracao, _, _ := resta(14, 6)

// 	fmt.Println("Resultado substracao:", substracao)

// 	// substracao, _, _ := resta(4, 6)
// 	// fmt.Println("Resultado substracao:", substracao)

// 	resultado := soma(10, 5)
// 	fmt.Println("O resultado da soma é:", resultado)
// }

// func soma(a, b int) int {
// 	return a + b
// }

// func resta(a, b int) (int, bool, string) {
// 	if a < b {
// 		return 0, true, "error"
// 	}

// 	return a - b, false, "ok"
// }

// func hello(nome, sobrenome string, idade int) {
// 	fmt.Println("Hello", nome, sobrenome, idade)
// }

func main() {
	// Criando um map onde a chave é string e o valor é int
	notas := make(map[string]int)

	// Adicionando elementos ao map
	notas["João"] = 8
	notas["Maria"] = 9

	// Acessando um valor pelo nome da chave
	fmt.Println("Nota do João:", notas["João"])

	// Percorrendo todos os elementos do map
	for nome, nota := range notas {
		fmt.Println(nome, "tirou", nota)
	}
}
