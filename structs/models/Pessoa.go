package models

type Pessoa struct {
	Nome     string
	Idade    int
	Endereco Endereco
}

func New(nome string, idade int) Pessoa {
	return Pessoa{
		Nome:  nome,
		Idade: idade,
	}
}

func (p Pessoa) GetIdade() int {
	return p.Idade
}

// func New() Pessoa {
// 	return Pessoa{
// 		nome:  "",
// 		idade: 0,
// 	}
// }

// func NewPessoa(nome string, idade int) (Pessoa, error) {
// 	if nome == "" || idade == 0 {
// 		return Pessoa{}, errors.New("error")
// 	}

// 	return Pessoa{
// 		nome:  nome,
// 		idade: idade,
// 	}, nil
// }

func obterCPF() int {
	return 5
}
