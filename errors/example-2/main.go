package main

import (
	"fmt"
)

// Função que tenta criar um usuário com um nome específico
func criarUsuario(nome string) error {
	if nome == "oscar" {
		// Criando um erro com uma mensagem dinâmica que inclui o nome do usuário
		return fmt.Errorf("erro: o nome de usuário '%s' não esta permitido", nome)
	}
	// Se o nome for válido, retornamos nil
	return nil
}

func main() {
	// Testando com nome vazio
	err := criarUsuario("oscar")
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Usuário criado com sucesso!")
	}

	// Testando com nome válido
	err = criarUsuario("joao")
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Usuário criado com sucesso!")
	}
}
