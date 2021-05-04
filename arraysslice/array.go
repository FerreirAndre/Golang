package main

import "fmt"

func main() {
	var array1 [5]int
	var i int
	for i = 0; i < 5; i++ {
		array1[i] = i + 1
	}
	fmt.Println(array1)

	slice1 := []int{}
	for i = 0; i < 10; i++ {
		slice1 = append(slice1, i+1)
	}
	fmt.Println(slice1)
	fmt.Println(cap(slice1))
}
