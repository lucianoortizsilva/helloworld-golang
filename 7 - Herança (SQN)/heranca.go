package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso      string
	mastricula string
}

func main() {

	p1 := pessoa{"", "", 10, 12}
	fmt.Println(p1)

	var estudante estudante
	estudante.nome = "Luciano"
	estudante.sobrenome = "Ortiz"
	estudante.idade = 38
	estudante.altura = 168
	estudante.curso = "Go"
	estudante.mastricula = "#52454556"

	fmt.Println(estudante)
}
