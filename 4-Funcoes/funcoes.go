package main

import "fmt"

//funções podem ter parametro e retorno.
//parametro - dados que colocamos
//retorno - o que a funcao retorna

func somar(n1, n2 int) int {
	return n1 + n2
}

func calculos(c1, c2 int32) (int32, int32) {
soma2 := c1 + c2
subtracao := c1 - c2
return soma2, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)
	
	var f = func() {
		fmt.Println("Função F")
	}
	f()

	var f2 = func(txt string) {
		fmt.Println(txt)
	}
	f2("Função F2")

	var f3 = func(txt2 string) string {
		return txt2
	}
	resultado := f3("Função F3")
	fmt.Println(resultado)

	//Quando tiver mais de um return, na hora de puxar no programa somente uma, usamos underline (_)
	
	resultadoSoma2, resultadoSubtraçao := calculos(30, 25)
	fmt.Println(resultadoSoma2, resultadoSubtraçao)
}
