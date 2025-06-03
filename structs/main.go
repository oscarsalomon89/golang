package main

import (
	"api/models"
	"fmt"
)

func main() {
	pessoa := models.Pessoa{
		Nome:  "Oscar",
		Idade: 35,
		Endereco: models.Endereco{
			Rua:    "rua 1",
			Cidade: "Sao Paulo",
		},
	}

	fmt.Println(eMaior(pessoa.Idade))

	idade := pessoa.GetIdade()
	fmt.Println(idade)

	//pessoa, _ := models.NewPessoa("Oscar", 35)
	fmt.Println(pessoa.Endereco.Cidade)
	//Apresentar(pessoa)
	// var pessoa2 models.Pessoa
	// pessoa2.Nome = "Fernando"
	// pessoa2.Idade = 20

	//fmt.Println(pessoa2)
}

func eMaior(idade int) bool {
	if idade > 18 {
		return true
	}

	return false
}

func Apresentar(p models.Pessoa) {
	fmt.Println(p)
}
