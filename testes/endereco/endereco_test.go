package endereco

import (
	"testing"
)

type cenarioTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	cenariosTeste := []cenarioTeste{
		{"Rua rio feio", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia Conego", "Rodovia"},
		{"Estrada das rosas", "Estrada"},
		{"RUA RIO FEIO", "Rua"},
		{"AVENIDA PAULISTA", "Avenida"},
		{"RODOVIA CONEGO", "Rodovia"},
		{"ESTRADA DAS ROSAS", "Estrada"},
		{"", "tipo invalido"},
	}
	for _, cenario := range cenariosTeste {
		enderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if enderecoRecebido != cenario.retornoEsperado {
			t.Errorf("Teste falho, resultado esperado %s, resultado recebido %s", cenario.retornoEsperado, enderecoRecebido)
		}
	}
}
