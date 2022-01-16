package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func httpErro(){
	res, err := http.Get("https://meuboteco.herokuapp.com/api/v1/category")

	if err != nil {
		log.Fatal(err.Error())
	}
		fmt.Print(res.Header)
}
func SomaErro(a,b int)(int, error){
	result:= a+b
	if result >10 {
		return 0, errors.New("Total Maior do que 10")
	}

	return result, nil
}

func main() {
	res, err := SomaErro(2,8)

	if err != nil {
		log.Fatal(err.Error())
	}
		fmt.Println(res)
}
