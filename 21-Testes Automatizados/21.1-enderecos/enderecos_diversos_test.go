// Teste de unidade
package	enderecos

type cenarioDeTeste struct {
	enderecoInserido string
	enderecoEsperado string
}

// go test ./... - roda todos os testes dentro da pasta
// go test -v - roda os testes verbose indicando o que o código rodou
// rodar testes em paralelo - t.Parallel()
// go test --cover
// go test --coverprofile cobertura.txt
// go test --func=cobertura.txt
// go test cover --html

func TestTipoDeEndereco(t. *testing.t) {

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia Sei la", "Rodovia"},
		{"Estrada velha", "Estrada"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				retornoRecebido,
				cenario.retornoEsperado,
			)			
		}
	}
}

	// Test seguido do nome da funcao
// 	enderecoParaTeste := "Avenida Paulista"

// 	tipoDeEnderecoEsperado := "Avenida"
// 	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

// 	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
// 		t.Error("O tipo recebido é diferente do esperado. Esperava %s e recebeu %s",
// 			tipoDeEnderecoEsperado,
// 			tipoDeEnderecoRecebido,
// 		)
// 	}

// }