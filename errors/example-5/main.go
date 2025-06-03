package main

import (
	"errors"
	"fmt"
)

// Definindo um tipo de erro personalizado
type CustomError struct {
	Code    int
	Message string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("Erro %d: %s", e.Code, e.Message)
}

// Função que retorna um erro personalizado
func gerarErro() error {
	return CustomError{Code: 404, Message: "Recurso não encontrado"}
}

// Função que envolve o erro personalizado com mais contexto
func processarRequisicao() error {
	err := gerarErro()
	if err != nil {
		// Envolvendo o erro com mais contexto
		return fmt.Errorf("erro ao processar a requisição: %w", err)
	}
	return nil
}

func main() {
	// Tentando processar uma requisição que gera um erro
	// if err := processarRequisicao(); err != nil {
	// 	// Verificando se o erro é do tipo CustomError usando errors.As()
	// 	var customErr CustomError
	// 	if errors.As(err, &customErr) {
	// 		fmt.Printf("Erro personalizado encontrado! Código: %d, Mensagem: %s\n", customErr.Code, customErr.Message)
	// 	} else {
	// 		// Se o erro não for do tipo CustomError, mostramos o erro completo
	// 		fmt.Println("Erro:", err)
	// 	}
	// }

	err := CustomError{
		Code:    401,
		Message: "error",
	}

	// target := new(CustomError)
	target := CustomError{}

	if errors.As(err, &target) {
		fmt.Println("equal")
	}

	fmt.Println("fin")
}
