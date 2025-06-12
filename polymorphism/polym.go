package main

import (
	"fmt"
)

type Forma interface {
	Area() int64
	Perimeter() float64
}

type Comportamento interface {
	Falar()
	Rodar()
}

type Quadrado struct {
	Lado float64
}

func (q Quadrado) Area() int64 {
	return int64(q.Lado * q.Lado)
}

func (c Quadrado) Perimeter() float64 {
	return 4 * c.Lado
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() int64 {
	return int64(3.1416 * c.Raio * c.Raio)
}

func (c Circulo) Perimeter() float64 {
	return 2 * 3.1416 * c.Raio
}

func (c Circulo) Falar() {
	fmt.Println("ola sou um circulo")
}

func (c Circulo) Rodar() {
	fmt.Println("rodando")
}

// Função que usa polimorfismo: aceita qualquer "Forma"
func imprimirInfo(f Forma) {
	fmt.Println("Área:", f.Area())
	fmt.Println("Perimeter:", f.Perimeter())
}

func imprimirInfo2(f Comportamento) {
	f.Falar()
	f.Rodar()
}

func main() {
	quadrado := Quadrado{Lado: 4}
	circulo := Circulo{Raio: 5}

	imprimirInfo(quadrado)
	fmt.Println("----------------")
	imprimirInfo(circulo)
	imprimirInfo2(circulo)
}
