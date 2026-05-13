package instrucciones

import "testing"

// ProbarTipoVector: prueba la función EsTipoVector con diferentes casos de entrada
func ProbarTipoVector(t *testing.T) {

	casos := map[string]bool{
		"[int]":   true,
		"int":     false,
		"[[int]]": false,
		"[]":      true,
	}

	for tipo, esperado := range casos {
		if EsTipoVector(tipo) != esperado {
			t.Errorf("EsTipoVector(%s) != %t", tipo, esperado)
		}
	}
}

// ProbarTipoMatriz: prueba la función EsTipoMatriz con diferentes casos de entrada
func ProbarTipoMatriz(t *testing.T) {
	casos := map[string]bool{
		"[int]":     false,
		"int":       false,
		"[[int]]":   true,
		"[]":        false,
		"[[[int]]]": true,
		"[":         false,
		"[[]":       false,
	}

	for tipo, esperado := range casos {
		if EsTipoMatriz(tipo) != esperado {
			t.Errorf("EsTipoMatriz(%s) != %t", tipo, esperado)
		}
	}
}
