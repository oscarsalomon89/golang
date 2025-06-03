package main

import (
	"errors"
	"fmt"
	"os"
)

// Função que tenta abrir o arquivo
func abrirArquivo(nome string) error {
	_, err := os.Open(nome)
	if err != nil {
		// Wrappeando o erro original com mais contexto
		return fmt.Errorf("erro ao tentar abrir o arquivo '%s': %w", nome, err)
	}
	return nil
}

// Função que chama 'abrirArquivo' e wrappeia o erro
func processarArquivo(nome string) error {
	err := abrirArquivo(nome)
	if err != nil {
		// Wrappeando o erro de 'abrirArquivo' com mais contexto
		return fmt.Errorf("falha ao processar o arquivo '%s': %w", nome, err)
	}
	return nil
}

// Função que chama 'processarArquivo' e wrappeia o erro
func iniciarProcessamento(nome string) error {
	err := processarArquivo(nome)
	if err != nil {
		// Wrappeando o erro de 'processarArquivo' com mais contexto
		return fmt.Errorf("erro ao iniciar o processamento do arquivo '%s': %w", nome, err)
	}
	return nil
}

func main() {
	// Tentando processar um arquivo inexistente
	err := iniciarProcessamento("arquivo_inexistente.txt")
	if err != nil {
		// Verificando se o erro é 'os.ErrNotExist' no final da cadeia
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("Erro: O arquivo não existe (erro original).")
			fmt.Println(err)
		} else {
			// Se o erro não for 'ErrNotExist', mostramos o erro completo
			fmt.Println("Erro:", err)
		}
	}
}
