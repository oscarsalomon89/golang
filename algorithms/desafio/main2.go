package main

import (
	"fmt"
)

type Livro struct {
	ID         int
	Titulo     string
	Autor      string
	Ano        int
	Quantidade int
	Disponivel bool
}

const disponivel = "Sim"
const naoDisponivel = "Não"

func (e Livro) Exibir() string {
	disponibilidade := disponivel
	if !e.Disponivel {
		disponibilidade = naoDisponivel
	}

	// msg := "ID:1, Titulo: Harry Potter e a Pedra Filosofal, Quantidade: 5. Disponivel: Sim"
	msg := fmt.Sprintf("ID: %d, Titulo: %s, Quantidade: %d. Disponivel: %s", e.ID, e.Titulo, e.Quantidade, disponibilidade)

	return msg
}

func (l Livro) LivroEmprestado() {
	if l.Quantidade > 0 {
		fmt.Println("O livro está disponivel")
	} else {
		fmt.Println("O livro não está disponivel")
	}
}

func (d *Livro) LivroDevolvido() {
	d.Quantidade += 1
	fmt.Println("Livro devolvido com sucesso:", d.Quantidade)
}

func listar(biblioteca []Livro) {
	for _, livro := range biblioteca {
		msg := livro.Exibir()
		fmt.Println(msg)
	}
}

func main() {
	Livro1 := Livro{
		ID:         1,
		Titulo:     "Harry Potter e a Pedra Filosofal",
		Autor:      "Oscar",
		Quantidade: 5,
		Disponivel: true,
	}

	Livro2 := Livro{
		ID:         2,
		Titulo:     "Entendendo Algoritmos",
		Autor:      "Daiana",
		Ano:        2015,
		Quantidade: 6,
		Disponivel: false,
	}

	biblioteca := []Livro{
		Livro1,
		Livro2,
		{ID: 3, Titulo: "O Cortiço", Autor: "Daniel", Ano: 2015, Quantidade: 8, Disponivel: true},
	}

	proximoID := 4

	for {
		fmt.Println("----MENU----")
		fmt.Println("1 - Listar Livros")
		fmt.Println("2 - Adicionar livros")
		fmt.Println("3 - Emprestar Livro")
		fmt.Println("4 - Devolver Livro")
		fmt.Println("0 - Sair")

		var opcao string
		fmt.Scan(&opcao)

		switch opcao {
		case "1":
			listar(biblioteca)
		case "2":
			var titulo, autor string
			var ano, quantidade int

			fmt.Print("Título: ")
			fmt.Scanln(&titulo)
			fmt.Print("Autor: ")
			fmt.Scanln(&autor)
			fmt.Print("Ano: ")
			fmt.Scanln(&ano)
			fmt.Print("Quantidade: ")
			fmt.Scanln(&quantidade)

			Livro := Livro{ID: proximoID, Titulo: titulo, Ano: ano, Quantidade: quantidade, Disponivel: true}
			biblioteca = append(biblioteca, Livro)
			proximoID++
			fmt.Println("Livro adicionado com sucesso.")

		case "3":
			var id int

			fmt.Print("ID: ")
			fmt.Scanln(&id)

			for livroID := range biblioteca {
				livroSelecionado := biblioteca[livroID]

				if livroSelecionado.ID == id {
					if livroSelecionado.Quantidade > 0 {
						livroSelecionado.LivroEmprestado()
						biblioteca[livroID].Quantidade--

						fmt.Println("Livro emprestado com suceso.")
					} else {
						fmt.Println("Livro não disponível.")
					}
				}
			}

		case "4":
			var id int
			fmt.Print("ID: ")
			fmt.Scanln(&id)

			for livroID := range biblioteca {
				if biblioteca[livroID].ID == id {
					if biblioteca[livroID].Quantidade >= 0 {
						(&biblioteca[livroID]).LivroDevolvido()
						fmt.Println("Livro devolvido com sucesso.")
					} else {
						fmt.Println("Livro não devolvido com sucesso")
					}
				}
			}

		}
	}
}
