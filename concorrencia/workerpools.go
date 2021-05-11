package main

import "fmt"

func main() {
	var posicao int
	fmt.Print("Digite a posicao: ")
	fmt.Scan(&posicao)
	tarefas := make(chan int, posicao)
	resultados := make(chan int, posicao)

	go worker(tarefas, resultados)
	//para aumentar a potencia, decomente as linhas abaixo.

	// go worker(tarefas, resultados)
	// go worker(tarefas, resultados)
	// go worker(tarefas, resultados)
	// go worker(tarefas, resultados)
	// go worker(tarefas, resultados)

	for i := 0; i < posicao; i++ {
		tarefas <- i
	}
	close(tarefas)
	for i := 0; i < posicao; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return 1
	}
	return fibonacci(posicao-1) + fibonacci(posicao-2)
}
