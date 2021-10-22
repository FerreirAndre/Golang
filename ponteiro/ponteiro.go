package main

import "fmt"

func main() {
	var (
		numero   int
		ponteiro *int
	)
	fmt.Print("Digite um numero: ")
	fmt.Scan(&numero)
	ponteiro = &numero
	fmt.Println("O numero digitado foi", *ponteiro)
}
