package main

import (
	"fmt"
	"my-api/repository"
	"my-api/service"
)

// Interface para um serviço de notificação
type NotificadorServico interface {
	Enviar(mensagem string)
}

// Implementação de um serviço de e-mail
type EmailServico struct{}

func (e EmailServico) Enviar(mensagem string) {
	fmt.Println("Enviando e-mail:", mensagem)
}

// Implementação de um serviço de SMS
type SMSServico struct{}

func (s SMSServico) Enviar(mensagem string) {
	fmt.Println("Enviando SMS:", mensagem)
}

// Usuário que depende de um serviço de notificação
type Usuario struct {
	Nome        string
	Notificador NotificadorServico
}

func main() {
	// emailServico := EmailServico{}
	// smsServico := SMSServico{}

	// usuario1 := Usuario{Nome: "João", Notificador: emailServico}

	// usuario2 := Usuario{Nome: "Maria", Notificador: smsServico}

	// usuario1.EnviarNotificacao("¡Hola, João!")
	// usuario2.EnviarNotificacao("¡Hola, Maria!")

	repo := repository.NewMongoRepository()
	userService := service.NewUserService(repo)

	fmt.Println(userService.FindAll())
}
