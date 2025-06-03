package pessoa

import (
	"modules-example/messagem"

	"github.com/google/uuid"
)

type Pessoa struct {
	ID                   string
	Nome                 string
	Idade                int
	MensagemDeBoasVindas string
}

func New(name string, age int) Pessoa {
	id := uuid.New().String()
	saudacao := messagem.GerarSaudacao(name)

	return Pessoa{
		ID:                   id,
		Nome:                 name,
		Idade:                age,
		MensagemDeBoasVindas: saudacao,
	}
}
