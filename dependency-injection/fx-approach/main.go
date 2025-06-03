package main

import (
	"fx-approach/notifier"
	"fx-approach/user"

	"go.uber.org/fx"
)

var userModule = fx.Options(
	fx.Provide(func() notifier.Notifier {
		return notifier.NewEmail()
	}),
	fx.Provide(func() string {
		return "DefaultUser"
	}),
	fx.Provide(user.NewUser),
)

var startApp = fx.Options(
	fx.Invoke(func(u user.User) {
		u.SendNotification("¡Hola desde fx!")
	}),
)

// Com FX, você faz isso de forma declarativa, dizendo:
// * Quem fornece cada coisa (Provide)
// * O que fazer com tudo isso (Invoke)
// * E o FX se encarrega de conectar tudo e executar (Run)
func main() {
	// fx.New(...) é como o contêiner principal da aplicação.
	// É onde você define todas as dependências que sua aplicação precisa:
	// o que vai ser injetado (fx.Provide), o que deve ser executado (fx.Invoke), etc.
	fx.New(
		userModule,
		startApp,
	).Run()
}

// notifierModule := fx.Options(
// 	fx.Provide(func() notifier.Notifier {
// 		return notifier.NewEmail()
// 	}),
// )

// userModule := fx.Options(
// 	fx.Provide(func() string {
// 		return "DefaultUser"
// 	}),
// 	fx.Provide(user.NewUser),
// )

// startApp := fx.Options(
// 	fx.Invoke(func(u user.User) {
// 		u.SendNotification("¡Hola desde fx!")
// 	}),
// )

// fx.New(
// 	notifierModule,
// 	userModule,
// 	startApp,
// ).Run()
