package main

import (
	"fmt"
)

func main() {

	fmt.Println(sum(1, 2, 3, 5, 4, 5, 4, 84, 78, 48, 5, 4))
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
