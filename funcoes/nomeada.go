package main

import "fmt"

func calculadora(num1 float64, num2 float64) (soma float64, subtracao float64, multiplicacao float64, divisao float64) {
	soma = num1 + num2
	subtracao = num1 - num2
	multiplicacao = num1 * num2
	divisao = num1 / num2
	return
}

func main() {
	fmt.Println("Calculadora.")

	for {
		var (
			num1     float64
			num2     float64
			operacao string
		)
		fmt.Println("Digite o primeiro numero, a operacao e o segundo numero")
		fmt.Scan(&num1, &operacao, &num2)
		if string(operacao) == "+" {
			resultado, _, _, _ := calculadora(num1, num2)
			fmt.Println(resultado)
		}
		if string(operacao) == "-" {
			_, resultado, _, _ := calculadora(num1, num2)
			fmt.Println(resultado)
		}
		if string(operacao) == "*" {
			_, _, resultado, _ := calculadora(num1, num2)
			fmt.Println(resultado)
		}
		if string(operacao) == "/" {
			_, _, _, resultado := calculadora(num1, num2)
			fmt.Println(resultado)
		}
	}

}
