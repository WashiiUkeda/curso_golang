package main

import "fmt"

var n int

func init() {
	fmt.Println("Executando func init")
	n = 10
}

func main () {
	fmt.Println("Executando func Main")
	fmt.Println(n)
}