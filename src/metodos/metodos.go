package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (user *usuario) aniversario() {
	user.idade++
}
func (idade usuario) maiorDeIdade()

func main() {
	usuario1 := usuario{"André", 20}
	fmt.Printf("%d\n", usuario1.idade)
	usuario1.aniversario()
	fmt.Printf("%d\n", usuario1.idade)
	usuario1.aniversario()
	fmt.Printf("%d\n", usuario1.idade)
}
