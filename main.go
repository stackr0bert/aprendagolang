package main

import "fmt"

func main() {
nomes := []string{"João", "Maria", "José", "Pedro", "Ana"}
for _, nome := range nomes {
	fmt.Println(nome)
}
}