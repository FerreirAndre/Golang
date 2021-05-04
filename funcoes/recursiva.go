package main

import "fmt"

func fibro(posicao uint) uint {
	if posicao <= 1 {
		return 1
	}
	return fibro(posicao-1) + fibro(posicao-2)
}

func main() {
	var posicao uint
	fmt.Print("Digite a posicao para saber o valor na lista fibonacci: ")
	fmt.Scan(&posicao)
	for i := uint(1); i <= posicao; i++ {
		fmt.Println(fibro(i))
	}
}
