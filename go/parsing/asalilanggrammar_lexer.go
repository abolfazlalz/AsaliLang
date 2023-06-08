// Code generated from AsaliLangGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parsing

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type AsaliLangGrammarLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var AsaliLangGrammarLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func asalilanggrammarlexerLexerInit() {
	staticData := &AsaliLangGrammarLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "','", "'||'", "'&&'", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='",
		"'+'", "'-'", "'*'", "'/'", "'%'", "'^'", "'!'", "';'", "':'", "'='",
		"'('", "')'", "'{'", "'}'", "'begin'", "'end'", "'do'", "'then'", "'true'",
		"'false'", "'nil'", "'if'", "'else'", "'while'", "'for'", "'loop'",
	}
	staticData.SymbolicNames = []string{
		"", "", "OR", "AND", "EQ", "NEQ", "GT", "LT", "GTEQ", "LTEQ", "PLUS",
		"MINUS", "MULT", "DIV", "MOD", "POW", "NOT", "SCOL", "COL", "ASSIGN",
		"OPAR", "CPAR", "OBRACE", "CBRACE", "BEGIN", "END", "DO", "THEN", "TRUE",
		"FALSE", "NIL", "IF", "ELSE", "WHILE", "FOR", "LOOP", "ID", "INT", "FLOAT",
		"STRING", "COMMENT", "SPACE", "OTHER",
	}
	staticData.RuleNames = []string{
		"T__0", "OR", "AND", "EQ", "NEQ", "GT", "LT", "GTEQ", "LTEQ", "PLUS",
		"MINUS", "MULT", "DIV", "MOD", "POW", "NOT", "SCOL", "COL", "ASSIGN",
		"OPAR", "CPAR", "OBRACE", "CBRACE", "BEGIN", "END", "DO", "THEN", "TRUE",
		"FALSE", "NIL", "IF", "ELSE", "WHILE", "FOR", "LOOP", "ID", "INT", "FLOAT",
		"STRING", "COMMENT", "SPACE", "OTHER",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 42, 251, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1,
		19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1,
		30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 35, 1, 35, 5, 35, 196, 8, 35, 10, 35, 12, 35, 199, 9, 35, 1, 36,
		4, 36, 202, 8, 36, 11, 36, 12, 36, 203, 1, 37, 4, 37, 207, 8, 37, 11, 37,
		12, 37, 208, 1, 37, 1, 37, 5, 37, 213, 8, 37, 10, 37, 12, 37, 216, 9, 37,
		1, 37, 1, 37, 4, 37, 220, 8, 37, 11, 37, 12, 37, 221, 3, 37, 224, 8, 37,
		1, 38, 1, 38, 1, 38, 1, 38, 5, 38, 230, 8, 38, 10, 38, 12, 38, 233, 9,
		38, 1, 38, 1, 38, 1, 39, 1, 39, 5, 39, 239, 8, 39, 10, 39, 12, 39, 242,
		9, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 0, 0, 42,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29,
		59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38,
		77, 39, 79, 40, 81, 41, 83, 42, 1, 0, 6, 3, 0, 65, 90, 95, 95, 97, 122,
		4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 3, 0, 10, 10, 13,
		13, 34, 34, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10, 13, 13, 32, 32, 259, 0,
		1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0,
		9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0,
		0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0,
		0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0,
		0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1,
		0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47,
		1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0,
		55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0,
		0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0,
		0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0,
		0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 1, 85, 1,
		0, 0, 0, 3, 87, 1, 0, 0, 0, 5, 90, 1, 0, 0, 0, 7, 93, 1, 0, 0, 0, 9, 96,
		1, 0, 0, 0, 11, 99, 1, 0, 0, 0, 13, 101, 1, 0, 0, 0, 15, 103, 1, 0, 0,
		0, 17, 106, 1, 0, 0, 0, 19, 109, 1, 0, 0, 0, 21, 111, 1, 0, 0, 0, 23, 113,
		1, 0, 0, 0, 25, 115, 1, 0, 0, 0, 27, 117, 1, 0, 0, 0, 29, 119, 1, 0, 0,
		0, 31, 121, 1, 0, 0, 0, 33, 123, 1, 0, 0, 0, 35, 125, 1, 0, 0, 0, 37, 127,
		1, 0, 0, 0, 39, 129, 1, 0, 0, 0, 41, 131, 1, 0, 0, 0, 43, 133, 1, 0, 0,
		0, 45, 135, 1, 0, 0, 0, 47, 137, 1, 0, 0, 0, 49, 143, 1, 0, 0, 0, 51, 147,
		1, 0, 0, 0, 53, 150, 1, 0, 0, 0, 55, 155, 1, 0, 0, 0, 57, 160, 1, 0, 0,
		0, 59, 166, 1, 0, 0, 0, 61, 170, 1, 0, 0, 0, 63, 173, 1, 0, 0, 0, 65, 178,
		1, 0, 0, 0, 67, 184, 1, 0, 0, 0, 69, 188, 1, 0, 0, 0, 71, 193, 1, 0, 0,
		0, 73, 201, 1, 0, 0, 0, 75, 223, 1, 0, 0, 0, 77, 225, 1, 0, 0, 0, 79, 236,
		1, 0, 0, 0, 81, 245, 1, 0, 0, 0, 83, 249, 1, 0, 0, 0, 85, 86, 5, 44, 0,
		0, 86, 2, 1, 0, 0, 0, 87, 88, 5, 124, 0, 0, 88, 89, 5, 124, 0, 0, 89, 4,
		1, 0, 0, 0, 90, 91, 5, 38, 0, 0, 91, 92, 5, 38, 0, 0, 92, 6, 1, 0, 0, 0,
		93, 94, 5, 61, 0, 0, 94, 95, 5, 61, 0, 0, 95, 8, 1, 0, 0, 0, 96, 97, 5,
		33, 0, 0, 97, 98, 5, 61, 0, 0, 98, 10, 1, 0, 0, 0, 99, 100, 5, 62, 0, 0,
		100, 12, 1, 0, 0, 0, 101, 102, 5, 60, 0, 0, 102, 14, 1, 0, 0, 0, 103, 104,
		5, 62, 0, 0, 104, 105, 5, 61, 0, 0, 105, 16, 1, 0, 0, 0, 106, 107, 5, 60,
		0, 0, 107, 108, 5, 61, 0, 0, 108, 18, 1, 0, 0, 0, 109, 110, 5, 43, 0, 0,
		110, 20, 1, 0, 0, 0, 111, 112, 5, 45, 0, 0, 112, 22, 1, 0, 0, 0, 113, 114,
		5, 42, 0, 0, 114, 24, 1, 0, 0, 0, 115, 116, 5, 47, 0, 0, 116, 26, 1, 0,
		0, 0, 117, 118, 5, 37, 0, 0, 118, 28, 1, 0, 0, 0, 119, 120, 5, 94, 0, 0,
		120, 30, 1, 0, 0, 0, 121, 122, 5, 33, 0, 0, 122, 32, 1, 0, 0, 0, 123, 124,
		5, 59, 0, 0, 124, 34, 1, 0, 0, 0, 125, 126, 5, 58, 0, 0, 126, 36, 1, 0,
		0, 0, 127, 128, 5, 61, 0, 0, 128, 38, 1, 0, 0, 0, 129, 130, 5, 40, 0, 0,
		130, 40, 1, 0, 0, 0, 131, 132, 5, 41, 0, 0, 132, 42, 1, 0, 0, 0, 133, 134,
		5, 123, 0, 0, 134, 44, 1, 0, 0, 0, 135, 136, 5, 125, 0, 0, 136, 46, 1,
		0, 0, 0, 137, 138, 5, 98, 0, 0, 138, 139, 5, 101, 0, 0, 139, 140, 5, 103,
		0, 0, 140, 141, 5, 105, 0, 0, 141, 142, 5, 110, 0, 0, 142, 48, 1, 0, 0,
		0, 143, 144, 5, 101, 0, 0, 144, 145, 5, 110, 0, 0, 145, 146, 5, 100, 0,
		0, 146, 50, 1, 0, 0, 0, 147, 148, 5, 100, 0, 0, 148, 149, 5, 111, 0, 0,
		149, 52, 1, 0, 0, 0, 150, 151, 5, 116, 0, 0, 151, 152, 5, 104, 0, 0, 152,
		153, 5, 101, 0, 0, 153, 154, 5, 110, 0, 0, 154, 54, 1, 0, 0, 0, 155, 156,
		5, 116, 0, 0, 156, 157, 5, 114, 0, 0, 157, 158, 5, 117, 0, 0, 158, 159,
		5, 101, 0, 0, 159, 56, 1, 0, 0, 0, 160, 161, 5, 102, 0, 0, 161, 162, 5,
		97, 0, 0, 162, 163, 5, 108, 0, 0, 163, 164, 5, 115, 0, 0, 164, 165, 5,
		101, 0, 0, 165, 58, 1, 0, 0, 0, 166, 167, 5, 110, 0, 0, 167, 168, 5, 105,
		0, 0, 168, 169, 5, 108, 0, 0, 169, 60, 1, 0, 0, 0, 170, 171, 5, 105, 0,
		0, 171, 172, 5, 102, 0, 0, 172, 62, 1, 0, 0, 0, 173, 174, 5, 101, 0, 0,
		174, 175, 5, 108, 0, 0, 175, 176, 5, 115, 0, 0, 176, 177, 5, 101, 0, 0,
		177, 64, 1, 0, 0, 0, 178, 179, 5, 119, 0, 0, 179, 180, 5, 104, 0, 0, 180,
		181, 5, 105, 0, 0, 181, 182, 5, 108, 0, 0, 182, 183, 5, 101, 0, 0, 183,
		66, 1, 0, 0, 0, 184, 185, 5, 102, 0, 0, 185, 186, 5, 111, 0, 0, 186, 187,
		5, 114, 0, 0, 187, 68, 1, 0, 0, 0, 188, 189, 5, 108, 0, 0, 189, 190, 5,
		111, 0, 0, 190, 191, 5, 111, 0, 0, 191, 192, 5, 112, 0, 0, 192, 70, 1,
		0, 0, 0, 193, 197, 7, 0, 0, 0, 194, 196, 7, 1, 0, 0, 195, 194, 1, 0, 0,
		0, 196, 199, 1, 0, 0, 0, 197, 195, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198,
		72, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 200, 202, 7, 2, 0, 0, 201, 200, 1,
		0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 203, 204, 1, 0, 0,
		0, 204, 74, 1, 0, 0, 0, 205, 207, 7, 2, 0, 0, 206, 205, 1, 0, 0, 0, 207,
		208, 1, 0, 0, 0, 208, 206, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209, 210,
		1, 0, 0, 0, 210, 214, 5, 46, 0, 0, 211, 213, 7, 2, 0, 0, 212, 211, 1, 0,
		0, 0, 213, 216, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0,
		215, 224, 1, 0, 0, 0, 216, 214, 1, 0, 0, 0, 217, 219, 5, 46, 0, 0, 218,
		220, 7, 2, 0, 0, 219, 218, 1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221, 219,
		1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 224, 1, 0, 0, 0, 223, 206, 1, 0,
		0, 0, 223, 217, 1, 0, 0, 0, 224, 76, 1, 0, 0, 0, 225, 231, 5, 34, 0, 0,
		226, 230, 8, 3, 0, 0, 227, 228, 5, 34, 0, 0, 228, 230, 5, 34, 0, 0, 229,
		226, 1, 0, 0, 0, 229, 227, 1, 0, 0, 0, 230, 233, 1, 0, 0, 0, 231, 229,
		1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 234, 1, 0, 0, 0, 233, 231, 1, 0,
		0, 0, 234, 235, 5, 34, 0, 0, 235, 78, 1, 0, 0, 0, 236, 240, 5, 35, 0, 0,
		237, 239, 8, 4, 0, 0, 238, 237, 1, 0, 0, 0, 239, 242, 1, 0, 0, 0, 240,
		238, 1, 0, 0, 0, 240, 241, 1, 0, 0, 0, 241, 243, 1, 0, 0, 0, 242, 240,
		1, 0, 0, 0, 243, 244, 6, 39, 0, 0, 244, 80, 1, 0, 0, 0, 245, 246, 7, 5,
		0, 0, 246, 247, 1, 0, 0, 0, 247, 248, 6, 40, 0, 0, 248, 82, 1, 0, 0, 0,
		249, 250, 9, 0, 0, 0, 250, 84, 1, 0, 0, 0, 10, 0, 197, 203, 208, 214, 221,
		223, 229, 231, 240, 1, 6, 0, 0,
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

// AsaliLangGrammarLexerInit initializes any static state used to implement AsaliLangGrammarLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewAsaliLangGrammarLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func AsaliLangGrammarLexerInit() {
	staticData := &AsaliLangGrammarLexerLexerStaticData
	staticData.once.Do(asalilanggrammarlexerLexerInit)
}

// NewAsaliLangGrammarLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewAsaliLangGrammarLexer(input antlr.CharStream) *AsaliLangGrammarLexer {
	AsaliLangGrammarLexerInit()
	l := new(AsaliLangGrammarLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &AsaliLangGrammarLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "AsaliLangGrammar.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AsaliLangGrammarLexer tokens.
const (
	AsaliLangGrammarLexerT__0    = 1
	AsaliLangGrammarLexerOR      = 2
	AsaliLangGrammarLexerAND     = 3
	AsaliLangGrammarLexerEQ      = 4
	AsaliLangGrammarLexerNEQ     = 5
	AsaliLangGrammarLexerGT      = 6
	AsaliLangGrammarLexerLT      = 7
	AsaliLangGrammarLexerGTEQ    = 8
	AsaliLangGrammarLexerLTEQ    = 9
	AsaliLangGrammarLexerPLUS    = 10
	AsaliLangGrammarLexerMINUS   = 11
	AsaliLangGrammarLexerMULT    = 12
	AsaliLangGrammarLexerDIV     = 13
	AsaliLangGrammarLexerMOD     = 14
	AsaliLangGrammarLexerPOW     = 15
	AsaliLangGrammarLexerNOT     = 16
	AsaliLangGrammarLexerSCOL    = 17
	AsaliLangGrammarLexerCOL     = 18
	AsaliLangGrammarLexerASSIGN  = 19
	AsaliLangGrammarLexerOPAR    = 20
	AsaliLangGrammarLexerCPAR    = 21
	AsaliLangGrammarLexerOBRACE  = 22
	AsaliLangGrammarLexerCBRACE  = 23
	AsaliLangGrammarLexerBEGIN   = 24
	AsaliLangGrammarLexerEND     = 25
	AsaliLangGrammarLexerDO      = 26
	AsaliLangGrammarLexerTHEN    = 27
	AsaliLangGrammarLexerTRUE    = 28
	AsaliLangGrammarLexerFALSE   = 29
	AsaliLangGrammarLexerNIL     = 30
	AsaliLangGrammarLexerIF      = 31
	AsaliLangGrammarLexerELSE    = 32
	AsaliLangGrammarLexerWHILE   = 33
	AsaliLangGrammarLexerFOR     = 34
	AsaliLangGrammarLexerLOOP    = 35
	AsaliLangGrammarLexerID      = 36
	AsaliLangGrammarLexerINT     = 37
	AsaliLangGrammarLexerFLOAT   = 38
	AsaliLangGrammarLexerSTRING  = 39
	AsaliLangGrammarLexerCOMMENT = 40
	AsaliLangGrammarLexerSPACE   = 41
	AsaliLangGrammarLexerOTHER   = 42
)
