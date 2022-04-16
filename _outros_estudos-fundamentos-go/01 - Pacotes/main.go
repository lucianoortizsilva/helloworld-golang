//para ser um arquivo executável
package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo no arquivo main")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("lucianoortizsilva@gmail.com")
	fmt.Println(erro)
}
