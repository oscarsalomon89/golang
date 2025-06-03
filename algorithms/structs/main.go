package main

import (
	"fmt"
)

type Livro struct {
	ID         int
	Titulo     string
	Autor      string
	Ano        int
	Disponivel bool
}

func (l Livro) Exibir() {
	disponibilidade := "Sim"
	if !l.Disponivel {
		disponibilidade = "Não"
	}
	fmt.Printf("ID: %d | Título: %s | Autor: %s | Ano: %d | Disponível: %s\n",
		l.ID, l.Titulo, l.Autor, l.Ano, disponibilidade)
}

func (l Livro) Emprestar() Livro {
	if l.Disponivel {
		l.Disponivel = false
		fmt.Println("Livro emprestado com sucesso.")
	} else {
		fmt.Println("Livro já está emprestado.")
	}
	return l
}

func (l Livro) Devolver() Livro {
	if !l.Disponivel {
		l.Disponivel = true
		fmt.Println("Livro devolvido com sucesso.")
	} else {
		fmt.Println("Livro já está disponível.")
	}
	return l
}

func main() {
	biblioteca := []Livro{
		{ID: 1, Titulo: "Dom Casmurro", Autor: "Machado de Assis", Ano: 1899, Disponivel: true},
		{ID: 2, Titulo: "Memórias Póstumas de Brás Cubas", Autor: "Machado de Assis", Ano: 1881, Disponivel: true},
		{ID: 3, Titulo: "O Cortiço", Autor: "Aluísio Azevedo", Ano: 1890, Disponivel: true},
	}
	proximoID := 4

	for {
		fmt.Println("\n--- Menu ---")
		fmt.Println("1. Listar livros")
		fmt.Println("2. Adicionar livro")
		fmt.Println("3. Emprestar livro")
		fmt.Println("4. Devolver livro")
		fmt.Println("5. Sair")
		fmt.Print("Escolha uma opção: ")

		var opcao string
		fmt.Scan(&opcao)

		switch opcao {
		case "1":
			for _, livro := range biblioteca {
				livro.Exibir()
			}
		case "2":
			var titulo, autor string
			var ano int
			fmt.Print("Título: ")
			fmt.Scanln(&titulo)
			fmt.Print("Autor: ")
			fmt.Scanln(&autor)
			fmt.Print("Ano: ")
			fmt.Scan(&ano)

			livro := Livro{ID: proximoID, Titulo: titulo, Autor: autor, Ano: ano, Disponivel: true}
			biblioteca = append(biblioteca, livro)
			proximoID++
			fmt.Println("Livro adicionado com sucesso.")
		case "3":
			var id int
			fmt.Print("ID do livro a emprestar: ")
			fmt.Scan(&id)
			for i, livro := range biblioteca {
				if livro.ID == id {
					biblioteca[i] = livro.Emprestar()
					break
				}
			}
		case "4":
			var id int
			fmt.Print("ID do livro a devolver: ")
			fmt.Scan(&id)
			for i, livro := range biblioteca {
				if livro.ID == id {
					biblioteca[i] = livro.Devolver()
					break
				}
			}
		case "5":
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção inválida.")
		}
	}
}
