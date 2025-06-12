package main

import (
	person "api-test/pessoa"
	"api-test/pessoa/endereco"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	p := person.Pessoa{
		ID:    uuid.New().String(),
		Nome:  "Salvador",
		Idade: 30,
	}

	p.SetCarro("123456789")

	p2 := person.Pessoa{
		ID:    uuid.New().String(),
		Nome:  "Oscar",
		Idade: 35,
		Endereco: endereco.Endereco{
			Logradouro: "Rua 1",
			Cidade:     "Cidade 1",
			Estado:     "Estado 1",
		},
	}

	carro := person.NewCarro("Carro", "Fiat", "Vermelho")

	fmt.Println(p)
	fmt.Println(p2)
	fmt.Println(carro)

	fmt.Println(Soma(1, 2, 3, 4, 5, 10, 8))
}

func Soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
