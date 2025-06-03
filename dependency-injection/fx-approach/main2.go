package main

// import (
// 	"fx-approach/notifier"
// 	"fx-approach/user"
// 	"os"

// 	"go.uber.org/fx"
// )

// // Módulo para a dependência do nome do usuário
// var stringDependency = fx.Module("username",
// 	fx.Provide(func() string {
// 		name = "DefaultUser"
// 	}),
// )

// // Módulo para a dependência do Notificador
// var notifierDependency = fx.Module("notifier",
// 	fx.Provide(func() notifier.Notifier {
// 		return notifier.NewEmail()
// 	}),
// )

// // Módulo para a criação do usuário
// var UserModule = fx.Module("user",
// 	stringDependency,
// 	notifierDependency,
// 	fx.Provide(user.NewUser),
// )

// // Módulo que contém a inicialização da aplicação
// var AppModule = fx.Module("app",
// 	UserModule,
// 	fx.Invoke(func(u *user.User) {
// 		u.SendNotification("¡Hola desde fx!")
// 	}),
// )

// func main() {
// 	fx.New(AppModule).Run()
// }
