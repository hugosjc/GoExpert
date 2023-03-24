package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	hugo := Cliente{
		Nome:  "Hugo",
		Idade: 38,
		Ativo: true,
	}

	hugo.Ativo = false

	fmt.Println(hugo.Ativo)
}
