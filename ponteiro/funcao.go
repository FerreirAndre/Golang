package main

import "fmt"

func inverterSinal(numero int) int {
	numero *= -1
	return numero
}
func inverterComPonteiro(numero *int) {
	*numero *= -1
}

func main() {
	numero := 20
	inverterSinal(numero)
	fmt.Println(numero)
	inverterComPonteiro(&numero)
	fmt.Println(*&numero)
}
