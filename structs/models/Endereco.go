package models

type Endereco struct {
	Rua    string
	Cidade string
}

func GetCPF() int {
	return obterCPF()
}
