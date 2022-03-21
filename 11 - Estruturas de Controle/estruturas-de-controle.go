package main

import "fmt"

func main() {

	numero := 100

	if numero >= 100 {
		fmt.Println("É MAIOR OU IGUAL A 100")
	} else {
		fmt.Println("É MENOR QUE 100")
	}

	//outroNumero TEM O ESCOPO APENAS DO "IF/ELSE"
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("OUTRO NUMERO É MAIOR QUE ZERO")
	} else if outroNumero < -8 {
		fmt.Println("< que - 8")
	} else {
		fmt.Println("acho que é isso!")
	}

}
