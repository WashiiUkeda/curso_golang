package main

import "fmt"

func funcao1() {
	fmt.Println("Executando função 1")
}

func funcao2() {
	fmt.Println("Executando função 2")
}

func aprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Segue resultado.")
	fmt.Println("Verificando notas do aluno.")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

func main() {

	funcao1()
	funcao2()
	fmt.Println(aprovado(1, 9))
}
