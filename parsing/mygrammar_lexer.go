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
		"", "'begin'", "'end'", "'if'", "'then'", "'else'", "'while'", "'for'",
		"'do'", "'{'", "'}'", "'loop'", "'print'", "", "", "", "'='", "':'",
		"','", "'+'", "'-'", "'*'", "'/'", "'^'", "'=='", "'!='", "'<'", "'>'",
		"'\"'", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "STRING", "IDENTIFIER",
		"INTEGER", "EQ", "COLON", "COMMA", "PLUS", "MINUS", "MULTI", "DIVIDE",
		"POWERBY", "EQUALBY", "NOTEQUALBY", "LT", "GT", "COT", "OPEN_PARAN",
		"CLOSE_PARAN", "CONDITION_OP", "NEWLINE", "NEXT_PARAM", "EMPTY",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "STRING", "IDENTIFIER", "INTEGER", "EQ", "COLON",
		"COMMA", "PLUS", "MINUS", "MULTI", "DIVIDE", "POWERBY", "EQUALBY", "NOTEQUALBY",
		"LT", "GT", "COT", "OPEN_PARAN", "CLOSE_PARAN", "CONDITION_OP", "NEWLINE",
		"NEXT_PARAM", "EMPTY",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 34, 194, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 5, 12, 123, 8, 12, 10, 12, 12, 12, 126, 9, 12, 1, 12, 1, 12, 1,
		13, 1, 13, 5, 13, 132, 8, 13, 10, 13, 12, 13, 135, 9, 13, 1, 14, 4, 14,
		138, 8, 14, 11, 14, 12, 14, 139, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1,
		17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22,
		1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1,
		27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30,
		178, 8, 30, 1, 31, 1, 31, 4, 31, 182, 8, 31, 11, 31, 12, 31, 183, 1, 32,
		4, 32, 187, 8, 32, 11, 32, 12, 32, 188, 1, 33, 1, 33, 1, 33, 1, 33, 0,
		0, 34, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10,
		21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19,
		39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28,
		57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 1, 0, 8, 1, 0, 34, 34,
		3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1,
		0, 48, 57, 1, 0, 59, 59, 2, 0, 10, 10, 13, 13, 1, 0, 44, 44, 2, 0, 9, 9,
		32, 32, 201, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7,
		1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0,
		15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0,
		0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0,
		0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0,
		0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1,
		0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53,
		1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0,
		61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0,
		1, 69, 1, 0, 0, 0, 3, 75, 1, 0, 0, 0, 5, 79, 1, 0, 0, 0, 7, 82, 1, 0, 0,
		0, 9, 87, 1, 0, 0, 0, 11, 92, 1, 0, 0, 0, 13, 98, 1, 0, 0, 0, 15, 102,
		1, 0, 0, 0, 17, 105, 1, 0, 0, 0, 19, 107, 1, 0, 0, 0, 21, 109, 1, 0, 0,
		0, 23, 114, 1, 0, 0, 0, 25, 120, 1, 0, 0, 0, 27, 129, 1, 0, 0, 0, 29, 137,
		1, 0, 0, 0, 31, 141, 1, 0, 0, 0, 33, 143, 1, 0, 0, 0, 35, 145, 1, 0, 0,
		0, 37, 147, 1, 0, 0, 0, 39, 149, 1, 0, 0, 0, 41, 151, 1, 0, 0, 0, 43, 153,
		1, 0, 0, 0, 45, 155, 1, 0, 0, 0, 47, 157, 1, 0, 0, 0, 49, 160, 1, 0, 0,
		0, 51, 163, 1, 0, 0, 0, 53, 165, 1, 0, 0, 0, 55, 167, 1, 0, 0, 0, 57, 169,
		1, 0, 0, 0, 59, 171, 1, 0, 0, 0, 61, 177, 1, 0, 0, 0, 63, 179, 1, 0, 0,
		0, 65, 186, 1, 0, 0, 0, 67, 190, 1, 0, 0, 0, 69, 70, 5, 98, 0, 0, 70, 71,
		5, 101, 0, 0, 71, 72, 5, 103, 0, 0, 72, 73, 5, 105, 0, 0, 73, 74, 5, 110,
		0, 0, 74, 2, 1, 0, 0, 0, 75, 76, 5, 101, 0, 0, 76, 77, 5, 110, 0, 0, 77,
		78, 5, 100, 0, 0, 78, 4, 1, 0, 0, 0, 79, 80, 5, 105, 0, 0, 80, 81, 5, 102,
		0, 0, 81, 6, 1, 0, 0, 0, 82, 83, 5, 116, 0, 0, 83, 84, 5, 104, 0, 0, 84,
		85, 5, 101, 0, 0, 85, 86, 5, 110, 0, 0, 86, 8, 1, 0, 0, 0, 87, 88, 5, 101,
		0, 0, 88, 89, 5, 108, 0, 0, 89, 90, 5, 115, 0, 0, 90, 91, 5, 101, 0, 0,
		91, 10, 1, 0, 0, 0, 92, 93, 5, 119, 0, 0, 93, 94, 5, 104, 0, 0, 94, 95,
		5, 105, 0, 0, 95, 96, 5, 108, 0, 0, 96, 97, 5, 101, 0, 0, 97, 12, 1, 0,
		0, 0, 98, 99, 5, 102, 0, 0, 99, 100, 5, 111, 0, 0, 100, 101, 5, 114, 0,
		0, 101, 14, 1, 0, 0, 0, 102, 103, 5, 100, 0, 0, 103, 104, 5, 111, 0, 0,
		104, 16, 1, 0, 0, 0, 105, 106, 5, 123, 0, 0, 106, 18, 1, 0, 0, 0, 107,
		108, 5, 125, 0, 0, 108, 20, 1, 0, 0, 0, 109, 110, 5, 108, 0, 0, 110, 111,
		5, 111, 0, 0, 111, 112, 5, 111, 0, 0, 112, 113, 5, 112, 0, 0, 113, 22,
		1, 0, 0, 0, 114, 115, 5, 112, 0, 0, 115, 116, 5, 114, 0, 0, 116, 117, 5,
		105, 0, 0, 117, 118, 5, 110, 0, 0, 118, 119, 5, 116, 0, 0, 119, 24, 1,
		0, 0, 0, 120, 124, 5, 34, 0, 0, 121, 123, 8, 0, 0, 0, 122, 121, 1, 0, 0,
		0, 123, 126, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125,
		127, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 127, 128, 5, 34, 0, 0, 128, 26,
		1, 0, 0, 0, 129, 133, 7, 1, 0, 0, 130, 132, 7, 2, 0, 0, 131, 130, 1, 0,
		0, 0, 132, 135, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0,
		134, 28, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 136, 138, 7, 3, 0, 0, 137, 136,
		1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0,
		0, 0, 140, 30, 1, 0, 0, 0, 141, 142, 5, 61, 0, 0, 142, 32, 1, 0, 0, 0,
		143, 144, 5, 58, 0, 0, 144, 34, 1, 0, 0, 0, 145, 146, 5, 44, 0, 0, 146,
		36, 1, 0, 0, 0, 147, 148, 5, 43, 0, 0, 148, 38, 1, 0, 0, 0, 149, 150, 5,
		45, 0, 0, 150, 40, 1, 0, 0, 0, 151, 152, 5, 42, 0, 0, 152, 42, 1, 0, 0,
		0, 153, 154, 5, 47, 0, 0, 154, 44, 1, 0, 0, 0, 155, 156, 5, 94, 0, 0, 156,
		46, 1, 0, 0, 0, 157, 158, 5, 61, 0, 0, 158, 159, 5, 61, 0, 0, 159, 48,
		1, 0, 0, 0, 160, 161, 5, 33, 0, 0, 161, 162, 5, 61, 0, 0, 162, 50, 1, 0,
		0, 0, 163, 164, 5, 60, 0, 0, 164, 52, 1, 0, 0, 0, 165, 166, 5, 62, 0, 0,
		166, 54, 1, 0, 0, 0, 167, 168, 5, 34, 0, 0, 168, 56, 1, 0, 0, 0, 169, 170,
		5, 40, 0, 0, 170, 58, 1, 0, 0, 0, 171, 172, 5, 41, 0, 0, 172, 60, 1, 0,
		0, 0, 173, 178, 3, 51, 25, 0, 174, 178, 3, 53, 26, 0, 175, 178, 3, 47,
		23, 0, 176, 178, 3, 49, 24, 0, 177, 173, 1, 0, 0, 0, 177, 174, 1, 0, 0,
		0, 177, 175, 1, 0, 0, 0, 177, 176, 1, 0, 0, 0, 178, 62, 1, 0, 0, 0, 179,
		181, 7, 4, 0, 0, 180, 182, 7, 5, 0, 0, 181, 180, 1, 0, 0, 0, 182, 183,
		1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 64, 1, 0,
		0, 0, 185, 187, 7, 6, 0, 0, 186, 185, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0,
		188, 186, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 66, 1, 0, 0, 0, 190, 191,
		7, 7, 0, 0, 191, 192, 1, 0, 0, 0, 192, 193, 6, 33, 0, 0, 193, 68, 1, 0,
		0, 0, 7, 0, 124, 133, 139, 177, 183, 188, 1, 0, 1, 0,
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
	MyGrammarLexerT__0         = 1
	MyGrammarLexerT__1         = 2
	MyGrammarLexerT__2         = 3
	MyGrammarLexerT__3         = 4
	MyGrammarLexerT__4         = 5
	MyGrammarLexerT__5         = 6
	MyGrammarLexerT__6         = 7
	MyGrammarLexerT__7         = 8
	MyGrammarLexerT__8         = 9
	MyGrammarLexerT__9         = 10
	MyGrammarLexerT__10        = 11
	MyGrammarLexerT__11        = 12
	MyGrammarLexerSTRING       = 13
	MyGrammarLexerIDENTIFIER   = 14
	MyGrammarLexerINTEGER      = 15
	MyGrammarLexerEQ           = 16
	MyGrammarLexerCOLON        = 17
	MyGrammarLexerCOMMA        = 18
	MyGrammarLexerPLUS         = 19
	MyGrammarLexerMINUS        = 20
	MyGrammarLexerMULTI        = 21
	MyGrammarLexerDIVIDE       = 22
	MyGrammarLexerPOWERBY      = 23
	MyGrammarLexerEQUALBY      = 24
	MyGrammarLexerNOTEQUALBY   = 25
	MyGrammarLexerLT           = 26
	MyGrammarLexerGT           = 27
	MyGrammarLexerCOT          = 28
	MyGrammarLexerOPEN_PARAN   = 29
	MyGrammarLexerCLOSE_PARAN  = 30
	MyGrammarLexerCONDITION_OP = 31
	MyGrammarLexerNEWLINE      = 32
	MyGrammarLexerNEXT_PARAM   = 33
	MyGrammarLexerEMPTY        = 34
)
