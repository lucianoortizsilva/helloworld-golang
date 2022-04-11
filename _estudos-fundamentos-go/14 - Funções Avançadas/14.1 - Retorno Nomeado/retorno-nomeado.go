package main

import "fmt"

func calculosMatematicos(valorA int, valorB int) (soma, subtracao int) {
	soma = valorA + valorB
	subtracao = valorA - valorB
	return
}

func main() {
	soma, subtracao := calculosMatematicos(10, 25)
	fmt.Println(soma, subtracao)
}
