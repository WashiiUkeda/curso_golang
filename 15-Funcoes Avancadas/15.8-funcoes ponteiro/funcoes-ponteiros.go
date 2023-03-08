package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func ponteiroInvertSinal(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvert := inverterSinal(numero)
	fmt.Println(numeroInvert)
	fmt.Println(numero)

	novoNumero := 40
	fmt.Println(novoNumero)
	ponteiroInvertSinal(&novoNumero)
	fmt.Println(novoNumero)
}