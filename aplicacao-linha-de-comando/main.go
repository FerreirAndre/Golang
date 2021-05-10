package main

import (
	"fmt"
	"linha-de-comando/application"
	"log"
	"os"
)

func Print(a ...interface{}) {
	fmt.Println(a...)
}

func main() {
	aplicacao := application.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
