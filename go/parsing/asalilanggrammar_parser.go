// Code generated from AsaliLangGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parsing // AsaliLangGrammar
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

type AsaliLangGrammarParser struct {
	*antlr.BaseParser
}

var AsaliLangGrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func asalilanggrammarParserInit() {
	staticData := &AsaliLangGrammarParserStaticData
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
		"parse", "block", "stat", "assignment", "ifStat", "conditionBlock",
		"statBlock", "whileStat", "forStat", "loopStat", "methodCallStat", "methodCall",
		"inlineMethodCall", "methodCallArguments", "expr", "atom",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 41, 174, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 37, 8, 1, 10, 1, 12, 1, 40, 9, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 48, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 60, 8, 4, 10, 4, 12, 4, 63, 9, 4, 1,
		4, 1, 4, 3, 4, 67, 8, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		3, 6, 77, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 86, 8,
		6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11,
		1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 5,
		13, 121, 8, 13, 10, 13, 12, 13, 124, 9, 13, 3, 13, 126, 8, 13, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 135, 8, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 158,
		8, 14, 10, 14, 12, 14, 161, 9, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 3, 15, 172, 8, 15, 1, 15, 0, 1, 28, 16, 0, 2,
		4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 0, 6, 1, 0, 12, 14,
		1, 0, 10, 11, 1, 0, 6, 9, 1, 0, 4, 5, 1, 0, 37, 38, 1, 0, 28, 29, 186,
		0, 32, 1, 0, 0, 0, 2, 38, 1, 0, 0, 0, 4, 47, 1, 0, 0, 0, 6, 49, 1, 0, 0,
		0, 8, 54, 1, 0, 0, 0, 10, 68, 1, 0, 0, 0, 12, 85, 1, 0, 0, 0, 14, 87, 1,
		0, 0, 0, 16, 91, 1, 0, 0, 0, 18, 99, 1, 0, 0, 0, 20, 105, 1, 0, 0, 0, 22,
		108, 1, 0, 0, 0, 24, 111, 1, 0, 0, 0, 26, 125, 1, 0, 0, 0, 28, 134, 1,
		0, 0, 0, 30, 171, 1, 0, 0, 0, 32, 33, 3, 2, 1, 0, 33, 34, 5, 0, 0, 1, 34,
		1, 1, 0, 0, 0, 35, 37, 3, 4, 2, 0, 36, 35, 1, 0, 0, 0, 37, 40, 1, 0, 0,
		0, 38, 36, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 3, 1, 0, 0, 0, 40, 38, 1,
		0, 0, 0, 41, 48, 3, 6, 3, 0, 42, 48, 3, 8, 4, 0, 43, 48, 3, 14, 7, 0, 44,
		48, 3, 20, 10, 0, 45, 48, 3, 16, 8, 0, 46, 48, 3, 18, 9, 0, 47, 41, 1,
		0, 0, 0, 47, 42, 1, 0, 0, 0, 47, 43, 1, 0, 0, 0, 47, 44, 1, 0, 0, 0, 47,
		45, 1, 0, 0, 0, 47, 46, 1, 0, 0, 0, 48, 5, 1, 0, 0, 0, 49, 50, 5, 36, 0,
		0, 50, 51, 5, 19, 0, 0, 51, 52, 3, 28, 14, 0, 52, 53, 5, 17, 0, 0, 53,
		7, 1, 0, 0, 0, 54, 55, 5, 31, 0, 0, 55, 61, 3, 10, 5, 0, 56, 57, 5, 32,
		0, 0, 57, 58, 5, 31, 0, 0, 58, 60, 3, 10, 5, 0, 59, 56, 1, 0, 0, 0, 60,
		63, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 66, 1, 0, 0,
		0, 63, 61, 1, 0, 0, 0, 64, 65, 5, 32, 0, 0, 65, 67, 3, 12, 6, 0, 66, 64,
		1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 9, 1, 0, 0, 0, 68, 69, 3, 28, 14, 0,
		69, 70, 3, 12, 6, 0, 70, 11, 1, 0, 0, 0, 71, 72, 5, 22, 0, 0, 72, 73, 3,
		2, 1, 0, 73, 74, 5, 23, 0, 0, 74, 86, 1, 0, 0, 0, 75, 77, 5, 26, 0, 0,
		76, 75, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 79, 5,
		24, 0, 0, 79, 80, 3, 2, 1, 0, 80, 81, 5, 25, 0, 0, 81, 86, 1, 0, 0, 0,
		82, 86, 3, 4, 2, 0, 83, 84, 5, 27, 0, 0, 84, 86, 3, 2, 1, 0, 85, 71, 1,
		0, 0, 0, 85, 76, 1, 0, 0, 0, 85, 82, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 86,
		13, 1, 0, 0, 0, 87, 88, 5, 33, 0, 0, 88, 89, 3, 28, 14, 0, 89, 90, 3, 12,
		6, 0, 90, 15, 1, 0, 0, 0, 91, 92, 5, 34, 0, 0, 92, 93, 5, 36, 0, 0, 93,
		94, 5, 19, 0, 0, 94, 95, 3, 28, 14, 0, 95, 96, 5, 18, 0, 0, 96, 97, 3,
		28, 14, 0, 97, 98, 3, 12, 6, 0, 98, 17, 1, 0, 0, 0, 99, 100, 5, 35, 0,
		0, 100, 101, 5, 36, 0, 0, 101, 102, 5, 18, 0, 0, 102, 103, 3, 28, 14, 0,
		103, 104, 3, 12, 6, 0, 104, 19, 1, 0, 0, 0, 105, 106, 3, 22, 11, 0, 106,
		107, 5, 17, 0, 0, 107, 21, 1, 0, 0, 0, 108, 109, 5, 36, 0, 0, 109, 110,
		3, 26, 13, 0, 110, 23, 1, 0, 0, 0, 111, 112, 5, 36, 0, 0, 112, 113, 5,
		20, 0, 0, 113, 114, 3, 26, 13, 0, 114, 115, 5, 21, 0, 0, 115, 25, 1, 0,
		0, 0, 116, 126, 1, 0, 0, 0, 117, 122, 3, 28, 14, 0, 118, 119, 5, 1, 0,
		0, 119, 121, 3, 28, 14, 0, 120, 118, 1, 0, 0, 0, 121, 124, 1, 0, 0, 0,
		122, 120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 126, 1, 0, 0, 0, 124,
		122, 1, 0, 0, 0, 125, 116, 1, 0, 0, 0, 125, 117, 1, 0, 0, 0, 126, 27, 1,
		0, 0, 0, 127, 128, 6, 14, -1, 0, 128, 135, 3, 24, 12, 0, 129, 130, 5, 11,
		0, 0, 130, 135, 3, 28, 14, 9, 131, 132, 5, 16, 0, 0, 132, 135, 3, 28, 14,
		8, 133, 135, 3, 30, 15, 0, 134, 127, 1, 0, 0, 0, 134, 129, 1, 0, 0, 0,
		134, 131, 1, 0, 0, 0, 134, 133, 1, 0, 0, 0, 135, 159, 1, 0, 0, 0, 136,
		137, 10, 10, 0, 0, 137, 138, 5, 15, 0, 0, 138, 158, 3, 28, 14, 10, 139,
		140, 10, 7, 0, 0, 140, 141, 7, 0, 0, 0, 141, 158, 3, 28, 14, 8, 142, 143,
		10, 6, 0, 0, 143, 144, 7, 1, 0, 0, 144, 158, 3, 28, 14, 7, 145, 146, 10,
		5, 0, 0, 146, 147, 7, 2, 0, 0, 147, 158, 3, 28, 14, 6, 148, 149, 10, 4,
		0, 0, 149, 150, 7, 3, 0, 0, 150, 158, 3, 28, 14, 5, 151, 152, 10, 3, 0,
		0, 152, 153, 5, 3, 0, 0, 153, 158, 3, 28, 14, 4, 154, 155, 10, 2, 0, 0,
		155, 156, 5, 2, 0, 0, 156, 158, 3, 28, 14, 3, 157, 136, 1, 0, 0, 0, 157,
		139, 1, 0, 0, 0, 157, 142, 1, 0, 0, 0, 157, 145, 1, 0, 0, 0, 157, 148,
		1, 0, 0, 0, 157, 151, 1, 0, 0, 0, 157, 154, 1, 0, 0, 0, 158, 161, 1, 0,
		0, 0, 159, 157, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 29, 1, 0, 0, 0,
		161, 159, 1, 0, 0, 0, 162, 163, 5, 20, 0, 0, 163, 164, 3, 28, 14, 0, 164,
		165, 5, 21, 0, 0, 165, 172, 1, 0, 0, 0, 166, 172, 7, 4, 0, 0, 167, 172,
		7, 5, 0, 0, 168, 172, 5, 36, 0, 0, 169, 172, 5, 39, 0, 0, 170, 172, 5,
		30, 0, 0, 171, 162, 1, 0, 0, 0, 171, 166, 1, 0, 0, 0, 171, 167, 1, 0, 0,
		0, 171, 168, 1, 0, 0, 0, 171, 169, 1, 0, 0, 0, 171, 170, 1, 0, 0, 0, 172,
		31, 1, 0, 0, 0, 12, 38, 47, 61, 66, 76, 85, 122, 125, 134, 157, 159, 171,
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

// AsaliLangGrammarParserInit initializes any static state used to implement AsaliLangGrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAsaliLangGrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AsaliLangGrammarParserInit() {
	staticData := &AsaliLangGrammarParserStaticData
	staticData.once.Do(asalilanggrammarParserInit)
}

// NewAsaliLangGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAsaliLangGrammarParser(input antlr.TokenStream) *AsaliLangGrammarParser {
	AsaliLangGrammarParserInit()
	this := new(AsaliLangGrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &AsaliLangGrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "AsaliLangGrammar.g4"

	return this
}

// AsaliLangGrammarParser tokens.
const (
	AsaliLangGrammarParserEOF     = antlr.TokenEOF
	AsaliLangGrammarParserT__0    = 1
	AsaliLangGrammarParserOR      = 2
	AsaliLangGrammarParserAND     = 3
	AsaliLangGrammarParserEQ      = 4
	AsaliLangGrammarParserNEQ     = 5
	AsaliLangGrammarParserGT      = 6
	AsaliLangGrammarParserLT      = 7
	AsaliLangGrammarParserGTEQ    = 8
	AsaliLangGrammarParserLTEQ    = 9
	AsaliLangGrammarParserPLUS    = 10
	AsaliLangGrammarParserMINUS   = 11
	AsaliLangGrammarParserMULT    = 12
	AsaliLangGrammarParserDIV     = 13
	AsaliLangGrammarParserMOD     = 14
	AsaliLangGrammarParserPOW     = 15
	AsaliLangGrammarParserNOT     = 16
	AsaliLangGrammarParserSCOL    = 17
	AsaliLangGrammarParserCOL     = 18
	AsaliLangGrammarParserASSIGN  = 19
	AsaliLangGrammarParserOPAR    = 20
	AsaliLangGrammarParserCPAR    = 21
	AsaliLangGrammarParserOBRACE  = 22
	AsaliLangGrammarParserCBRACE  = 23
	AsaliLangGrammarParserBEGIN   = 24
	AsaliLangGrammarParserEND     = 25
	AsaliLangGrammarParserDO      = 26
	AsaliLangGrammarParserTHEN    = 27
	AsaliLangGrammarParserTRUE    = 28
	AsaliLangGrammarParserFALSE   = 29
	AsaliLangGrammarParserNIL     = 30
	AsaliLangGrammarParserIF      = 31
	AsaliLangGrammarParserELSE    = 32
	AsaliLangGrammarParserWHILE   = 33
	AsaliLangGrammarParserFOR     = 34
	AsaliLangGrammarParserLOOP    = 35
	AsaliLangGrammarParserID      = 36
	AsaliLangGrammarParserINT     = 37
	AsaliLangGrammarParserFLOAT   = 38
	AsaliLangGrammarParserSTRING  = 39
	AsaliLangGrammarParserCOMMENT = 40
	AsaliLangGrammarParserSPACE   = 41
)

// AsaliLangGrammarParser rules.
const (
	AsaliLangGrammarParserRULE_parse               = 0
	AsaliLangGrammarParserRULE_block               = 1
	AsaliLangGrammarParserRULE_stat                = 2
	AsaliLangGrammarParserRULE_assignment          = 3
	AsaliLangGrammarParserRULE_ifStat              = 4
	AsaliLangGrammarParserRULE_conditionBlock      = 5
	AsaliLangGrammarParserRULE_statBlock           = 6
	AsaliLangGrammarParserRULE_whileStat           = 7
	AsaliLangGrammarParserRULE_forStat             = 8
	AsaliLangGrammarParserRULE_loopStat            = 9
	AsaliLangGrammarParserRULE_methodCallStat      = 10
	AsaliLangGrammarParserRULE_methodCall          = 11
	AsaliLangGrammarParserRULE_inlineMethodCall    = 12
	AsaliLangGrammarParserRULE_methodCallArguments = 13
	AsaliLangGrammarParserRULE_expr                = 14
	AsaliLangGrammarParserRULE_atom                = 15
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	EOF() antlr.TerminalNode

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsaliLangGrammarParserRULE_parse
	return p
}

func InitEmptyParseContext(p *ParseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsaliLangGrammarParserRULE_parse
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsaliLangGrammarParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ParseContext) EOF() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserEOF, 0)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.EnterParse(s)
	}
}

func (s *ParseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.ExitParse(s)
	}
}

func (s *ParseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AsaliLangGrammarVisitor:
		return t.VisitParse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AsaliLangGrammarParser) Parse() (localctx IParseContext) {
	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AsaliLangGrammarParserRULE_parse)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.Block()
	}
	{
		p.SetState(33)
		p.Match(AsaliLangGrammarParserEOF)
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStat() []IStatContext
	Stat(i int) IStatContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsaliLangGrammarParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsaliLangGrammarParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsaliLangGrammarParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStat() []IStatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatContext); ok {
			len++
		}
	}

	tst := make([]IStatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatContext); ok {
			tst[i] = t.(IStatContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Stat(i int) IStatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatContext); ok {
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

	return t.(IStatContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AsaliLangGrammarVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AsaliLangGrammarParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AsaliLangGrammarParserRULE_block)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(38)
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
				p.SetState(35)
				p.Stat()
			}

		}
		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
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

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assignment() IAssignmentContext
	IfStat() IIfStatContext
	WhileStat() IWhileStatContext
	MethodCallStat() IMethodCallStatContext
	ForStat() IForStatContext
	LoopStat() ILoopStatContext

	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsaliLangGrammarParserRULE_stat
	return p
}

func InitEmptyStatContext(p *StatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsaliLangGrammarParserRULE_stat
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsaliLangGrammarParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatContext) IfStat() IIfStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatContext)
}

func (s *StatContext) WhileStat() IWhileStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStatContext)
}

func (s *StatContext) MethodCallStat() IMethodCallStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodCallStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodCallStatContext)
}

func (s *StatContext) ForStat() IForStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStatContext)
}

func (s *StatContext) LoopStat() ILoopStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopStatContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.EnterStat(s)
	}
}

func (s *StatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.ExitStat(s)
	}
}

func (s *StatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AsaliLangGrammarVisitor:
		return t.VisitStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AsaliLangGrammarParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AsaliLangGrammarParserRULE_stat)
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(41)
			p.Assignment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(42)
			p.IfStat()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(43)
			p.WhileStat()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(44)
			p.MethodCallStat()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(45)
			p.ForStat()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(46)
			p.LoopStat()
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

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expr() IExprContext
	SCOL() antlr.TerminalNode

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsaliLangGrammarParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsaliLangGrammarParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsaliLangGrammarParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) ID() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserID, 0)
}

func (s *AssignmentContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserASSIGN, 0)
}

func (s *AssignmentContext) Expr() IExprContext {
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

func (s *AssignmentContext) SCOL() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserSCOL, 0)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AsaliLangGrammarVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AsaliLangGrammarParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AsaliLangGrammarParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		p.Match(AsaliLangGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
		p.Match(AsaliLangGrammarParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(51)
		p.expr(0)
	}
	{
		p.SetState(52)
		p.Match(AsaliLangGrammarParserSCOL)
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

// IIfStatContext is an interface to support dynamic dispatch.
type IIfStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIF() []antlr.TerminalNode
	IF(i int) antlr.TerminalNode
	AllConditionBlock() []IConditionBlockContext
	ConditionBlock(i int) IConditionBlockContext
	AllELSE() []antlr.TerminalNode
	ELSE(i int) antlr.TerminalNode
	StatBlock() IStatBlockContext

	// IsIfStatContext differentiates from other interfaces.
	IsIfStatContext()
}

type IfStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatContext() *IfStatContext {
	var p = new(IfStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsaliLangGrammarParserRULE_ifStat
	return p
}

func InitEmptyIfStatContext(p *IfStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsaliLangGrammarParserRULE_ifStat
}

func (*IfStatContext) IsIfStatContext() {}

func NewIfStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatContext {
	var p = new(IfStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsaliLangGrammarParserRULE_ifStat

	return p
}

func (s *IfStatContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatContext) AllIF() []antlr.TerminalNode {
	return s.GetTokens(AsaliLangGrammarParserIF)
}

func (s *IfStatContext) IF(i int) antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserIF, i)
}

func (s *IfStatContext) AllConditionBlock() []IConditionBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionBlockContext); ok {
			len++
		}
	}

	tst := make([]IConditionBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionBlockContext); ok {
			tst[i] = t.(IConditionBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfStatContext) ConditionBlock(i int) IConditionBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionBlockContext); ok {
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

	return t.(IConditionBlockContext)
}

func (s *IfStatContext) AllELSE() []antlr.TerminalNode {
	return s.GetTokens(AsaliLangGrammarParserELSE)
}

func (s *IfStatContext) ELSE(i int) antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserELSE, i)
}

func (s *IfStatContext) StatBlock() IStatBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatBlockContext)
}

func (s *IfStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.EnterIfStat(s)
	}
}

func (s *IfStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.ExitIfStat(s)
	}
}

func (s *IfStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AsaliLangGrammarVisitor:
		return t.VisitIfStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AsaliLangGrammarParser) IfStat() (localctx IIfStatContext) {
	localctx = NewIfStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AsaliLangGrammarParserRULE_ifStat)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		p.Match(AsaliLangGrammarParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(55)
		p.ConditionBlock()
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(56)
				p.Match(AsaliLangGrammarParserELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(57)
				p.Match(AsaliLangGrammarParserIF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(58)
				p.ConditionBlock()
			}

		}
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(66)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(64)
			p.Match(AsaliLangGrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)
			p.StatBlock()
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

// IConditionBlockContext is an interface to support dynamic dispatch.
type IConditionBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	StatBlock() IStatBlockContext

	// IsConditionBlockContext differentiates from other interfaces.
	IsConditionBlockContext()
}

type ConditionBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionBlockContext() *ConditionBlockContext {
	var p = new(ConditionBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsaliLangGrammarParserRULE_conditionBlock
	return p
}

func InitEmptyConditionBlockContext(p *ConditionBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsaliLangGrammarParserRULE_conditionBlock
}

func (*ConditionBlockContext) IsConditionBlockContext() {}

func NewConditionBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionBlockContext {
	var p = new(ConditionBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsaliLangGrammarParserRULE_conditionBlock

	return p
}

func (s *ConditionBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionBlockContext) Expr() IExprContext {
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

func (s *ConditionBlockContext) StatBlock() IStatBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatBlockContext)
}

func (s *ConditionBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.EnterConditionBlock(s)
	}
}

func (s *ConditionBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.ExitConditionBlock(s)
	}
}

func (s *ConditionBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AsaliLangGrammarVisitor:
		return t.VisitConditionBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AsaliLangGrammarParser) ConditionBlock() (localctx IConditionBlockContext) {
	localctx = NewConditionBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AsaliLangGrammarParserRULE_conditionBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.expr(0)
	}
	{
		p.SetState(69)
		p.StatBlock()
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

// IStatBlockContext is an interface to support dynamic dispatch.
type IStatBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OBRACE() antlr.TerminalNode
	Block() IBlockContext
	CBRACE() antlr.TerminalNode
	BEGIN() antlr.TerminalNode
	END() antlr.TerminalNode
	DO() antlr.TerminalNode
	Stat() IStatContext
	THEN() antlr.TerminalNode

	// IsStatBlockContext differentiates from other interfaces.
	IsStatBlockContext()
}

type StatBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatBlockContext() *StatBlockContext {
	var p = new(StatBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsaliLangGrammarParserRULE_statBlock
	return p
}

func InitEmptyStatBlockContext(p *StatBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsaliLangGrammarParserRULE_statBlock
}

func (*StatBlockContext) IsStatBlockContext() {}

func NewStatBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatBlockContext {
	var p = new(StatBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsaliLangGrammarParserRULE_statBlock

	return p
}

func (s *StatBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *StatBlockContext) OBRACE() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserOBRACE, 0)
}

func (s *StatBlockContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatBlockContext) CBRACE() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserCBRACE, 0)
}

func (s *StatBlockContext) BEGIN() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserBEGIN, 0)
}

func (s *StatBlockContext) END() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserEND, 0)
}

func (s *StatBlockContext) DO() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserDO, 0)
}

func (s *StatBlockContext) Stat() IStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *StatBlockContext) THEN() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserTHEN, 0)
}

func (s *StatBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.EnterStatBlock(s)
	}
}

func (s *StatBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.ExitStatBlock(s)
	}
}

func (s *StatBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AsaliLangGrammarVisitor:
		return t.VisitStatBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AsaliLangGrammarParser) StatBlock() (localctx IStatBlockContext) {
	localctx = NewStatBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AsaliLangGrammarParserRULE_statBlock)
	var _la int

	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AsaliLangGrammarParserOBRACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(71)
			p.Match(AsaliLangGrammarParserOBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(72)
			p.Block()
		}
		{
			p.SetState(73)
			p.Match(AsaliLangGrammarParserCBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AsaliLangGrammarParserBEGIN, AsaliLangGrammarParserDO:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AsaliLangGrammarParserDO {
			{
				p.SetState(75)
				p.Match(AsaliLangGrammarParserDO)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(78)
			p.Match(AsaliLangGrammarParserBEGIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)
			p.Block()
		}
		{
			p.SetState(80)
			p.Match(AsaliLangGrammarParserEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AsaliLangGrammarParserIF, AsaliLangGrammarParserWHILE, AsaliLangGrammarParserFOR, AsaliLangGrammarParserLOOP, AsaliLangGrammarParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(82)
			p.Stat()
		}

	case AsaliLangGrammarParserTHEN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(83)
			p.Match(AsaliLangGrammarParserTHEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(84)
			p.Block()
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

// IWhileStatContext is an interface to support dynamic dispatch.
type IWhileStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHILE() antlr.TerminalNode
	Expr() IExprContext
	StatBlock() IStatBlockContext

	// IsWhileStatContext differentiates from other interfaces.
	IsWhileStatContext()
}

type WhileStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStatContext() *WhileStatContext {
	var p = new(WhileStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsaliLangGrammarParserRULE_whileStat
	return p
}

func InitEmptyWhileStatContext(p *WhileStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsaliLangGrammarParserRULE_whileStat
}

func (*WhileStatContext) IsWhileStatContext() {}

func NewWhileStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStatContext {
	var p = new(WhileStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsaliLangGrammarParserRULE_whileStat

	return p
}

func (s *WhileStatContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStatContext) WHILE() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserWHILE, 0)
}

func (s *WhileStatContext) Expr() IExprContext {
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

func (s *WhileStatContext) StatBlock() IStatBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatBlockContext)
}

func (s *WhileStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.EnterWhileStat(s)
	}
}

func (s *WhileStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.ExitWhileStat(s)
	}
}

func (s *WhileStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AsaliLangGrammarVisitor:
		return t.VisitWhileStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AsaliLangGrammarParser) WhileStat() (localctx IWhileStatContext) {
	localctx = NewWhileStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AsaliLangGrammarParserRULE_whileStat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(87)
		p.Match(AsaliLangGrammarParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(88)
		p.expr(0)
	}
	{
		p.SetState(89)
		p.StatBlock()
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

// IForStatContext is an interface to support dynamic dispatch.
type IForStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FOR() antlr.TerminalNode
	ID() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	COL() antlr.TerminalNode
	StatBlock() IStatBlockContext

	// IsForStatContext differentiates from other interfaces.
	IsForStatContext()
}

type ForStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStatContext() *ForStatContext {
	var p = new(ForStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsaliLangGrammarParserRULE_forStat
	return p
}

func InitEmptyForStatContext(p *ForStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsaliLangGrammarParserRULE_forStat
}

func (*ForStatContext) IsForStatContext() {}

func NewForStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStatContext {
	var p = new(ForStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsaliLangGrammarParserRULE_forStat

	return p
}

func (s *ForStatContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStatContext) FOR() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserFOR, 0)
}

func (s *ForStatContext) ID() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserID, 0)
}

func (s *ForStatContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserASSIGN, 0)
}

func (s *ForStatContext) AllExpr() []IExprContext {
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

func (s *ForStatContext) Expr(i int) IExprContext {
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

func (s *ForStatContext) COL() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserCOL, 0)
}

func (s *ForStatContext) StatBlock() IStatBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatBlockContext)
}

func (s *ForStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.EnterForStat(s)
	}
}

func (s *ForStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.ExitForStat(s)
	}
}

func (s *ForStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AsaliLangGrammarVisitor:
		return t.VisitForStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AsaliLangGrammarParser) ForStat() (localctx IForStatContext) {
	localctx = NewForStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AsaliLangGrammarParserRULE_forStat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Match(AsaliLangGrammarParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(92)
		p.Match(AsaliLangGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)
		p.Match(AsaliLangGrammarParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)
		p.expr(0)
	}
	{
		p.SetState(95)
		p.Match(AsaliLangGrammarParserCOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(96)
		p.expr(0)
	}
	{
		p.SetState(97)
		p.StatBlock()
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

// ILoopStatContext is an interface to support dynamic dispatch.
type ILoopStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LOOP() antlr.TerminalNode
	ID() antlr.TerminalNode
	COL() antlr.TerminalNode
	Expr() IExprContext
	StatBlock() IStatBlockContext

	// IsLoopStatContext differentiates from other interfaces.
	IsLoopStatContext()
}

type LoopStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopStatContext() *LoopStatContext {
	var p = new(LoopStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsaliLangGrammarParserRULE_loopStat
	return p
}

func InitEmptyLoopStatContext(p *LoopStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsaliLangGrammarParserRULE_loopStat
}

func (*LoopStatContext) IsLoopStatContext() {}

func NewLoopStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopStatContext {
	var p = new(LoopStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsaliLangGrammarParserRULE_loopStat

	return p
}

func (s *LoopStatContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopStatContext) LOOP() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserLOOP, 0)
}

func (s *LoopStatContext) ID() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserID, 0)
}

func (s *LoopStatContext) COL() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserCOL, 0)
}

func (s *LoopStatContext) Expr() IExprContext {
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

func (s *LoopStatContext) StatBlock() IStatBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatBlockContext)
}

func (s *LoopStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.EnterLoopStat(s)
	}
}

func (s *LoopStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.ExitLoopStat(s)
	}
}

func (s *LoopStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AsaliLangGrammarVisitor:
		return t.VisitLoopStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AsaliLangGrammarParser) LoopStat() (localctx ILoopStatContext) {
	localctx = NewLoopStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AsaliLangGrammarParserRULE_loopStat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Match(AsaliLangGrammarParserLOOP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(100)
		p.Match(AsaliLangGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(101)
		p.Match(AsaliLangGrammarParserCOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(102)
		p.expr(0)
	}
	{
		p.SetState(103)
		p.StatBlock()
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

// IMethodCallStatContext is an interface to support dynamic dispatch.
type IMethodCallStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MethodCall() IMethodCallContext
	SCOL() antlr.TerminalNode

	// IsMethodCallStatContext differentiates from other interfaces.
	IsMethodCallStatContext()
}

type MethodCallStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodCallStatContext() *MethodCallStatContext {
	var p = new(MethodCallStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsaliLangGrammarParserRULE_methodCallStat
	return p
}

func InitEmptyMethodCallStatContext(p *MethodCallStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsaliLangGrammarParserRULE_methodCallStat
}

func (*MethodCallStatContext) IsMethodCallStatContext() {}

func NewMethodCallStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodCallStatContext {
	var p = new(MethodCallStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsaliLangGrammarParserRULE_methodCallStat

	return p
}

func (s *MethodCallStatContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodCallStatContext) MethodCall() IMethodCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodCallContext)
}

func (s *MethodCallStatContext) SCOL() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserSCOL, 0)
}

func (s *MethodCallStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodCallStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodCallStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.EnterMethodCallStat(s)
	}
}

func (s *MethodCallStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.ExitMethodCallStat(s)
	}
}

func (s *MethodCallStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AsaliLangGrammarVisitor:
		return t.VisitMethodCallStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AsaliLangGrammarParser) MethodCallStat() (localctx IMethodCallStatContext) {
	localctx = NewMethodCallStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, AsaliLangGrammarParserRULE_methodCallStat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.MethodCall()
	}
	{
		p.SetState(106)
		p.Match(AsaliLangGrammarParserSCOL)
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

// IMethodCallContext is an interface to support dynamic dispatch.
type IMethodCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	MethodCallArguments() IMethodCallArgumentsContext

	// IsMethodCallContext differentiates from other interfaces.
	IsMethodCallContext()
}

type MethodCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodCallContext() *MethodCallContext {
	var p = new(MethodCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsaliLangGrammarParserRULE_methodCall
	return p
}

func InitEmptyMethodCallContext(p *MethodCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsaliLangGrammarParserRULE_methodCall
}

func (*MethodCallContext) IsMethodCallContext() {}

func NewMethodCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodCallContext {
	var p = new(MethodCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsaliLangGrammarParserRULE_methodCall

	return p
}

func (s *MethodCallContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodCallContext) ID() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserID, 0)
}

func (s *MethodCallContext) MethodCallArguments() IMethodCallArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodCallArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodCallArgumentsContext)
}

func (s *MethodCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.EnterMethodCall(s)
	}
}

func (s *MethodCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.ExitMethodCall(s)
	}
}

func (s *MethodCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AsaliLangGrammarVisitor:
		return t.VisitMethodCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AsaliLangGrammarParser) MethodCall() (localctx IMethodCallContext) {
	localctx = NewMethodCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, AsaliLangGrammarParserRULE_methodCall)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(AsaliLangGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
		p.MethodCallArguments()
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

// IInlineMethodCallContext is an interface to support dynamic dispatch.
type IInlineMethodCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	OPAR() antlr.TerminalNode
	MethodCallArguments() IMethodCallArgumentsContext
	CPAR() antlr.TerminalNode

	// IsInlineMethodCallContext differentiates from other interfaces.
	IsInlineMethodCallContext()
}

type InlineMethodCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlineMethodCallContext() *InlineMethodCallContext {
	var p = new(InlineMethodCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsaliLangGrammarParserRULE_inlineMethodCall
	return p
}

func InitEmptyInlineMethodCallContext(p *InlineMethodCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsaliLangGrammarParserRULE_inlineMethodCall
}

func (*InlineMethodCallContext) IsInlineMethodCallContext() {}

func NewInlineMethodCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineMethodCallContext {
	var p = new(InlineMethodCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsaliLangGrammarParserRULE_inlineMethodCall

	return p
}

func (s *InlineMethodCallContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineMethodCallContext) ID() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserID, 0)
}

func (s *InlineMethodCallContext) OPAR() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserOPAR, 0)
}

func (s *InlineMethodCallContext) MethodCallArguments() IMethodCallArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodCallArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodCallArgumentsContext)
}

func (s *InlineMethodCallContext) CPAR() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserCPAR, 0)
}

func (s *InlineMethodCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineMethodCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InlineMethodCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.EnterInlineMethodCall(s)
	}
}

func (s *InlineMethodCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.ExitInlineMethodCall(s)
	}
}

func (s *InlineMethodCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AsaliLangGrammarVisitor:
		return t.VisitInlineMethodCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AsaliLangGrammarParser) InlineMethodCall() (localctx IInlineMethodCallContext) {
	localctx = NewInlineMethodCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, AsaliLangGrammarParserRULE_inlineMethodCall)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(AsaliLangGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(112)
		p.Match(AsaliLangGrammarParserOPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
		p.MethodCallArguments()
	}
	{
		p.SetState(114)
		p.Match(AsaliLangGrammarParserCPAR)
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

// IMethodCallArgumentsContext is an interface to support dynamic dispatch.
type IMethodCallArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsMethodCallArgumentsContext differentiates from other interfaces.
	IsMethodCallArgumentsContext()
}

type MethodCallArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodCallArgumentsContext() *MethodCallArgumentsContext {
	var p = new(MethodCallArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsaliLangGrammarParserRULE_methodCallArguments
	return p
}

func InitEmptyMethodCallArgumentsContext(p *MethodCallArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsaliLangGrammarParserRULE_methodCallArguments
}

func (*MethodCallArgumentsContext) IsMethodCallArgumentsContext() {}

func NewMethodCallArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodCallArgumentsContext {
	var p = new(MethodCallArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsaliLangGrammarParserRULE_methodCallArguments

	return p
}

func (s *MethodCallArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodCallArgumentsContext) AllExpr() []IExprContext {
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

func (s *MethodCallArgumentsContext) Expr(i int) IExprContext {
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

func (s *MethodCallArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodCallArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodCallArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.EnterMethodCallArguments(s)
	}
}

func (s *MethodCallArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.ExitMethodCallArguments(s)
	}
}

func (s *MethodCallArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AsaliLangGrammarVisitor:
		return t.VisitMethodCallArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AsaliLangGrammarParser) MethodCallArguments() (localctx IMethodCallArgumentsContext) {
	localctx = NewMethodCallArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, AsaliLangGrammarParserRULE_methodCallArguments)
	var _la int

	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AsaliLangGrammarParserSCOL, AsaliLangGrammarParserCPAR:
		p.EnterOuterAlt(localctx, 1)

	case AsaliLangGrammarParserMINUS, AsaliLangGrammarParserNOT, AsaliLangGrammarParserOPAR, AsaliLangGrammarParserTRUE, AsaliLangGrammarParserFALSE, AsaliLangGrammarParserNIL, AsaliLangGrammarParserID, AsaliLangGrammarParserINT, AsaliLangGrammarParserFLOAT, AsaliLangGrammarParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
			p.expr(0)
		}
		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == AsaliLangGrammarParserT__0 {
			{
				p.SetState(118)
				p.Match(AsaliLangGrammarParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(119)
				p.expr(0)
			}

			p.SetState(124)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = AsaliLangGrammarParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsaliLangGrammarParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsaliLangGrammarParserRULE_expr

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

type MethodCallExprContext struct {
	ExprContext
}

func NewMethodCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MethodCallExprContext {
	var p = new(MethodCallExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MethodCallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodCallExprContext) InlineMethodCall() IInlineMethodCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInlineMethodCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInlineMethodCallContext)
}

func (s *MethodCallExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.EnterMethodCallExpr(s)
	}
}

func (s *MethodCallExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.ExitMethodCallExpr(s)
	}
}

func (s *MethodCallExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AsaliLangGrammarVisitor:
		return t.VisitMethodCallExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotExprContext struct {
	ExprContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserNOT, 0)
}

func (s *NotExprContext) Expr() IExprContext {
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

func (s *NotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.EnterNotExpr(s)
	}
}

func (s *NotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.ExitNotExpr(s)
	}
}

func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AsaliLangGrammarVisitor:
		return t.VisitNotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryMinusExprContext struct {
	ExprContext
}

func NewUnaryMinusExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryMinusExprContext {
	var p = new(UnaryMinusExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *UnaryMinusExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryMinusExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserMINUS, 0)
}

func (s *UnaryMinusExprContext) Expr() IExprContext {
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

func (s *UnaryMinusExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.EnterUnaryMinusExpr(s)
	}
}

func (s *UnaryMinusExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.ExitUnaryMinusExpr(s)
	}
}

func (s *UnaryMinusExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AsaliLangGrammarVisitor:
		return t.VisitUnaryMinusExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type MultiplicationExprContext struct {
	ExprContext
	op antlr.Token
}

func NewMultiplicationExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicationExprContext {
	var p = new(MultiplicationExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MultiplicationExprContext) GetOp() antlr.Token { return s.op }

func (s *MultiplicationExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *MultiplicationExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicationExprContext) AllExpr() []IExprContext {
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

func (s *MultiplicationExprContext) Expr(i int) IExprContext {
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

func (s *MultiplicationExprContext) MULT() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserMULT, 0)
}

func (s *MultiplicationExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserDIV, 0)
}

func (s *MultiplicationExprContext) MOD() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserMOD, 0)
}

func (s *MultiplicationExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.EnterMultiplicationExpr(s)
	}
}

func (s *MultiplicationExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.ExitMultiplicationExpr(s)
	}
}

func (s *MultiplicationExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AsaliLangGrammarVisitor:
		return t.VisitMultiplicationExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AtomExprContext struct {
	ExprContext
}

func NewAtomExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtomExprContext {
	var p = new(AtomExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AtomExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomExprContext) Atom() IAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *AtomExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.EnterAtomExpr(s)
	}
}

func (s *AtomExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.ExitAtomExpr(s)
	}
}

func (s *AtomExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AsaliLangGrammarVisitor:
		return t.VisitAtomExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrExprContext struct {
	ExprContext
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) AllExpr() []IExprContext {
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

func (s *OrExprContext) Expr(i int) IExprContext {
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

func (s *OrExprContext) OR() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserOR, 0)
}

func (s *OrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.EnterOrExpr(s)
	}
}

func (s *OrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.ExitOrExpr(s)
	}
}

func (s *OrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AsaliLangGrammarVisitor:
		return t.VisitOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AdditiveExprContext struct {
	ExprContext
	op antlr.Token
}

func NewAdditiveExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditiveExprContext {
	var p = new(AdditiveExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AdditiveExprContext) GetOp() antlr.Token { return s.op }

func (s *AdditiveExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *AdditiveExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExprContext) AllExpr() []IExprContext {
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

func (s *AdditiveExprContext) Expr(i int) IExprContext {
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

func (s *AdditiveExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserPLUS, 0)
}

func (s *AdditiveExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserMINUS, 0)
}

func (s *AdditiveExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.EnterAdditiveExpr(s)
	}
}

func (s *AdditiveExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.ExitAdditiveExpr(s)
	}
}

func (s *AdditiveExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AsaliLangGrammarVisitor:
		return t.VisitAdditiveExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type PowExprContext struct {
	ExprContext
}

func NewPowExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowExprContext {
	var p = new(PowExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *PowExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowExprContext) AllExpr() []IExprContext {
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

func (s *PowExprContext) Expr(i int) IExprContext {
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

func (s *PowExprContext) POW() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserPOW, 0)
}

func (s *PowExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.EnterPowExpr(s)
	}
}

func (s *PowExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.ExitPowExpr(s)
	}
}

func (s *PowExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AsaliLangGrammarVisitor:
		return t.VisitPowExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type RelationalExprContext struct {
	ExprContext
	op antlr.Token
}

func NewRelationalExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelationalExprContext {
	var p = new(RelationalExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *RelationalExprContext) GetOp() antlr.Token { return s.op }

func (s *RelationalExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *RelationalExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExprContext) AllExpr() []IExprContext {
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

func (s *RelationalExprContext) Expr(i int) IExprContext {
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

func (s *RelationalExprContext) LTEQ() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserLTEQ, 0)
}

func (s *RelationalExprContext) GTEQ() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserGTEQ, 0)
}

func (s *RelationalExprContext) LT() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserLT, 0)
}

func (s *RelationalExprContext) GT() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserGT, 0)
}

func (s *RelationalExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.EnterRelationalExpr(s)
	}
}

func (s *RelationalExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.ExitRelationalExpr(s)
	}
}

func (s *RelationalExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AsaliLangGrammarVisitor:
		return t.VisitRelationalExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqualityExprContext struct {
	ExprContext
	op antlr.Token
}

func NewEqualityExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualityExprContext {
	var p = new(EqualityExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EqualityExprContext) GetOp() antlr.Token { return s.op }

func (s *EqualityExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *EqualityExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExprContext) AllExpr() []IExprContext {
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

func (s *EqualityExprContext) Expr(i int) IExprContext {
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

func (s *EqualityExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserEQ, 0)
}

func (s *EqualityExprContext) NEQ() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserNEQ, 0)
}

func (s *EqualityExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.EnterEqualityExpr(s)
	}
}

func (s *EqualityExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.ExitEqualityExpr(s)
	}
}

func (s *EqualityExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AsaliLangGrammarVisitor:
		return t.VisitEqualityExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndExprContext struct {
	ExprContext
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) AllExpr() []IExprContext {
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

func (s *AndExprContext) Expr(i int) IExprContext {
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

func (s *AndExprContext) AND() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserAND, 0)
}

func (s *AndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.EnterAndExpr(s)
	}
}

func (s *AndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.ExitAndExpr(s)
	}
}

func (s *AndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AsaliLangGrammarVisitor:
		return t.VisitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AsaliLangGrammarParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *AsaliLangGrammarParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 28
	p.EnterRecursionRule(localctx, 28, AsaliLangGrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewMethodCallExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(128)
			p.InlineMethodCall()
		}

	case 2:
		localctx = NewUnaryMinusExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(129)
			p.Match(AsaliLangGrammarParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.expr(9)
		}

	case 3:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(131)
			p.Match(AsaliLangGrammarParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)
			p.expr(8)
		}

	case 4:
		localctx = NewAtomExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(133)
			p.Atom()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(157)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPowExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AsaliLangGrammarParserRULE_expr)
				p.SetState(136)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(137)
					p.Match(AsaliLangGrammarParserPOW)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(138)
					p.expr(10)
				}

			case 2:
				localctx = NewMultiplicationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AsaliLangGrammarParserRULE_expr)
				p.SetState(139)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(140)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MultiplicationExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&28672) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MultiplicationExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(141)
					p.expr(8)
				}

			case 3:
				localctx = NewAdditiveExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AsaliLangGrammarParserRULE_expr)
				p.SetState(142)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(143)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AdditiveExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == AsaliLangGrammarParserPLUS || _la == AsaliLangGrammarParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AdditiveExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(144)
					p.expr(7)
				}

			case 4:
				localctx = NewRelationalExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AsaliLangGrammarParserRULE_expr)
				p.SetState(145)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(146)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*RelationalExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&960) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*RelationalExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(147)
					p.expr(6)
				}

			case 5:
				localctx = NewEqualityExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AsaliLangGrammarParserRULE_expr)
				p.SetState(148)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(149)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EqualityExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == AsaliLangGrammarParserEQ || _la == AsaliLangGrammarParserNEQ) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EqualityExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(150)
					p.expr(5)
				}

			case 6:
				localctx = NewAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AsaliLangGrammarParserRULE_expr)
				p.SetState(151)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(152)
					p.Match(AsaliLangGrammarParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(153)
					p.expr(4)
				}

			case 7:
				localctx = NewOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AsaliLangGrammarParserRULE_expr)
				p.SetState(154)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(155)
					p.Match(AsaliLangGrammarParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(156)
					p.expr(3)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
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

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsaliLangGrammarParserRULE_atom
	return p
}

func InitEmptyAtomContext(p *AtomContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AsaliLangGrammarParserRULE_atom
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AsaliLangGrammarParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) CopyAll(ctx *AtomContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParExprContext struct {
	AtomContext
}

func NewParExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParExprContext {
	var p = new(ParExprContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *ParExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParExprContext) OPAR() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserOPAR, 0)
}

func (s *ParExprContext) Expr() IExprContext {
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

func (s *ParExprContext) CPAR() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserCPAR, 0)
}

func (s *ParExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.EnterParExpr(s)
	}
}

func (s *ParExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.ExitParExpr(s)
	}
}

func (s *ParExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AsaliLangGrammarVisitor:
		return t.VisitParExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BooleanAtomContext struct {
	AtomContext
}

func NewBooleanAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanAtomContext {
	var p = new(BooleanAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *BooleanAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanAtomContext) TRUE() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserTRUE, 0)
}

func (s *BooleanAtomContext) FALSE() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserFALSE, 0)
}

func (s *BooleanAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.EnterBooleanAtom(s)
	}
}

func (s *BooleanAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.ExitBooleanAtom(s)
	}
}

func (s *BooleanAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AsaliLangGrammarVisitor:
		return t.VisitBooleanAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdAtomContext struct {
	AtomContext
}

func NewIdAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdAtomContext {
	var p = new(IdAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *IdAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdAtomContext) ID() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserID, 0)
}

func (s *IdAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.EnterIdAtom(s)
	}
}

func (s *IdAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.ExitIdAtom(s)
	}
}

func (s *IdAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AsaliLangGrammarVisitor:
		return t.VisitIdAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringAtomContext struct {
	AtomContext
}

func NewStringAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringAtomContext {
	var p = new(StringAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *StringAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringAtomContext) STRING() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserSTRING, 0)
}

func (s *StringAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.EnterStringAtom(s)
	}
}

func (s *StringAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.ExitStringAtom(s)
	}
}

func (s *StringAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AsaliLangGrammarVisitor:
		return t.VisitStringAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type NilAtomContext struct {
	AtomContext
}

func NewNilAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NilAtomContext {
	var p = new(NilAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *NilAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilAtomContext) NIL() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserNIL, 0)
}

func (s *NilAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.EnterNilAtom(s)
	}
}

func (s *NilAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.ExitNilAtom(s)
	}
}

func (s *NilAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AsaliLangGrammarVisitor:
		return t.VisitNilAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberAtomContext struct {
	AtomContext
}

func NewNumberAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberAtomContext {
	var p = new(NumberAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *NumberAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberAtomContext) INT() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserINT, 0)
}

func (s *NumberAtomContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(AsaliLangGrammarParserFLOAT, 0)
}

func (s *NumberAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.EnterNumberAtom(s)
	}
}

func (s *NumberAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AsaliLangGrammarListener); ok {
		listenerT.ExitNumberAtom(s)
	}
}

func (s *NumberAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AsaliLangGrammarVisitor:
		return t.VisitNumberAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AsaliLangGrammarParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, AsaliLangGrammarParserRULE_atom)
	var _la int

	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AsaliLangGrammarParserOPAR:
		localctx = NewParExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(162)
			p.Match(AsaliLangGrammarParserOPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(163)
			p.expr(0)
		}
		{
			p.SetState(164)
			p.Match(AsaliLangGrammarParserCPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AsaliLangGrammarParserINT, AsaliLangGrammarParserFLOAT:
		localctx = NewNumberAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(166)
			_la = p.GetTokenStream().LA(1)

			if !(_la == AsaliLangGrammarParserINT || _la == AsaliLangGrammarParserFLOAT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case AsaliLangGrammarParserTRUE, AsaliLangGrammarParserFALSE:
		localctx = NewBooleanAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(167)
			_la = p.GetTokenStream().LA(1)

			if !(_la == AsaliLangGrammarParserTRUE || _la == AsaliLangGrammarParserFALSE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case AsaliLangGrammarParserID:
		localctx = NewIdAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(168)
			p.Match(AsaliLangGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AsaliLangGrammarParserSTRING:
		localctx = NewStringAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(169)
			p.Match(AsaliLangGrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AsaliLangGrammarParserNIL:
		localctx = NewNilAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(170)
			p.Match(AsaliLangGrammarParserNIL)
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

func (p *AsaliLangGrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 14:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *AsaliLangGrammarParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
