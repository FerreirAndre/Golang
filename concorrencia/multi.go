package main

import (
	"fmt"
	"time"
)

func main() {
	canal := multiplexar(escrever("testando"), escrever("slk maneirokkk"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexar(canal1, canal2 <-chan string) <-chan string {
	canalSaida := make(chan string)
	go func() {
		for {
			select {
			case mensagem := <-canal1:
				canalSaida <- mensagem
			case mensagem := <-canal2:
				canalSaida <- mensagem
			}
		}
	}()
	return canalSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)
	go func() {
		for {
			canal <- texto

			time.Sleep(time.Millisecond * 500)
		}
	}()
	return canal
}
