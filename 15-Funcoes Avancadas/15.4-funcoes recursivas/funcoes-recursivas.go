package main

import "fmt"

//funcoes recursivas chama ela mesma
//depende de uma outra execução dela mesma\

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	posicao := uint(12)
	fmt.Println(fibonacci(posicao))

	for i := uint(4); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}
}
