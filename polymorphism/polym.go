package main

import (
	"exemplo/forma"
	"fmt"
)

type Quadrado struct {
	Lado float64
}

func (q Quadrado) Area() float64 {
	return q.Lado * q.Lado
}

func (c Quadrado) Perimeter() float64 {
	return 4 * c.Lado
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return 3.1416 * c.Raio * c.Raio
}

func (c Circulo) Perimeter() float64 {
	return 2 * 3.1416 * c.Raio
}

// Função que usa polimorfismo: aceita qualquer "Forma"
func imprimirArea(f forma.Forma) {
	fmt.Println("Área:", f.Area())
}

func main() {
	quadrado := Quadrado{Lado: 4}
	circulo := Circulo{Raio: 5}

	// Ambos os tipos implementam a interface "Forma"
	imprimirArea(quadrado)
	imprimirArea(circulo)
}
