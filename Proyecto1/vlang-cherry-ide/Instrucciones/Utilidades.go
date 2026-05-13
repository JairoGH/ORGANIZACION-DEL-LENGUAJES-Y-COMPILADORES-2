package instrucciones

import (
	"main/tiposDeDato"
	"regexp"
	"strings"
)

// CadenaAVector: convierte una cadena en un vector de caracteres individuales
func CadenaAVector(s *tiposDeDato.ValorCadena) *TipoVector {
	items := make([]tiposDeDato.ValorInterno, 0, len(s.ValorInterno))
	for _, rune := range s.ValorInterno {
		items = append(items, &tiposDeDato.ValorCadena{
			ValorInterno: string(rune),
		})
	}

	tag := "[" + tiposDeDato.TIPO_CADENA + "]"
	return NewTipoVector(items, tag, tiposDeDato.TIPO_CADENA)
}

// EsTipoVector: verifica si un tipo es vector (un solo nivel de corchetes)
// Distingue entre [tipo] (vector) y [[tipo]] (matriz)
func EsTipoVector(tipo string) bool {
	// Vector: inicia con un solo [ y termina con un solo ]
	formatoVector := "^\\[.*\\]$"

	// Matriz: inicia con al menos [[ y termina con al menos ]]
	formatoMatriz := "^\\[\\[.*\\]\\](\\[.*\\]\\])*$"

	match, _ := regexp.MatchString(formatoVector, tipo)
	match2, _ := regexp.MatchString(formatoMatriz, tipo)

	return match && !match2
}

// EliminarCorchetes: remueve los corchetes de apertura y cierre de un tipo
func EliminarCorchetes(s string) string {
	return strings.Trim(s, "[]")
}

// EsTipoMatriz: verifica si un tipo es matriz (múltiples niveles de corchetes)
func EsTipoMatriz(tipo string) bool {
	// Matriz: inicia con al menos [[ y termina con al menos ]]
	formatoMatriz := "^\\[\\[.*\\]\\](\\[.*\\]\\])*$"

	match, _ := regexp.MatchString(formatoMatriz, tipo)

	return match
}
