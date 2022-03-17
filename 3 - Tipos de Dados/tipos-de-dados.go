package main

import (
	"errors"
	"fmt"
)

func main() {
	// int = int8, int16, int32, int64
	var numero int64 = -10000000
	fmt.Println(numero)

	var numero2 uint32 = 1000
	fmt.Println(numero2)

	//alias - Geralmente usados para caracteres da tabela ASCII
	//INT32 = RUNE
	var numero3 rune = 12345
	fmt.Println(numero3)

	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal float32 = 1235.45
	fmt.Println(numeroReal)

	var numeroReal2 float64 = 12356456456
	fmt.Println(numeroReal2)

	var nome string = "Luciano"
	fmt.Println(nome)

	//char tabela ascii
	char := 'A'
	fmt.Println(char)

	var texto string
	fmt.Println(texto)

	var booleano bool
	fmt.Println(booleano)

	var erro error = errors.New("erro interno")
	fmt.Println(erro)
}
