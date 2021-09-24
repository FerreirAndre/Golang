package main

import "fmt"

type endereco struct {
	rua    string
	numero uint16
}

type pessoa struct {
	nome  string
	idade uint8
	endereco
}

func main() {
	endereco1 := endereco{"Rio Feio", 4}
	pessoa := pessoa{"André Ferreira", 20, endereco1}
	fmt.Println(pessoa)
}
