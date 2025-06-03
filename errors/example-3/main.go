package main

import (
	"errors"
	"fmt"
	"os"
)

// Função que tenta abrir um arquivo
func abrirArquivo(nome string) error {
	// Tentamos abrir o arquivo
	_, err := os.Open(nome)
	if err != nil {
		// Se ocorrer um erro, wrappeamos o erro com mais contexto
		return fmt.Errorf("não foi possível abrir o arquivo '%s': %w", nome, err)
	}
	return nil
}

func main() {
	// Tentamos abrir um arquivo que não existe
	err := abrirArquivo("arquivo_inexistente.txt")
	if err != nil {
		// Verificamos o erro e se for o tipo de erro específico, podemos tratá-lo
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("Erro específico: o arquivo não existe")
			fmt.Println(err)
		} else {
			// Caso contrário, mostramos o erro genérico
			fmt.Println("Erro:", err)
		}
	} else {
		fmt.Println("Arquivo aberto com sucesso!")
	}
}
