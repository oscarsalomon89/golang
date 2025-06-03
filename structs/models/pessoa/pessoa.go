package pessoa

type Pessoa struct {
	Nome  string
	Idade int
}

func GetIdade(p Pessoa) int {
	return p.Idade
}
