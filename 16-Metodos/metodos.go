package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.nome)
}

func (u usuario) underAge() bool {
	return u.idade >= 18
}

func (u *usuario) bday() {
	u.idade++
}

func main() {
	usuario1 := usuario{"usuario 1", 20}
	usuario1.salvar()

	usuario2 := usuario{"Davi", 30}
	usuario2.salvar()

	underAge := usuario2.underAge()
	fmt.Println(underAge)

	usuario2.bday()
	fmt.Println(usuario2.idade)
}
