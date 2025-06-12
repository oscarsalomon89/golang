package main

import "fmt"

// ******************Exemplo 1
// func main() {
// 	x := 42
// 	fmt.Println("Valor de x:", x)
// 	fmt.Println("Endereço de x:", &x)
// }

// ******************Exemplo 2
// func main() {
// 	// x guarda o valor 10.
// 	x := 10

// 	// &x pega o endereço de memória de x (p guarda esse endereço.)
// 	p := &x

// 	fmt.Println("Valor de x:", x)
// 	fmt.Println("Endereço de x:", &x)
// 	fmt.Println("Ponteiro p:", p)
// 	// *p acessa o valor guardado no endereço que p está apontando.
// 	fmt.Println("Valor apontado por p:", *p)
// }

// ******************Exemplo 3
// func main() {
// 	x := 10
// 	sumar(x)
// 	fmt.Println("x após sumar 5:", x)
// }

// func sumar(valor *int) {
// 	*valor = *valor + 5
// }

// ******************Exemplo 4
// type Pessoa struct {
// 	nome  string
// 	idade int
// }

// func fazerAniversario(p *Pessoa) {
// 	p.idade++
// }

// func main() {
// 	p := Pessoa{nome: "Ana", idade: 30}
// 	fazerAniversario(&p)

// 	fmt.Println(p.idade)
// }

// type Frota struct {
// 	Veiculos []Veiculo
// }

// func (f *Frota) AdicionarVeiculo(veiculo Veiculo) {
// 	f.Veiculos = append(f.Veiculos, veiculo)
// }

// func fazerAniversario(p Pessoa) Pessoa {
// 	p.idade++

// 	return p
// }

// func main() {
// 	p := Pessoa{nome: "Ana", idade: 30}

// 	novaPessoa := fazerAniversario(p)

// 	fmt.Println(novaPessoa.idade)
// }

// func mudarNome(nome string) {
// 	nome = "Fernando"
// 	fmt.Println(nome)
// }

// 16 name = "Oscar"
// 17 chamamos func mudarNomePointer
// 35 nome = "xamaoijda5484d45a" ["xamaoijda5484d45a" : "Oscar"]
// 36 name = "Fernando"
// 37 imprime Fernando
// 18
// 19 Imprime valor do name (Fernando)

// func mudarNomePointer(nome *string) {
// 	*nome = "Fernando"
// 	fmt.Println(*nome)
// }

// 16 name = "oscar"
// 17 chamamos func mudarNome
// 22 nome = "oscar"
// 23 nome = "Fernando"
// 24 imprime Fernando
// 18
// 19 Imprime valor do name (Oscar)

type CarrinhoDeCompras struct {
	produtos []string
}

// Método com receiver sem ponteiro (não altera o carrinho original)
func (c CarrinhoDeCompras) AdicionarProdutoSemPonteiro(produto string) {
	c.produtos = append(c.produtos, produto)
	fmt.Println("Dentro de AdicionarProdutoSemPonteiro:", c.produtos)
}

// Método com receiver com ponteiro (altera o carrinho original)
func (c *CarrinhoDeCompras) AdicionarProdutoComPonteiro(produto string) {
	c.produtos = append(c.produtos, produto)
	fmt.Println("Dentro de AdicionarProdutoComPonteiro:", c.produtos)
}

func main() {
	carrinho := CarrinhoDeCompras{
		produtos: []string{"Arroz"},
	}

	// Usando sem ponteiro
	// carrinho.AdicionarProdutoSemPonteiro("Feijão")
	// fmt.Println("Depois de AdicionarProdutoSemPonteiro:", carrinho.produtos)

	// Usando com ponteiro
	carrinho.AdicionarProdutoComPonteiro("Farinha")
	fmt.Println("Depois de AdicionarProdutoComPonteiro:", carrinho.produtos)
}
