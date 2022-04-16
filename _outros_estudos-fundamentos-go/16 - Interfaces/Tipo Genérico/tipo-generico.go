package main

import "fmt"

func generica(qualquerInterface interface{}) {
	fmt.Println(qualquerInterface)
}

func main() {

	generica("string")
	generica(123)
	generica(true)
}
