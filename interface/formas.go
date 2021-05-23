package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	raio float64
}
type Retangulo struct {
	largura float64
	altura  float64
}
type Forma interface {
	Area() float64
}

func EscreverArea(f Forma) {
	fmt.Printf("A area é: %.2f\n", f.Area())
}

func (r Retangulo) Area() float64 {
	return r.altura * r.largura
}
func (c Circulo) Area() float64 {
	area := math.Pi * math.Pow(c.raio, 2)
	return area
}

func main() {
	r := Retangulo{15, 20}
	c := Circulo{5}
	EscreverArea(r)
	EscreverArea(c)
}
