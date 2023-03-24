package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {
	hugo := Cliente{
		Nome:  "Hugo",
		Idade: 38,
		Ativo: true,
	}

	hugo.Ativo = false
	hugo.Logradouro = "Praca Assis Chateaubriand"

	fmt.Println(hugo.Ativo)
	fmt.Println(hugo.Logradouro)
}
