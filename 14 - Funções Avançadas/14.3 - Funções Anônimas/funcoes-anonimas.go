package main

import "fmt"

func main() {

	fmt.Println("Funções Anônimas - Funções sem nome")

	func() {
		fmt.Println("Olá mundo")
	}()

	func(nome string) {
		fmt.Printf("Meu nome é: %s", nome)
	}("Luciano Ortiz")

	retorno := func(linguagem string) string {
		return fmt.Sprintf("\nLinguagem de programação: %s", linguagem)
	}("Golang")
	fmt.Println(retorno)

}
