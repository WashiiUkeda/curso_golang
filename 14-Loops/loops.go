package main

import (
	"fmt"
	"time"
)

func main() {

	// só tem o for

	//i := 10
	//for i < 20 {
	//	i++
	//	fmt.Println("Incrementando i")
	//	time.Sleep(time.Second)
	//}
	//fmt.Println(i)

	//for j := 0; j < 10; j++ {
	//	fmt.Println("Incrementando j", j)
	//	time.Sleep(time.Second)
	//}

	nomes := [3]string{"João", "David", "Lucas"}
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	
	}

	usuario := map[string]string{
		"nome": "Leonardo",
		"Sobrenome": "Silva",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for {
		fmt.Println("Loop Infinito")
		time.Sleep(time.Second)
	}
}
