package main

import (
	"fmt"
)

type Forma interface {
	Area() float64
	Perimeter() float64
}

type Comportamento interface {
	Falar()
	Rodar()
}

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

func criarForma(tipo string, medida int) Forma {
	switch tipo {
	case "quadrado":
		return Quadrado{Lado: float64(medida)}
	case "circulo":
		return Circulo{Raio: float64(medida)}
	default:
		return nil
	}
}

type Biblioteca struct {
	Formas []Forma
}

func main() {
	quadrado := criarForma("quadrado", 4)
	if quadrado == nil {
		fmt.Println("error")
	}
	circulo := Circulo{Raio: 5}

	var biblioteca []Forma

	biblioteca = append(biblioteca, quadrado, circulo)

	fmt.Println(biblioteca)

	// Ambos os tipos implementam a interface "Forma"
	imprimirInfo(quadrado)
	imprimirInfo(circulo)
}
