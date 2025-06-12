package main

import (
	"errors"
	"fmt"
)

type ErroSaldoInsuficiente struct{}

func (e ErroSaldoInsuficiente) Error() string {
	return "saldo insuficiente para a operação"
}

func sacar(valor, saldo float64) error {
	if valor > saldo {
		return ErroSaldoInsuficiente{}
	}
	return nil
}

func main() {
	err := sacar(100, 50)
	if err != nil {
		if errors.Is(err, ErroSaldoInsuficiente{}) {
			fmt.Println("Erro: Saldo insuficiente (erro original).")
		} else {
			fmt.Println("Erro:", err)
		}
	}

	fmt.Println("fin")
}
