package main

import "fmt"

func main() {
	//Maps precisa ter o mesmo tipo
	usuario := map[string]string{
		"nome": "joão",
		"sobrenome": "Pedro",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Pedro",
			"ultimo": "Alencar",
		},
		"curso": {
			"nome": "Engenharia",
			"periodo": "Noturno",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

}