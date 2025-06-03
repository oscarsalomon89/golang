package main

import (
	"fmt"
	"modules-example/pessoa"
)

func main() {
	// usuario := user.User{
	// 	Username: "oscar",
	// 	Password: "12345",
	// }

	// fmt.Println(usuario)

	p := pessoa.New("Oscar", 35)

	fmt.Println(p.MensagemDeBoasVindas)
}
