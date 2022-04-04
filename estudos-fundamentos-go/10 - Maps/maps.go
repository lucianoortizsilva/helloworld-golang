package main

import "fmt"

func main() {

	var usuario = map[string]string{
		"nome":      "Luciano",
		"sobrenome": "Ortiz",
	}

	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"nome": {
			"first": "Luciano",
			"last":  "Silva",
		},
		"curso": {
			"name":   "ADS",
			"campus": "Zona Sul",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "curso") //funcao nativa do go
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "aquário",
	}
	fmt.Println(usuario2)
}
