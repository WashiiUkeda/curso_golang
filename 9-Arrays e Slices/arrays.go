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
}
