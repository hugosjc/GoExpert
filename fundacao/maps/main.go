package main

import "fmt"

func main() {
	salarios := map[string]int{"Hugo": 1000, "João": 1000, "Maria": 3000}
	//fmt.Println(salarios["Hugo"])
	delete(salarios, "Hugo")
	salarios["Hugo"] = 5000
	//fmt.Println(salarios["Hugo"])

	//sal := make(map[string]int)
	//sal1 := map[string]int{}
	//sal1["Hugo"] = 1000

	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}
}
