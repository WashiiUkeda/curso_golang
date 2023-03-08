package main

import "fmt"

//Funcoes variaticas pode receber N parametros.
//Não precisa informar quantos vai receber

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	somatotal := soma(1, 2, 5, 6, 7)
	somatotal2 := soma()
	fmt.Println(somatotal)
	fmt.Println(somatotal2)

	escrever("Olá Mundo", 20, 22)
}
