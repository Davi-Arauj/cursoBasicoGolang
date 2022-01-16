package main

import "fmt"


func somaArray(arr []int) int {
	total := 0
	for i := 0; i < len(arr); i++ {
		total += arr[i]
	}
	total2:= 0
	for _, v := range arr {
		total2 += v
	}
	return total2
}

func main() {

	array := []int{
		1, 2, 3, 4,5,6,
	}

	result := somaArray(array)

	fmt.Println(result)

}
