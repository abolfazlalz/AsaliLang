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
		"STRING", "COMMENT", "SPACE",
	}
	staticData.RuleNames = []string{
		"T__0", "OR", "AND", "EQ", "NEQ", "GT", "LT", "GTEQ", "LTEQ", "PLUS",
		"MINUS", "MULT", "DIV", "MOD", "POW", "NOT", "SCOL", "COL", "ASSIGN",
		"OPAR", "CPAR", "OBRACE", "CBRACE", "BEGIN", "END", "DO", "THEN", "TRUE",
		"FALSE", "NIL", "IF", "ELSE", "WHILE", "FOR", "LOOP", "ID", "INT", "FLOAT",
		"STRING", "COMMENT", "SPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 41, 247, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1,
		9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1,
		20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32,
		1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1,
		35, 5, 35, 194, 8, 35, 10, 35, 12, 35, 197, 9, 35, 1, 36, 4, 36, 200, 8,
		36, 11, 36, 12, 36, 201, 1, 37, 4, 37, 205, 8, 37, 11, 37, 12, 37, 206,
		1, 37, 1, 37, 5, 37, 211, 8, 37, 10, 37, 12, 37, 214, 9, 37, 1, 37, 1,
		37, 4, 37, 218, 8, 37, 11, 37, 12, 37, 219, 3, 37, 222, 8, 37, 1, 38, 1,
		38, 1, 38, 1, 38, 5, 38, 228, 8, 38, 10, 38, 12, 38, 231, 9, 38, 1, 38,
		1, 38, 1, 39, 1, 39, 5, 39, 237, 8, 39, 10, 39, 12, 39, 240, 9, 39, 1,
		39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 0, 0, 41, 1, 1, 3, 2, 5, 3, 7, 4,
		9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23,
		47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32,
		65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41,
		1, 0, 6, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97,
		122, 1, 0, 48, 57, 3, 0, 10, 10, 13, 13, 34, 34, 2, 0, 10, 10, 13, 13,
		3, 0, 9, 10, 13, 13, 32, 32, 255, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0,
		5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0,
		13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0,
		0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0,
		0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0,
		0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1,
		0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51,
		1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0,
		59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0,
		0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0,
		0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0,
		0, 0, 1, 83, 1, 0, 0, 0, 3, 85, 1, 0, 0, 0, 5, 88, 1, 0, 0, 0, 7, 91, 1,
		0, 0, 0, 9, 94, 1, 0, 0, 0, 11, 97, 1, 0, 0, 0, 13, 99, 1, 0, 0, 0, 15,
		101, 1, 0, 0, 0, 17, 104, 1, 0, 0, 0, 19, 107, 1, 0, 0, 0, 21, 109, 1,
		0, 0, 0, 23, 111, 1, 0, 0, 0, 25, 113, 1, 0, 0, 0, 27, 115, 1, 0, 0, 0,
		29, 117, 1, 0, 0, 0, 31, 119, 1, 0, 0, 0, 33, 121, 1, 0, 0, 0, 35, 123,
		1, 0, 0, 0, 37, 125, 1, 0, 0, 0, 39, 127, 1, 0, 0, 0, 41, 129, 1, 0, 0,
		0, 43, 131, 1, 0, 0, 0, 45, 133, 1, 0, 0, 0, 47, 135, 1, 0, 0, 0, 49, 141,
		1, 0, 0, 0, 51, 145, 1, 0, 0, 0, 53, 148, 1, 0, 0, 0, 55, 153, 1, 0, 0,
		0, 57, 158, 1, 0, 0, 0, 59, 164, 1, 0, 0, 0, 61, 168, 1, 0, 0, 0, 63, 171,
		1, 0, 0, 0, 65, 176, 1, 0, 0, 0, 67, 182, 1, 0, 0, 0, 69, 186, 1, 0, 0,
		0, 71, 191, 1, 0, 0, 0, 73, 199, 1, 0, 0, 0, 75, 221, 1, 0, 0, 0, 77, 223,
		1, 0, 0, 0, 79, 234, 1, 0, 0, 0, 81, 243, 1, 0, 0, 0, 83, 84, 5, 44, 0,
		0, 84, 2, 1, 0, 0, 0, 85, 86, 5, 124, 0, 0, 86, 87, 5, 124, 0, 0, 87, 4,
		1, 0, 0, 0, 88, 89, 5, 38, 0, 0, 89, 90, 5, 38, 0, 0, 90, 6, 1, 0, 0, 0,
		91, 92, 5, 61, 0, 0, 92, 93, 5, 61, 0, 0, 93, 8, 1, 0, 0, 0, 94, 95, 5,
		33, 0, 0, 95, 96, 5, 61, 0, 0, 96, 10, 1, 0, 0, 0, 97, 98, 5, 62, 0, 0,
		98, 12, 1, 0, 0, 0, 99, 100, 5, 60, 0, 0, 100, 14, 1, 0, 0, 0, 101, 102,
		5, 62, 0, 0, 102, 103, 5, 61, 0, 0, 103, 16, 1, 0, 0, 0, 104, 105, 5, 60,
		0, 0, 105, 106, 5, 61, 0, 0, 106, 18, 1, 0, 0, 0, 107, 108, 5, 43, 0, 0,
		108, 20, 1, 0, 0, 0, 109, 110, 5, 45, 0, 0, 110, 22, 1, 0, 0, 0, 111, 112,
		5, 42, 0, 0, 112, 24, 1, 0, 0, 0, 113, 114, 5, 47, 0, 0, 114, 26, 1, 0,
		0, 0, 115, 116, 5, 37, 0, 0, 116, 28, 1, 0, 0, 0, 117, 118, 5, 94, 0, 0,
		118, 30, 1, 0, 0, 0, 119, 120, 5, 33, 0, 0, 120, 32, 1, 0, 0, 0, 121, 122,
		5, 59, 0, 0, 122, 34, 1, 0, 0, 0, 123, 124, 5, 58, 0, 0, 124, 36, 1, 0,
		0, 0, 125, 126, 5, 61, 0, 0, 126, 38, 1, 0, 0, 0, 127, 128, 5, 40, 0, 0,
		128, 40, 1, 0, 0, 0, 129, 130, 5, 41, 0, 0, 130, 42, 1, 0, 0, 0, 131, 132,
		5, 123, 0, 0, 132, 44, 1, 0, 0, 0, 133, 134, 5, 125, 0, 0, 134, 46, 1,
		0, 0, 0, 135, 136, 5, 98, 0, 0, 136, 137, 5, 101, 0, 0, 137, 138, 5, 103,
		0, 0, 138, 139, 5, 105, 0, 0, 139, 140, 5, 110, 0, 0, 140, 48, 1, 0, 0,
		0, 141, 142, 5, 101, 0, 0, 142, 143, 5, 110, 0, 0, 143, 144, 5, 100, 0,
		0, 144, 50, 1, 0, 0, 0, 145, 146, 5, 100, 0, 0, 146, 147, 5, 111, 0, 0,
		147, 52, 1, 0, 0, 0, 148, 149, 5, 116, 0, 0, 149, 150, 5, 104, 0, 0, 150,
		151, 5, 101, 0, 0, 151, 152, 5, 110, 0, 0, 152, 54, 1, 0, 0, 0, 153, 154,
		5, 116, 0, 0, 154, 155, 5, 114, 0, 0, 155, 156, 5, 117, 0, 0, 156, 157,
		5, 101, 0, 0, 157, 56, 1, 0, 0, 0, 158, 159, 5, 102, 0, 0, 159, 160, 5,
		97, 0, 0, 160, 161, 5, 108, 0, 0, 161, 162, 5, 115, 0, 0, 162, 163, 5,
		101, 0, 0, 163, 58, 1, 0, 0, 0, 164, 165, 5, 110, 0, 0, 165, 166, 5, 105,
		0, 0, 166, 167, 5, 108, 0, 0, 167, 60, 1, 0, 0, 0, 168, 169, 5, 105, 0,
		0, 169, 170, 5, 102, 0, 0, 170, 62, 1, 0, 0, 0, 171, 172, 5, 101, 0, 0,
		172, 173, 5, 108, 0, 0, 173, 174, 5, 115, 0, 0, 174, 175, 5, 101, 0, 0,
		175, 64, 1, 0, 0, 0, 176, 177, 5, 119, 0, 0, 177, 178, 5, 104, 0, 0, 178,
		179, 5, 105, 0, 0, 179, 180, 5, 108, 0, 0, 180, 181, 5, 101, 0, 0, 181,
		66, 1, 0, 0, 0, 182, 183, 5, 102, 0, 0, 183, 184, 5, 111, 0, 0, 184, 185,
		5, 114, 0, 0, 185, 68, 1, 0, 0, 0, 186, 187, 5, 108, 0, 0, 187, 188, 5,
		111, 0, 0, 188, 189, 5, 111, 0, 0, 189, 190, 5, 112, 0, 0, 190, 70, 1,
		0, 0, 0, 191, 195, 7, 0, 0, 0, 192, 194, 7, 1, 0, 0, 193, 192, 1, 0, 0,
		0, 194, 197, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196,
		72, 1, 0, 0, 0, 197, 195, 1, 0, 0, 0, 198, 200, 7, 2, 0, 0, 199, 198, 1,
		0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 201, 202, 1, 0, 0,
		0, 202, 74, 1, 0, 0, 0, 203, 205, 7, 2, 0, 0, 204, 203, 1, 0, 0, 0, 205,
		206, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207, 208,
		1, 0, 0, 0, 208, 212, 5, 46, 0, 0, 209, 211, 7, 2, 0, 0, 210, 209, 1, 0,
		0, 0, 211, 214, 1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0,
		213, 222, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0, 215, 217, 5, 46, 0, 0, 216,
		218, 7, 2, 0, 0, 217, 216, 1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219, 217,
		1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 222, 1, 0, 0, 0, 221, 204, 1, 0,
		0, 0, 221, 215, 1, 0, 0, 0, 222, 76, 1, 0, 0, 0, 223, 229, 5, 34, 0, 0,
		224, 228, 8, 3, 0, 0, 225, 226, 5, 34, 0, 0, 226, 228, 5, 34, 0, 0, 227,
		224, 1, 0, 0, 0, 227, 225, 1, 0, 0, 0, 228, 231, 1, 0, 0, 0, 229, 227,
		1, 0, 0, 0, 229, 230, 1, 0, 0, 0, 230, 232, 1, 0, 0, 0, 231, 229, 1, 0,
		0, 0, 232, 233, 5, 34, 0, 0, 233, 78, 1, 0, 0, 0, 234, 238, 5, 35, 0, 0,
		235, 237, 8, 4, 0, 0, 236, 235, 1, 0, 0, 0, 237, 240, 1, 0, 0, 0, 238,
		236, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 241, 1, 0, 0, 0, 240, 238,
		1, 0, 0, 0, 241, 242, 6, 39, 0, 0, 242, 80, 1, 0, 0, 0, 243, 244, 7, 5,
		0, 0, 244, 245, 1, 0, 0, 0, 245, 246, 6, 40, 0, 0, 246, 82, 1, 0, 0, 0,
		10, 0, 195, 201, 206, 212, 219, 221, 227, 229, 238, 1, 6, 0, 0,
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
)
