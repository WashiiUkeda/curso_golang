//Multiplexador - Pegar 2 canais e juntar em um canal só.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	canal := multiplexar(escrever("Olá mundo!"), escrever("Programando em go."))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexar(entrada1, entrada2 <-chan string) <-chan string {
	saida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-entrada1:
				saida <- mensagem
			case mensagem := <-entrada2:
				saida <- mensagem
			}
		}
	}()
	return saida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()
	return canal
}
