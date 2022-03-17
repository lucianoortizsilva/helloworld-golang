package auxiliar

import "fmt"

// Visível em todos os pacotes, pois função inicia com letra maiuscula
func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")
	escrever2()
}
