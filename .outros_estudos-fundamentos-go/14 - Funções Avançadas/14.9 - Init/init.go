package main

import "fmt"

var salario float64

func init() {
	fmt.Println("INICIALIZANDO")
	salario = 100000000.00
}

func main() {
	fmt.Printf("O salário é %f", salario)
}
