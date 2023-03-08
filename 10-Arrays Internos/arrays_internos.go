package main

import "fmt"

func main() {
	// Array = Lista de valores "irredutivel"

	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Pos1", "Pos2", "Pos3", "Pos4", "Pos5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	//slices são mais dinamicos
	
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(slice)

	slice = append(slice, 8)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "PosAlt"
	fmt.Println(slice2)

	// Arrays Internos
	fmt.Println("------")
	slice3 := make([]float32, 10, 11) //tipo, tamanho, capacidade
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)
	//Não informar parametro de capacidade, ela assume o tamanho.
	fmt.Println(slice4)
	fmt.Println(cap(slice4))
	slice4 = append(slice4, 10)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
