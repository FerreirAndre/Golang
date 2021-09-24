package main

import "fmt"

func main() {
	frase := "analises e os carai"
	func(frase string) {
		fmt.Println(frase)
	}(frase)
}
