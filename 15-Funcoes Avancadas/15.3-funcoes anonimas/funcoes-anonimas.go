package main

import "fmt"

func main () {

	func () {
		fmt.Println("Olá Mundo")
	}()

	func (texto string) {
		fmt.Println(texto)
	}("Aleatório")

	retorno := func (texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Um Parametro")
	fmt.Println(retorno)
}