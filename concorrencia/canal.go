package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("porra eh essa", canal)
	for {
		mensagem, aberto := <-canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}
	novoCan := make(chan string)
	go escrever("testando e os cacete", novoCan)
	for porra := range novoCan {
		fmt.Println(porra)
	}
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 4; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}
