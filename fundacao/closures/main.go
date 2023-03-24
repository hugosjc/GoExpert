package main

import "fmt"

func main() {
	total := func() int {
		return sum(1, 3, 45, 5, 7, 89, 9, 9, 997, 56, 45, 45, 457, 54) * 2
	}()
	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
