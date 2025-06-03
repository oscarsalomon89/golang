package user

import (
	"dependency-2/notifier"
	"fmt"
)

type user struct {
	Name     string
	Notifier notifier.Notifier
}

func NewUser(name string, notifier notifier.Notifier) user {
	return user{
		Name:     name,
		Notifier: notifier,
	}
}

func (u user) SendNotification(msg string) {
	u.Notifier.Send(fmt.Sprintf("%s:%s", u.Name, msg))
}
