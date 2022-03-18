package main

import "fmt"

func main() {
	//ARITMÉTICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2
	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)
	/////////////////////////
	fmt.Println("------------------")
	//ATRIBUIÇÃO
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	fmt.Println("------------------")
	//RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	fmt.Println("------------------")
	//LÓGICOS
	fmt.Println(true && false)
	fmt.Println(false || true)
	fmt.Println(!false)
	fmt.Println("------------------")
	//UNÁRIOS
	numero := 10
	numero++
	numero += 20
	numero *= 3
	fmt.Println(numero)
	fmt.Println("------------------")

	fmt.Println("NAO TEM TERNARIO EM GO")
}
