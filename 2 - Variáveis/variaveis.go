package main

import "fmt"

func main() {
	var variavel1 string = "variavel1"
	variavel2 := "variavel2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "aaaaa"
		variavel4 string = "bbbbb"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variavel 5", "variavel 6"
	fmt.Println(variavel5, variavel6)

	const constant1 string = "CONSTANTE1"
	fmt.Println(constant1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
