package main

import "fmt"

func main() {

	var variavelA int = 10
	var variavelB int = 20

	fmt.Println(variavelA, variavelB)

	var variavel3 int
	var variavel4 *int

	variavel3 = 100
	variavel4 = &variavel3
	fmt.Println(variavel3, variavel4)
	variavel3 = 999
	fmt.Println(variavel3, *variavel4) //desreferenciação
}
