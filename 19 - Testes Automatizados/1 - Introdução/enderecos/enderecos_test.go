package enderecos

import "testing"

//TESTE UNITÁRIO

type cenarioTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	t.Parallel()

	enderecoParaTeste := "Avenida Paulista"
	tipoEnderecoEsperado := "Avenida"
	tipoEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)
	if tipoEnderecoRecebido != tipoEnderecoEsperado {
		t.Errorf("#1 - O tipo recebido é diferente do tipo esperado. Esperava: %s, Recebido: %s", tipoEnderecoEsperado, tipoEnderecoRecebido)
	}
}

func TestVariosTipoDeEndereco(t *testing.T) {
	
	t.Parallel()

	cenariosDeTeste := []cenarioTeste{
		{"Rua Carlos Von Koseritz", "Rua"},
		{"Avenida Brasil", "Avenida"},
		//{"Estrada Tatuapé", "Estrada"},
		{"Rodovia João Dias", "Rodovia"},
		//{"Beco do Adelar", "Tipo Inválido"},
	}

	for _, cenarioTeste := range cenariosDeTeste {
		retorno := TipoDeEndereco(cenarioTeste.enderecoInserido)
		if retorno != cenarioTeste.retornoEsperado {
			t.Errorf("#2 - O tipo recebido é diferente do tipo esperado. Esperava: %s, Recebido: %s", cenarioTeste.retornoEsperado, retorno)
		}
	}

}
