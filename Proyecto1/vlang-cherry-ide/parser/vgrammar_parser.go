// Code generated from parser/VGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // VGrammar
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type VGrammar struct {
	*antlr.BaseParser
}

var VGrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func vgrammarParserInit() {
	staticData := &VGrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "'mut'", "':='", "'+='", "'-='", "'='", "'int'", "",
		"'string'", "'bool'", "", "", "", "", "'nil'", "'+'", "'-'", "'*'",
		"'/'", "'%'", "'=='", "'!='", "'<'", "'>'", "'<='", "'>='", "'&&'",
		"'||'", "'!'", "'main'", "'fn'", "'struct'", "'if'", "'else'", "'switch'",
		"'case'", "'default'", "'for'", "'while'", "'break'", "'continue'",
		"'return'", "'in'", "", "'('", "')'", "'{'", "'}'", "'['", "']'", "'.'",
		"','", "';'", "':'",
	}
	staticData.SymbolicNames = []string{
		"", "WS", "COMENTARIO", "COMENTARIOMULT", "RW_MUT", "OP_ASSIGN", "OP_INCREMENTO",
		"OP_DECREMENTO", "OP_DECLARACION", "RW_INT", "RW_FLOAT64", "RW_STRING",
		"RW_BOOL", "INT_LITERAL", "FLOAT_LITERAL", "STRING_LITERAL", "BOOL_LITERAL",
		"NIL_LITERAL", "OP_SUMA", "OP_RESTA", "OP_MULT", "OP_DIV", "OP_MOD",
		"OP_IGUAL", "OP_DIFERENTE", "OP_MENORQ", "OP_MAYORQ", "OP_MENORIGUAL",
		"OP_MAYORIGUAL", "OP_AND", "OP_OR", "OP_NOT", "RW_MAIN", "RW_FN", "RW_STRUCT",
		"RW_IF", "RW_ELSE", "RW_SWITCH", "RW_CASE", "RW_DEFAULT", "RW_FOR",
		"RW_WHILE", "RW_BREAK", "RW_CONTINUE", "RW_RETURN", "RW_IN", "ID", "PAR_IZQ",
		"PAR_DER", "LLAVE_IZQ", "LLAVE_DER", "CORCHETE_IZQ", "CORCHETE_DER",
		"PUNTO", "COMA", "PUNTO_Y_COMA", "DOS_PUNTOS",
	}
	staticData.RuleNames = []string{
		"program", "main_func", "stmt", "stmt_declaracion", "stmt_asignar",
		"stmt_transferencia", "if_stmt", "if_chain", "else_stmt", "switch_stmt",
		"switch_case", "default_case", "while_stmt", "for_stmt", "range", "declarar_funcion",
		"llamar_funcion", "lista_parametros", "parametro_fun", "lista_argumentos",
		"argumento_fun", "declarar_struct", "propiedad_struct", "crear_struct",
		"lista_parametros_init", "parametros_init_struct", "lista_slice", "item_slice",
		"prop_slice", "fun_slice", "tipo", "tipos_slices", "expr", "literal",
		"patronId",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 56, 470, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 1, 0, 5, 0, 72, 8, 0, 10,
		0, 12, 0, 75, 9, 0, 1, 0, 3, 0, 78, 8, 0, 1, 0, 3, 0, 81, 8, 0, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 89, 8, 1, 10, 1, 12, 1, 92, 9, 1, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		3, 2, 107, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 3, 3, 131, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 137, 8, 3, 3, 3, 139,
		8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 157, 8, 4, 1, 5, 1, 5, 3, 5, 161, 8,
		5, 1, 5, 1, 5, 3, 5, 165, 8, 5, 1, 6, 1, 6, 1, 6, 5, 6, 170, 8, 6, 10,
		6, 12, 6, 173, 9, 6, 1, 6, 3, 6, 176, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 5,
		7, 182, 8, 7, 10, 7, 12, 7, 185, 9, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 5,
		8, 192, 8, 8, 10, 8, 12, 8, 195, 9, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1,
		9, 5, 9, 203, 8, 9, 10, 9, 12, 9, 206, 9, 9, 1, 9, 3, 9, 209, 8, 9, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 217, 8, 10, 10, 10, 12, 10,
		220, 9, 10, 1, 11, 1, 11, 1, 11, 5, 11, 225, 8, 11, 10, 11, 12, 11, 228,
		9, 11, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 234, 8, 12, 10, 12, 12, 12, 237,
		9, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 246, 8,
		13, 1, 13, 1, 13, 5, 13, 250, 8, 13, 10, 13, 12, 13, 253, 9, 13, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1,
		15, 3, 15, 267, 8, 15, 1, 15, 1, 15, 3, 15, 271, 8, 15, 1, 15, 1, 15, 5,
		15, 275, 8, 15, 10, 15, 12, 15, 278, 9, 15, 1, 15, 1, 15, 1, 16, 1, 16,
		1, 16, 3, 16, 285, 8, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 5, 17, 292,
		8, 17, 10, 17, 12, 17, 295, 9, 17, 1, 17, 3, 17, 298, 8, 17, 1, 18, 3,
		18, 301, 8, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 5, 19, 309, 8,
		19, 10, 19, 12, 19, 312, 9, 19, 1, 19, 3, 19, 315, 8, 19, 1, 20, 1, 20,
		3, 20, 319, 8, 20, 1, 20, 1, 20, 3, 20, 323, 8, 20, 1, 21, 1, 21, 1, 21,
		1, 21, 5, 21, 329, 8, 21, 10, 21, 12, 21, 332, 9, 21, 1, 21, 1, 21, 1,
		22, 1, 22, 1, 22, 1, 22, 3, 22, 340, 8, 22, 1, 23, 1, 23, 1, 23, 3, 23,
		345, 8, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 5, 24, 352, 8, 24, 10, 24,
		12, 24, 355, 9, 24, 1, 24, 3, 24, 358, 8, 24, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 26, 1, 26, 1, 26, 1, 26, 5, 26, 368, 8, 26, 10, 26, 12, 26, 371, 9,
		26, 3, 26, 373, 8, 26, 1, 26, 3, 26, 376, 8, 26, 1, 26, 1, 26, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 4, 27, 385, 8, 27, 11, 27, 12, 27, 386, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 30, 3, 30, 403, 8, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 3, 31, 413, 8, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32,
		3, 32, 430, 8, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32,
		5, 32, 450, 8, 32, 10, 32, 12, 32, 453, 9, 32, 1, 33, 1, 33, 1, 33, 1,
		33, 1, 33, 3, 33, 460, 8, 33, 1, 34, 1, 34, 1, 34, 5, 34, 465, 8, 34, 10,
		34, 12, 34, 468, 9, 34, 1, 34, 0, 1, 64, 35, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
		52, 54, 56, 58, 60, 62, 64, 66, 68, 0, 7, 1, 0, 6, 7, 1, 0, 6, 8, 2, 0,
		19, 19, 31, 31, 1, 0, 20, 22, 1, 0, 18, 19, 1, 0, 25, 28, 1, 0, 23, 24,
		518, 0, 73, 1, 0, 0, 0, 2, 82, 1, 0, 0, 0, 4, 106, 1, 0, 0, 0, 6, 138,
		1, 0, 0, 0, 8, 156, 1, 0, 0, 0, 10, 164, 1, 0, 0, 0, 12, 166, 1, 0, 0,
		0, 14, 177, 1, 0, 0, 0, 16, 188, 1, 0, 0, 0, 18, 198, 1, 0, 0, 0, 20, 212,
		1, 0, 0, 0, 22, 221, 1, 0, 0, 0, 24, 229, 1, 0, 0, 0, 26, 240, 1, 0, 0,
		0, 28, 256, 1, 0, 0, 0, 30, 262, 1, 0, 0, 0, 32, 281, 1, 0, 0, 0, 34, 288,
		1, 0, 0, 0, 36, 300, 1, 0, 0, 0, 38, 305, 1, 0, 0, 0, 40, 318, 1, 0, 0,
		0, 42, 324, 1, 0, 0, 0, 44, 335, 1, 0, 0, 0, 46, 341, 1, 0, 0, 0, 48, 348,
		1, 0, 0, 0, 50, 359, 1, 0, 0, 0, 52, 363, 1, 0, 0, 0, 54, 379, 1, 0, 0,
		0, 56, 388, 1, 0, 0, 0, 58, 392, 1, 0, 0, 0, 60, 402, 1, 0, 0, 0, 62, 412,
		1, 0, 0, 0, 64, 429, 1, 0, 0, 0, 66, 459, 1, 0, 0, 0, 68, 461, 1, 0, 0,
		0, 70, 72, 3, 4, 2, 0, 71, 70, 1, 0, 0, 0, 72, 75, 1, 0, 0, 0, 73, 71,
		1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 77, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0,
		76, 78, 3, 2, 1, 0, 77, 76, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 80, 1,
		0, 0, 0, 79, 81, 5, 0, 0, 1, 80, 79, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81,
		1, 1, 0, 0, 0, 82, 83, 5, 33, 0, 0, 83, 84, 5, 32, 0, 0, 84, 85, 5, 47,
		0, 0, 85, 86, 5, 48, 0, 0, 86, 90, 5, 49, 0, 0, 87, 89, 3, 4, 2, 0, 88,
		87, 1, 0, 0, 0, 89, 92, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0,
		0, 91, 93, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 93, 94, 5, 50, 0, 0, 94, 3,
		1, 0, 0, 0, 95, 107, 3, 6, 3, 0, 96, 107, 3, 8, 4, 0, 97, 107, 3, 10, 5,
		0, 98, 107, 3, 12, 6, 0, 99, 107, 3, 18, 9, 0, 100, 107, 3, 24, 12, 0,
		101, 107, 3, 26, 13, 0, 102, 107, 3, 32, 16, 0, 103, 107, 3, 58, 29, 0,
		104, 107, 3, 30, 15, 0, 105, 107, 3, 42, 21, 0, 106, 95, 1, 0, 0, 0, 106,
		96, 1, 0, 0, 0, 106, 97, 1, 0, 0, 0, 106, 98, 1, 0, 0, 0, 106, 99, 1, 0,
		0, 0, 106, 100, 1, 0, 0, 0, 106, 101, 1, 0, 0, 0, 106, 102, 1, 0, 0, 0,
		106, 103, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 106, 105, 1, 0, 0, 0, 107,
		5, 1, 0, 0, 0, 108, 109, 5, 4, 0, 0, 109, 110, 5, 46, 0, 0, 110, 111, 5,
		5, 0, 0, 111, 139, 3, 64, 32, 0, 112, 113, 5, 46, 0, 0, 113, 114, 5, 5,
		0, 0, 114, 139, 3, 64, 32, 0, 115, 116, 5, 4, 0, 0, 116, 117, 5, 46, 0,
		0, 117, 118, 3, 60, 30, 0, 118, 119, 5, 8, 0, 0, 119, 120, 3, 64, 32, 0,
		120, 139, 1, 0, 0, 0, 121, 122, 5, 4, 0, 0, 122, 123, 5, 46, 0, 0, 123,
		139, 3, 60, 30, 0, 124, 125, 5, 46, 0, 0, 125, 126, 3, 60, 30, 0, 126,
		127, 5, 8, 0, 0, 127, 128, 3, 64, 32, 0, 128, 139, 1, 0, 0, 0, 129, 131,
		5, 4, 0, 0, 130, 129, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 132, 1, 0,
		0, 0, 132, 133, 5, 46, 0, 0, 133, 134, 5, 8, 0, 0, 134, 136, 3, 60, 30,
		0, 135, 137, 3, 64, 32, 0, 136, 135, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0,
		137, 139, 1, 0, 0, 0, 138, 108, 1, 0, 0, 0, 138, 112, 1, 0, 0, 0, 138,
		115, 1, 0, 0, 0, 138, 121, 1, 0, 0, 0, 138, 124, 1, 0, 0, 0, 138, 130,
		1, 0, 0, 0, 139, 7, 1, 0, 0, 0, 140, 141, 3, 54, 27, 0, 141, 142, 5, 8,
		0, 0, 142, 143, 3, 64, 32, 0, 143, 157, 1, 0, 0, 0, 144, 145, 3, 68, 34,
		0, 145, 146, 5, 8, 0, 0, 146, 147, 3, 64, 32, 0, 147, 157, 1, 0, 0, 0,
		148, 149, 3, 68, 34, 0, 149, 150, 7, 0, 0, 0, 150, 151, 3, 64, 32, 0, 151,
		157, 1, 0, 0, 0, 152, 153, 3, 54, 27, 0, 153, 154, 7, 1, 0, 0, 154, 155,
		3, 64, 32, 0, 155, 157, 1, 0, 0, 0, 156, 140, 1, 0, 0, 0, 156, 144, 1,
		0, 0, 0, 156, 148, 1, 0, 0, 0, 156, 152, 1, 0, 0, 0, 157, 9, 1, 0, 0, 0,
		158, 160, 5, 44, 0, 0, 159, 161, 3, 64, 32, 0, 160, 159, 1, 0, 0, 0, 160,
		161, 1, 0, 0, 0, 161, 165, 1, 0, 0, 0, 162, 165, 5, 42, 0, 0, 163, 165,
		5, 43, 0, 0, 164, 158, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 164, 163, 1, 0,
		0, 0, 165, 11, 1, 0, 0, 0, 166, 171, 3, 14, 7, 0, 167, 168, 5, 36, 0, 0,
		168, 170, 3, 14, 7, 0, 169, 167, 1, 0, 0, 0, 170, 173, 1, 0, 0, 0, 171,
		169, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 175, 1, 0, 0, 0, 173, 171,
		1, 0, 0, 0, 174, 176, 3, 16, 8, 0, 175, 174, 1, 0, 0, 0, 175, 176, 1, 0,
		0, 0, 176, 13, 1, 0, 0, 0, 177, 178, 5, 35, 0, 0, 178, 179, 3, 64, 32,
		0, 179, 183, 5, 49, 0, 0, 180, 182, 3, 4, 2, 0, 181, 180, 1, 0, 0, 0, 182,
		185, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 186,
		1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 186, 187, 5, 50, 0, 0, 187, 15, 1, 0,
		0, 0, 188, 189, 5, 36, 0, 0, 189, 193, 5, 49, 0, 0, 190, 192, 3, 4, 2,
		0, 191, 190, 1, 0, 0, 0, 192, 195, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 193,
		194, 1, 0, 0, 0, 194, 196, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 196, 197,
		5, 50, 0, 0, 197, 17, 1, 0, 0, 0, 198, 199, 5, 37, 0, 0, 199, 200, 3, 64,
		32, 0, 200, 204, 5, 49, 0, 0, 201, 203, 3, 20, 10, 0, 202, 201, 1, 0, 0,
		0, 203, 206, 1, 0, 0, 0, 204, 202, 1, 0, 0, 0, 204, 205, 1, 0, 0, 0, 205,
		208, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 207, 209, 3, 22, 11, 0, 208, 207,
		1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 211, 5, 50,
		0, 0, 211, 19, 1, 0, 0, 0, 212, 213, 5, 38, 0, 0, 213, 214, 3, 64, 32,
		0, 214, 218, 5, 56, 0, 0, 215, 217, 3, 4, 2, 0, 216, 215, 1, 0, 0, 0, 217,
		220, 1, 0, 0, 0, 218, 216, 1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219, 21, 1,
		0, 0, 0, 220, 218, 1, 0, 0, 0, 221, 222, 5, 39, 0, 0, 222, 226, 5, 56,
		0, 0, 223, 225, 3, 4, 2, 0, 224, 223, 1, 0, 0, 0, 225, 228, 1, 0, 0, 0,
		226, 224, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 23, 1, 0, 0, 0, 228, 226,
		1, 0, 0, 0, 229, 230, 5, 40, 0, 0, 230, 231, 3, 64, 32, 0, 231, 235, 5,
		49, 0, 0, 232, 234, 3, 4, 2, 0, 233, 232, 1, 0, 0, 0, 234, 237, 1, 0, 0,
		0, 235, 233, 1, 0, 0, 0, 235, 236, 1, 0, 0, 0, 236, 238, 1, 0, 0, 0, 237,
		235, 1, 0, 0, 0, 238, 239, 5, 50, 0, 0, 239, 25, 1, 0, 0, 0, 240, 241,
		5, 40, 0, 0, 241, 242, 5, 46, 0, 0, 242, 245, 5, 45, 0, 0, 243, 246, 3,
		64, 32, 0, 244, 246, 3, 28, 14, 0, 245, 243, 1, 0, 0, 0, 245, 244, 1, 0,
		0, 0, 246, 247, 1, 0, 0, 0, 247, 251, 5, 49, 0, 0, 248, 250, 3, 4, 2, 0,
		249, 248, 1, 0, 0, 0, 250, 253, 1, 0, 0, 0, 251, 249, 1, 0, 0, 0, 251,
		252, 1, 0, 0, 0, 252, 254, 1, 0, 0, 0, 253, 251, 1, 0, 0, 0, 254, 255,
		5, 50, 0, 0, 255, 27, 1, 0, 0, 0, 256, 257, 3, 64, 32, 0, 257, 258, 5,
		53, 0, 0, 258, 259, 5, 53, 0, 0, 259, 260, 5, 53, 0, 0, 260, 261, 3, 64,
		32, 0, 261, 29, 1, 0, 0, 0, 262, 263, 5, 33, 0, 0, 263, 264, 5, 46, 0,
		0, 264, 266, 5, 47, 0, 0, 265, 267, 3, 34, 17, 0, 266, 265, 1, 0, 0, 0,
		266, 267, 1, 0, 0, 0, 267, 268, 1, 0, 0, 0, 268, 270, 5, 48, 0, 0, 269,
		271, 3, 60, 30, 0, 270, 269, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 272,
		1, 0, 0, 0, 272, 276, 5, 49, 0, 0, 273, 275, 3, 4, 2, 0, 274, 273, 1, 0,
		0, 0, 275, 278, 1, 0, 0, 0, 276, 274, 1, 0, 0, 0, 276, 277, 1, 0, 0, 0,
		277, 279, 1, 0, 0, 0, 278, 276, 1, 0, 0, 0, 279, 280, 5, 50, 0, 0, 280,
		31, 1, 0, 0, 0, 281, 282, 3, 68, 34, 0, 282, 284, 5, 47, 0, 0, 283, 285,
		3, 38, 19, 0, 284, 283, 1, 0, 0, 0, 284, 285, 1, 0, 0, 0, 285, 286, 1,
		0, 0, 0, 286, 287, 5, 48, 0, 0, 287, 33, 1, 0, 0, 0, 288, 293, 3, 36, 18,
		0, 289, 290, 5, 54, 0, 0, 290, 292, 3, 36, 18, 0, 291, 289, 1, 0, 0, 0,
		292, 295, 1, 0, 0, 0, 293, 291, 1, 0, 0, 0, 293, 294, 1, 0, 0, 0, 294,
		297, 1, 0, 0, 0, 295, 293, 1, 0, 0, 0, 296, 298, 5, 54, 0, 0, 297, 296,
		1, 0, 0, 0, 297, 298, 1, 0, 0, 0, 298, 35, 1, 0, 0, 0, 299, 301, 5, 46,
		0, 0, 300, 299, 1, 0, 0, 0, 300, 301, 1, 0, 0, 0, 301, 302, 1, 0, 0, 0,
		302, 303, 5, 46, 0, 0, 303, 304, 3, 60, 30, 0, 304, 37, 1, 0, 0, 0, 305,
		310, 3, 40, 20, 0, 306, 307, 5, 54, 0, 0, 307, 309, 3, 40, 20, 0, 308,
		306, 1, 0, 0, 0, 309, 312, 1, 0, 0, 0, 310, 308, 1, 0, 0, 0, 310, 311,
		1, 0, 0, 0, 311, 314, 1, 0, 0, 0, 312, 310, 1, 0, 0, 0, 313, 315, 5, 54,
		0, 0, 314, 313, 1, 0, 0, 0, 314, 315, 1, 0, 0, 0, 315, 39, 1, 0, 0, 0,
		316, 317, 5, 46, 0, 0, 317, 319, 5, 56, 0, 0, 318, 316, 1, 0, 0, 0, 318,
		319, 1, 0, 0, 0, 319, 322, 1, 0, 0, 0, 320, 323, 3, 68, 34, 0, 321, 323,
		3, 64, 32, 0, 322, 320, 1, 0, 0, 0, 322, 321, 1, 0, 0, 0, 323, 41, 1, 0,
		0, 0, 324, 325, 5, 34, 0, 0, 325, 326, 5, 46, 0, 0, 326, 330, 5, 49, 0,
		0, 327, 329, 3, 44, 22, 0, 328, 327, 1, 0, 0, 0, 329, 332, 1, 0, 0, 0,
		330, 328, 1, 0, 0, 0, 330, 331, 1, 0, 0, 0, 331, 333, 1, 0, 0, 0, 332,
		330, 1, 0, 0, 0, 333, 334, 5, 50, 0, 0, 334, 43, 1, 0, 0, 0, 335, 336,
		3, 60, 30, 0, 336, 339, 5, 46, 0, 0, 337, 338, 5, 8, 0, 0, 338, 340, 3,
		64, 32, 0, 339, 337, 1, 0, 0, 0, 339, 340, 1, 0, 0, 0, 340, 45, 1, 0, 0,
		0, 341, 342, 5, 46, 0, 0, 342, 344, 5, 49, 0, 0, 343, 345, 3, 48, 24, 0,
		344, 343, 1, 0, 0, 0, 344, 345, 1, 0, 0, 0, 345, 346, 1, 0, 0, 0, 346,
		347, 5, 50, 0, 0, 347, 47, 1, 0, 0, 0, 348, 353, 3, 50, 25, 0, 349, 350,
		5, 54, 0, 0, 350, 352, 3, 50, 25, 0, 351, 349, 1, 0, 0, 0, 352, 355, 1,
		0, 0, 0, 353, 351, 1, 0, 0, 0, 353, 354, 1, 0, 0, 0, 354, 357, 1, 0, 0,
		0, 355, 353, 1, 0, 0, 0, 356, 358, 5, 54, 0, 0, 357, 356, 1, 0, 0, 0, 357,
		358, 1, 0, 0, 0, 358, 49, 1, 0, 0, 0, 359, 360, 5, 46, 0, 0, 360, 361,
		5, 56, 0, 0, 361, 362, 3, 64, 32, 0, 362, 51, 1, 0, 0, 0, 363, 372, 5,
		49, 0, 0, 364, 369, 3, 64, 32, 0, 365, 366, 5, 54, 0, 0, 366, 368, 3, 64,
		32, 0, 367, 365, 1, 0, 0, 0, 368, 371, 1, 0, 0, 0, 369, 367, 1, 0, 0, 0,
		369, 370, 1, 0, 0, 0, 370, 373, 1, 0, 0, 0, 371, 369, 1, 0, 0, 0, 372,
		364, 1, 0, 0, 0, 372, 373, 1, 0, 0, 0, 373, 375, 1, 0, 0, 0, 374, 376,
		5, 54, 0, 0, 375, 374, 1, 0, 0, 0, 375, 376, 1, 0, 0, 0, 376, 377, 1, 0,
		0, 0, 377, 378, 5, 50, 0, 0, 378, 53, 1, 0, 0, 0, 379, 384, 3, 68, 34,
		0, 380, 381, 5, 51, 0, 0, 381, 382, 3, 64, 32, 0, 382, 383, 5, 52, 0, 0,
		383, 385, 1, 0, 0, 0, 384, 380, 1, 0, 0, 0, 385, 386, 1, 0, 0, 0, 386,
		384, 1, 0, 0, 0, 386, 387, 1, 0, 0, 0, 387, 55, 1, 0, 0, 0, 388, 389, 3,
		54, 27, 0, 389, 390, 5, 53, 0, 0, 390, 391, 3, 68, 34, 0, 391, 57, 1, 0,
		0, 0, 392, 393, 3, 54, 27, 0, 393, 394, 5, 53, 0, 0, 394, 395, 3, 32, 16,
		0, 395, 59, 1, 0, 0, 0, 396, 403, 5, 46, 0, 0, 397, 403, 5, 9, 0, 0, 398,
		403, 5, 10, 0, 0, 399, 403, 5, 11, 0, 0, 400, 403, 5, 12, 0, 0, 401, 403,
		3, 62, 31, 0, 402, 396, 1, 0, 0, 0, 402, 397, 1, 0, 0, 0, 402, 398, 1,
		0, 0, 0, 402, 399, 1, 0, 0, 0, 402, 400, 1, 0, 0, 0, 402, 401, 1, 0, 0,
		0, 403, 61, 1, 0, 0, 0, 404, 405, 5, 51, 0, 0, 405, 406, 5, 52, 0, 0, 406,
		413, 3, 60, 30, 0, 407, 408, 5, 51, 0, 0, 408, 409, 5, 52, 0, 0, 409, 410,
		5, 51, 0, 0, 410, 411, 5, 52, 0, 0, 411, 413, 3, 60, 30, 0, 412, 404, 1,
		0, 0, 0, 412, 407, 1, 0, 0, 0, 413, 63, 1, 0, 0, 0, 414, 415, 6, 32, -1,
		0, 415, 416, 5, 47, 0, 0, 416, 417, 3, 64, 32, 0, 417, 418, 5, 48, 0, 0,
		418, 430, 1, 0, 0, 0, 419, 430, 3, 32, 16, 0, 420, 430, 3, 68, 34, 0, 421,
		430, 3, 54, 27, 0, 422, 430, 3, 56, 28, 0, 423, 430, 3, 58, 29, 0, 424,
		430, 3, 66, 33, 0, 425, 430, 3, 52, 26, 0, 426, 430, 3, 46, 23, 0, 427,
		428, 7, 2, 0, 0, 428, 430, 3, 64, 32, 7, 429, 414, 1, 0, 0, 0, 429, 419,
		1, 0, 0, 0, 429, 420, 1, 0, 0, 0, 429, 421, 1, 0, 0, 0, 429, 422, 1, 0,
		0, 0, 429, 423, 1, 0, 0, 0, 429, 424, 1, 0, 0, 0, 429, 425, 1, 0, 0, 0,
		429, 426, 1, 0, 0, 0, 429, 427, 1, 0, 0, 0, 430, 451, 1, 0, 0, 0, 431,
		432, 10, 6, 0, 0, 432, 433, 7, 3, 0, 0, 433, 450, 3, 64, 32, 7, 434, 435,
		10, 5, 0, 0, 435, 436, 7, 4, 0, 0, 436, 450, 3, 64, 32, 6, 437, 438, 10,
		4, 0, 0, 438, 439, 7, 5, 0, 0, 439, 450, 3, 64, 32, 5, 440, 441, 10, 3,
		0, 0, 441, 442, 7, 6, 0, 0, 442, 450, 3, 64, 32, 4, 443, 444, 10, 2, 0,
		0, 444, 445, 5, 29, 0, 0, 445, 450, 3, 64, 32, 3, 446, 447, 10, 1, 0, 0,
		447, 448, 5, 30, 0, 0, 448, 450, 3, 64, 32, 2, 449, 431, 1, 0, 0, 0, 449,
		434, 1, 0, 0, 0, 449, 437, 1, 0, 0, 0, 449, 440, 1, 0, 0, 0, 449, 443,
		1, 0, 0, 0, 449, 446, 1, 0, 0, 0, 450, 453, 1, 0, 0, 0, 451, 449, 1, 0,
		0, 0, 451, 452, 1, 0, 0, 0, 452, 65, 1, 0, 0, 0, 453, 451, 1, 0, 0, 0,
		454, 460, 5, 13, 0, 0, 455, 460, 5, 14, 0, 0, 456, 460, 5, 15, 0, 0, 457,
		460, 5, 16, 0, 0, 458, 460, 5, 17, 0, 0, 459, 454, 1, 0, 0, 0, 459, 455,
		1, 0, 0, 0, 459, 456, 1, 0, 0, 0, 459, 457, 1, 0, 0, 0, 459, 458, 1, 0,
		0, 0, 460, 67, 1, 0, 0, 0, 461, 466, 5, 46, 0, 0, 462, 463, 5, 53, 0, 0,
		463, 465, 5, 46, 0, 0, 464, 462, 1, 0, 0, 0, 465, 468, 1, 0, 0, 0, 466,
		464, 1, 0, 0, 0, 466, 467, 1, 0, 0, 0, 467, 69, 1, 0, 0, 0, 468, 466, 1,
		0, 0, 0, 49, 73, 77, 80, 90, 106, 130, 136, 138, 156, 160, 164, 171, 175,
		183, 193, 204, 208, 218, 226, 235, 245, 251, 266, 270, 276, 284, 293, 297,
		300, 310, 314, 318, 322, 330, 339, 344, 353, 357, 369, 372, 375, 386, 402,
		412, 429, 449, 451, 459, 466,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// VGrammarInit initializes any static state used to implement VGrammar. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewVGrammar(). You can call this function if you wish to initialize the static state ahead
// of time.
func VGrammarInit() {
	staticData := &VGrammarParserStaticData
	staticData.once.Do(vgrammarParserInit)
}

// NewVGrammar produces a new parser instance for the optional input antlr.TokenStream.
func NewVGrammar(input antlr.TokenStream) *VGrammar {
	VGrammarInit()
	this := new(VGrammar)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &VGrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "VGrammar.g4"

	return this
}

// VGrammar tokens.
const (
	VGrammarEOF            = antlr.TokenEOF
	VGrammarWS             = 1
	VGrammarCOMENTARIO     = 2
	VGrammarCOMENTARIOMULT = 3
	VGrammarRW_MUT         = 4
	VGrammarOP_ASSIGN      = 5
	VGrammarOP_INCREMENTO  = 6
	VGrammarOP_DECREMENTO  = 7
	VGrammarOP_DECLARACION = 8
	VGrammarRW_INT         = 9
	VGrammarRW_FLOAT64     = 10
	VGrammarRW_STRING      = 11
	VGrammarRW_BOOL        = 12
	VGrammarINT_LITERAL    = 13
	VGrammarFLOAT_LITERAL  = 14
	VGrammarSTRING_LITERAL = 15
	VGrammarBOOL_LITERAL   = 16
	VGrammarNIL_LITERAL    = 17
	VGrammarOP_SUMA        = 18
	VGrammarOP_RESTA       = 19
	VGrammarOP_MULT        = 20
	VGrammarOP_DIV         = 21
	VGrammarOP_MOD         = 22
	VGrammarOP_IGUAL       = 23
	VGrammarOP_DIFERENTE   = 24
	VGrammarOP_MENORQ      = 25
	VGrammarOP_MAYORQ      = 26
	VGrammarOP_MENORIGUAL  = 27
	VGrammarOP_MAYORIGUAL  = 28
	VGrammarOP_AND         = 29
	VGrammarOP_OR          = 30
	VGrammarOP_NOT         = 31
	VGrammarRW_MAIN        = 32
	VGrammarRW_FN          = 33
	VGrammarRW_STRUCT      = 34
	VGrammarRW_IF          = 35
	VGrammarRW_ELSE        = 36
	VGrammarRW_SWITCH      = 37
	VGrammarRW_CASE        = 38
	VGrammarRW_DEFAULT     = 39
	VGrammarRW_FOR         = 40
	VGrammarRW_WHILE       = 41
	VGrammarRW_BREAK       = 42
	VGrammarRW_CONTINUE    = 43
	VGrammarRW_RETURN      = 44
	VGrammarRW_IN          = 45
	VGrammarID             = 46
	VGrammarPAR_IZQ        = 47
	VGrammarPAR_DER        = 48
	VGrammarLLAVE_IZQ      = 49
	VGrammarLLAVE_DER      = 50
	VGrammarCORCHETE_IZQ   = 51
	VGrammarCORCHETE_DER   = 52
	VGrammarPUNTO          = 53
	VGrammarCOMA           = 54
	VGrammarPUNTO_Y_COMA   = 55
	VGrammarDOS_PUNTOS     = 56
)

// VGrammar rules.
const (
	VGrammarRULE_program                = 0
	VGrammarRULE_main_func              = 1
	VGrammarRULE_stmt                   = 2
	VGrammarRULE_stmt_declaracion       = 3
	VGrammarRULE_stmt_asignar           = 4
	VGrammarRULE_stmt_transferencia     = 5
	VGrammarRULE_if_stmt                = 6
	VGrammarRULE_if_chain               = 7
	VGrammarRULE_else_stmt              = 8
	VGrammarRULE_switch_stmt            = 9
	VGrammarRULE_switch_case            = 10
	VGrammarRULE_default_case           = 11
	VGrammarRULE_while_stmt             = 12
	VGrammarRULE_for_stmt               = 13
	VGrammarRULE_range                  = 14
	VGrammarRULE_declarar_funcion       = 15
	VGrammarRULE_llamar_funcion         = 16
	VGrammarRULE_lista_parametros       = 17
	VGrammarRULE_parametro_fun          = 18
	VGrammarRULE_lista_argumentos       = 19
	VGrammarRULE_argumento_fun          = 20
	VGrammarRULE_declarar_struct        = 21
	VGrammarRULE_propiedad_struct       = 22
	VGrammarRULE_crear_struct           = 23
	VGrammarRULE_lista_parametros_init  = 24
	VGrammarRULE_parametros_init_struct = 25
	VGrammarRULE_lista_slice            = 26
	VGrammarRULE_item_slice             = 27
	VGrammarRULE_prop_slice             = 28
	VGrammarRULE_fun_slice              = 29
	VGrammarRULE_tipo                   = 30
	VGrammarRULE_tipos_slices           = 31
	VGrammarRULE_expr                   = 32
	VGrammarRULE_literal                = 33
	VGrammarRULE_patronId               = 34
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext
	Main_func() IMain_funcContext
	EOF() antlr.TerminalNode

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ProgramContext) Main_func() IMain_funcContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMain_funcContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMain_funcContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(VGrammarEOF, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VGrammarRULE_program)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(70)
				p.Stmt()
			}

		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarRW_FN {
		{
			p.SetState(76)
			p.Main_func()
		}

	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(79)
			p.Match(VGrammarEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMain_funcContext is an interface to support dynamic dispatch.
type IMain_funcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMain_funcContext differentiates from other interfaces.
	IsMain_funcContext()
}

type Main_funcContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMain_funcContext() *Main_funcContext {
	var p = new(Main_funcContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_main_func
	return p
}

func InitEmptyMain_funcContext(p *Main_funcContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_main_func
}

func (*Main_funcContext) IsMain_funcContext() {}

func NewMain_funcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Main_funcContext {
	var p = new(Main_funcContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_main_func

	return p
}

func (s *Main_funcContext) GetParser() antlr.Parser { return s.parser }

func (s *Main_funcContext) CopyAll(ctx *Main_funcContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Main_funcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Main_funcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncionMainContext struct {
	Main_funcContext
}

func NewFuncionMainContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncionMainContext {
	var p = new(FuncionMainContext)

	InitEmptyMain_funcContext(&p.Main_funcContext)
	p.parser = parser
	p.CopyAll(ctx.(*Main_funcContext))

	return p
}

func (s *FuncionMainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionMainContext) RW_FN() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_FN, 0)
}

func (s *FuncionMainContext) RW_MAIN() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_MAIN, 0)
}

func (s *FuncionMainContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_IZQ, 0)
}

func (s *FuncionMainContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_DER, 0)
}

func (s *FuncionMainContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *FuncionMainContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *FuncionMainContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *FuncionMainContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *FuncionMainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterFuncionMain(s)
	}
}

func (s *FuncionMainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitFuncionMain(s)
	}
}

func (s *FuncionMainContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitFuncionMain(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Main_func() (localctx IMain_funcContext) {
	localctx = NewMain_funcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, VGrammarRULE_main_func)
	var _la int

	localctx = NewFuncionMainContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Match(VGrammarRW_FN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(83)
		p.Match(VGrammarRW_MAIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(84)
		p.Match(VGrammarPAR_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(85)
		p.Match(VGrammarPAR_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(86)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&102452149878800) != 0 {
		{
			p.SetState(87)
			p.Stmt()
		}

		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(93)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Stmt_declaracion() IStmt_declaracionContext
	Stmt_asignar() IStmt_asignarContext
	Stmt_transferencia() IStmt_transferenciaContext
	If_stmt() IIf_stmtContext
	Switch_stmt() ISwitch_stmtContext
	While_stmt() IWhile_stmtContext
	For_stmt() IFor_stmtContext
	Llamar_funcion() ILlamar_funcionContext
	Fun_slice() IFun_sliceContext
	Declarar_funcion() IDeclarar_funcionContext
	Declarar_struct() IDeclarar_structContext

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Stmt_declaracion() IStmt_declaracionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmt_declaracionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmt_declaracionContext)
}

func (s *StmtContext) Stmt_asignar() IStmt_asignarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmt_asignarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmt_asignarContext)
}

func (s *StmtContext) Stmt_transferencia() IStmt_transferenciaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmt_transferenciaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmt_transferenciaContext)
}

func (s *StmtContext) If_stmt() IIf_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_stmtContext)
}

func (s *StmtContext) Switch_stmt() ISwitch_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitch_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitch_stmtContext)
}

func (s *StmtContext) While_stmt() IWhile_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhile_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhile_stmtContext)
}

func (s *StmtContext) For_stmt() IFor_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFor_stmtContext)
}

func (s *StmtContext) Llamar_funcion() ILlamar_funcionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamar_funcionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamar_funcionContext)
}

func (s *StmtContext) Fun_slice() IFun_sliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFun_sliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFun_sliceContext)
}

func (s *StmtContext) Declarar_funcion() IDeclarar_funcionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarar_funcionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarar_funcionContext)
}

func (s *StmtContext) Declarar_struct() IDeclarar_structContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarar_structContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarar_structContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, VGrammarRULE_stmt)
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(95)
			p.Stmt_declaracion()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(96)
			p.Stmt_asignar()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(97)
			p.Stmt_transferencia()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(98)
			p.If_stmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(99)
			p.Switch_stmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(100)
			p.While_stmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(101)
			p.For_stmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(102)
			p.Llamar_funcion()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(103)
			p.Fun_slice()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(104)
			p.Declarar_funcion()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(105)
			p.Declarar_struct()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmt_declaracionContext is an interface to support dynamic dispatch.
type IStmt_declaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStmt_declaracionContext differentiates from other interfaces.
	IsStmt_declaracionContext()
}

type Stmt_declaracionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmt_declaracionContext() *Stmt_declaracionContext {
	var p = new(Stmt_declaracionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_stmt_declaracion
	return p
}

func InitEmptyStmt_declaracionContext(p *Stmt_declaracionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_stmt_declaracion
}

func (*Stmt_declaracionContext) IsStmt_declaracionContext() {}

func NewStmt_declaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Stmt_declaracionContext {
	var p = new(Stmt_declaracionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_stmt_declaracion

	return p
}

func (s *Stmt_declaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *Stmt_declaracionContext) CopyAll(ctx *Stmt_declaracionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Stmt_declaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stmt_declaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DeclararSliceContext struct {
	Stmt_declaracionContext
}

func NewDeclararSliceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclararSliceContext {
	var p = new(DeclararSliceContext)

	InitEmptyStmt_declaracionContext(&p.Stmt_declaracionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Stmt_declaracionContext))

	return p
}

func (s *DeclararSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclararSliceContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *DeclararSliceContext) OP_DECLARACION() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DECLARACION, 0)
}

func (s *DeclararSliceContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *DeclararSliceContext) RW_MUT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_MUT, 0)
}

func (s *DeclararSliceContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclararSliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterDeclararSlice(s)
	}
}

func (s *DeclararSliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitDeclararSlice(s)
	}
}

func (s *DeclararSliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitDeclararSlice(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclararInferenciaContext struct {
	Stmt_declaracionContext
}

func NewDeclararInferenciaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclararInferenciaContext {
	var p = new(DeclararInferenciaContext)

	InitEmptyStmt_declaracionContext(&p.Stmt_declaracionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Stmt_declaracionContext))

	return p
}

func (s *DeclararInferenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclararInferenciaContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *DeclararInferenciaContext) OP_ASSIGN() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_ASSIGN, 0)
}

func (s *DeclararInferenciaContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclararInferenciaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterDeclararInferencia(s)
	}
}

func (s *DeclararInferenciaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitDeclararInferencia(s)
	}
}

func (s *DeclararInferenciaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitDeclararInferencia(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclaraTipoValorContext struct {
	Stmt_declaracionContext
}

func NewDeclaraTipoValorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclaraTipoValorContext {
	var p = new(DeclaraTipoValorContext)

	InitEmptyStmt_declaracionContext(&p.Stmt_declaracionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Stmt_declaracionContext))

	return p
}

func (s *DeclaraTipoValorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaraTipoValorContext) RW_MUT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_MUT, 0)
}

func (s *DeclaraTipoValorContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *DeclaraTipoValorContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *DeclaraTipoValorContext) OP_DECLARACION() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DECLARACION, 0)
}

func (s *DeclaraTipoValorContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclaraTipoValorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterDeclaraTipoValor(s)
	}
}

func (s *DeclaraTipoValorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitDeclaraTipoValor(s)
	}
}

func (s *DeclaraTipoValorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitDeclaraTipoValor(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclararSinMutValorContext struct {
	Stmt_declaracionContext
}

func NewDeclararSinMutValorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclararSinMutValorContext {
	var p = new(DeclararSinMutValorContext)

	InitEmptyStmt_declaracionContext(&p.Stmt_declaracionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Stmt_declaracionContext))

	return p
}

func (s *DeclararSinMutValorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclararSinMutValorContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *DeclararSinMutValorContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *DeclararSinMutValorContext) OP_DECLARACION() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DECLARACION, 0)
}

func (s *DeclararSinMutValorContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclararSinMutValorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterDeclararSinMutValor(s)
	}
}

func (s *DeclararSinMutValorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitDeclararSinMutValor(s)
	}
}

func (s *DeclararSinMutValorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitDeclararSinMutValor(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclararInferenciaMutContext struct {
	Stmt_declaracionContext
}

func NewDeclararInferenciaMutContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclararInferenciaMutContext {
	var p = new(DeclararInferenciaMutContext)

	InitEmptyStmt_declaracionContext(&p.Stmt_declaracionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Stmt_declaracionContext))

	return p
}

func (s *DeclararInferenciaMutContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclararInferenciaMutContext) RW_MUT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_MUT, 0)
}

func (s *DeclararInferenciaMutContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *DeclararInferenciaMutContext) OP_ASSIGN() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_ASSIGN, 0)
}

func (s *DeclararInferenciaMutContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclararInferenciaMutContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterDeclararInferenciaMut(s)
	}
}

func (s *DeclararInferenciaMutContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitDeclararInferenciaMut(s)
	}
}

func (s *DeclararInferenciaMutContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitDeclararInferenciaMut(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclararTipoContext struct {
	Stmt_declaracionContext
}

func NewDeclararTipoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclararTipoContext {
	var p = new(DeclararTipoContext)

	InitEmptyStmt_declaracionContext(&p.Stmt_declaracionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Stmt_declaracionContext))

	return p
}

func (s *DeclararTipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclararTipoContext) RW_MUT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_MUT, 0)
}

func (s *DeclararTipoContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *DeclararTipoContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *DeclararTipoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterDeclararTipo(s)
	}
}

func (s *DeclararTipoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitDeclararTipo(s)
	}
}

func (s *DeclararTipoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitDeclararTipo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Stmt_declaracion() (localctx IStmt_declaracionContext) {
	localctx = NewStmt_declaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, VGrammarRULE_stmt_declaracion)
	var _la int

	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDeclararInferenciaMutContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)
			p.Match(VGrammarRW_MUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(110)
			p.Match(VGrammarOP_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.expr(0)
		}

	case 2:
		localctx = NewDeclararInferenciaContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)
			p.Match(VGrammarOP_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)
			p.expr(0)
		}

	case 3:
		localctx = NewDeclaraTipoValorContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(115)
			p.Match(VGrammarRW_MUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(116)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)
			p.Tipo()
		}
		{
			p.SetState(118)
			p.Match(VGrammarOP_DECLARACION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(119)
			p.expr(0)
		}

	case 4:
		localctx = NewDeclararTipoContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(121)
			p.Match(VGrammarRW_MUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)
			p.Tipo()
		}

	case 5:
		localctx = NewDeclararSinMutValorContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(124)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.Tipo()
		}
		{
			p.SetState(126)
			p.Match(VGrammarOP_DECLARACION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)
			p.expr(0)
		}

	case 6:
		localctx = NewDeclararSliceContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == VGrammarRW_MUT {
			{
				p.SetState(129)
				p.Match(VGrammarRW_MUT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(132)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.Match(VGrammarOP_DECLARACION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)
			p.Tipo()
		}
		p.SetState(136)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(135)
				p.expr(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmt_asignarContext is an interface to support dynamic dispatch.
type IStmt_asignarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStmt_asignarContext differentiates from other interfaces.
	IsStmt_asignarContext()
}

type Stmt_asignarContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmt_asignarContext() *Stmt_asignarContext {
	var p = new(Stmt_asignarContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_stmt_asignar
	return p
}

func InitEmptyStmt_asignarContext(p *Stmt_asignarContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_stmt_asignar
}

func (*Stmt_asignarContext) IsStmt_asignarContext() {}

func NewStmt_asignarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Stmt_asignarContext {
	var p = new(Stmt_asignarContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_stmt_asignar

	return p
}

func (s *Stmt_asignarContext) GetParser() antlr.Parser { return s.parser }

func (s *Stmt_asignarContext) CopyAll(ctx *Stmt_asignarContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Stmt_asignarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stmt_asignarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AsignacionSliceItemContext struct {
	Stmt_asignarContext
}

func NewAsignacionSliceItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignacionSliceItemContext {
	var p = new(AsignacionSliceItemContext)

	InitEmptyStmt_asignarContext(&p.Stmt_asignarContext)
	p.parser = parser
	p.CopyAll(ctx.(*Stmt_asignarContext))

	return p
}

func (s *AsignacionSliceItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionSliceItemContext) Item_slice() IItem_sliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IItem_sliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IItem_sliceContext)
}

func (s *AsignacionSliceItemContext) OP_DECLARACION() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DECLARACION, 0)
}

func (s *AsignacionSliceItemContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AsignacionSliceItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterAsignacionSliceItem(s)
	}
}

func (s *AsignacionSliceItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitAsignacionSliceItem(s)
	}
}

func (s *AsignacionSliceItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitAsignacionSliceItem(s)

	default:
		return t.VisitChildren(s)
	}
}

type AsignacionDirectaContext struct {
	Stmt_asignarContext
}

func NewAsignacionDirectaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignacionDirectaContext {
	var p = new(AsignacionDirectaContext)

	InitEmptyStmt_asignarContext(&p.Stmt_asignarContext)
	p.parser = parser
	p.CopyAll(ctx.(*Stmt_asignarContext))

	return p
}

func (s *AsignacionDirectaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionDirectaContext) PatronId() IPatronIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatronIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatronIdContext)
}

func (s *AsignacionDirectaContext) OP_DECLARACION() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DECLARACION, 0)
}

func (s *AsignacionDirectaContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AsignacionDirectaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterAsignacionDirecta(s)
	}
}

func (s *AsignacionDirectaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitAsignacionDirecta(s)
	}
}

func (s *AsignacionDirectaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitAsignacionDirecta(s)

	default:
		return t.VisitChildren(s)
	}
}

type AsignacionSliceContext struct {
	Stmt_asignarContext
	op antlr.Token
}

func NewAsignacionSliceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignacionSliceContext {
	var p = new(AsignacionSliceContext)

	InitEmptyStmt_asignarContext(&p.Stmt_asignarContext)
	p.parser = parser
	p.CopyAll(ctx.(*Stmt_asignarContext))

	return p
}

func (s *AsignacionSliceContext) GetOp() antlr.Token { return s.op }

func (s *AsignacionSliceContext) SetOp(v antlr.Token) { s.op = v }

func (s *AsignacionSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionSliceContext) Item_slice() IItem_sliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IItem_sliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IItem_sliceContext)
}

func (s *AsignacionSliceContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AsignacionSliceContext) OP_INCREMENTO() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_INCREMENTO, 0)
}

func (s *AsignacionSliceContext) OP_DECREMENTO() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DECREMENTO, 0)
}

func (s *AsignacionSliceContext) OP_DECLARACION() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DECLARACION, 0)
}

func (s *AsignacionSliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterAsignacionSlice(s)
	}
}

func (s *AsignacionSliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitAsignacionSlice(s)
	}
}

func (s *AsignacionSliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitAsignacionSlice(s)

	default:
		return t.VisitChildren(s)
	}
}

type AsignacionAritmeticaContext struct {
	Stmt_asignarContext
	op antlr.Token
}

func NewAsignacionAritmeticaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignacionAritmeticaContext {
	var p = new(AsignacionAritmeticaContext)

	InitEmptyStmt_asignarContext(&p.Stmt_asignarContext)
	p.parser = parser
	p.CopyAll(ctx.(*Stmt_asignarContext))

	return p
}

func (s *AsignacionAritmeticaContext) GetOp() antlr.Token { return s.op }

func (s *AsignacionAritmeticaContext) SetOp(v antlr.Token) { s.op = v }

func (s *AsignacionAritmeticaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionAritmeticaContext) PatronId() IPatronIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatronIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatronIdContext)
}

func (s *AsignacionAritmeticaContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AsignacionAritmeticaContext) OP_INCREMENTO() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_INCREMENTO, 0)
}

func (s *AsignacionAritmeticaContext) OP_DECREMENTO() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DECREMENTO, 0)
}

func (s *AsignacionAritmeticaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterAsignacionAritmetica(s)
	}
}

func (s *AsignacionAritmeticaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitAsignacionAritmetica(s)
	}
}

func (s *AsignacionAritmeticaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitAsignacionAritmetica(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Stmt_asignar() (localctx IStmt_asignarContext) {
	localctx = NewStmt_asignarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, VGrammarRULE_stmt_asignar)
	var _la int

	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAsignacionSliceItemContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(140)
			p.Item_slice()
		}
		{
			p.SetState(141)
			p.Match(VGrammarOP_DECLARACION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.expr(0)
		}

	case 2:
		localctx = NewAsignacionDirectaContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(144)
			p.PatronId()
		}
		{
			p.SetState(145)
			p.Match(VGrammarOP_DECLARACION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.expr(0)
		}

	case 3:
		localctx = NewAsignacionAritmeticaContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(148)
			p.PatronId()
		}
		{
			p.SetState(149)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AsignacionAritmeticaContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == VGrammarOP_INCREMENTO || _la == VGrammarOP_DECREMENTO) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AsignacionAritmeticaContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(150)
			p.expr(0)
		}

	case 4:
		localctx = NewAsignacionSliceContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(152)
			p.Item_slice()
		}
		{
			p.SetState(153)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AsignacionSliceContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&448) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AsignacionSliceContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(154)
			p.expr(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmt_transferenciaContext is an interface to support dynamic dispatch.
type IStmt_transferenciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStmt_transferenciaContext differentiates from other interfaces.
	IsStmt_transferenciaContext()
}

type Stmt_transferenciaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmt_transferenciaContext() *Stmt_transferenciaContext {
	var p = new(Stmt_transferenciaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_stmt_transferencia
	return p
}

func InitEmptyStmt_transferenciaContext(p *Stmt_transferenciaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_stmt_transferencia
}

func (*Stmt_transferenciaContext) IsStmt_transferenciaContext() {}

func NewStmt_transferenciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Stmt_transferenciaContext {
	var p = new(Stmt_transferenciaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_stmt_transferencia

	return p
}

func (s *Stmt_transferenciaContext) GetParser() antlr.Parser { return s.parser }

func (s *Stmt_transferenciaContext) CopyAll(ctx *Stmt_transferenciaContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Stmt_transferenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stmt_transferenciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ContinueStmtContext struct {
	Stmt_transferenciaContext
}

func NewContinueStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueStmtContext {
	var p = new(ContinueStmtContext)

	InitEmptyStmt_transferenciaContext(&p.Stmt_transferenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*Stmt_transferenciaContext))

	return p
}

func (s *ContinueStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStmtContext) RW_CONTINUE() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_CONTINUE, 0)
}

func (s *ContinueStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterContinueStmt(s)
	}
}

func (s *ContinueStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitContinueStmt(s)
	}
}

func (s *ContinueStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitContinueStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type BreakStmtContext struct {
	Stmt_transferenciaContext
}

func NewBreakStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakStmtContext {
	var p = new(BreakStmtContext)

	InitEmptyStmt_transferenciaContext(&p.Stmt_transferenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*Stmt_transferenciaContext))

	return p
}

func (s *BreakStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStmtContext) RW_BREAK() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_BREAK, 0)
}

func (s *BreakStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterBreakStmt(s)
	}
}

func (s *BreakStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitBreakStmt(s)
	}
}

func (s *BreakStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitBreakStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnStmtContext struct {
	Stmt_transferenciaContext
}

func NewReturnStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	InitEmptyStmt_transferenciaContext(&p.Stmt_transferenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*Stmt_transferenciaContext))

	return p
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) RW_RETURN() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_RETURN, 0)
}

func (s *ReturnStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterReturnStmt(s)
	}
}

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitReturnStmt(s)
	}
}

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Stmt_transferencia() (localctx IStmt_transferenciaContext) {
	localctx = NewStmt_transferenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, VGrammarRULE_stmt_transferencia)
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VGrammarRW_RETURN:
		localctx = NewReturnStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(158)
			p.Match(VGrammarRW_RETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(160)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(159)
				p.expr(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case VGrammarRW_BREAK:
		localctx = NewBreakStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(162)
			p.Match(VGrammarRW_BREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarRW_CONTINUE:
		localctx = NewContinueStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(163)
			p.Match(VGrammarRW_CONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIf_stmtContext is an interface to support dynamic dispatch.
type IIf_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIf_stmtContext differentiates from other interfaces.
	IsIf_stmtContext()
}

type If_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_stmtContext() *If_stmtContext {
	var p = new(If_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_if_stmt
	return p
}

func InitEmptyIf_stmtContext(p *If_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_if_stmt
}

func (*If_stmtContext) IsIf_stmtContext() {}

func NewIf_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_stmtContext {
	var p = new(If_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_if_stmt

	return p
}

func (s *If_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *If_stmtContext) CopyAll(ctx *If_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *If_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IFstmtContext struct {
	If_stmtContext
}

func NewIFstmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IFstmtContext {
	var p = new(IFstmtContext)

	InitEmptyIf_stmtContext(&p.If_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*If_stmtContext))

	return p
}

func (s *IFstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IFstmtContext) AllIf_chain() []IIf_chainContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIf_chainContext); ok {
			len++
		}
	}

	tst := make([]IIf_chainContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIf_chainContext); ok {
			tst[i] = t.(IIf_chainContext)
			i++
		}
	}

	return tst
}

func (s *IFstmtContext) If_chain(i int) IIf_chainContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_chainContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_chainContext)
}

func (s *IFstmtContext) AllRW_ELSE() []antlr.TerminalNode {
	return s.GetTokens(VGrammarRW_ELSE)
}

func (s *IFstmtContext) RW_ELSE(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarRW_ELSE, i)
}

func (s *IFstmtContext) Else_stmt() IElse_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElse_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElse_stmtContext)
}

func (s *IFstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterIFstmt(s)
	}
}

func (s *IFstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitIFstmt(s)
	}
}

func (s *IFstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitIFstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) If_stmt() (localctx IIf_stmtContext) {
	localctx = NewIf_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, VGrammarRULE_if_stmt)
	var _la int

	var _alt int

	localctx = NewIFstmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(166)
		p.If_chain()
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(167)
				p.Match(VGrammarRW_ELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(168)
				p.If_chain()
			}

		}
		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarRW_ELSE {
		{
			p.SetState(174)
			p.Else_stmt()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIf_chainContext is an interface to support dynamic dispatch.
type IIf_chainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIf_chainContext differentiates from other interfaces.
	IsIf_chainContext()
}

type If_chainContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_chainContext() *If_chainContext {
	var p = new(If_chainContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_if_chain
	return p
}

func InitEmptyIf_chainContext(p *If_chainContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_if_chain
}

func (*If_chainContext) IsIf_chainContext() {}

func NewIf_chainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_chainContext {
	var p = new(If_chainContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_if_chain

	return p
}

func (s *If_chainContext) GetParser() antlr.Parser { return s.parser }

func (s *If_chainContext) CopyAll(ctx *If_chainContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *If_chainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_chainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IFcadenaContext struct {
	If_chainContext
}

func NewIFcadenaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IFcadenaContext {
	var p = new(IFcadenaContext)

	InitEmptyIf_chainContext(&p.If_chainContext)
	p.parser = parser
	p.CopyAll(ctx.(*If_chainContext))

	return p
}

func (s *IFcadenaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IFcadenaContext) RW_IF() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_IF, 0)
}

func (s *IFcadenaContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IFcadenaContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *IFcadenaContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *IFcadenaContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *IFcadenaContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *IFcadenaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterIFcadena(s)
	}
}

func (s *IFcadenaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitIFcadena(s)
	}
}

func (s *IFcadenaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitIFcadena(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) If_chain() (localctx IIf_chainContext) {
	localctx = NewIf_chainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, VGrammarRULE_if_chain)
	var _la int

	localctx = NewIFcadenaContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)
		p.Match(VGrammarRW_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(178)
		p.expr(0)
	}
	{
		p.SetState(179)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&102452149878800) != 0 {
		{
			p.SetState(180)
			p.Stmt()
		}

		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(186)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElse_stmtContext is an interface to support dynamic dispatch.
type IElse_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsElse_stmtContext differentiates from other interfaces.
	IsElse_stmtContext()
}

type Else_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElse_stmtContext() *Else_stmtContext {
	var p = new(Else_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_else_stmt
	return p
}

func InitEmptyElse_stmtContext(p *Else_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_else_stmt
}

func (*Else_stmtContext) IsElse_stmtContext() {}

func NewElse_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Else_stmtContext {
	var p = new(Else_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_else_stmt

	return p
}

func (s *Else_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Else_stmtContext) CopyAll(ctx *Else_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Else_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Else_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ElseStmtContext struct {
	Else_stmtContext
}

func NewElseStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ElseStmtContext {
	var p = new(ElseStmtContext)

	InitEmptyElse_stmtContext(&p.Else_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Else_stmtContext))

	return p
}

func (s *ElseStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseStmtContext) RW_ELSE() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_ELSE, 0)
}

func (s *ElseStmtContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *ElseStmtContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *ElseStmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ElseStmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ElseStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterElseStmt(s)
	}
}

func (s *ElseStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitElseStmt(s)
	}
}

func (s *ElseStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitElseStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Else_stmt() (localctx IElse_stmtContext) {
	localctx = NewElse_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, VGrammarRULE_else_stmt)
	var _la int

	localctx = NewElseStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Match(VGrammarRW_ELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(189)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&102452149878800) != 0 {
		{
			p.SetState(190)
			p.Stmt()
		}

		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(196)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitch_stmtContext is an interface to support dynamic dispatch.
type ISwitch_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSwitch_stmtContext differentiates from other interfaces.
	IsSwitch_stmtContext()
}

type Switch_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitch_stmtContext() *Switch_stmtContext {
	var p = new(Switch_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_switch_stmt
	return p
}

func InitEmptySwitch_stmtContext(p *Switch_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_switch_stmt
}

func (*Switch_stmtContext) IsSwitch_stmtContext() {}

func NewSwitch_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_stmtContext {
	var p = new(Switch_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_switch_stmt

	return p
}

func (s *Switch_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Switch_stmtContext) CopyAll(ctx *Switch_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Switch_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SwitchStmtContext struct {
	Switch_stmtContext
}

func NewSwitchStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchStmtContext {
	var p = new(SwitchStmtContext)

	InitEmptySwitch_stmtContext(&p.Switch_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Switch_stmtContext))

	return p
}

func (s *SwitchStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStmtContext) RW_SWITCH() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_SWITCH, 0)
}

func (s *SwitchStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwitchStmtContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *SwitchStmtContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *SwitchStmtContext) AllSwitch_case() []ISwitch_caseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISwitch_caseContext); ok {
			len++
		}
	}

	tst := make([]ISwitch_caseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISwitch_caseContext); ok {
			tst[i] = t.(ISwitch_caseContext)
			i++
		}
	}

	return tst
}

func (s *SwitchStmtContext) Switch_case(i int) ISwitch_caseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitch_caseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitch_caseContext)
}

func (s *SwitchStmtContext) Default_case() IDefault_caseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefault_caseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefault_caseContext)
}

func (s *SwitchStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterSwitchStmt(s)
	}
}

func (s *SwitchStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitSwitchStmt(s)
	}
}

func (s *SwitchStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitSwitchStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Switch_stmt() (localctx ISwitch_stmtContext) {
	localctx = NewSwitch_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, VGrammarRULE_switch_stmt)
	var _la int

	localctx = NewSwitchStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		p.Match(VGrammarRW_SWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(199)
		p.expr(0)
	}
	{
		p.SetState(200)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VGrammarRW_CASE {
		{
			p.SetState(201)
			p.Switch_case()
		}

		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarRW_DEFAULT {
		{
			p.SetState(207)
			p.Default_case()
		}

	}
	{
		p.SetState(210)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitch_caseContext is an interface to support dynamic dispatch.
type ISwitch_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSwitch_caseContext differentiates from other interfaces.
	IsSwitch_caseContext()
}

type Switch_caseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitch_caseContext() *Switch_caseContext {
	var p = new(Switch_caseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_switch_case
	return p
}

func InitEmptySwitch_caseContext(p *Switch_caseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_switch_case
}

func (*Switch_caseContext) IsSwitch_caseContext() {}

func NewSwitch_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_caseContext {
	var p = new(Switch_caseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_switch_case

	return p
}

func (s *Switch_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *Switch_caseContext) CopyAll(ctx *Switch_caseContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Switch_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SwitchCaseContext struct {
	Switch_caseContext
}

func NewSwitchCaseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchCaseContext {
	var p = new(SwitchCaseContext)

	InitEmptySwitch_caseContext(&p.Switch_caseContext)
	p.parser = parser
	p.CopyAll(ctx.(*Switch_caseContext))

	return p
}

func (s *SwitchCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchCaseContext) RW_CASE() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_CASE, 0)
}

func (s *SwitchCaseContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwitchCaseContext) DOS_PUNTOS() antlr.TerminalNode {
	return s.GetToken(VGrammarDOS_PUNTOS, 0)
}

func (s *SwitchCaseContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *SwitchCaseContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *SwitchCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterSwitchCase(s)
	}
}

func (s *SwitchCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitSwitchCase(s)
	}
}

func (s *SwitchCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitSwitchCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Switch_case() (localctx ISwitch_caseContext) {
	localctx = NewSwitch_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, VGrammarRULE_switch_case)
	var _la int

	localctx = NewSwitchCaseContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		p.Match(VGrammarRW_CASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(213)
		p.expr(0)
	}
	{
		p.SetState(214)
		p.Match(VGrammarDOS_PUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&102452149878800) != 0 {
		{
			p.SetState(215)
			p.Stmt()
		}

		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefault_caseContext is an interface to support dynamic dispatch.
type IDefault_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDefault_caseContext differentiates from other interfaces.
	IsDefault_caseContext()
}

type Default_caseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefault_caseContext() *Default_caseContext {
	var p = new(Default_caseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_default_case
	return p
}

func InitEmptyDefault_caseContext(p *Default_caseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_default_case
}

func (*Default_caseContext) IsDefault_caseContext() {}

func NewDefault_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Default_caseContext {
	var p = new(Default_caseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_default_case

	return p
}

func (s *Default_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *Default_caseContext) CopyAll(ctx *Default_caseContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Default_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Default_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DefaultCaseContext struct {
	Default_caseContext
}

func NewDefaultCaseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DefaultCaseContext {
	var p = new(DefaultCaseContext)

	InitEmptyDefault_caseContext(&p.Default_caseContext)
	p.parser = parser
	p.CopyAll(ctx.(*Default_caseContext))

	return p
}

func (s *DefaultCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultCaseContext) RW_DEFAULT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_DEFAULT, 0)
}

func (s *DefaultCaseContext) DOS_PUNTOS() antlr.TerminalNode {
	return s.GetToken(VGrammarDOS_PUNTOS, 0)
}

func (s *DefaultCaseContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *DefaultCaseContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *DefaultCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterDefaultCase(s)
	}
}

func (s *DefaultCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitDefaultCase(s)
	}
}

func (s *DefaultCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitDefaultCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Default_case() (localctx IDefault_caseContext) {
	localctx = NewDefault_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, VGrammarRULE_default_case)
	var _la int

	localctx = NewDefaultCaseContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)
		p.Match(VGrammarRW_DEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(222)
		p.Match(VGrammarDOS_PUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&102452149878800) != 0 {
		{
			p.SetState(223)
			p.Stmt()
		}

		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhile_stmtContext is an interface to support dynamic dispatch.
type IWhile_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsWhile_stmtContext differentiates from other interfaces.
	IsWhile_stmtContext()
}

type While_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhile_stmtContext() *While_stmtContext {
	var p = new(While_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_while_stmt
	return p
}

func InitEmptyWhile_stmtContext(p *While_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_while_stmt
}

func (*While_stmtContext) IsWhile_stmtContext() {}

func NewWhile_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_stmtContext {
	var p = new(While_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_while_stmt

	return p
}

func (s *While_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *While_stmtContext) CopyAll(ctx *While_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *While_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type WhileStmtContext struct {
	While_stmtContext
}

func NewWhileStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileStmtContext {
	var p = new(WhileStmtContext)

	InitEmptyWhile_stmtContext(&p.While_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*While_stmtContext))

	return p
}

func (s *WhileStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStmtContext) RW_FOR() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_FOR, 0)
}

func (s *WhileStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhileStmtContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *WhileStmtContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *WhileStmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *WhileStmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *WhileStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterWhileStmt(s)
	}
}

func (s *WhileStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitWhileStmt(s)
	}
}

func (s *WhileStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitWhileStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) While_stmt() (localctx IWhile_stmtContext) {
	localctx = NewWhile_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, VGrammarRULE_while_stmt)
	var _la int

	localctx = NewWhileStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)
		p.Match(VGrammarRW_FOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(230)
		p.expr(0)
	}
	{
		p.SetState(231)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&102452149878800) != 0 {
		{
			p.SetState(232)
			p.Stmt()
		}

		p.SetState(237)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(238)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFor_stmtContext is an interface to support dynamic dispatch.
type IFor_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFor_stmtContext differentiates from other interfaces.
	IsFor_stmtContext()
}

type For_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_stmtContext() *For_stmtContext {
	var p = new(For_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_for_stmt
	return p
}

func InitEmptyFor_stmtContext(p *For_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_for_stmt
}

func (*For_stmtContext) IsFor_stmtContext() {}

func NewFor_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_stmtContext {
	var p = new(For_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_for_stmt

	return p
}

func (s *For_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *For_stmtContext) CopyAll(ctx *For_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *For_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForStmtContext struct {
	For_stmtContext
}

func NewForStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForStmtContext {
	var p = new(ForStmtContext)

	InitEmptyFor_stmtContext(&p.For_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*For_stmtContext))

	return p
}

func (s *ForStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStmtContext) RW_FOR() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_FOR, 0)
}

func (s *ForStmtContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *ForStmtContext) RW_IN() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_IN, 0)
}

func (s *ForStmtContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *ForStmtContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *ForStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForStmtContext) Range_() IRangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangeContext)
}

func (s *ForStmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ForStmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ForStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterForStmt(s)
	}
}

func (s *ForStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitForStmt(s)
	}
}

func (s *ForStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitForStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) For_stmt() (localctx IFor_stmtContext) {
	localctx = NewFor_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, VGrammarRULE_for_stmt)
	var _la int

	localctx = NewForStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(240)
		p.Match(VGrammarRW_FOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(241)
		p.Match(VGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(242)
		p.Match(VGrammarRW_IN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(243)
			p.expr(0)
		}

	case 2:
		{
			p.SetState(244)
			p.Range_()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(247)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&102452149878800) != 0 {
		{
			p.SetState(248)
			p.Stmt()
		}

		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(254)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRangeContext is an interface to support dynamic dispatch.
type IRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRangeContext differentiates from other interfaces.
	IsRangeContext()
}

type RangeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeContext() *RangeContext {
	var p = new(RangeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_range
	return p
}

func InitEmptyRangeContext(p *RangeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_range
}

func (*RangeContext) IsRangeContext() {}

func NewRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeContext {
	var p = new(RangeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_range

	return p
}

func (s *RangeContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeContext) CopyAll(ctx *RangeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *RangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RangoNumContext struct {
	RangeContext
}

func NewRangoNumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RangoNumContext {
	var p = new(RangoNumContext)

	InitEmptyRangeContext(&p.RangeContext)
	p.parser = parser
	p.CopyAll(ctx.(*RangeContext))

	return p
}

func (s *RangoNumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangoNumContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *RangoNumContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RangoNumContext) AllPUNTO() []antlr.TerminalNode {
	return s.GetTokens(VGrammarPUNTO)
}

func (s *RangoNumContext) PUNTO(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarPUNTO, i)
}

func (s *RangoNumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterRangoNum(s)
	}
}

func (s *RangoNumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitRangoNum(s)
	}
}

func (s *RangoNumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitRangoNum(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Range_() (localctx IRangeContext) {
	localctx = NewRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, VGrammarRULE_range)
	localctx = NewRangoNumContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)
		p.expr(0)
	}
	{
		p.SetState(257)
		p.Match(VGrammarPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(258)
		p.Match(VGrammarPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(259)
		p.Match(VGrammarPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(260)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarar_funcionContext is an interface to support dynamic dispatch.
type IDeclarar_funcionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDeclarar_funcionContext differentiates from other interfaces.
	IsDeclarar_funcionContext()
}

type Declarar_funcionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarar_funcionContext() *Declarar_funcionContext {
	var p = new(Declarar_funcionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_declarar_funcion
	return p
}

func InitEmptyDeclarar_funcionContext(p *Declarar_funcionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_declarar_funcion
}

func (*Declarar_funcionContext) IsDeclarar_funcionContext() {}

func NewDeclarar_funcionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declarar_funcionContext {
	var p = new(Declarar_funcionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_declarar_funcion

	return p
}

func (s *Declarar_funcionContext) GetParser() antlr.Parser { return s.parser }

func (s *Declarar_funcionContext) CopyAll(ctx *Declarar_funcionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Declarar_funcionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declarar_funcionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncionDecleradaContext struct {
	Declarar_funcionContext
}

func NewFuncionDecleradaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncionDecleradaContext {
	var p = new(FuncionDecleradaContext)

	InitEmptyDeclarar_funcionContext(&p.Declarar_funcionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Declarar_funcionContext))

	return p
}

func (s *FuncionDecleradaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionDecleradaContext) RW_FN() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_FN, 0)
}

func (s *FuncionDecleradaContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *FuncionDecleradaContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_IZQ, 0)
}

func (s *FuncionDecleradaContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_DER, 0)
}

func (s *FuncionDecleradaContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *FuncionDecleradaContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *FuncionDecleradaContext) Lista_parametros() ILista_parametrosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_parametrosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_parametrosContext)
}

func (s *FuncionDecleradaContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *FuncionDecleradaContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *FuncionDecleradaContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *FuncionDecleradaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterFuncionDeclerada(s)
	}
}

func (s *FuncionDecleradaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitFuncionDeclerada(s)
	}
}

func (s *FuncionDecleradaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitFuncionDeclerada(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Declarar_funcion() (localctx IDeclarar_funcionContext) {
	localctx = NewDeclarar_funcionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, VGrammarRULE_declarar_funcion)
	var _la int

	localctx = NewFuncionDecleradaContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(262)
		p.Match(VGrammarRW_FN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(263)
		p.Match(VGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(264)
		p.Match(VGrammarPAR_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarID {
		{
			p.SetState(265)
			p.Lista_parametros()
		}

	}
	{
		p.SetState(268)
		p.Match(VGrammarPAR_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2322168557870592) != 0 {
		{
			p.SetState(269)
			p.Tipo()
		}

	}
	{
		p.SetState(272)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(276)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&102452149878800) != 0 {
		{
			p.SetState(273)
			p.Stmt()
		}

		p.SetState(278)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(279)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILlamar_funcionContext is an interface to support dynamic dispatch.
type ILlamar_funcionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLlamar_funcionContext differentiates from other interfaces.
	IsLlamar_funcionContext()
}

type Llamar_funcionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLlamar_funcionContext() *Llamar_funcionContext {
	var p = new(Llamar_funcionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_llamar_funcion
	return p
}

func InitEmptyLlamar_funcionContext(p *Llamar_funcionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_llamar_funcion
}

func (*Llamar_funcionContext) IsLlamar_funcionContext() {}

func NewLlamar_funcionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Llamar_funcionContext {
	var p = new(Llamar_funcionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_llamar_funcion

	return p
}

func (s *Llamar_funcionContext) GetParser() antlr.Parser { return s.parser }

func (s *Llamar_funcionContext) CopyAll(ctx *Llamar_funcionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Llamar_funcionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Llamar_funcionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LlamarFuncionContext struct {
	Llamar_funcionContext
}

func NewLlamarFuncionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LlamarFuncionContext {
	var p = new(LlamarFuncionContext)

	InitEmptyLlamar_funcionContext(&p.Llamar_funcionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Llamar_funcionContext))

	return p
}

func (s *LlamarFuncionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LlamarFuncionContext) PatronId() IPatronIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatronIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatronIdContext)
}

func (s *LlamarFuncionContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_IZQ, 0)
}

func (s *LlamarFuncionContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_DER, 0)
}

func (s *LlamarFuncionContext) Lista_argumentos() ILista_argumentosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_argumentosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_argumentosContext)
}

func (s *LlamarFuncionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterLlamarFuncion(s)
	}
}

func (s *LlamarFuncionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitLlamarFuncion(s)
	}
}

func (s *LlamarFuncionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitLlamarFuncion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Llamar_funcion() (localctx ILlamar_funcionContext) {
	localctx = NewLlamar_funcionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, VGrammarRULE_llamar_funcion)
	var _la int

	localctx = NewLlamarFuncionContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(281)
		p.PatronId()
	}
	{
		p.SetState(282)
		p.Match(VGrammarPAR_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&774058334216192) != 0 {
		{
			p.SetState(283)
			p.Lista_argumentos()
		}

	}
	{
		p.SetState(286)
		p.Match(VGrammarPAR_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILista_parametrosContext is an interface to support dynamic dispatch.
type ILista_parametrosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLista_parametrosContext differentiates from other interfaces.
	IsLista_parametrosContext()
}

type Lista_parametrosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLista_parametrosContext() *Lista_parametrosContext {
	var p = new(Lista_parametrosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_lista_parametros
	return p
}

func InitEmptyLista_parametrosContext(p *Lista_parametrosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_lista_parametros
}

func (*Lista_parametrosContext) IsLista_parametrosContext() {}

func NewLista_parametrosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lista_parametrosContext {
	var p = new(Lista_parametrosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_lista_parametros

	return p
}

func (s *Lista_parametrosContext) GetParser() antlr.Parser { return s.parser }

func (s *Lista_parametrosContext) CopyAll(ctx *Lista_parametrosContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Lista_parametrosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lista_parametrosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ListaParametrosContext struct {
	Lista_parametrosContext
}

func NewListaParametrosContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListaParametrosContext {
	var p = new(ListaParametrosContext)

	InitEmptyLista_parametrosContext(&p.Lista_parametrosContext)
	p.parser = parser
	p.CopyAll(ctx.(*Lista_parametrosContext))

	return p
}

func (s *ListaParametrosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaParametrosContext) AllParametro_fun() []IParametro_funContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametro_funContext); ok {
			len++
		}
	}

	tst := make([]IParametro_funContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametro_funContext); ok {
			tst[i] = t.(IParametro_funContext)
			i++
		}
	}

	return tst
}

func (s *ListaParametrosContext) Parametro_fun(i int) IParametro_funContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametro_funContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametro_funContext)
}

func (s *ListaParametrosContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(VGrammarCOMA)
}

func (s *ListaParametrosContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarCOMA, i)
}

func (s *ListaParametrosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterListaParametros(s)
	}
}

func (s *ListaParametrosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitListaParametros(s)
	}
}

func (s *ListaParametrosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitListaParametros(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Lista_parametros() (localctx ILista_parametrosContext) {
	localctx = NewLista_parametrosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, VGrammarRULE_lista_parametros)
	var _la int

	var _alt int

	localctx = NewListaParametrosContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(288)
		p.Parametro_fun()
	}
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(289)
				p.Match(VGrammarCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(290)
				p.Parametro_fun()
			}

		}
		p.SetState(295)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarCOMA {
		{
			p.SetState(296)
			p.Match(VGrammarCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParametro_funContext is an interface to support dynamic dispatch.
type IParametro_funContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsParametro_funContext differentiates from other interfaces.
	IsParametro_funContext()
}

type Parametro_funContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametro_funContext() *Parametro_funContext {
	var p = new(Parametro_funContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_parametro_fun
	return p
}

func InitEmptyParametro_funContext(p *Parametro_funContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_parametro_fun
}

func (*Parametro_funContext) IsParametro_funContext() {}

func NewParametro_funContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Parametro_funContext {
	var p = new(Parametro_funContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_parametro_fun

	return p
}

func (s *Parametro_funContext) GetParser() antlr.Parser { return s.parser }

func (s *Parametro_funContext) CopyAll(ctx *Parametro_funContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Parametro_funContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Parametro_funContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParametroFunContext struct {
	Parametro_funContext
}

func NewParametroFunContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParametroFunContext {
	var p = new(ParametroFunContext)

	InitEmptyParametro_funContext(&p.Parametro_funContext)
	p.parser = parser
	p.CopyAll(ctx.(*Parametro_funContext))

	return p
}

func (s *ParametroFunContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametroFunContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(VGrammarID)
}

func (s *ParametroFunContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarID, i)
}

func (s *ParametroFunContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *ParametroFunContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterParametroFun(s)
	}
}

func (s *ParametroFunContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitParametroFun(s)
	}
}

func (s *ParametroFunContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitParametroFun(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Parametro_fun() (localctx IParametro_funContext) {
	localctx = NewParametro_funContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, VGrammarRULE_parametro_fun)
	localctx = NewParametroFunContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(300)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(299)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(302)
		p.Match(VGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(303)
		p.Tipo()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILista_argumentosContext is an interface to support dynamic dispatch.
type ILista_argumentosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLista_argumentosContext differentiates from other interfaces.
	IsLista_argumentosContext()
}

type Lista_argumentosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLista_argumentosContext() *Lista_argumentosContext {
	var p = new(Lista_argumentosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_lista_argumentos
	return p
}

func InitEmptyLista_argumentosContext(p *Lista_argumentosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_lista_argumentos
}

func (*Lista_argumentosContext) IsLista_argumentosContext() {}

func NewLista_argumentosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lista_argumentosContext {
	var p = new(Lista_argumentosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_lista_argumentos

	return p
}

func (s *Lista_argumentosContext) GetParser() antlr.Parser { return s.parser }

func (s *Lista_argumentosContext) CopyAll(ctx *Lista_argumentosContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Lista_argumentosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lista_argumentosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ListaArgumentosContext struct {
	Lista_argumentosContext
}

func NewListaArgumentosContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListaArgumentosContext {
	var p = new(ListaArgumentosContext)

	InitEmptyLista_argumentosContext(&p.Lista_argumentosContext)
	p.parser = parser
	p.CopyAll(ctx.(*Lista_argumentosContext))

	return p
}

func (s *ListaArgumentosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaArgumentosContext) AllArgumento_fun() []IArgumento_funContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgumento_funContext); ok {
			len++
		}
	}

	tst := make([]IArgumento_funContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgumento_funContext); ok {
			tst[i] = t.(IArgumento_funContext)
			i++
		}
	}

	return tst
}

func (s *ListaArgumentosContext) Argumento_fun(i int) IArgumento_funContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumento_funContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumento_funContext)
}

func (s *ListaArgumentosContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(VGrammarCOMA)
}

func (s *ListaArgumentosContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarCOMA, i)
}

func (s *ListaArgumentosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterListaArgumentos(s)
	}
}

func (s *ListaArgumentosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitListaArgumentos(s)
	}
}

func (s *ListaArgumentosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitListaArgumentos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Lista_argumentos() (localctx ILista_argumentosContext) {
	localctx = NewLista_argumentosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, VGrammarRULE_lista_argumentos)
	var _la int

	var _alt int

	localctx = NewListaArgumentosContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(305)
		p.Argumento_fun()
	}
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(306)
				p.Match(VGrammarCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(307)
				p.Argumento_fun()
			}

		}
		p.SetState(312)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarCOMA {
		{
			p.SetState(313)
			p.Match(VGrammarCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumento_funContext is an interface to support dynamic dispatch.
type IArgumento_funContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsArgumento_funContext differentiates from other interfaces.
	IsArgumento_funContext()
}

type Argumento_funContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumento_funContext() *Argumento_funContext {
	var p = new(Argumento_funContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_argumento_fun
	return p
}

func InitEmptyArgumento_funContext(p *Argumento_funContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_argumento_fun
}

func (*Argumento_funContext) IsArgumento_funContext() {}

func NewArgumento_funContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Argumento_funContext {
	var p = new(Argumento_funContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_argumento_fun

	return p
}

func (s *Argumento_funContext) GetParser() antlr.Parser { return s.parser }

func (s *Argumento_funContext) CopyAll(ctx *Argumento_funContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Argumento_funContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Argumento_funContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncionArgContext struct {
	Argumento_funContext
}

func NewFuncionArgContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncionArgContext {
	var p = new(FuncionArgContext)

	InitEmptyArgumento_funContext(&p.Argumento_funContext)
	p.parser = parser
	p.CopyAll(ctx.(*Argumento_funContext))

	return p
}

func (s *FuncionArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionArgContext) PatronId() IPatronIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatronIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatronIdContext)
}

func (s *FuncionArgContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FuncionArgContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *FuncionArgContext) DOS_PUNTOS() antlr.TerminalNode {
	return s.GetToken(VGrammarDOS_PUNTOS, 0)
}

func (s *FuncionArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterFuncionArg(s)
	}
}

func (s *FuncionArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitFuncionArg(s)
	}
}

func (s *FuncionArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitFuncionArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Argumento_fun() (localctx IArgumento_funContext) {
	localctx = NewArgumento_funContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, VGrammarRULE_argumento_fun)
	localctx = NewFuncionArgContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(318)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(316)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(317)
			p.Match(VGrammarDOS_PUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(322)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(320)
			p.PatronId()
		}

	case 2:
		{
			p.SetState(321)
			p.expr(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarar_structContext is an interface to support dynamic dispatch.
type IDeclarar_structContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDeclarar_structContext differentiates from other interfaces.
	IsDeclarar_structContext()
}

type Declarar_structContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarar_structContext() *Declarar_structContext {
	var p = new(Declarar_structContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_declarar_struct
	return p
}

func InitEmptyDeclarar_structContext(p *Declarar_structContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_declarar_struct
}

func (*Declarar_structContext) IsDeclarar_structContext() {}

func NewDeclarar_structContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declarar_structContext {
	var p = new(Declarar_structContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_declarar_struct

	return p
}

func (s *Declarar_structContext) GetParser() antlr.Parser { return s.parser }

func (s *Declarar_structContext) CopyAll(ctx *Declarar_structContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Declarar_structContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declarar_structContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DeclararStructContext struct {
	Declarar_structContext
}

func NewDeclararStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclararStructContext {
	var p = new(DeclararStructContext)

	InitEmptyDeclarar_structContext(&p.Declarar_structContext)
	p.parser = parser
	p.CopyAll(ctx.(*Declarar_structContext))

	return p
}

func (s *DeclararStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclararStructContext) RW_STRUCT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_STRUCT, 0)
}

func (s *DeclararStructContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *DeclararStructContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *DeclararStructContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *DeclararStructContext) AllPropiedad_struct() []IPropiedad_structContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPropiedad_structContext); ok {
			len++
		}
	}

	tst := make([]IPropiedad_structContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPropiedad_structContext); ok {
			tst[i] = t.(IPropiedad_structContext)
			i++
		}
	}

	return tst
}

func (s *DeclararStructContext) Propiedad_struct(i int) IPropiedad_structContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropiedad_structContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropiedad_structContext)
}

func (s *DeclararStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterDeclararStruct(s)
	}
}

func (s *DeclararStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitDeclararStruct(s)
	}
}

func (s *DeclararStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitDeclararStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Declarar_struct() (localctx IDeclarar_structContext) {
	localctx = NewDeclarar_structContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, VGrammarRULE_declarar_struct)
	var _la int

	localctx = NewDeclararStructContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(324)
		p.Match(VGrammarRW_STRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(325)
		p.Match(VGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(326)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(330)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2322168557870592) != 0 {
		{
			p.SetState(327)
			p.Propiedad_struct()
		}

		p.SetState(332)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(333)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPropiedad_structContext is an interface to support dynamic dispatch.
type IPropiedad_structContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPropiedad_structContext differentiates from other interfaces.
	IsPropiedad_structContext()
}

type Propiedad_structContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropiedad_structContext() *Propiedad_structContext {
	var p = new(Propiedad_structContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_propiedad_struct
	return p
}

func InitEmptyPropiedad_structContext(p *Propiedad_structContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_propiedad_struct
}

func (*Propiedad_structContext) IsPropiedad_structContext() {}

func NewPropiedad_structContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Propiedad_structContext {
	var p = new(Propiedad_structContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_propiedad_struct

	return p
}

func (s *Propiedad_structContext) GetParser() antlr.Parser { return s.parser }

func (s *Propiedad_structContext) CopyAll(ctx *Propiedad_structContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Propiedad_structContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Propiedad_structContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PropiedadStructContext struct {
	Propiedad_structContext
}

func NewPropiedadStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropiedadStructContext {
	var p = new(PropiedadStructContext)

	InitEmptyPropiedad_structContext(&p.Propiedad_structContext)
	p.parser = parser
	p.CopyAll(ctx.(*Propiedad_structContext))

	return p
}

func (s *PropiedadStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropiedadStructContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *PropiedadStructContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *PropiedadStructContext) OP_DECLARACION() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DECLARACION, 0)
}

func (s *PropiedadStructContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PropiedadStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterPropiedadStruct(s)
	}
}

func (s *PropiedadStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitPropiedadStruct(s)
	}
}

func (s *PropiedadStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitPropiedadStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Propiedad_struct() (localctx IPropiedad_structContext) {
	localctx = NewPropiedad_structContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, VGrammarRULE_propiedad_struct)
	var _la int

	localctx = NewPropiedadStructContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(335)
		p.Tipo()
	}
	{
		p.SetState(336)
		p.Match(VGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(339)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarOP_DECLARACION {
		{
			p.SetState(337)
			p.Match(VGrammarOP_DECLARACION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(338)
			p.expr(0)
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICrear_structContext is an interface to support dynamic dispatch.
type ICrear_structContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCrear_structContext differentiates from other interfaces.
	IsCrear_structContext()
}

type Crear_structContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCrear_structContext() *Crear_structContext {
	var p = new(Crear_structContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_crear_struct
	return p
}

func InitEmptyCrear_structContext(p *Crear_structContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_crear_struct
}

func (*Crear_structContext) IsCrear_structContext() {}

func NewCrear_structContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Crear_structContext {
	var p = new(Crear_structContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_crear_struct

	return p
}

func (s *Crear_structContext) GetParser() antlr.Parser { return s.parser }

func (s *Crear_structContext) CopyAll(ctx *Crear_structContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Crear_structContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Crear_structContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CrearStructContext struct {
	Crear_structContext
}

func NewCrearStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CrearStructContext {
	var p = new(CrearStructContext)

	InitEmptyCrear_structContext(&p.Crear_structContext)
	p.parser = parser
	p.CopyAll(ctx.(*Crear_structContext))

	return p
}

func (s *CrearStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CrearStructContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *CrearStructContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *CrearStructContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *CrearStructContext) Lista_parametros_init() ILista_parametros_initContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_parametros_initContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_parametros_initContext)
}

func (s *CrearStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterCrearStruct(s)
	}
}

func (s *CrearStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitCrearStruct(s)
	}
}

func (s *CrearStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitCrearStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Crear_struct() (localctx ICrear_structContext) {
	localctx = NewCrear_structContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, VGrammarRULE_crear_struct)
	var _la int

	localctx = NewCrearStructContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(341)
		p.Match(VGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(342)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarID {
		{
			p.SetState(343)
			p.Lista_parametros_init()
		}

	}
	{
		p.SetState(346)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILista_parametros_initContext is an interface to support dynamic dispatch.
type ILista_parametros_initContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLista_parametros_initContext differentiates from other interfaces.
	IsLista_parametros_initContext()
}

type Lista_parametros_initContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLista_parametros_initContext() *Lista_parametros_initContext {
	var p = new(Lista_parametros_initContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_lista_parametros_init
	return p
}

func InitEmptyLista_parametros_initContext(p *Lista_parametros_initContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_lista_parametros_init
}

func (*Lista_parametros_initContext) IsLista_parametros_initContext() {}

func NewLista_parametros_initContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lista_parametros_initContext {
	var p = new(Lista_parametros_initContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_lista_parametros_init

	return p
}

func (s *Lista_parametros_initContext) GetParser() antlr.Parser { return s.parser }

func (s *Lista_parametros_initContext) CopyAll(ctx *Lista_parametros_initContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Lista_parametros_initContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lista_parametros_initContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ListaParametrosInitContext struct {
	Lista_parametros_initContext
}

func NewListaParametrosInitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListaParametrosInitContext {
	var p = new(ListaParametrosInitContext)

	InitEmptyLista_parametros_initContext(&p.Lista_parametros_initContext)
	p.parser = parser
	p.CopyAll(ctx.(*Lista_parametros_initContext))

	return p
}

func (s *ListaParametrosInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaParametrosInitContext) AllParametros_init_struct() []IParametros_init_structContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametros_init_structContext); ok {
			len++
		}
	}

	tst := make([]IParametros_init_structContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametros_init_structContext); ok {
			tst[i] = t.(IParametros_init_structContext)
			i++
		}
	}

	return tst
}

func (s *ListaParametrosInitContext) Parametros_init_struct(i int) IParametros_init_structContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametros_init_structContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametros_init_structContext)
}

func (s *ListaParametrosInitContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(VGrammarCOMA)
}

func (s *ListaParametrosInitContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarCOMA, i)
}

func (s *ListaParametrosInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterListaParametrosInit(s)
	}
}

func (s *ListaParametrosInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitListaParametrosInit(s)
	}
}

func (s *ListaParametrosInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitListaParametrosInit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Lista_parametros_init() (localctx ILista_parametros_initContext) {
	localctx = NewLista_parametros_initContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, VGrammarRULE_lista_parametros_init)
	var _la int

	var _alt int

	localctx = NewListaParametrosInitContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(348)
		p.Parametros_init_struct()
	}
	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(349)
				p.Match(VGrammarCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(350)
				p.Parametros_init_struct()
			}

		}
		p.SetState(355)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(357)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarCOMA {
		{
			p.SetState(356)
			p.Match(VGrammarCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParametros_init_structContext is an interface to support dynamic dispatch.
type IParametros_init_structContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsParametros_init_structContext differentiates from other interfaces.
	IsParametros_init_structContext()
}

type Parametros_init_structContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametros_init_structContext() *Parametros_init_structContext {
	var p = new(Parametros_init_structContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_parametros_init_struct
	return p
}

func InitEmptyParametros_init_structContext(p *Parametros_init_structContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_parametros_init_struct
}

func (*Parametros_init_structContext) IsParametros_init_structContext() {}

func NewParametros_init_structContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Parametros_init_structContext {
	var p = new(Parametros_init_structContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_parametros_init_struct

	return p
}

func (s *Parametros_init_structContext) GetParser() antlr.Parser { return s.parser }

func (s *Parametros_init_structContext) CopyAll(ctx *Parametros_init_structContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Parametros_init_structContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Parametros_init_structContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParametrosInitStructContext struct {
	Parametros_init_structContext
}

func NewParametrosInitStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParametrosInitStructContext {
	var p = new(ParametrosInitStructContext)

	InitEmptyParametros_init_structContext(&p.Parametros_init_structContext)
	p.parser = parser
	p.CopyAll(ctx.(*Parametros_init_structContext))

	return p
}

func (s *ParametrosInitStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametrosInitStructContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *ParametrosInitStructContext) DOS_PUNTOS() antlr.TerminalNode {
	return s.GetToken(VGrammarDOS_PUNTOS, 0)
}

func (s *ParametrosInitStructContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParametrosInitStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterParametrosInitStruct(s)
	}
}

func (s *ParametrosInitStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitParametrosInitStruct(s)
	}
}

func (s *ParametrosInitStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitParametrosInitStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Parametros_init_struct() (localctx IParametros_init_structContext) {
	localctx = NewParametros_init_structContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, VGrammarRULE_parametros_init_struct)
	localctx = NewParametrosInitStructContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(359)
		p.Match(VGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(360)
		p.Match(VGrammarDOS_PUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(361)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILista_sliceContext is an interface to support dynamic dispatch.
type ILista_sliceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLista_sliceContext differentiates from other interfaces.
	IsLista_sliceContext()
}

type Lista_sliceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLista_sliceContext() *Lista_sliceContext {
	var p = new(Lista_sliceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_lista_slice
	return p
}

func InitEmptyLista_sliceContext(p *Lista_sliceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_lista_slice
}

func (*Lista_sliceContext) IsLista_sliceContext() {}

func NewLista_sliceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lista_sliceContext {
	var p = new(Lista_sliceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_lista_slice

	return p
}

func (s *Lista_sliceContext) GetParser() antlr.Parser { return s.parser }

func (s *Lista_sliceContext) CopyAll(ctx *Lista_sliceContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Lista_sliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lista_sliceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ListaSliceContext struct {
	Lista_sliceContext
}

func NewListaSliceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListaSliceContext {
	var p = new(ListaSliceContext)

	InitEmptyLista_sliceContext(&p.Lista_sliceContext)
	p.parser = parser
	p.CopyAll(ctx.(*Lista_sliceContext))

	return p
}

func (s *ListaSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaSliceContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *ListaSliceContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *ListaSliceContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ListaSliceContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListaSliceContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(VGrammarCOMA)
}

func (s *ListaSliceContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarCOMA, i)
}

func (s *ListaSliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterListaSlice(s)
	}
}

func (s *ListaSliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitListaSlice(s)
	}
}

func (s *ListaSliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitListaSlice(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Lista_slice() (localctx ILista_sliceContext) {
	localctx = NewLista_sliceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, VGrammarRULE_lista_slice)
	var _la int

	var _alt int

	localctx = NewListaSliceContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(363)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&774058334216192) != 0 {
		{
			p.SetState(364)
			p.expr(0)
		}
		p.SetState(369)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(365)
					p.Match(VGrammarCOMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(366)
					p.expr(0)
				}

			}
			p.SetState(371)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	}
	p.SetState(375)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarCOMA {
		{
			p.SetState(374)
			p.Match(VGrammarCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(377)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IItem_sliceContext is an interface to support dynamic dispatch.
type IItem_sliceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsItem_sliceContext differentiates from other interfaces.
	IsItem_sliceContext()
}

type Item_sliceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyItem_sliceContext() *Item_sliceContext {
	var p = new(Item_sliceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_item_slice
	return p
}

func InitEmptyItem_sliceContext(p *Item_sliceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_item_slice
}

func (*Item_sliceContext) IsItem_sliceContext() {}

func NewItem_sliceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Item_sliceContext {
	var p = new(Item_sliceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_item_slice

	return p
}

func (s *Item_sliceContext) GetParser() antlr.Parser { return s.parser }

func (s *Item_sliceContext) CopyAll(ctx *Item_sliceContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Item_sliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Item_sliceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ItemSliceContext struct {
	Item_sliceContext
}

func NewItemSliceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ItemSliceContext {
	var p = new(ItemSliceContext)

	InitEmptyItem_sliceContext(&p.Item_sliceContext)
	p.parser = parser
	p.CopyAll(ctx.(*Item_sliceContext))

	return p
}

func (s *ItemSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ItemSliceContext) PatronId() IPatronIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatronIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatronIdContext)
}

func (s *ItemSliceContext) AllCORCHETE_IZQ() []antlr.TerminalNode {
	return s.GetTokens(VGrammarCORCHETE_IZQ)
}

func (s *ItemSliceContext) CORCHETE_IZQ(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarCORCHETE_IZQ, i)
}

func (s *ItemSliceContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ItemSliceContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ItemSliceContext) AllCORCHETE_DER() []antlr.TerminalNode {
	return s.GetTokens(VGrammarCORCHETE_DER)
}

func (s *ItemSliceContext) CORCHETE_DER(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarCORCHETE_DER, i)
}

func (s *ItemSliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterItemSlice(s)
	}
}

func (s *ItemSliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitItemSlice(s)
	}
}

func (s *ItemSliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitItemSlice(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Item_slice() (localctx IItem_sliceContext) {
	localctx = NewItem_sliceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, VGrammarRULE_item_slice)
	var _alt int

	localctx = NewItemSliceContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(379)
		p.PatronId()
	}
	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(380)
				p.Match(VGrammarCORCHETE_IZQ)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(381)
				p.expr(0)
			}
			{
				p.SetState(382)
				p.Match(VGrammarCORCHETE_DER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(386)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProp_sliceContext is an interface to support dynamic dispatch.
type IProp_sliceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsProp_sliceContext differentiates from other interfaces.
	IsProp_sliceContext()
}

type Prop_sliceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProp_sliceContext() *Prop_sliceContext {
	var p = new(Prop_sliceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_prop_slice
	return p
}

func InitEmptyProp_sliceContext(p *Prop_sliceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_prop_slice
}

func (*Prop_sliceContext) IsProp_sliceContext() {}

func NewProp_sliceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Prop_sliceContext {
	var p = new(Prop_sliceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_prop_slice

	return p
}

func (s *Prop_sliceContext) GetParser() antlr.Parser { return s.parser }

func (s *Prop_sliceContext) CopyAll(ctx *Prop_sliceContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Prop_sliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Prop_sliceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PropSliceContext struct {
	Prop_sliceContext
}

func NewPropSliceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropSliceContext {
	var p = new(PropSliceContext)

	InitEmptyProp_sliceContext(&p.Prop_sliceContext)
	p.parser = parser
	p.CopyAll(ctx.(*Prop_sliceContext))

	return p
}

func (s *PropSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropSliceContext) Item_slice() IItem_sliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IItem_sliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IItem_sliceContext)
}

func (s *PropSliceContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(VGrammarPUNTO, 0)
}

func (s *PropSliceContext) PatronId() IPatronIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatronIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatronIdContext)
}

func (s *PropSliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterPropSlice(s)
	}
}

func (s *PropSliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitPropSlice(s)
	}
}

func (s *PropSliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitPropSlice(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Prop_slice() (localctx IProp_sliceContext) {
	localctx = NewProp_sliceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, VGrammarRULE_prop_slice)
	localctx = NewPropSliceContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(388)
		p.Item_slice()
	}
	{
		p.SetState(389)
		p.Match(VGrammarPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(390)
		p.PatronId()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFun_sliceContext is an interface to support dynamic dispatch.
type IFun_sliceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFun_sliceContext differentiates from other interfaces.
	IsFun_sliceContext()
}

type Fun_sliceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFun_sliceContext() *Fun_sliceContext {
	var p = new(Fun_sliceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_fun_slice
	return p
}

func InitEmptyFun_sliceContext(p *Fun_sliceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_fun_slice
}

func (*Fun_sliceContext) IsFun_sliceContext() {}

func NewFun_sliceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fun_sliceContext {
	var p = new(Fun_sliceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_fun_slice

	return p
}

func (s *Fun_sliceContext) GetParser() antlr.Parser { return s.parser }

func (s *Fun_sliceContext) CopyAll(ctx *Fun_sliceContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Fun_sliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fun_sliceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncionSliceContext struct {
	Fun_sliceContext
}

func NewFuncionSliceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncionSliceContext {
	var p = new(FuncionSliceContext)

	InitEmptyFun_sliceContext(&p.Fun_sliceContext)
	p.parser = parser
	p.CopyAll(ctx.(*Fun_sliceContext))

	return p
}

func (s *FuncionSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionSliceContext) Item_slice() IItem_sliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IItem_sliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IItem_sliceContext)
}

func (s *FuncionSliceContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(VGrammarPUNTO, 0)
}

func (s *FuncionSliceContext) Llamar_funcion() ILlamar_funcionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamar_funcionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamar_funcionContext)
}

func (s *FuncionSliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterFuncionSlice(s)
	}
}

func (s *FuncionSliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitFuncionSlice(s)
	}
}

func (s *FuncionSliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitFuncionSlice(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Fun_slice() (localctx IFun_sliceContext) {
	localctx = NewFun_sliceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, VGrammarRULE_fun_slice)
	localctx = NewFuncionSliceContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(392)
		p.Item_slice()
	}
	{
		p.SetState(393)
		p.Match(VGrammarPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(394)
		p.Llamar_funcion()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITipoContext is an interface to support dynamic dispatch.
type ITipoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	RW_INT() antlr.TerminalNode
	RW_FLOAT64() antlr.TerminalNode
	RW_STRING() antlr.TerminalNode
	RW_BOOL() antlr.TerminalNode
	Tipos_slices() ITipos_slicesContext

	// IsTipoContext differentiates from other interfaces.
	IsTipoContext()
}

type TipoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipoContext() *TipoContext {
	var p = new(TipoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_tipo
	return p
}

func InitEmptyTipoContext(p *TipoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_tipo
}

func (*TipoContext) IsTipoContext() {}

func NewTipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipoContext {
	var p = new(TipoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_tipo

	return p
}

func (s *TipoContext) GetParser() antlr.Parser { return s.parser }

func (s *TipoContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *TipoContext) RW_INT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_INT, 0)
}

func (s *TipoContext) RW_FLOAT64() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_FLOAT64, 0)
}

func (s *TipoContext) RW_STRING() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_STRING, 0)
}

func (s *TipoContext) RW_BOOL() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_BOOL, 0)
}

func (s *TipoContext) Tipos_slices() ITipos_slicesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipos_slicesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipos_slicesContext)
}

func (s *TipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterTipo(s)
	}
}

func (s *TipoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitTipo(s)
	}
}

func (s *TipoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitTipo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Tipo() (localctx ITipoContext) {
	localctx = NewTipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, VGrammarRULE_tipo)
	p.SetState(402)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VGrammarID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(396)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarRW_INT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(397)
			p.Match(VGrammarRW_INT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarRW_FLOAT64:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(398)
			p.Match(VGrammarRW_FLOAT64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarRW_STRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(399)
			p.Match(VGrammarRW_STRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarRW_BOOL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(400)
			p.Match(VGrammarRW_BOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarCORCHETE_IZQ:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(401)
			p.Tipos_slices()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITipos_slicesContext is an interface to support dynamic dispatch.
type ITipos_slicesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTipos_slicesContext differentiates from other interfaces.
	IsTipos_slicesContext()
}

type Tipos_slicesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipos_slicesContext() *Tipos_slicesContext {
	var p = new(Tipos_slicesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_tipos_slices
	return p
}

func InitEmptyTipos_slicesContext(p *Tipos_slicesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_tipos_slices
}

func (*Tipos_slicesContext) IsTipos_slicesContext() {}

func NewTipos_slicesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tipos_slicesContext {
	var p = new(Tipos_slicesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_tipos_slices

	return p
}

func (s *Tipos_slicesContext) GetParser() antlr.Parser { return s.parser }

func (s *Tipos_slicesContext) CopyAll(ctx *Tipos_slicesContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Tipos_slicesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipos_slicesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VectorSimpleContext struct {
	Tipos_slicesContext
}

func NewVectorSimpleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorSimpleContext {
	var p = new(VectorSimpleContext)

	InitEmptyTipos_slicesContext(&p.Tipos_slicesContext)
	p.parser = parser
	p.CopyAll(ctx.(*Tipos_slicesContext))

	return p
}

func (s *VectorSimpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorSimpleContext) CORCHETE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarCORCHETE_IZQ, 0)
}

func (s *VectorSimpleContext) CORCHETE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarCORCHETE_DER, 0)
}

func (s *VectorSimpleContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *VectorSimpleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterVectorSimple(s)
	}
}

func (s *VectorSimpleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitVectorSimple(s)
	}
}

func (s *VectorSimpleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitVectorSimple(s)

	default:
		return t.VisitChildren(s)
	}
}

type MatrizDobleContext struct {
	Tipos_slicesContext
}

func NewMatrizDobleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MatrizDobleContext {
	var p = new(MatrizDobleContext)

	InitEmptyTipos_slicesContext(&p.Tipos_slicesContext)
	p.parser = parser
	p.CopyAll(ctx.(*Tipos_slicesContext))

	return p
}

func (s *MatrizDobleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatrizDobleContext) AllCORCHETE_IZQ() []antlr.TerminalNode {
	return s.GetTokens(VGrammarCORCHETE_IZQ)
}

func (s *MatrizDobleContext) CORCHETE_IZQ(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarCORCHETE_IZQ, i)
}

func (s *MatrizDobleContext) AllCORCHETE_DER() []antlr.TerminalNode {
	return s.GetTokens(VGrammarCORCHETE_DER)
}

func (s *MatrizDobleContext) CORCHETE_DER(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarCORCHETE_DER, i)
}

func (s *MatrizDobleContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *MatrizDobleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterMatrizDoble(s)
	}
}

func (s *MatrizDobleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitMatrizDoble(s)
	}
}

func (s *MatrizDobleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitMatrizDoble(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Tipos_slices() (localctx ITipos_slicesContext) {
	localctx = NewTipos_slicesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, VGrammarRULE_tipos_slices)
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVectorSimpleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(404)
			p.Match(VGrammarCORCHETE_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(405)
			p.Match(VGrammarCORCHETE_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(406)
			p.Tipo()
		}

	case 2:
		localctx = NewMatrizDobleContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(407)
			p.Match(VGrammarCORCHETE_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(408)
			p.Match(VGrammarCORCHETE_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(409)
			p.Match(VGrammarCORCHETE_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(410)
			p.Match(VGrammarCORCHETE_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(411)
			p.Tipo()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LiteralExpContext struct {
	ExprContext
}

func NewLiteralExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExpContext {
	var p = new(LiteralExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LiteralExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExpContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterLiteralExp(s)
	}
}

func (s *LiteralExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitLiteralExp(s)
	}
}

func (s *LiteralExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitLiteralExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdExpContext struct {
	ExprContext
}

func NewIdExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdExpContext {
	var p = new(IdExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IdExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdExpContext) PatronId() IPatronIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatronIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatronIdContext)
}

func (s *IdExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterIdExp(s)
	}
}

func (s *IdExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitIdExp(s)
	}
}

func (s *IdExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitIdExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type CrearStructExpContext struct {
	ExprContext
}

func NewCrearStructExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CrearStructExpContext {
	var p = new(CrearStructExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CrearStructExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CrearStructExpContext) Crear_struct() ICrear_structContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICrear_structContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICrear_structContext)
}

func (s *CrearStructExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterCrearStructExp(s)
	}
}

func (s *CrearStructExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitCrearStructExp(s)
	}
}

func (s *CrearStructExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitCrearStructExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnarioExpContext struct {
	ExprContext
	op antlr.Token
}

func NewUnarioExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnarioExpContext {
	var p = new(UnarioExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *UnarioExpContext) GetOp() antlr.Token { return s.op }

func (s *UnarioExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnarioExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnarioExpContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UnarioExpContext) OP_NOT() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_NOT, 0)
}

func (s *UnarioExpContext) OP_RESTA() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_RESTA, 0)
}

func (s *UnarioExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterUnarioExp(s)
	}
}

func (s *UnarioExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitUnarioExp(s)
	}
}

func (s *UnarioExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitUnarioExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ItemSliceExpContext struct {
	ExprContext
}

func NewItemSliceExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ItemSliceExpContext {
	var p = new(ItemSliceExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ItemSliceExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ItemSliceExpContext) Item_slice() IItem_sliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IItem_sliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IItem_sliceContext)
}

func (s *ItemSliceExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterItemSliceExp(s)
	}
}

func (s *ItemSliceExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitItemSliceExp(s)
	}
}

func (s *ItemSliceExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitItemSliceExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type VectorFuncExpContext struct {
	ExprContext
}

func NewVectorFuncExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorFuncExpContext {
	var p = new(VectorFuncExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *VectorFuncExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorFuncExpContext) Fun_slice() IFun_sliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFun_sliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFun_sliceContext)
}

func (s *VectorFuncExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterVectorFuncExp(s)
	}
}

func (s *VectorFuncExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitVectorFuncExp(s)
	}
}

func (s *VectorFuncExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitVectorFuncExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type SliceExpContext struct {
	ExprContext
}

func NewSliceExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceExpContext {
	var p = new(SliceExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *SliceExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceExpContext) Lista_slice() ILista_sliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_sliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_sliceContext)
}

func (s *SliceExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterSliceExp(s)
	}
}

func (s *SliceExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitSliceExp(s)
	}
}

func (s *SliceExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitSliceExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type PropSliceExpContext struct {
	ExprContext
}

func NewPropSliceExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropSliceExpContext {
	var p = new(PropSliceExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *PropSliceExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropSliceExpContext) Prop_slice() IProp_sliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProp_sliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProp_sliceContext)
}

func (s *PropSliceExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterPropSliceExp(s)
	}
}

func (s *PropSliceExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitPropSliceExp(s)
	}
}

func (s *PropSliceExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitPropSliceExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type LlamarFuncionExpContext struct {
	ExprContext
}

func NewLlamarFuncionExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LlamarFuncionExpContext {
	var p = new(LlamarFuncionExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LlamarFuncionExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LlamarFuncionExpContext) Llamar_funcion() ILlamar_funcionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamar_funcionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamar_funcionContext)
}

func (s *LlamarFuncionExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterLlamarFuncionExp(s)
	}
}

func (s *LlamarFuncionExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitLlamarFuncionExp(s)
	}
}

func (s *LlamarFuncionExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitLlamarFuncionExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type BinarioExpContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewBinarioExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinarioExpContext {
	var p = new(BinarioExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BinarioExpContext) GetOp() antlr.Token { return s.op }

func (s *BinarioExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *BinarioExpContext) GetLeft() IExprContext { return s.left }

func (s *BinarioExpContext) GetRight() IExprContext { return s.right }

func (s *BinarioExpContext) SetLeft(v IExprContext) { s.left = v }

func (s *BinarioExpContext) SetRight(v IExprContext) { s.right = v }

func (s *BinarioExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinarioExpContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *BinarioExpContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BinarioExpContext) OP_MULT() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_MULT, 0)
}

func (s *BinarioExpContext) OP_DIV() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DIV, 0)
}

func (s *BinarioExpContext) OP_MOD() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_MOD, 0)
}

func (s *BinarioExpContext) OP_SUMA() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_SUMA, 0)
}

func (s *BinarioExpContext) OP_RESTA() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_RESTA, 0)
}

func (s *BinarioExpContext) OP_MENORQ() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_MENORQ, 0)
}

func (s *BinarioExpContext) OP_MENORIGUAL() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_MENORIGUAL, 0)
}

func (s *BinarioExpContext) OP_MAYORQ() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_MAYORQ, 0)
}

func (s *BinarioExpContext) OP_MAYORIGUAL() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_MAYORIGUAL, 0)
}

func (s *BinarioExpContext) OP_IGUAL() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_IGUAL, 0)
}

func (s *BinarioExpContext) OP_DIFERENTE() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DIFERENTE, 0)
}

func (s *BinarioExpContext) OP_AND() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_AND, 0)
}

func (s *BinarioExpContext) OP_OR() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_OR, 0)
}

func (s *BinarioExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterBinarioExp(s)
	}
}

func (s *BinarioExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitBinarioExp(s)
	}
}

func (s *BinarioExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitBinarioExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParentecisExpContext struct {
	ExprContext
}

func NewParentecisExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentecisExpContext {
	var p = new(ParentecisExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParentecisExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentecisExpContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_IZQ, 0)
}

func (s *ParentecisExpContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParentecisExpContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_DER, 0)
}

func (s *ParentecisExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterParentecisExp(s)
	}
}

func (s *ParentecisExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitParentecisExp(s)
	}
}

func (s *ParentecisExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitParentecisExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *VGrammar) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 64
	p.EnterRecursionRule(localctx, 64, VGrammarRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(429)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParentecisExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(415)
			p.Match(VGrammarPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(416)
			p.expr(0)
		}
		{
			p.SetState(417)
			p.Match(VGrammarPAR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewLlamarFuncionExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(419)
			p.Llamar_funcion()
		}

	case 3:
		localctx = NewIdExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(420)
			p.PatronId()
		}

	case 4:
		localctx = NewItemSliceExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(421)
			p.Item_slice()
		}

	case 5:
		localctx = NewPropSliceExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(422)
			p.Prop_slice()
		}

	case 6:
		localctx = NewVectorFuncExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(423)
			p.Fun_slice()
		}

	case 7:
		localctx = NewLiteralExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(424)
			p.Literal()
		}

	case 8:
		localctx = NewSliceExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(425)
			p.Lista_slice()
		}

	case 9:
		localctx = NewCrearStructExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(426)
			p.Crear_struct()
		}

	case 10:
		localctx = NewUnarioExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(427)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnarioExpContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == VGrammarOP_RESTA || _la == VGrammarOP_NOT) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnarioExpContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(428)
			p.expr(7)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(451)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(449)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinarioExpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinarioExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(431)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(432)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinarioExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7340032) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinarioExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(433)

					var _x = p.expr(7)

					localctx.(*BinarioExpContext).right = _x
				}

			case 2:
				localctx = NewBinarioExpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinarioExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(434)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(435)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinarioExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == VGrammarOP_SUMA || _la == VGrammarOP_RESTA) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinarioExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(436)

					var _x = p.expr(6)

					localctx.(*BinarioExpContext).right = _x
				}

			case 3:
				localctx = NewBinarioExpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinarioExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(437)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(438)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinarioExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&503316480) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinarioExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(439)

					var _x = p.expr(5)

					localctx.(*BinarioExpContext).right = _x
				}

			case 4:
				localctx = NewBinarioExpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinarioExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(440)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(441)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinarioExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == VGrammarOP_IGUAL || _la == VGrammarOP_DIFERENTE) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinarioExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(442)

					var _x = p.expr(4)

					localctx.(*BinarioExpContext).right = _x
				}

			case 5:
				localctx = NewBinarioExpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinarioExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(443)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(444)

					var _m = p.Match(VGrammarOP_AND)

					localctx.(*BinarioExpContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(445)

					var _x = p.expr(3)

					localctx.(*BinarioExpContext).right = _x
				}

			case 6:
				localctx = NewBinarioExpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinarioExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(446)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(447)

					var _m = p.Match(VGrammarOP_OR)

					localctx.(*BinarioExpContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(448)

					var _x = p.expr(2)

					localctx.(*BinarioExpContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(453)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyAll(ctx *LiteralContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StringLiteralContext struct {
	LiteralContext
}

func NewStringLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLiteralContext {
	var p = new(StringLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(VGrammarSTRING_LITERAL, 0)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (s *StringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitStringLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoolLiteralContext struct {
	LiteralContext
}

func NewBoolLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolLiteralContext {
	var p = new(BoolLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *BoolLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolLiteralContext) BOOL_LITERAL() antlr.TerminalNode {
	return s.GetToken(VGrammarBOOL_LITERAL, 0)
}

func (s *BoolLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterBoolLiteral(s)
	}
}

func (s *BoolLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitBoolLiteral(s)
	}
}

func (s *BoolLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitBoolLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatLiteralContext struct {
	LiteralContext
}

func NewFloatLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatLiteralContext {
	var p = new(FloatLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *FloatLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLiteralContext) FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(VGrammarFLOAT_LITERAL, 0)
}

func (s *FloatLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitFloatLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type NilLiteralContext struct {
	LiteralContext
}

func NewNilLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NilLiteralContext {
	var p = new(NilLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *NilLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilLiteralContext) NIL_LITERAL() antlr.TerminalNode {
	return s.GetToken(VGrammarNIL_LITERAL, 0)
}

func (s *NilLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterNilLiteral(s)
	}
}

func (s *NilLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitNilLiteral(s)
	}
}

func (s *NilLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitNilLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntLiteralContext struct {
	LiteralContext
}

func NewIntLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntLiteralContext {
	var p = new(IntLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *IntLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntLiteralContext) INT_LITERAL() antlr.TerminalNode {
	return s.GetToken(VGrammarINT_LITERAL, 0)
}

func (s *IntLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterIntLiteral(s)
	}
}

func (s *IntLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitIntLiteral(s)
	}
}

func (s *IntLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitIntLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, VGrammarRULE_literal)
	p.SetState(459)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VGrammarINT_LITERAL:
		localctx = NewIntLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(454)
			p.Match(VGrammarINT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarFLOAT_LITERAL:
		localctx = NewFloatLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(455)
			p.Match(VGrammarFLOAT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarSTRING_LITERAL:
		localctx = NewStringLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(456)
			p.Match(VGrammarSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarBOOL_LITERAL:
		localctx = NewBoolLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(457)
			p.Match(VGrammarBOOL_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarNIL_LITERAL:
		localctx = NewNilLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(458)
			p.Match(VGrammarNIL_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPatronIdContext is an interface to support dynamic dispatch.
type IPatronIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPatronIdContext differentiates from other interfaces.
	IsPatronIdContext()
}

type PatronIdContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPatronIdContext() *PatronIdContext {
	var p = new(PatronIdContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_patronId
	return p
}

func InitEmptyPatronIdContext(p *PatronIdContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_patronId
}

func (*PatronIdContext) IsPatronIdContext() {}

func NewPatronIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatronIdContext {
	var p = new(PatronIdContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_patronId

	return p
}

func (s *PatronIdContext) GetParser() antlr.Parser { return s.parser }

func (s *PatronIdContext) CopyAll(ctx *PatronIdContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PatronIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatronIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ID_PatronContext struct {
	PatronIdContext
}

func NewID_PatronContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ID_PatronContext {
	var p = new(ID_PatronContext)

	InitEmptyPatronIdContext(&p.PatronIdContext)
	p.parser = parser
	p.CopyAll(ctx.(*PatronIdContext))

	return p
}

func (s *ID_PatronContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ID_PatronContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(VGrammarID)
}

func (s *ID_PatronContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarID, i)
}

func (s *ID_PatronContext) AllPUNTO() []antlr.TerminalNode {
	return s.GetTokens(VGrammarPUNTO)
}

func (s *ID_PatronContext) PUNTO(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarPUNTO, i)
}

func (s *ID_PatronContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterID_Patron(s)
	}
}

func (s *ID_PatronContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitID_Patron(s)
	}
}

func (s *ID_PatronContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitID_Patron(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) PatronId() (localctx IPatronIdContext) {
	localctx = NewPatronIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, VGrammarRULE_patronId)
	var _alt int

	localctx = NewID_PatronContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(461)
		p.Match(VGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(466)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(462)
				p.Match(VGrammarPUNTO)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(463)
				p.Match(VGrammarID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(468)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *VGrammar) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 32:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *VGrammar) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
