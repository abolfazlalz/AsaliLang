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

type MyGrammarLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var MyGrammarLexerLexerStaticData struct {
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

func mygrammarlexerLexerInit() {
	staticData := &MyGrammarLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "", "", "'='", "':'", "','", "'+'", "'-'", "'*'", "'/'", "'^'",
		"'if'", "'else'", "'do'", "'begin'", "'end'", "'then'", "'while'", "'for'",
		"'loop'", "'print'", "'true'", "'false'", "'=='", "'!='", "'<'", "'>'",
		"'\"'", "';'", "'('", "')'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "STRING", "IDENTIFIER", "INTEGER", "ASSIGN", "COLON", "COMMA", "PLUS",
		"MINUS", "MULTI", "DIVIDE", "POWERBY", "IF", "ElSE", "DO", "BEGIN",
		"END", "THEN", "WHILE", "FOR", "LOOP", "PRINT", "TRUE", "FALSE", "EQ",
		"NOTEQUALBY", "LT", "GT", "COT", "SEMICOLON", "OPARAN", "CPARAN", "OBRAC",
		"CBRAC", "CONDITION_OP", "NEXT_PARAM", "EMPTY",
	}
	staticData.RuleNames = []string{
		"STRING", "IDENTIFIER", "INTEGER", "ASSIGN", "COLON", "COMMA", "PLUS",
		"MINUS", "MULTI", "DIVIDE", "POWERBY", "IF", "ElSE", "DO", "BEGIN",
		"END", "THEN", "WHILE", "FOR", "LOOP", "PRINT", "TRUE", "FALSE", "EQ",
		"NOTEQUALBY", "LT", "GT", "COT", "SEMICOLON", "OPARAN", "CPARAN", "OBRAC",
		"CBRAC", "CONDITION_OP", "NEXT_PARAM", "EMPTY",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 36, 205, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 1, 0,
		1, 0, 5, 0, 76, 8, 0, 10, 0, 12, 0, 79, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 5,
		1, 85, 8, 1, 10, 1, 12, 1, 88, 9, 1, 1, 2, 4, 2, 91, 8, 2, 11, 2, 12, 2,
		92, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27,
		1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1,
		32, 1, 33, 1, 33, 1, 33, 1, 33, 3, 33, 195, 8, 33, 1, 34, 4, 34, 198, 8,
		34, 11, 34, 12, 34, 199, 1, 35, 1, 35, 1, 35, 1, 35, 0, 0, 36, 1, 1, 3,
		2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12,
		25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21,
		43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30,
		61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 1, 0, 6, 1, 0, 34, 34,
		3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1,
		0, 48, 57, 1, 0, 44, 44, 2, 0, 9, 10, 32, 32, 211, 0, 1, 1, 0, 0, 0, 0,
		3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0,
		11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0,
		0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0,
		0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0,
		0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1,
		0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49,
		1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0,
		57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0,
		0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0,
		0, 1, 73, 1, 0, 0, 0, 3, 82, 1, 0, 0, 0, 5, 90, 1, 0, 0, 0, 7, 94, 1, 0,
		0, 0, 9, 96, 1, 0, 0, 0, 11, 98, 1, 0, 0, 0, 13, 100, 1, 0, 0, 0, 15, 102,
		1, 0, 0, 0, 17, 104, 1, 0, 0, 0, 19, 106, 1, 0, 0, 0, 21, 108, 1, 0, 0,
		0, 23, 110, 1, 0, 0, 0, 25, 113, 1, 0, 0, 0, 27, 118, 1, 0, 0, 0, 29, 121,
		1, 0, 0, 0, 31, 127, 1, 0, 0, 0, 33, 131, 1, 0, 0, 0, 35, 136, 1, 0, 0,
		0, 37, 142, 1, 0, 0, 0, 39, 146, 1, 0, 0, 0, 41, 151, 1, 0, 0, 0, 43, 157,
		1, 0, 0, 0, 45, 162, 1, 0, 0, 0, 47, 168, 1, 0, 0, 0, 49, 171, 1, 0, 0,
		0, 51, 174, 1, 0, 0, 0, 53, 176, 1, 0, 0, 0, 55, 178, 1, 0, 0, 0, 57, 180,
		1, 0, 0, 0, 59, 182, 1, 0, 0, 0, 61, 184, 1, 0, 0, 0, 63, 186, 1, 0, 0,
		0, 65, 188, 1, 0, 0, 0, 67, 194, 1, 0, 0, 0, 69, 197, 1, 0, 0, 0, 71, 201,
		1, 0, 0, 0, 73, 77, 5, 34, 0, 0, 74, 76, 8, 0, 0, 0, 75, 74, 1, 0, 0, 0,
		76, 79, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 80, 1,
		0, 0, 0, 79, 77, 1, 0, 0, 0, 80, 81, 5, 34, 0, 0, 81, 2, 1, 0, 0, 0, 82,
		86, 7, 1, 0, 0, 83, 85, 7, 2, 0, 0, 84, 83, 1, 0, 0, 0, 85, 88, 1, 0, 0,
		0, 86, 84, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 4, 1, 0, 0, 0, 88, 86, 1,
		0, 0, 0, 89, 91, 7, 3, 0, 0, 90, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92,
		90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 6, 1, 0, 0, 0, 94, 95, 5, 61, 0,
		0, 95, 8, 1, 0, 0, 0, 96, 97, 5, 58, 0, 0, 97, 10, 1, 0, 0, 0, 98, 99,
		5, 44, 0, 0, 99, 12, 1, 0, 0, 0, 100, 101, 5, 43, 0, 0, 101, 14, 1, 0,
		0, 0, 102, 103, 5, 45, 0, 0, 103, 16, 1, 0, 0, 0, 104, 105, 5, 42, 0, 0,
		105, 18, 1, 0, 0, 0, 106, 107, 5, 47, 0, 0, 107, 20, 1, 0, 0, 0, 108, 109,
		5, 94, 0, 0, 109, 22, 1, 0, 0, 0, 110, 111, 5, 105, 0, 0, 111, 112, 5,
		102, 0, 0, 112, 24, 1, 0, 0, 0, 113, 114, 5, 101, 0, 0, 114, 115, 5, 108,
		0, 0, 115, 116, 5, 115, 0, 0, 116, 117, 5, 101, 0, 0, 117, 26, 1, 0, 0,
		0, 118, 119, 5, 100, 0, 0, 119, 120, 5, 111, 0, 0, 120, 28, 1, 0, 0, 0,
		121, 122, 5, 98, 0, 0, 122, 123, 5, 101, 0, 0, 123, 124, 5, 103, 0, 0,
		124, 125, 5, 105, 0, 0, 125, 126, 5, 110, 0, 0, 126, 30, 1, 0, 0, 0, 127,
		128, 5, 101, 0, 0, 128, 129, 5, 110, 0, 0, 129, 130, 5, 100, 0, 0, 130,
		32, 1, 0, 0, 0, 131, 132, 5, 116, 0, 0, 132, 133, 5, 104, 0, 0, 133, 134,
		5, 101, 0, 0, 134, 135, 5, 110, 0, 0, 135, 34, 1, 0, 0, 0, 136, 137, 5,
		119, 0, 0, 137, 138, 5, 104, 0, 0, 138, 139, 5, 105, 0, 0, 139, 140, 5,
		108, 0, 0, 140, 141, 5, 101, 0, 0, 141, 36, 1, 0, 0, 0, 142, 143, 5, 102,
		0, 0, 143, 144, 5, 111, 0, 0, 144, 145, 5, 114, 0, 0, 145, 38, 1, 0, 0,
		0, 146, 147, 5, 108, 0, 0, 147, 148, 5, 111, 0, 0, 148, 149, 5, 111, 0,
		0, 149, 150, 5, 112, 0, 0, 150, 40, 1, 0, 0, 0, 151, 152, 5, 112, 0, 0,
		152, 153, 5, 114, 0, 0, 153, 154, 5, 105, 0, 0, 154, 155, 5, 110, 0, 0,
		155, 156, 5, 116, 0, 0, 156, 42, 1, 0, 0, 0, 157, 158, 5, 116, 0, 0, 158,
		159, 5, 114, 0, 0, 159, 160, 5, 117, 0, 0, 160, 161, 5, 101, 0, 0, 161,
		44, 1, 0, 0, 0, 162, 163, 5, 102, 0, 0, 163, 164, 5, 97, 0, 0, 164, 165,
		5, 108, 0, 0, 165, 166, 5, 115, 0, 0, 166, 167, 5, 101, 0, 0, 167, 46,
		1, 0, 0, 0, 168, 169, 5, 61, 0, 0, 169, 170, 5, 61, 0, 0, 170, 48, 1, 0,
		0, 0, 171, 172, 5, 33, 0, 0, 172, 173, 5, 61, 0, 0, 173, 50, 1, 0, 0, 0,
		174, 175, 5, 60, 0, 0, 175, 52, 1, 0, 0, 0, 176, 177, 5, 62, 0, 0, 177,
		54, 1, 0, 0, 0, 178, 179, 5, 34, 0, 0, 179, 56, 1, 0, 0, 0, 180, 181, 5,
		59, 0, 0, 181, 58, 1, 0, 0, 0, 182, 183, 5, 40, 0, 0, 183, 60, 1, 0, 0,
		0, 184, 185, 5, 41, 0, 0, 185, 62, 1, 0, 0, 0, 186, 187, 5, 123, 0, 0,
		187, 64, 1, 0, 0, 0, 188, 189, 5, 125, 0, 0, 189, 66, 1, 0, 0, 0, 190,
		195, 3, 51, 25, 0, 191, 195, 3, 53, 26, 0, 192, 195, 3, 47, 23, 0, 193,
		195, 3, 49, 24, 0, 194, 190, 1, 0, 0, 0, 194, 191, 1, 0, 0, 0, 194, 192,
		1, 0, 0, 0, 194, 193, 1, 0, 0, 0, 195, 68, 1, 0, 0, 0, 196, 198, 7, 4,
		0, 0, 197, 196, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0,
		199, 200, 1, 0, 0, 0, 200, 70, 1, 0, 0, 0, 201, 202, 7, 5, 0, 0, 202, 203,
		1, 0, 0, 0, 203, 204, 6, 35, 0, 0, 204, 72, 1, 0, 0, 0, 6, 0, 77, 86, 92,
		194, 199, 1, 0, 1, 0,
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

// MyGrammarLexerInit initializes any static state used to implement MyGrammarLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewMyGrammarLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func MyGrammarLexerInit() {
	staticData := &MyGrammarLexerLexerStaticData
	staticData.once.Do(mygrammarlexerLexerInit)
}

// NewMyGrammarLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewMyGrammarLexer(input antlr.CharStream) *MyGrammarLexer {
	MyGrammarLexerInit()
	l := new(MyGrammarLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &MyGrammarLexerLexerStaticData
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

// MyGrammarLexer tokens.
const (
	MyGrammarLexerSTRING       = 1
	MyGrammarLexerIDENTIFIER   = 2
	MyGrammarLexerINTEGER      = 3
	MyGrammarLexerASSIGN       = 4
	MyGrammarLexerCOLON        = 5
	MyGrammarLexerCOMMA        = 6
	MyGrammarLexerPLUS         = 7
	MyGrammarLexerMINUS        = 8
	MyGrammarLexerMULTI        = 9
	MyGrammarLexerDIVIDE       = 10
	MyGrammarLexerPOWERBY      = 11
	MyGrammarLexerIF           = 12
	MyGrammarLexerElSE         = 13
	MyGrammarLexerDO           = 14
	MyGrammarLexerBEGIN        = 15
	MyGrammarLexerEND          = 16
	MyGrammarLexerTHEN         = 17
	MyGrammarLexerWHILE        = 18
	MyGrammarLexerFOR          = 19
	MyGrammarLexerLOOP         = 20
	MyGrammarLexerPRINT        = 21
	MyGrammarLexerTRUE         = 22
	MyGrammarLexerFALSE        = 23
	MyGrammarLexerEQ           = 24
	MyGrammarLexerNOTEQUALBY   = 25
	MyGrammarLexerLT           = 26
	MyGrammarLexerGT           = 27
	MyGrammarLexerCOT          = 28
	MyGrammarLexerSEMICOLON    = 29
	MyGrammarLexerOPARAN       = 30
	MyGrammarLexerCPARAN       = 31
	MyGrammarLexerOBRAC        = 32
	MyGrammarLexerCBRAC        = 33
	MyGrammarLexerCONDITION_OP = 34
	MyGrammarLexerNEXT_PARAM   = 35
	MyGrammarLexerEMPTY        = 36
)
