package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {

	numero := 20
	sinalInvertido := inverterSinal(numero)
	fmt.Println(numero)
	fmt.Println(sinalInvertido)
	fmt.Println(numero)
	fmt.Println("------------")

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)

}
