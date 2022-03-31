package enderecos

import "testing"

//TESTE UNITÁRIO

func TestTipoDeEndereco(t *testing.T) {

	enderecoParaTeste := "Avenida Paulista"
	tipoEnderecoEsperado := "Rua"
	tipoEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)
	if tipoEnderecoRecebido != tipoEnderecoEsperado {
		t.Errorf("O tipo recebido é diferente do tipo esperado. Esperava: %s, Recebido: %s", tipoEnderecoEsperado, tipoEnderecoRecebido)
	}
}
