package main

import "fmt"

var( 
	nome string
	n1 int
	n2 int32
)
func main(){
	mensagem := "Aula de Go"
	fmt.Println(mensagem)

	var b, f, s = true, 2.3, "Olá"
	fmt.Println(b, f, s)

	var total int
	total++
	fmt.Println("total:", total)

	nome = "r0bert"
	fmt.Println("Hello", nome, "!")

	var x,y = 10, 20
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)

}