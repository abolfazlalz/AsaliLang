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
		"", "'begin'", "'end'", "'if'", "'then'", "'else'", "'while'", "'do'",
		"'for'", "'loop'", "'print'", "", "", "", "'='", "':'", "','", "'+'",
		"'-'", "'*'", "'/'", "'^'", "'=='", "'!='", "'<'", "'>'", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "IDENTIFIER", "STRING_LITERAL",
		"INTEGER", "EQ", "COLON", "COMMA", "PLUS", "MINUS", "MULTI", "DIVIDE",
		"POWERBY", "EQUALBY", "NOTEQUALBY", "LT", "GT", "OPEN_PARAN", "CLOSE_PARAN",
		"CONDITION_OP", "NEWLINE", "EMPTY",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "IDENTIFIER", "STRING_LITERAL", "INTEGER", "EQ", "COLON", "COMMA",
		"PLUS", "MINUS", "MULTI", "DIVIDE", "POWERBY", "EQUALBY", "NOTEQUALBY",
		"LT", "GT", "OPEN_PARAN", "CLOSE_PARAN", "CONDITION_OP", "NEWLINE",
		"EMPTY",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 30, 174, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 5, 10,
		111, 8, 10, 10, 10, 12, 10, 114, 9, 10, 1, 11, 1, 11, 5, 11, 118, 8, 11,
		10, 11, 12, 11, 121, 9, 11, 1, 11, 1, 11, 1, 12, 4, 12, 126, 8, 12, 11,
		12, 12, 12, 127, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16,
		1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1,
		21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26,
		1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 164, 8, 27, 1, 28, 4, 28, 167,
		8, 28, 11, 28, 12, 28, 168, 1, 29, 1, 29, 1, 29, 1, 29, 0, 0, 30, 1, 1,
		3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23,
		12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41,
		21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59,
		30, 1, 0, 7, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95,
		97, 122, 1, 0, 34, 34, 1, 0, 42, 42, 1, 0, 48, 57, 2, 0, 10, 10, 13, 13,
		2, 0, 9, 9, 32, 32, 180, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0,
		0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1,
		0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21,
		1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0,
		29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0,
		0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0,
		0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0,
		0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1,
		0, 0, 0, 1, 61, 1, 0, 0, 0, 3, 67, 1, 0, 0, 0, 5, 71, 1, 0, 0, 0, 7, 74,
		1, 0, 0, 0, 9, 79, 1, 0, 0, 0, 11, 84, 1, 0, 0, 0, 13, 90, 1, 0, 0, 0,
		15, 93, 1, 0, 0, 0, 17, 97, 1, 0, 0, 0, 19, 102, 1, 0, 0, 0, 21, 108, 1,
		0, 0, 0, 23, 115, 1, 0, 0, 0, 25, 125, 1, 0, 0, 0, 27, 129, 1, 0, 0, 0,
		29, 131, 1, 0, 0, 0, 31, 133, 1, 0, 0, 0, 33, 135, 1, 0, 0, 0, 35, 137,
		1, 0, 0, 0, 37, 139, 1, 0, 0, 0, 39, 141, 1, 0, 0, 0, 41, 143, 1, 0, 0,
		0, 43, 145, 1, 0, 0, 0, 45, 148, 1, 0, 0, 0, 47, 151, 1, 0, 0, 0, 49, 153,
		1, 0, 0, 0, 51, 155, 1, 0, 0, 0, 53, 157, 1, 0, 0, 0, 55, 163, 1, 0, 0,
		0, 57, 166, 1, 0, 0, 0, 59, 170, 1, 0, 0, 0, 61, 62, 5, 98, 0, 0, 62, 63,
		5, 101, 0, 0, 63, 64, 5, 103, 0, 0, 64, 65, 5, 105, 0, 0, 65, 66, 5, 110,
		0, 0, 66, 2, 1, 0, 0, 0, 67, 68, 5, 101, 0, 0, 68, 69, 5, 110, 0, 0, 69,
		70, 5, 100, 0, 0, 70, 4, 1, 0, 0, 0, 71, 72, 5, 105, 0, 0, 72, 73, 5, 102,
		0, 0, 73, 6, 1, 0, 0, 0, 74, 75, 5, 116, 0, 0, 75, 76, 5, 104, 0, 0, 76,
		77, 5, 101, 0, 0, 77, 78, 5, 110, 0, 0, 78, 8, 1, 0, 0, 0, 79, 80, 5, 101,
		0, 0, 80, 81, 5, 108, 0, 0, 81, 82, 5, 115, 0, 0, 82, 83, 5, 101, 0, 0,
		83, 10, 1, 0, 0, 0, 84, 85, 5, 119, 0, 0, 85, 86, 5, 104, 0, 0, 86, 87,
		5, 105, 0, 0, 87, 88, 5, 108, 0, 0, 88, 89, 5, 101, 0, 0, 89, 12, 1, 0,
		0, 0, 90, 91, 5, 100, 0, 0, 91, 92, 5, 111, 0, 0, 92, 14, 1, 0, 0, 0, 93,
		94, 5, 102, 0, 0, 94, 95, 5, 111, 0, 0, 95, 96, 5, 114, 0, 0, 96, 16, 1,
		0, 0, 0, 97, 98, 5, 108, 0, 0, 98, 99, 5, 111, 0, 0, 99, 100, 5, 111, 0,
		0, 100, 101, 5, 112, 0, 0, 101, 18, 1, 0, 0, 0, 102, 103, 5, 112, 0, 0,
		103, 104, 5, 114, 0, 0, 104, 105, 5, 105, 0, 0, 105, 106, 5, 110, 0, 0,
		106, 107, 5, 116, 0, 0, 107, 20, 1, 0, 0, 0, 108, 112, 7, 0, 0, 0, 109,
		111, 7, 1, 0, 0, 110, 109, 1, 0, 0, 0, 111, 114, 1, 0, 0, 0, 112, 110,
		1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 22, 1, 0, 0, 0, 114, 112, 1, 0,
		0, 0, 115, 119, 7, 2, 0, 0, 116, 118, 7, 3, 0, 0, 117, 116, 1, 0, 0, 0,
		118, 121, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120,
		122, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0, 122, 123, 7, 2, 0, 0, 123, 24, 1,
		0, 0, 0, 124, 126, 7, 4, 0, 0, 125, 124, 1, 0, 0, 0, 126, 127, 1, 0, 0,
		0, 127, 125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 26, 1, 0, 0, 0, 129,
		130, 5, 61, 0, 0, 130, 28, 1, 0, 0, 0, 131, 132, 5, 58, 0, 0, 132, 30,
		1, 0, 0, 0, 133, 134, 5, 44, 0, 0, 134, 32, 1, 0, 0, 0, 135, 136, 5, 43,
		0, 0, 136, 34, 1, 0, 0, 0, 137, 138, 5, 45, 0, 0, 138, 36, 1, 0, 0, 0,
		139, 140, 5, 42, 0, 0, 140, 38, 1, 0, 0, 0, 141, 142, 5, 47, 0, 0, 142,
		40, 1, 0, 0, 0, 143, 144, 5, 94, 0, 0, 144, 42, 1, 0, 0, 0, 145, 146, 5,
		61, 0, 0, 146, 147, 5, 61, 0, 0, 147, 44, 1, 0, 0, 0, 148, 149, 5, 33,
		0, 0, 149, 150, 5, 61, 0, 0, 150, 46, 1, 0, 0, 0, 151, 152, 5, 60, 0, 0,
		152, 48, 1, 0, 0, 0, 153, 154, 5, 62, 0, 0, 154, 50, 1, 0, 0, 0, 155, 156,
		5, 40, 0, 0, 156, 52, 1, 0, 0, 0, 157, 158, 5, 41, 0, 0, 158, 54, 1, 0,
		0, 0, 159, 164, 3, 47, 23, 0, 160, 164, 3, 49, 24, 0, 161, 164, 3, 43,
		21, 0, 162, 164, 3, 45, 22, 0, 163, 159, 1, 0, 0, 0, 163, 160, 1, 0, 0,
		0, 163, 161, 1, 0, 0, 0, 163, 162, 1, 0, 0, 0, 164, 56, 1, 0, 0, 0, 165,
		167, 7, 5, 0, 0, 166, 165, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 166,
		1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 58, 1, 0, 0, 0, 170, 171, 7, 6,
		0, 0, 171, 172, 1, 0, 0, 0, 172, 173, 6, 29, 0, 0, 173, 60, 1, 0, 0, 0,
		6, 0, 112, 119, 127, 163, 168, 1, 0, 1, 0,
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
	MyGrammarLexerT__0           = 1
	MyGrammarLexerT__1           = 2
	MyGrammarLexerT__2           = 3
	MyGrammarLexerT__3           = 4
	MyGrammarLexerT__4           = 5
	MyGrammarLexerT__5           = 6
	MyGrammarLexerT__6           = 7
	MyGrammarLexerT__7           = 8
	MyGrammarLexerT__8           = 9
	MyGrammarLexerT__9           = 10
	MyGrammarLexerIDENTIFIER     = 11
	MyGrammarLexerSTRING_LITERAL = 12
	MyGrammarLexerINTEGER        = 13
	MyGrammarLexerEQ             = 14
	MyGrammarLexerCOLON          = 15
	MyGrammarLexerCOMMA          = 16
	MyGrammarLexerPLUS           = 17
	MyGrammarLexerMINUS          = 18
	MyGrammarLexerMULTI          = 19
	MyGrammarLexerDIVIDE         = 20
	MyGrammarLexerPOWERBY        = 21
	MyGrammarLexerEQUALBY        = 22
	MyGrammarLexerNOTEQUALBY     = 23
	MyGrammarLexerLT             = 24
	MyGrammarLexerGT             = 25
	MyGrammarLexerOPEN_PARAN     = 26
	MyGrammarLexerCLOSE_PARAN    = 27
	MyGrammarLexerCONDITION_OP   = 28
	MyGrammarLexerNEWLINE        = 29
	MyGrammarLexerEMPTY          = 30
)
