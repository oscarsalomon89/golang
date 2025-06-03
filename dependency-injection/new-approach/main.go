package main

import (
	"dependency-2/notifier"
	"dependency-2/user"
)

func main() {
	emailService := notifier.NewEmail()
	usuario1 := user.NewUser("João", emailService)
	usuario1.SendNotification("¡Ola, João!")
}
