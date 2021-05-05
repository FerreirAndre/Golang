package main

import (
	"fmt"
	"math"
)

type circulo struct {
	raio float64
}
type retangulo struct {
	largura float64
	altura  float64
}
type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A area é: %.2f\n", f.area())
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}
func (c circulo) area() float64 {
	area := math.Pi * math.Pow(c.raio, 2)
	return area
}

func main() {
	r := retangulo{15, 20}
	c := circulo{5}
	escreverArea(r)
	escreverArea(c)
}
