package main

import (
	"fmt"
	"time"
)

// canal de comunicação, pode enviar ou receberr dados para sincronizar.

func main() {
	canal := make(chan string)
	go escrever("Olá Mundo", canal)

	fmt.Println("Após execução")

	//for {
		// mensagem, aberto := <-canal // <- Canal = canal esperando receber o valor. Aberto é para verificar se o canal ainda está aberto.
		// if !aberto {
			//break // parar a execuçao do loop
		//}
		//fmt.Println(mensagem)
	//}

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do Programa")

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // Canal <- = Canal recebe, o Texto está sendo atribuido ao canal.
		time.Sleep(time.Second)
	}

	close(canal) // fechar o canal
}
