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
		"", "'||'", "'&&'", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='", "'+'",
		"'-'", "'*'", "'/'", "'%'", "'^'", "'!'", "';'", "':'", "'='", "'('",
		"')'", "'{'", "'}'", "'begin'", "'end'", "'do'", "'then'", "'true'",
		"'false'", "'nil'", "'if'", "'else'", "'while'", "'for'", "'loop'",
		"'log'",
	}
	staticData.SymbolicNames = []string{
		"", "OR", "AND", "EQ", "NEQ", "GT", "LT", "GTEQ", "LTEQ", "PLUS", "MINUS",
		"MULT", "DIV", "MOD", "POW", "NOT", "SCOL", "COL", "ASSIGN", "OPAR",
		"CPAR", "OBRACE", "CBRACE", "BEGIN", "END", "DO", "THEN", "TRUE", "FALSE",
		"NIL", "IF", "ELSE", "WHILE", "FOR", "LOOP", "LOG", "ID", "INT", "FLOAT",
		"STRING", "COMMENT", "SPACE", "OTHER",
	}
	staticData.RuleNames = []string{
		"OR", "AND", "EQ", "NEQ", "GT", "LT", "GTEQ", "LTEQ", "PLUS", "MINUS",
		"MULT", "DIV", "MOD", "POW", "NOT", "SCOL", "COL", "ASSIGN", "OPAR",
		"CPAR", "OBRACE", "CBRACE", "BEGIN", "END", "DO", "THEN", "TRUE", "FALSE",
		"NIL", "IF", "ELSE", "WHILE", "FOR", "LOOP", "LOG", "ID", "INT", "FLOAT",
		"STRING", "COMMENT", "SPACE", "OTHER",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 42, 253, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1,
		19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 35, 1, 35, 5, 35, 198, 8, 35, 10, 35, 12, 35, 201,
		9, 35, 1, 36, 4, 36, 204, 8, 36, 11, 36, 12, 36, 205, 1, 37, 4, 37, 209,
		8, 37, 11, 37, 12, 37, 210, 1, 37, 1, 37, 5, 37, 215, 8, 37, 10, 37, 12,
		37, 218, 9, 37, 1, 37, 1, 37, 4, 37, 222, 8, 37, 11, 37, 12, 37, 223, 3,
		37, 226, 8, 37, 1, 38, 1, 38, 1, 38, 1, 38, 5, 38, 232, 8, 38, 10, 38,
		12, 38, 235, 9, 38, 1, 38, 1, 38, 1, 39, 1, 39, 5, 39, 241, 8, 39, 10,
		39, 12, 39, 244, 9, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41,
		1, 41, 0, 0, 42, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17,
		9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35,
		18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53,
		27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71,
		36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 1, 0, 6, 3, 0, 65,
		90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57,
		3, 0, 10, 10, 13, 13, 34, 34, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10, 13, 13,
		32, 32, 261, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7,
		1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0,
		15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0,
		0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0,
		0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0,
		0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1,
		0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53,
		1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0,
		61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0,
		0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0,
		0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0,
		0, 0, 1, 85, 1, 0, 0, 0, 3, 88, 1, 0, 0, 0, 5, 91, 1, 0, 0, 0, 7, 94, 1,
		0, 0, 0, 9, 97, 1, 0, 0, 0, 11, 99, 1, 0, 0, 0, 13, 101, 1, 0, 0, 0, 15,
		104, 1, 0, 0, 0, 17, 107, 1, 0, 0, 0, 19, 109, 1, 0, 0, 0, 21, 111, 1,
		0, 0, 0, 23, 113, 1, 0, 0, 0, 25, 115, 1, 0, 0, 0, 27, 117, 1, 0, 0, 0,
		29, 119, 1, 0, 0, 0, 31, 121, 1, 0, 0, 0, 33, 123, 1, 0, 0, 0, 35, 125,
		1, 0, 0, 0, 37, 127, 1, 0, 0, 0, 39, 129, 1, 0, 0, 0, 41, 131, 1, 0, 0,
		0, 43, 133, 1, 0, 0, 0, 45, 135, 1, 0, 0, 0, 47, 141, 1, 0, 0, 0, 49, 145,
		1, 0, 0, 0, 51, 148, 1, 0, 0, 0, 53, 153, 1, 0, 0, 0, 55, 158, 1, 0, 0,
		0, 57, 164, 1, 0, 0, 0, 59, 168, 1, 0, 0, 0, 61, 171, 1, 0, 0, 0, 63, 176,
		1, 0, 0, 0, 65, 182, 1, 0, 0, 0, 67, 186, 1, 0, 0, 0, 69, 191, 1, 0, 0,
		0, 71, 195, 1, 0, 0, 0, 73, 203, 1, 0, 0, 0, 75, 225, 1, 0, 0, 0, 77, 227,
		1, 0, 0, 0, 79, 238, 1, 0, 0, 0, 81, 247, 1, 0, 0, 0, 83, 251, 1, 0, 0,
		0, 85, 86, 5, 124, 0, 0, 86, 87, 5, 124, 0, 0, 87, 2, 1, 0, 0, 0, 88, 89,
		5, 38, 0, 0, 89, 90, 5, 38, 0, 0, 90, 4, 1, 0, 0, 0, 91, 92, 5, 61, 0,
		0, 92, 93, 5, 61, 0, 0, 93, 6, 1, 0, 0, 0, 94, 95, 5, 33, 0, 0, 95, 96,
		5, 61, 0, 0, 96, 8, 1, 0, 0, 0, 97, 98, 5, 62, 0, 0, 98, 10, 1, 0, 0, 0,
		99, 100, 5, 60, 0, 0, 100, 12, 1, 0, 0, 0, 101, 102, 5, 62, 0, 0, 102,
		103, 5, 61, 0, 0, 103, 14, 1, 0, 0, 0, 104, 105, 5, 60, 0, 0, 105, 106,
		5, 61, 0, 0, 106, 16, 1, 0, 0, 0, 107, 108, 5, 43, 0, 0, 108, 18, 1, 0,
		0, 0, 109, 110, 5, 45, 0, 0, 110, 20, 1, 0, 0, 0, 111, 112, 5, 42, 0, 0,
		112, 22, 1, 0, 0, 0, 113, 114, 5, 47, 0, 0, 114, 24, 1, 0, 0, 0, 115, 116,
		5, 37, 0, 0, 116, 26, 1, 0, 0, 0, 117, 118, 5, 94, 0, 0, 118, 28, 1, 0,
		0, 0, 119, 120, 5, 33, 0, 0, 120, 30, 1, 0, 0, 0, 121, 122, 5, 59, 0, 0,
		122, 32, 1, 0, 0, 0, 123, 124, 5, 58, 0, 0, 124, 34, 1, 0, 0, 0, 125, 126,
		5, 61, 0, 0, 126, 36, 1, 0, 0, 0, 127, 128, 5, 40, 0, 0, 128, 38, 1, 0,
		0, 0, 129, 130, 5, 41, 0, 0, 130, 40, 1, 0, 0, 0, 131, 132, 5, 123, 0,
		0, 132, 42, 1, 0, 0, 0, 133, 134, 5, 125, 0, 0, 134, 44, 1, 0, 0, 0, 135,
		136, 5, 98, 0, 0, 136, 137, 5, 101, 0, 0, 137, 138, 5, 103, 0, 0, 138,
		139, 5, 105, 0, 0, 139, 140, 5, 110, 0, 0, 140, 46, 1, 0, 0, 0, 141, 142,
		5, 101, 0, 0, 142, 143, 5, 110, 0, 0, 143, 144, 5, 100, 0, 0, 144, 48,
		1, 0, 0, 0, 145, 146, 5, 100, 0, 0, 146, 147, 5, 111, 0, 0, 147, 50, 1,
		0, 0, 0, 148, 149, 5, 116, 0, 0, 149, 150, 5, 104, 0, 0, 150, 151, 5, 101,
		0, 0, 151, 152, 5, 110, 0, 0, 152, 52, 1, 0, 0, 0, 153, 154, 5, 116, 0,
		0, 154, 155, 5, 114, 0, 0, 155, 156, 5, 117, 0, 0, 156, 157, 5, 101, 0,
		0, 157, 54, 1, 0, 0, 0, 158, 159, 5, 102, 0, 0, 159, 160, 5, 97, 0, 0,
		160, 161, 5, 108, 0, 0, 161, 162, 5, 115, 0, 0, 162, 163, 5, 101, 0, 0,
		163, 56, 1, 0, 0, 0, 164, 165, 5, 110, 0, 0, 165, 166, 5, 105, 0, 0, 166,
		167, 5, 108, 0, 0, 167, 58, 1, 0, 0, 0, 168, 169, 5, 105, 0, 0, 169, 170,
		5, 102, 0, 0, 170, 60, 1, 0, 0, 0, 171, 172, 5, 101, 0, 0, 172, 173, 5,
		108, 0, 0, 173, 174, 5, 115, 0, 0, 174, 175, 5, 101, 0, 0, 175, 62, 1,
		0, 0, 0, 176, 177, 5, 119, 0, 0, 177, 178, 5, 104, 0, 0, 178, 179, 5, 105,
		0, 0, 179, 180, 5, 108, 0, 0, 180, 181, 5, 101, 0, 0, 181, 64, 1, 0, 0,
		0, 182, 183, 5, 102, 0, 0, 183, 184, 5, 111, 0, 0, 184, 185, 5, 114, 0,
		0, 185, 66, 1, 0, 0, 0, 186, 187, 5, 108, 0, 0, 187, 188, 5, 111, 0, 0,
		188, 189, 5, 111, 0, 0, 189, 190, 5, 112, 0, 0, 190, 68, 1, 0, 0, 0, 191,
		192, 5, 108, 0, 0, 192, 193, 5, 111, 0, 0, 193, 194, 5, 103, 0, 0, 194,
		70, 1, 0, 0, 0, 195, 199, 7, 0, 0, 0, 196, 198, 7, 1, 0, 0, 197, 196, 1,
		0, 0, 0, 198, 201, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 199, 200, 1, 0, 0,
		0, 200, 72, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 202, 204, 7, 2, 0, 0, 203,
		202, 1, 0, 0, 0, 204, 205, 1, 0, 0, 0, 205, 203, 1, 0, 0, 0, 205, 206,
		1, 0, 0, 0, 206, 74, 1, 0, 0, 0, 207, 209, 7, 2, 0, 0, 208, 207, 1, 0,
		0, 0, 209, 210, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0,
		211, 212, 1, 0, 0, 0, 212, 216, 5, 46, 0, 0, 213, 215, 7, 2, 0, 0, 214,
		213, 1, 0, 0, 0, 215, 218, 1, 0, 0, 0, 216, 214, 1, 0, 0, 0, 216, 217,
		1, 0, 0, 0, 217, 226, 1, 0, 0, 0, 218, 216, 1, 0, 0, 0, 219, 221, 5, 46,
		0, 0, 220, 222, 7, 2, 0, 0, 221, 220, 1, 0, 0, 0, 222, 223, 1, 0, 0, 0,
		223, 221, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 226, 1, 0, 0, 0, 225,
		208, 1, 0, 0, 0, 225, 219, 1, 0, 0, 0, 226, 76, 1, 0, 0, 0, 227, 233, 5,
		34, 0, 0, 228, 232, 8, 3, 0, 0, 229, 230, 5, 34, 0, 0, 230, 232, 5, 34,
		0, 0, 231, 228, 1, 0, 0, 0, 231, 229, 1, 0, 0, 0, 232, 235, 1, 0, 0, 0,
		233, 231, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 236, 1, 0, 0, 0, 235,
		233, 1, 0, 0, 0, 236, 237, 5, 34, 0, 0, 237, 78, 1, 0, 0, 0, 238, 242,
		5, 35, 0, 0, 239, 241, 8, 4, 0, 0, 240, 239, 1, 0, 0, 0, 241, 244, 1, 0,
		0, 0, 242, 240, 1, 0, 0, 0, 242, 243, 1, 0, 0, 0, 243, 245, 1, 0, 0, 0,
		244, 242, 1, 0, 0, 0, 245, 246, 6, 39, 0, 0, 246, 80, 1, 0, 0, 0, 247,
		248, 7, 5, 0, 0, 248, 249, 1, 0, 0, 0, 249, 250, 6, 40, 0, 0, 250, 82,
		1, 0, 0, 0, 251, 252, 9, 0, 0, 0, 252, 84, 1, 0, 0, 0, 10, 0, 199, 205,
		210, 216, 223, 225, 231, 233, 242, 1, 6, 0, 0,
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
	MyGrammarLexerOR      = 1
	MyGrammarLexerAND     = 2
	MyGrammarLexerEQ      = 3
	MyGrammarLexerNEQ     = 4
	MyGrammarLexerGT      = 5
	MyGrammarLexerLT      = 6
	MyGrammarLexerGTEQ    = 7
	MyGrammarLexerLTEQ    = 8
	MyGrammarLexerPLUS    = 9
	MyGrammarLexerMINUS   = 10
	MyGrammarLexerMULT    = 11
	MyGrammarLexerDIV     = 12
	MyGrammarLexerMOD     = 13
	MyGrammarLexerPOW     = 14
	MyGrammarLexerNOT     = 15
	MyGrammarLexerSCOL    = 16
	MyGrammarLexerCOL     = 17
	MyGrammarLexerASSIGN  = 18
	MyGrammarLexerOPAR    = 19
	MyGrammarLexerCPAR    = 20
	MyGrammarLexerOBRACE  = 21
	MyGrammarLexerCBRACE  = 22
	MyGrammarLexerBEGIN   = 23
	MyGrammarLexerEND     = 24
	MyGrammarLexerDO      = 25
	MyGrammarLexerTHEN    = 26
	MyGrammarLexerTRUE    = 27
	MyGrammarLexerFALSE   = 28
	MyGrammarLexerNIL     = 29
	MyGrammarLexerIF      = 30
	MyGrammarLexerELSE    = 31
	MyGrammarLexerWHILE   = 32
	MyGrammarLexerFOR     = 33
	MyGrammarLexerLOOP    = 34
	MyGrammarLexerLOG     = 35
	MyGrammarLexerID      = 36
	MyGrammarLexerINT     = 37
	MyGrammarLexerFLOAT   = 38
	MyGrammarLexerSTRING  = 39
	MyGrammarLexerCOMMENT = 40
	MyGrammarLexerSPACE   = 41
	MyGrammarLexerOTHER   = 42
)
