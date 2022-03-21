package main

import "fmt"

func main() {

	fmt.Println("### Arrays") //lista com tamanho fixo

	var frutas [3]string
	frutas[0] = "melância"
	frutas[1] = "pêssego"
	frutas[2] = "abacaxi"
	fmt.Println(frutas)

	carros := [3]string{"fox", "gol", "uno"}
	carros[2] = "bmw"
	fmt.Println(carros)

	numeros := [...]int{1, 2, 3, 6, 4, 9}
	fmt.Println(numeros)

	fmt.Println("### Slice") //lista sem tamanho fixo
	letras := []string{"A", "B", "C"}
	fmt.Println(letras)
	letras = append(letras, "L")
	fmt.Println(letras)

	outroSlice := frutas[1:2]
	fmt.Println(outroSlice)

	fmt.Println("### Arrays Internos")
	valores := make([]int, 3, 5)
	fmt.Println(len(valores)) //tamanho
	fmt.Println(cap(valores)) //capacidade
	fmt.Println(valores)
}
