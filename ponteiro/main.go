package main

import "fmt"

type Carro struct {
	Name string
}

func (c *Carro) andou() {
	c.Name = "BMW"
}
func main() {

	car := Carro{
		"Palio",
	}
	car.andou()
	fmt.Println(car.Name)
}
