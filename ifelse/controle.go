package main

import "fmt"

func main() {
	var num int
	fmt.Print("Digite um numero: ")
	fmt.Scan(&num)
	if num > 10 {
		fmt.Println("numero maior que 10")
	} else {
		fmt.Println("Menor que 10")
	}
}
