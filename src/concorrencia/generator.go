package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("caralho")
	for i := 0; i < 5; i++ {
		fmt.Println(<-canal)
	}
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
