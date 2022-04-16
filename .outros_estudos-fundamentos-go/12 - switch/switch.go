package main

import "fmt"

func diaSemana(dia int8) string {
	switch dia {
	case 0:
		return "domingo"
	case 1:
		return "segunda-feira"
	case 2:
		return "terça-feira"
	case 3:
		return "quarta-feira"
	case 4:
		return "quinta-feira"
	case 5:
		return "sexta-feira"
	case 6:
		return "sábado"
	default:
		return "inválido"
	}
}

func diaSemanaOutro(dia int8) string {
	var diaSemana string
	switch {
	case dia == 0:
		diaSemana = "domingo"
	case dia == 1:
		diaSemana = "segunda-feira"
		break
	case dia == 2:
		diaSemana = "terça-feira"
	case dia == 3:
		diaSemana = "quarta-feira"
	case dia == 4:
		diaSemana = "quinta-feira"
		fallthrough //entra direto no próximo case, sem avaliar
	case dia == 5:
		diaSemana = "sexta-feira"
	case dia == 6:
		diaSemana = "sábado"
	default:
		diaSemana = "inválido"
	}
	return diaSemana
}

func main() {

	dia := diaSemana(1)
	diaOutro := diaSemanaOutro(4)
	fmt.Println(dia)
	fmt.Println(diaOutro)

}
