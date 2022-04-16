package main

import "fmt"

func somar(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func imprimirIdades(texto string, idades ...int8) {
	for _, idade := range idades {
		fmt.Println(texto, idade)
	}
}

func main() {
	total := somar(162, 2, 14, 3, 6)
	fmt.Println(total)

	imprimirIdades("A idade é: ", 12, 16, 32, 65, 99, 12, 10)
}
