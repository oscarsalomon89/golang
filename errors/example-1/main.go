package main

import (
	"errors"
	"fmt"
)

// errors.New(): Se utiliza para crear errores simples com uma mensagem fixa. Es más directo y adecuado para errores simples, sin formato.

// Uma função que retorna um erro se o número for negativo
func verificarNumero(n int) error {
	if n < 0 {
		// Se o número for negativo, retornamos um erro
		return errors.New("o número não pode ser negativo")
	}
	// Se não houver erro, retornamos nil
	return nil
}

func main() {
	// Testamos com um número negativo
	if err := verificarNumero(-5); err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("O número é válido")
	}

	// Testamos com um número positivo
	if err := verificarNumero(10); err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("O número é válido")
	}
}
