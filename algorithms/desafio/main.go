package main

// import (
// 	"fmt"
// )

// type Livro struct {
// 	ID         int
// 	Titulo     string
// 	Autor      string
// 	Ano        int
// 	Disponivel bool
// }

// func (l Livro) Exibir() string {
// 	disponibilidade := "Sim"
// 	if !l.Disponivel {
// 		disponibilidade = "Nao"
// 	}

// 	msg := fmt.Sprintf("ID: %d, Titulo %s. Disponivel: %s", l.ID, l.Titulo, disponibilidade)
// 	//msg := "ID: " + strconv.Itoa(l.ID) + "Titulo: " + l.Titulo + "Disponivel: " + disponibilidade
// 	return msg
// }

// func main() {
// 	livro := Livro{
// 		ID:         1,
// 		Titulo:     "Ola Golang!",
// 		Autor:      "Oscar",
// 		Disponivel: false,
// 	}

// 	livro2 := Livro{
// 		ID:         2,
// 		Titulo:     "Memórias Póstumas de Brás Cubas",
// 		Autor:      "Machado de Assis",
// 		Ano:        1881,
// 		Disponivel: true,
// 	}

// 	biblioteca := []Livro{
// 		livro,
// 		livro2,
// 		{ID: 3, Titulo: "O Cortiço", Autor: "Aluísio Azevedo", Ano: 1890, Disponivel: true},
// 	}

// 	proximoID := 4

// 	for {
// 		fmt.Println("--- Menu ---")
// 		fmt.Println("1. Listar livros")
// 		fmt.Println("2. Adicionar livro")
// 		fmt.Println("5. Sair")
// 		fmt.Print("Escolha uma opção: ")

// 		var opcao string
// 		fmt.Scan(&opcao)

// 		switch opcao {
// 		case "1":
// 			listar(biblioteca)
// 		case "2":
// 			var titulo, autor string
// 			var ano int
// 			fmt.Print("Título: ")
// 			fmt.Scanln(&titulo)
// 			fmt.Print("Autor: ")
// 			fmt.Scanln(&autor)
// 			fmt.Print("Ano: ")
// 			fmt.Scan(&ano)

// 			livro := Livro{ID: proximoID, Titulo: titulo, Autor: autor, Ano: ano, Disponivel: true}
// 			biblioteca = append(biblioteca, livro)
// 			proximoID++
// 			fmt.Println("Livro adicionado com sucesso.")
// 		case "5":
// 			fmt.Println("Saindo...")
// 			return
// 		default:
// 			fmt.Println("Opção inválida.")
// 		}
// 	}
// }

// func listar(biblioteca []Livro) {
// 	// for i := 0; i < len(biblioteca); i++ {
// 	// 	fmt.Println(biblioteca[i].Exibir())
// 	// }

// 	for _, livro := range biblioteca {
// 		msg := livro.Exibir()
// 		fmt.Println(msg)
// 	}
// }
