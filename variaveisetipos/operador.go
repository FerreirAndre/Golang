package main

import "fmt"

func maiorOuMenor(valorA int, valorB int) int {
	if valorA > valorB {
		return valorA
	} else if valorB > valorB {
		return valorB
	} else {
		return valorA
	}
}

func main() {
	n1, n2 := 40, 30

	fmt.Println(maiorOuMenor(n1, n2))

}
