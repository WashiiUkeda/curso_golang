package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")
	var variavel1 int = 100
	var variavel2 int = variavel1
	fmt.Println(variavel1, variavel2)

	variavel2++
	fmt.Println(variavel1, variavel2)

	// Ponteiro é uma referencia de memória
	fmt.Println("---------")
	var variavel3 string = "sei la"
	var ponteiro *string
	

	ponteiro = &variavel3
	fmt.Println(variavel3, ponteiro)

	variavel3 = "vasco"
	fmt.Println(variavel3, ponteiro, *ponteiro)
}