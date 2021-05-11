package main

import (
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "canal unzinho"
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "canal dois"
		}
	}()

	for {
		select {
		case mensagem1 := <-canal1:
			println(mensagem1)
		case mensagem2 := <-canal2:
			println(mensagem2)
		}
	}
}
