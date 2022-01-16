package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func (p Pessoa) envelheceu() int {
	p.Idade++
	return p.Idade
}
func main() {
	pessoa := Pessoa{
		"Davi",
		29,
	}
	fmt.Println(pessoa)
	fmt.Println("Ano novo")
	pessoa.Idade = pessoa.envelheceu()
	fmt.Println(pessoa)

}
