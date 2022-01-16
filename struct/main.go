package main

import "fmt"

type Cliente struct {
	Name  string
	Email string
	CPF   int
}

type ClienteInternacional struct {
	Cliente
	Pais string
}

func (c Cliente) exibe() {
	fmt.Println(c.Name)
}

func main() {
	cliente1 := Cliente{
		"José",
		"j@j.com",
		1565498632,
	}

	fmt.Println(cliente1)

	clienteGringo := ClienteInternacional{
		Cliente{
			"Davi",
			"d@d.com",
			14665851,
		},
		"ARGENTINA",
	}

	fmt.Println(clienteGringo)
	cliente1.exibe()
	clienteGringo.exibe()
}
