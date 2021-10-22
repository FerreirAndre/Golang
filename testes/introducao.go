package main

import (
	"bufio"
	"fmt"
	"introducao-teste/endereco"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digita o endereco: ")
	text, _ := reader.ReadString('\n')
	tipoEndereco := endereco.TipoDeEndereco(text)
	fmt.Println(tipoEndereco)
}
