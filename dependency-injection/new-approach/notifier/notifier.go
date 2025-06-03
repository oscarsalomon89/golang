package notifier

import "fmt"

type Notifier interface {
	Send(message string)
}

// ///////////////////////////////////////////
type Email struct{}

func NewEmail() Email {
	return Email{}
}

func (e Email) Send(message string) {
	fmt.Println("Enviando e-mail:", message)
}

// ///////////////////////////////////////////
type SMS struct{}

func NewSMS() SMS {
	return SMS{}
}

func (s SMS) Send(mensagem string) {
	fmt.Println("Enviando SMS:", mensagem)
}
