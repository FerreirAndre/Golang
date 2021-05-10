package main

import "fmt"

func main() {
	canal := make(chan string, 1)
	canal <- "olá viado"
	msg := <-canal
	fmt.Println(msg)
}
