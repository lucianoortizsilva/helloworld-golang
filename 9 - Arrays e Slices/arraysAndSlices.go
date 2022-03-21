package main

import "fmt"

func main() {

	fmt.Println("### Arrays")

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

	fmt.Println("### Slice")
	letras := []string{"A", "B", "C"}
	fmt.Println(letras)
	letras = append(letras, "L")
	fmt.Println(letras)

	outroSlice := frutas[1:2]
	fmt.Println(outroSlice)
}
