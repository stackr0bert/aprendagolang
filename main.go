package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     uint8
	Status    bool
}

type Categoria struct {
	Nome string
	Pai  string
}

func main() {
	p := Pessoa{
		Nome:      "Robert",
		Sobrenome: "Araújo",
		Idade:     27,
		Status:    true,
	}

	fmt.Println(p, p.Nome)
}
