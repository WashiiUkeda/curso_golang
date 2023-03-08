package main

import "fmt"

// Closure - serve para "manter" a variavél em memória e ser utilizado

func closure() func() {
	texto := "Dentro da Closure"
	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

func main() {
	texto := "Dentro da Main"
	fmt.Println(texto)

	funcNova := closure()
	funcNova()
}
