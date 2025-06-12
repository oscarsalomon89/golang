package pessoa

import (
	"api-test/pessoa/endereco"
	"fmt"
)

type Pessoa struct {
	ID       string
	Nome     string
	Idade    int
	Endereco endereco.Endereco
	Carro    Carro
}

func (p Pessoa) SetCarro(matricula string) {
	fmt.Println(p.Nome)

	p.Carro = Carro{
		Name: matricula,
	}
}

type Carro struct {
	Name   string
	Modelo string
	Color  string
}

func NewCarro(nome, modelo, cor string) Carro {

	return Carro{
		Name:   nome,
		Modelo: modelo,
		Color:  cor,
	}
}
