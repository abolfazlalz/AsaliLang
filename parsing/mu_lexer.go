// Code generated from MyGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

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

type MuLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var MuLexerLexerStaticData struct {
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

func mulexerLexerInit() {
	staticData := &MuLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'||'", "'&&'", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='", "'+'",
		"'-'", "'*'", "'/'", "'%'", "'^'", "'!'", "';'", "'='", "'('", "')'",
		"'{'", "'}'", "'true'", "'false'", "'nil'", "'if'", "'else'", "'while'",
		"'log'",
	}
	staticData.SymbolicNames = []string{
		"", "OR", "AND", "EQ", "NEQ", "GT", "LT", "GTEQ", "LTEQ", "PLUS", "MINUS",
		"MULT", "DIV", "MOD", "POW", "NOT", "SCOL", "ASSIGN", "OPAR", "CPAR",
		"OBRACE", "CBRACE", "TRUE", "FALSE", "NIL", "IF", "ELSE", "WHILE", "LOG",
		"ID", "INT", "FLOAT", "STRING", "COMMENT", "SPACE", "OTHER",
	}
	staticData.RuleNames = []string{
		"OR", "AND", "EQ", "NEQ", "GT", "LT", "GTEQ", "LTEQ", "PLUS", "MINUS",
		"MULT", "DIV", "MOD", "POW", "NOT", "SCOL", "ASSIGN", "OPAR", "CPAR",
		"OBRACE", "CBRACE", "TRUE", "FALSE", "NIL", "IF", "ELSE", "WHILE", "LOG",
		"ID", "INT", "FLOAT", "STRING", "COMMENT", "SPACE", "OTHER",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 35, 210, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 1, 0, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 28, 1, 28, 5, 28, 155, 8, 28, 10, 28, 12, 28, 158, 9, 28,
		1, 29, 4, 29, 161, 8, 29, 11, 29, 12, 29, 162, 1, 30, 4, 30, 166, 8, 30,
		11, 30, 12, 30, 167, 1, 30, 1, 30, 5, 30, 172, 8, 30, 10, 30, 12, 30, 175,
		9, 30, 1, 30, 1, 30, 4, 30, 179, 8, 30, 11, 30, 12, 30, 180, 3, 30, 183,
		8, 30, 1, 31, 1, 31, 1, 31, 1, 31, 5, 31, 189, 8, 31, 10, 31, 12, 31, 192,
		9, 31, 1, 31, 1, 31, 1, 32, 1, 32, 5, 32, 198, 8, 32, 10, 32, 12, 32, 201,
		9, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 0, 0, 35,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29,
		59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 1, 0, 6, 3, 0, 65, 90,
		95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 3,
		0, 10, 10, 13, 13, 34, 34, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10, 13, 13, 32,
		32, 218, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0,
		23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0,
		0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0,
		0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0,
		0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1,
		0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61,
		1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0,
		69, 1, 0, 0, 0, 1, 71, 1, 0, 0, 0, 3, 74, 1, 0, 0, 0, 5, 77, 1, 0, 0, 0,
		7, 80, 1, 0, 0, 0, 9, 83, 1, 0, 0, 0, 11, 85, 1, 0, 0, 0, 13, 87, 1, 0,
		0, 0, 15, 90, 1, 0, 0, 0, 17, 93, 1, 0, 0, 0, 19, 95, 1, 0, 0, 0, 21, 97,
		1, 0, 0, 0, 23, 99, 1, 0, 0, 0, 25, 101, 1, 0, 0, 0, 27, 103, 1, 0, 0,
		0, 29, 105, 1, 0, 0, 0, 31, 107, 1, 0, 0, 0, 33, 109, 1, 0, 0, 0, 35, 111,
		1, 0, 0, 0, 37, 113, 1, 0, 0, 0, 39, 115, 1, 0, 0, 0, 41, 117, 1, 0, 0,
		0, 43, 119, 1, 0, 0, 0, 45, 124, 1, 0, 0, 0, 47, 130, 1, 0, 0, 0, 49, 134,
		1, 0, 0, 0, 51, 137, 1, 0, 0, 0, 53, 142, 1, 0, 0, 0, 55, 148, 1, 0, 0,
		0, 57, 152, 1, 0, 0, 0, 59, 160, 1, 0, 0, 0, 61, 182, 1, 0, 0, 0, 63, 184,
		1, 0, 0, 0, 65, 195, 1, 0, 0, 0, 67, 204, 1, 0, 0, 0, 69, 208, 1, 0, 0,
		0, 71, 72, 5, 124, 0, 0, 72, 73, 5, 124, 0, 0, 73, 2, 1, 0, 0, 0, 74, 75,
		5, 38, 0, 0, 75, 76, 5, 38, 0, 0, 76, 4, 1, 0, 0, 0, 77, 78, 5, 61, 0,
		0, 78, 79, 5, 61, 0, 0, 79, 6, 1, 0, 0, 0, 80, 81, 5, 33, 0, 0, 81, 82,
		5, 61, 0, 0, 82, 8, 1, 0, 0, 0, 83, 84, 5, 62, 0, 0, 84, 10, 1, 0, 0, 0,
		85, 86, 5, 60, 0, 0, 86, 12, 1, 0, 0, 0, 87, 88, 5, 62, 0, 0, 88, 89, 5,
		61, 0, 0, 89, 14, 1, 0, 0, 0, 90, 91, 5, 60, 0, 0, 91, 92, 5, 61, 0, 0,
		92, 16, 1, 0, 0, 0, 93, 94, 5, 43, 0, 0, 94, 18, 1, 0, 0, 0, 95, 96, 5,
		45, 0, 0, 96, 20, 1, 0, 0, 0, 97, 98, 5, 42, 0, 0, 98, 22, 1, 0, 0, 0,
		99, 100, 5, 47, 0, 0, 100, 24, 1, 0, 0, 0, 101, 102, 5, 37, 0, 0, 102,
		26, 1, 0, 0, 0, 103, 104, 5, 94, 0, 0, 104, 28, 1, 0, 0, 0, 105, 106, 5,
		33, 0, 0, 106, 30, 1, 0, 0, 0, 107, 108, 5, 59, 0, 0, 108, 32, 1, 0, 0,
		0, 109, 110, 5, 61, 0, 0, 110, 34, 1, 0, 0, 0, 111, 112, 5, 40, 0, 0, 112,
		36, 1, 0, 0, 0, 113, 114, 5, 41, 0, 0, 114, 38, 1, 0, 0, 0, 115, 116, 5,
		123, 0, 0, 116, 40, 1, 0, 0, 0, 117, 118, 5, 125, 0, 0, 118, 42, 1, 0,
		0, 0, 119, 120, 5, 116, 0, 0, 120, 121, 5, 114, 0, 0, 121, 122, 5, 117,
		0, 0, 122, 123, 5, 101, 0, 0, 123, 44, 1, 0, 0, 0, 124, 125, 5, 102, 0,
		0, 125, 126, 5, 97, 0, 0, 126, 127, 5, 108, 0, 0, 127, 128, 5, 115, 0,
		0, 128, 129, 5, 101, 0, 0, 129, 46, 1, 0, 0, 0, 130, 131, 5, 110, 0, 0,
		131, 132, 5, 105, 0, 0, 132, 133, 5, 108, 0, 0, 133, 48, 1, 0, 0, 0, 134,
		135, 5, 105, 0, 0, 135, 136, 5, 102, 0, 0, 136, 50, 1, 0, 0, 0, 137, 138,
		5, 101, 0, 0, 138, 139, 5, 108, 0, 0, 139, 140, 5, 115, 0, 0, 140, 141,
		5, 101, 0, 0, 141, 52, 1, 0, 0, 0, 142, 143, 5, 119, 0, 0, 143, 144, 5,
		104, 0, 0, 144, 145, 5, 105, 0, 0, 145, 146, 5, 108, 0, 0, 146, 147, 5,
		101, 0, 0, 147, 54, 1, 0, 0, 0, 148, 149, 5, 108, 0, 0, 149, 150, 5, 111,
		0, 0, 150, 151, 5, 103, 0, 0, 151, 56, 1, 0, 0, 0, 152, 156, 7, 0, 0, 0,
		153, 155, 7, 1, 0, 0, 154, 153, 1, 0, 0, 0, 155, 158, 1, 0, 0, 0, 156,
		154, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 58, 1, 0, 0, 0, 158, 156, 1,
		0, 0, 0, 159, 161, 7, 2, 0, 0, 160, 159, 1, 0, 0, 0, 161, 162, 1, 0, 0,
		0, 162, 160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 60, 1, 0, 0, 0, 164,
		166, 7, 2, 0, 0, 165, 164, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 165,
		1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 173, 5, 46,
		0, 0, 170, 172, 7, 2, 0, 0, 171, 170, 1, 0, 0, 0, 172, 175, 1, 0, 0, 0,
		173, 171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 183, 1, 0, 0, 0, 175,
		173, 1, 0, 0, 0, 176, 178, 5, 46, 0, 0, 177, 179, 7, 2, 0, 0, 178, 177,
		1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 180, 181, 1, 0,
		0, 0, 181, 183, 1, 0, 0, 0, 182, 165, 1, 0, 0, 0, 182, 176, 1, 0, 0, 0,
		183, 62, 1, 0, 0, 0, 184, 190, 5, 34, 0, 0, 185, 189, 8, 3, 0, 0, 186,
		187, 5, 34, 0, 0, 187, 189, 5, 34, 0, 0, 188, 185, 1, 0, 0, 0, 188, 186,
		1, 0, 0, 0, 189, 192, 1, 0, 0, 0, 190, 188, 1, 0, 0, 0, 190, 191, 1, 0,
		0, 0, 191, 193, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 193, 194, 5, 34, 0, 0,
		194, 64, 1, 0, 0, 0, 195, 199, 5, 35, 0, 0, 196, 198, 8, 4, 0, 0, 197,
		196, 1, 0, 0, 0, 198, 201, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 199, 200,
		1, 0, 0, 0, 200, 202, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 202, 203, 6, 32,
		0, 0, 203, 66, 1, 0, 0, 0, 204, 205, 7, 5, 0, 0, 205, 206, 1, 0, 0, 0,
		206, 207, 6, 33, 0, 0, 207, 68, 1, 0, 0, 0, 208, 209, 9, 0, 0, 0, 209,
		70, 1, 0, 0, 0, 10, 0, 156, 162, 167, 173, 180, 182, 188, 190, 199, 1,
		6, 0, 0,
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

// MuLexerInit initializes any static state used to implement MuLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewMuLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func MuLexerInit() {
	staticData := &MuLexerLexerStaticData
	staticData.once.Do(mulexerLexerInit)
}

// NewMuLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewMuLexer(input antlr.CharStream) *MuLexer {
	MuLexerInit()
	l := new(MuLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &MuLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "MyGrammar.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MuLexer tokens.
const (
	MuLexerOR      = 1
	MuLexerAND     = 2
	MuLexerEQ      = 3
	MuLexerNEQ     = 4
	MuLexerGT      = 5
	MuLexerLT      = 6
	MuLexerGTEQ    = 7
	MuLexerLTEQ    = 8
	MuLexerPLUS    = 9
	MuLexerMINUS   = 10
	MuLexerMULT    = 11
	MuLexerDIV     = 12
	MuLexerMOD     = 13
	MuLexerPOW     = 14
	MuLexerNOT     = 15
	MuLexerSCOL    = 16
	MuLexerASSIGN  = 17
	MuLexerOPAR    = 18
	MuLexerCPAR    = 19
	MuLexerOBRACE  = 20
	MuLexerCBRACE  = 21
	MuLexerTRUE    = 22
	MuLexerFALSE   = 23
	MuLexerNIL     = 24
	MuLexerIF      = 25
	MuLexerELSE    = 26
	MuLexerWHILE   = 27
	MuLexerLOG     = 28
	MuLexerID      = 29
	MuLexerINT     = 30
	MuLexerFLOAT   = 31
	MuLexerSTRING  = 32
	MuLexerCOMMENT = 33
	MuLexerSPACE   = 34
	MuLexerOTHER   = 35
)