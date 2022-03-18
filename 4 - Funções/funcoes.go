package main

import "fmt"

func somar(numero1 int8, numero2 int8) int8 {
	return numero1 + numero2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {

	resultado := somar(20, 25)
	fmt.Println(resultado)

	var f = func(txt string) {
		fmt.Println("Função f: " + txt)
	}

	f("ok ok ok")

	resultadoSoma, resultadoSubtracao := calculosMatematicos(20, 6)
	resultadoSoma2, _ := calculosMatematicos(21, 9) //com _ para nao usar o retorno "2"
	fmt.Println(resultadoSoma, resultadoSubtracao)
	fmt.Println(resultadoSoma2)
}
