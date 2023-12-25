package main

import (
	"fmt"
	"math"
)
type Geometria interface{
	area() float64
}
type Retangulo struct{
	altura, largura float64
}

type Circulo struct{
	radius float64
}

func (c Circulo) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (r Retangulo) area() float64 {
	return r.largura * r.altura 
}

func ExibeArea(g Geometria){
	fmt.Println(g.area())
}

func main(){
	fmt.Println("Tudo Certo Para começar")

	retangulo := Retangulo{
		largura: 1,
		altura: 2,
	}

	circulo := Circulo{
		radius: 3,
	}

	ExibeArea(retangulo)
	ExibeArea(circulo)

}