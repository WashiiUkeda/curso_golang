package main

import (
	"errors"
	"fmt"
)

func main() {
	//int8, int16, int32 e int64 retorna um número inteiro baseado na quantidade de bits
	//int retorna um número inteiro baseado na arquitetura do processador da máquina (32 - 64 bits)
	var numero int = -100000000
	fmt.Println(numero)

	//uint - unsigned integers. Mesmas caracteristicas do int, porém sem sinais.
	numero2 := 20341
	fmt.Println(numero2)

	//Alias - resumindo, é um "apelido" utilizado para chamar declarar algum tipo.
	//rune = int32
	//byte = uint8
	var (
		numero3 rune = 4322
		numero4 byte = 33
	)
	fmt.Println(numero3, numero4)

//Números reais - float32 e float64. Não pode declarar somente Float.
	var (
		numeroReal1 float32 = 231.345
		numeroReal2 float64 = 4325.436576
	 )
	 fmt.Println(numeroReal1, numeroReal2)

	 //string - cadeia de caracteres
	 var str string = "texto"
	 str2 := "texto2"
	 fmt.Println(str, str2)

	 char := '$'
	 fmt.Println(char)

	 //Os dados em GO tem um valor inicial quando declarados com VAR, que é chamdo de valor zero.
	 //Em strings retorna em branco, em int/float retorna Zero.

	 var texto int32
	 fmt.Println(texto)

	 //boleanos - True, False
	 var boleano bool = true
	 fmt.Println(boleano)

	 //erro é um tipo em Go. Error. 
	 var erro error = errors.New("Erro interno")
	 fmt.Println(erro)

}