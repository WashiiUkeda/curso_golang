package main

import (
	"fmt"
)

// canal com buffer só vai bloquear o tráfego quando o canal atingir a capacidade máxima informada.

func main() {
	canal := make(chan string, 3)
	canal <- "Olá Mundo"
	canal <- "Programando em Go"
	canal <- "sei la"

	mensagem := <-canal
	mensagem2 := <-canal
	mensagem3 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
	fmt.Println(mensagem3)
}
