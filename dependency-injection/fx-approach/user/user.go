package user

import (
	"fmt"
	"fx-approach/notifier"
)

type User struct {
	Name     string
	Notifier notifier.Notifier
}

func NewUser(name string, notifier notifier.Notifier) User {
	return User{
		Name:     name,
		Notifier: notifier,
	}
}

func (u User) SendNotification(msg string) {
	u.Notifier.Send(fmt.Sprintf("%s:%s", u.Name, msg))
}
