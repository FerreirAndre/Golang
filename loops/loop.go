package main

import "fmt"

func main() {
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}

	abc := "abc"
	ABC := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, letra := range abc {
		fmt.Println(letra)
	}
	for _, letra := range ABC {
		fmt.Println(letra)
	}
}
