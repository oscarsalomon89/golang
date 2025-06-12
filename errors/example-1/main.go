package main

import (
	"errors"
	"fmt"
)

// errors.New(): Se utiliza para crear errores simples com uma mensagem fixa. Es más directo y adecuado para errores simples, sin formato.
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("não é possível dividir por zero")
	}
	return a / b, nil
}

func main() {
	resultado, err := dividir(10, 2)
	if err != nil {
		fmt.Println("Erro:", err)
	}

	fmt.Printf("Resultado: %.2f\n", resultado)

}
