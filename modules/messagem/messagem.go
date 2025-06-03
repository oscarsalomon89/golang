package messagem

import (
	"fmt"
)

func GerarSaudacao(nome string) string {
	//message := "Bem-vindo, "+ nome +"! Esperamos que você tenha um ótimo dia!"
	return fmt.Sprintf("Bem-vindo, %s! Esperamos que você tenha um ótimo dia!", nome)
}
