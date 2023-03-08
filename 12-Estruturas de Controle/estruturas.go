package main

import "fmt"

func main() {

	numero := -2

	if numero > 15 {
		fmt.Println("Maior que 15.")
	} else {
		fmt.Println("Menor ou igual a 15.")
	}

	if numero2 := numero; numero2 > 0 {
		fmt.Println("Numero é maior que zero.")
	} else if numero < -10 {
		fmt.Println("Número é menor que -10")
	} else {
		fmt.Println("Número entre 0 e -10.")
	}
}
