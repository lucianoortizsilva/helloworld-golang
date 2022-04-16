package main

import "fmt"

func minhaFuncaoClosure() func() {

	texto := "TEXTO DENTRO DA minhaFuncaoClosure()"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {

	texto := "TEXTO DENTRO DO MAIN"
	fmt.Println(texto)

	funcaoNova := minhaFuncaoClosure()

	funcaoNova()
}
