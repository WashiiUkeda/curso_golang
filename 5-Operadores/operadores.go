package main

import "fmt"

func main() {
	//Operadores Aritiméticos
	// + soma - subtração * multiplicacao
	// / divisão % mod(resto da divisao)

	soma := 1 + 2
	sub := 3 - 4
	div := 10 / 5
	mult := 8 * 30
	restDiv := 27 % 7

	fmt.Println(soma, sub, div, mult, restDiv)

	var numero1 int16 = 10
	var numero2 int16 = 25

	soma2 := numero1 + numero2
	fmt.Println(soma2)

	// Atribuição
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	// Ralacionais (true or false)
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// Operadores Lógicos (E &&, OU || e Negação !variavel)
	fmt.Println("---------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// Operadores Unários
	numero := 10
	numero++ // numero = numero + 1
	numero += 31 // numero = numero + 1 + valor atribuido
	numero-- // numero = numero - 1
	numero -= 4 // numero = numero - 1 - valor atribuido
	numero *= 2 // numero = numero * valor atribuido
	numero /= 2 // numero = numero / valor atribuido
	numero %= 10 // numero = numero % valor atribuido
	fmt.Println(numero)

}