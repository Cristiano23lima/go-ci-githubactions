package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(15, 15)

	if total != 30 {
		t.Errorf("Error, resultado inválido. \nResultado esperado: %d;\nresultado retornado: %d", 30, total)
	}
}
