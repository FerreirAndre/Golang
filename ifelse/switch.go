package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Segunda-feira"
	case 2:
		return "Terça-feira"
	case 3:
		return "Quarta-feira"
	case 4:
		return "Quinta-feira"
	case 5:
		return "Sexta-feira"
	case 6:
		return "Sabado"
	case 7:
		return "Domingo"
	default:
		return "Numero invalido"
	}
}
func main() {
	var num int
	fmt.Print("Digite um numero para saber o dia da semana: ")
	fmt.Scan(&num)
	fmt.Println(diaDaSemana(num))
}
