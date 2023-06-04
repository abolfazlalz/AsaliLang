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
		"'do'", "'loop'", "'print'", "", "", "", "'='", "':'", "','", "'+'",
		"'-'", "'*'", "'/'", "'^'", "'=='", "'!='", "'<'", "'>'", "'\"'", "'('",
		"')'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "STRING", "IDENTIFIER",
		"INTEGER", "EQ", "COLON", "COMMA", "PLUS", "MINUS", "MULTI", "DIVIDE",
		"POWERBY", "EQUALBY", "NOTEQUALBY", "LT", "GT", "COT", "OPEN_PARAN",
		"CLOSE_PARAN", "CONDITION_OP", "NEWLINE", "NEXT_PARAM", "EMPTY",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "STRING", "IDENTIFIER", "INTEGER", "EQ", "COLON", "COMMA", "PLUS",
		"MINUS", "MULTI", "DIVIDE", "POWERBY", "EQUALBY", "NOTEQUALBY", "LT",
		"GT", "COT", "OPEN_PARAN", "CLOSE_PARAN", "CONDITION_OP", "NEWLINE",
		"NEXT_PARAM", "EMPTY",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 32, 186, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 10, 1, 10, 5, 10, 115, 8, 10, 10, 10, 12, 10, 118, 9, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 5, 11, 124, 8, 11, 10, 11, 12, 11, 127, 9, 11, 1,
		12, 4, 12, 130, 8, 12, 11, 12, 12, 12, 131, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1,
		20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24,
		1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1,
		28, 3, 28, 170, 8, 28, 1, 29, 1, 29, 4, 29, 174, 8, 29, 11, 29, 12, 29,
		175, 1, 30, 4, 30, 179, 8, 30, 11, 30, 12, 30, 180, 1, 31, 1, 31, 1, 31,
		1, 31, 0, 0, 32, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17,
		9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35,
		18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53,
		27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 1, 0, 8, 1, 0, 34, 34, 3, 0,
		65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48,
		57, 1, 0, 59, 59, 2, 0, 10, 10, 13, 13, 1, 0, 44, 44, 2, 0, 9, 9, 32, 32,
		193, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0,
		0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1,
		0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23,
		1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0,
		31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0,
		0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0,
		0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0,
		0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1,
		0, 0, 0, 0, 63, 1, 0, 0, 0, 1, 65, 1, 0, 0, 0, 3, 71, 1, 0, 0, 0, 5, 75,
		1, 0, 0, 0, 7, 78, 1, 0, 0, 0, 9, 83, 1, 0, 0, 0, 11, 88, 1, 0, 0, 0, 13,
		94, 1, 0, 0, 0, 15, 98, 1, 0, 0, 0, 17, 101, 1, 0, 0, 0, 19, 106, 1, 0,
		0, 0, 21, 112, 1, 0, 0, 0, 23, 121, 1, 0, 0, 0, 25, 129, 1, 0, 0, 0, 27,
		133, 1, 0, 0, 0, 29, 135, 1, 0, 0, 0, 31, 137, 1, 0, 0, 0, 33, 139, 1,
		0, 0, 0, 35, 141, 1, 0, 0, 0, 37, 143, 1, 0, 0, 0, 39, 145, 1, 0, 0, 0,
		41, 147, 1, 0, 0, 0, 43, 149, 1, 0, 0, 0, 45, 152, 1, 0, 0, 0, 47, 155,
		1, 0, 0, 0, 49, 157, 1, 0, 0, 0, 51, 159, 1, 0, 0, 0, 53, 161, 1, 0, 0,
		0, 55, 163, 1, 0, 0, 0, 57, 169, 1, 0, 0, 0, 59, 171, 1, 0, 0, 0, 61, 178,
		1, 0, 0, 0, 63, 182, 1, 0, 0, 0, 65, 66, 5, 98, 0, 0, 66, 67, 5, 101, 0,
		0, 67, 68, 5, 103, 0, 0, 68, 69, 5, 105, 0, 0, 69, 70, 5, 110, 0, 0, 70,
		2, 1, 0, 0, 0, 71, 72, 5, 101, 0, 0, 72, 73, 5, 110, 0, 0, 73, 74, 5, 100,
		0, 0, 74, 4, 1, 0, 0, 0, 75, 76, 5, 105, 0, 0, 76, 77, 5, 102, 0, 0, 77,
		6, 1, 0, 0, 0, 78, 79, 5, 116, 0, 0, 79, 80, 5, 104, 0, 0, 80, 81, 5, 101,
		0, 0, 81, 82, 5, 110, 0, 0, 82, 8, 1, 0, 0, 0, 83, 84, 5, 101, 0, 0, 84,
		85, 5, 108, 0, 0, 85, 86, 5, 115, 0, 0, 86, 87, 5, 101, 0, 0, 87, 10, 1,
		0, 0, 0, 88, 89, 5, 119, 0, 0, 89, 90, 5, 104, 0, 0, 90, 91, 5, 105, 0,
		0, 91, 92, 5, 108, 0, 0, 92, 93, 5, 101, 0, 0, 93, 12, 1, 0, 0, 0, 94,
		95, 5, 102, 0, 0, 95, 96, 5, 111, 0, 0, 96, 97, 5, 114, 0, 0, 97, 14, 1,
		0, 0, 0, 98, 99, 5, 100, 0, 0, 99, 100, 5, 111, 0, 0, 100, 16, 1, 0, 0,
		0, 101, 102, 5, 108, 0, 0, 102, 103, 5, 111, 0, 0, 103, 104, 5, 111, 0,
		0, 104, 105, 5, 112, 0, 0, 105, 18, 1, 0, 0, 0, 106, 107, 5, 112, 0, 0,
		107, 108, 5, 114, 0, 0, 108, 109, 5, 105, 0, 0, 109, 110, 5, 110, 0, 0,
		110, 111, 5, 116, 0, 0, 111, 20, 1, 0, 0, 0, 112, 116, 5, 34, 0, 0, 113,
		115, 8, 0, 0, 0, 114, 113, 1, 0, 0, 0, 115, 118, 1, 0, 0, 0, 116, 114,
		1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 119, 1, 0, 0, 0, 118, 116, 1, 0,
		0, 0, 119, 120, 5, 34, 0, 0, 120, 22, 1, 0, 0, 0, 121, 125, 7, 1, 0, 0,
		122, 124, 7, 2, 0, 0, 123, 122, 1, 0, 0, 0, 124, 127, 1, 0, 0, 0, 125,
		123, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 24, 1, 0, 0, 0, 127, 125, 1,
		0, 0, 0, 128, 130, 7, 3, 0, 0, 129, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0,
		0, 131, 129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 26, 1, 0, 0, 0, 133,
		134, 5, 61, 0, 0, 134, 28, 1, 0, 0, 0, 135, 136, 5, 58, 0, 0, 136, 30,
		1, 0, 0, 0, 137, 138, 5, 44, 0, 0, 138, 32, 1, 0, 0, 0, 139, 140, 5, 43,
		0, 0, 140, 34, 1, 0, 0, 0, 141, 142, 5, 45, 0, 0, 142, 36, 1, 0, 0, 0,
		143, 144, 5, 42, 0, 0, 144, 38, 1, 0, 0, 0, 145, 146, 5, 47, 0, 0, 146,
		40, 1, 0, 0, 0, 147, 148, 5, 94, 0, 0, 148, 42, 1, 0, 0, 0, 149, 150, 5,
		61, 0, 0, 150, 151, 5, 61, 0, 0, 151, 44, 1, 0, 0, 0, 152, 153, 5, 33,
		0, 0, 153, 154, 5, 61, 0, 0, 154, 46, 1, 0, 0, 0, 155, 156, 5, 60, 0, 0,
		156, 48, 1, 0, 0, 0, 157, 158, 5, 62, 0, 0, 158, 50, 1, 0, 0, 0, 159, 160,
		5, 34, 0, 0, 160, 52, 1, 0, 0, 0, 161, 162, 5, 40, 0, 0, 162, 54, 1, 0,
		0, 0, 163, 164, 5, 41, 0, 0, 164, 56, 1, 0, 0, 0, 165, 170, 3, 47, 23,
		0, 166, 170, 3, 49, 24, 0, 167, 170, 3, 43, 21, 0, 168, 170, 3, 45, 22,
		0, 169, 165, 1, 0, 0, 0, 169, 166, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 169,
		168, 1, 0, 0, 0, 170, 58, 1, 0, 0, 0, 171, 173, 7, 4, 0, 0, 172, 174, 7,
		5, 0, 0, 173, 172, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 173, 1, 0, 0,
		0, 175, 176, 1, 0, 0, 0, 176, 60, 1, 0, 0, 0, 177, 179, 7, 6, 0, 0, 178,
		177, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 180, 181,
		1, 0, 0, 0, 181, 62, 1, 0, 0, 0, 182, 183, 7, 7, 0, 0, 183, 184, 1, 0,
		0, 0, 184, 185, 6, 31, 0, 0, 185, 64, 1, 0, 0, 0, 7, 0, 116, 125, 131,
		169, 175, 180, 1, 0, 1, 0,
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
	MyGrammarLexerSTRING       = 11
	MyGrammarLexerIDENTIFIER   = 12
	MyGrammarLexerINTEGER      = 13
	MyGrammarLexerEQ           = 14
	MyGrammarLexerCOLON        = 15
	MyGrammarLexerCOMMA        = 16
	MyGrammarLexerPLUS         = 17
	MyGrammarLexerMINUS        = 18
	MyGrammarLexerMULTI        = 19
	MyGrammarLexerDIVIDE       = 20
	MyGrammarLexerPOWERBY      = 21
	MyGrammarLexerEQUALBY      = 22
	MyGrammarLexerNOTEQUALBY   = 23
	MyGrammarLexerLT           = 24
	MyGrammarLexerGT           = 25
	MyGrammarLexerCOT          = 26
	MyGrammarLexerOPEN_PARAN   = 27
	MyGrammarLexerCLOSE_PARAN  = 28
	MyGrammarLexerCONDITION_OP = 29
	MyGrammarLexerNEWLINE      = 30
	MyGrammarLexerNEXT_PARAM   = 31
	MyGrammarLexerEMPTY        = 32
)
