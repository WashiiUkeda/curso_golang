package main

import "fmt"

func recExe() {
	if r := recover(); r != nil {
		fmt.Println("Exe rec with success")
	}
}

func aprovado(n1, n2 float64) bool {
	defer recExe()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media > 6 {
		return false
	}
	panic("A média é exatamente 6.")
}

func main() {
	fmt.Println(aprovado(5, 9))
	fmt.Println("estou salvo")
}