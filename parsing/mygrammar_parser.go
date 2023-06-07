// Code generated from MyGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parsing // MyGrammar
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

type MyGrammarParser struct {
	*antlr.BaseParser
}

var MyGrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func mygrammarParserInit() {
	staticData := &MyGrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "'||'", "'&&'", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='", "'+'",
		"'-'", "'*'", "'/'", "'%'", "'^'", "'!'", "';'", "':'", "'='", "'('",
		"')'", "'{'", "'}'", "'begin'", "'end'", "'do'", "'true'", "'false'",
		"'nil'", "'if'", "'else'", "'while'", "'for'", "'loop'", "'log'",
	}
	staticData.SymbolicNames = []string{
		"", "OR", "AND", "EQ", "NEQ", "GT", "LT", "GTEQ", "LTEQ", "PLUS", "MINUS",
		"MULT", "DIV", "MOD", "POW", "NOT", "SCOL", "COL", "ASSIGN", "OPAR",
		"CPAR", "OBRACE", "CBRACE", "BEGIN", "END", "DO", "TRUE", "FALSE", "NIL",
		"IF", "ELSE", "WHILE", "FOR", "LOOP", "LOG", "ID", "INT", "FLOAT", "STRING",
		"COMMENT", "SPACE", "OTHER",
	}
	staticData.RuleNames = []string{
		"parse", "block", "stat", "assignment", "if_stat", "condition_block",
		"stat_block", "whileStat", "forStat", "loopStat", "log", "expr", "atom",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 41, 149, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 31, 8, 1,
		10, 1, 12, 1, 34, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		3, 2, 44, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 5, 4, 56, 8, 4, 10, 4, 12, 4, 59, 9, 4, 1, 4, 1, 4, 3, 4, 63, 8, 4,
		1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 73, 8, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 3, 6, 80, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3,
		11, 110, 8, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 5, 11, 133, 8, 11, 10, 11, 12, 11, 136, 9, 11, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 147, 8,
		12, 1, 12, 0, 1, 22, 13, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		0, 6, 1, 0, 11, 13, 1, 0, 9, 10, 1, 0, 5, 8, 1, 0, 3, 4, 1, 0, 36, 37,
		1, 0, 26, 27, 161, 0, 26, 1, 0, 0, 0, 2, 32, 1, 0, 0, 0, 4, 43, 1, 0, 0,
		0, 6, 45, 1, 0, 0, 0, 8, 50, 1, 0, 0, 0, 10, 64, 1, 0, 0, 0, 12, 79, 1,
		0, 0, 0, 14, 81, 1, 0, 0, 0, 16, 85, 1, 0, 0, 0, 18, 93, 1, 0, 0, 0, 20,
		99, 1, 0, 0, 0, 22, 109, 1, 0, 0, 0, 24, 146, 1, 0, 0, 0, 26, 27, 3, 2,
		1, 0, 27, 28, 5, 0, 0, 1, 28, 1, 1, 0, 0, 0, 29, 31, 3, 4, 2, 0, 30, 29,
		1, 0, 0, 0, 31, 34, 1, 0, 0, 0, 32, 30, 1, 0, 0, 0, 32, 33, 1, 0, 0, 0,
		33, 3, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 35, 44, 3, 6, 3, 0, 36, 44, 3, 8,
		4, 0, 37, 44, 3, 14, 7, 0, 38, 44, 3, 20, 10, 0, 39, 44, 3, 16, 8, 0, 40,
		44, 3, 18, 9, 0, 41, 42, 5, 41, 0, 0, 42, 44, 6, 2, -1, 0, 43, 35, 1, 0,
		0, 0, 43, 36, 1, 0, 0, 0, 43, 37, 1, 0, 0, 0, 43, 38, 1, 0, 0, 0, 43, 39,
		1, 0, 0, 0, 43, 40, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0, 44, 5, 1, 0, 0, 0,
		45, 46, 5, 35, 0, 0, 46, 47, 5, 18, 0, 0, 47, 48, 3, 22, 11, 0, 48, 49,
		5, 16, 0, 0, 49, 7, 1, 0, 0, 0, 50, 51, 5, 29, 0, 0, 51, 57, 3, 10, 5,
		0, 52, 53, 5, 30, 0, 0, 53, 54, 5, 29, 0, 0, 54, 56, 3, 10, 5, 0, 55, 52,
		1, 0, 0, 0, 56, 59, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0,
		58, 62, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 60, 61, 5, 30, 0, 0, 61, 63, 3,
		12, 6, 0, 62, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 9, 1, 0, 0, 0, 64,
		65, 3, 22, 11, 0, 65, 66, 3, 12, 6, 0, 66, 11, 1, 0, 0, 0, 67, 68, 5, 21,
		0, 0, 68, 69, 3, 2, 1, 0, 69, 70, 5, 22, 0, 0, 70, 80, 1, 0, 0, 0, 71,
		73, 5, 25, 0, 0, 72, 71, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 74, 1, 0,
		0, 0, 74, 75, 5, 23, 0, 0, 75, 76, 3, 2, 1, 0, 76, 77, 5, 24, 0, 0, 77,
		80, 1, 0, 0, 0, 78, 80, 3, 4, 2, 0, 79, 67, 1, 0, 0, 0, 79, 72, 1, 0, 0,
		0, 79, 78, 1, 0, 0, 0, 80, 13, 1, 0, 0, 0, 81, 82, 5, 31, 0, 0, 82, 83,
		3, 22, 11, 0, 83, 84, 3, 12, 6, 0, 84, 15, 1, 0, 0, 0, 85, 86, 5, 32, 0,
		0, 86, 87, 5, 35, 0, 0, 87, 88, 5, 18, 0, 0, 88, 89, 3, 22, 11, 0, 89,
		90, 5, 17, 0, 0, 90, 91, 3, 22, 11, 0, 91, 92, 3, 12, 6, 0, 92, 17, 1,
		0, 0, 0, 93, 94, 5, 33, 0, 0, 94, 95, 5, 35, 0, 0, 95, 96, 5, 17, 0, 0,
		96, 97, 3, 22, 11, 0, 97, 98, 3, 12, 6, 0, 98, 19, 1, 0, 0, 0, 99, 100,
		5, 34, 0, 0, 100, 101, 3, 22, 11, 0, 101, 102, 5, 16, 0, 0, 102, 21, 1,
		0, 0, 0, 103, 104, 6, 11, -1, 0, 104, 105, 5, 10, 0, 0, 105, 110, 3, 22,
		11, 9, 106, 107, 5, 15, 0, 0, 107, 110, 3, 22, 11, 8, 108, 110, 3, 24,
		12, 0, 109, 103, 1, 0, 0, 0, 109, 106, 1, 0, 0, 0, 109, 108, 1, 0, 0, 0,
		110, 134, 1, 0, 0, 0, 111, 112, 10, 10, 0, 0, 112, 113, 5, 14, 0, 0, 113,
		133, 3, 22, 11, 10, 114, 115, 10, 7, 0, 0, 115, 116, 7, 0, 0, 0, 116, 133,
		3, 22, 11, 8, 117, 118, 10, 6, 0, 0, 118, 119, 7, 1, 0, 0, 119, 133, 3,
		22, 11, 7, 120, 121, 10, 5, 0, 0, 121, 122, 7, 2, 0, 0, 122, 133, 3, 22,
		11, 6, 123, 124, 10, 4, 0, 0, 124, 125, 7, 3, 0, 0, 125, 133, 3, 22, 11,
		5, 126, 127, 10, 3, 0, 0, 127, 128, 5, 2, 0, 0, 128, 133, 3, 22, 11, 4,
		129, 130, 10, 2, 0, 0, 130, 131, 5, 1, 0, 0, 131, 133, 3, 22, 11, 3, 132,
		111, 1, 0, 0, 0, 132, 114, 1, 0, 0, 0, 132, 117, 1, 0, 0, 0, 132, 120,
		1, 0, 0, 0, 132, 123, 1, 0, 0, 0, 132, 126, 1, 0, 0, 0, 132, 129, 1, 0,
		0, 0, 133, 136, 1, 0, 0, 0, 134, 132, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0,
		135, 23, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 137, 138, 5, 19, 0, 0, 138,
		139, 3, 22, 11, 0, 139, 140, 5, 20, 0, 0, 140, 147, 1, 0, 0, 0, 141, 147,
		7, 4, 0, 0, 142, 147, 7, 5, 0, 0, 143, 147, 5, 35, 0, 0, 144, 147, 5, 38,
		0, 0, 145, 147, 5, 28, 0, 0, 146, 137, 1, 0, 0, 0, 146, 141, 1, 0, 0, 0,
		146, 142, 1, 0, 0, 0, 146, 143, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 146,
		145, 1, 0, 0, 0, 147, 25, 1, 0, 0, 0, 10, 32, 43, 57, 62, 72, 79, 109,
		132, 134, 146,
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

// MyGrammarParserInit initializes any static state used to implement MyGrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMyGrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MyGrammarParserInit() {
	staticData := &MyGrammarParserStaticData
	staticData.once.Do(mygrammarParserInit)
}

// NewMyGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMyGrammarParser(input antlr.TokenStream) *MyGrammarParser {
	MyGrammarParserInit()
	this := new(MyGrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &MyGrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "MyGrammar.g4"

	return this
}

// MyGrammarParser tokens.
const (
	MyGrammarParserEOF     = antlr.TokenEOF
	MyGrammarParserOR      = 1
	MyGrammarParserAND     = 2
	MyGrammarParserEQ      = 3
	MyGrammarParserNEQ     = 4
	MyGrammarParserGT      = 5
	MyGrammarParserLT      = 6
	MyGrammarParserGTEQ    = 7
	MyGrammarParserLTEQ    = 8
	MyGrammarParserPLUS    = 9
	MyGrammarParserMINUS   = 10
	MyGrammarParserMULT    = 11
	MyGrammarParserDIV     = 12
	MyGrammarParserMOD     = 13
	MyGrammarParserPOW     = 14
	MyGrammarParserNOT     = 15
	MyGrammarParserSCOL    = 16
	MyGrammarParserCOL     = 17
	MyGrammarParserASSIGN  = 18
	MyGrammarParserOPAR    = 19
	MyGrammarParserCPAR    = 20
	MyGrammarParserOBRACE  = 21
	MyGrammarParserCBRACE  = 22
	MyGrammarParserBEGIN   = 23
	MyGrammarParserEND     = 24
	MyGrammarParserDO      = 25
	MyGrammarParserTRUE    = 26
	MyGrammarParserFALSE   = 27
	MyGrammarParserNIL     = 28
	MyGrammarParserIF      = 29
	MyGrammarParserELSE    = 30
	MyGrammarParserWHILE   = 31
	MyGrammarParserFOR     = 32
	MyGrammarParserLOOP    = 33
	MyGrammarParserLOG     = 34
	MyGrammarParserID      = 35
	MyGrammarParserINT     = 36
	MyGrammarParserFLOAT   = 37
	MyGrammarParserSTRING  = 38
	MyGrammarParserCOMMENT = 39
	MyGrammarParserSPACE   = 40
	MyGrammarParserOTHER   = 41
)

// MyGrammarParser rules.
const (
	MyGrammarParserRULE_parse           = 0
	MyGrammarParserRULE_block           = 1
	MyGrammarParserRULE_stat            = 2
	MyGrammarParserRULE_assignment      = 3
	MyGrammarParserRULE_if_stat         = 4
	MyGrammarParserRULE_condition_block = 5
	MyGrammarParserRULE_stat_block      = 6
	MyGrammarParserRULE_whileStat       = 7
	MyGrammarParserRULE_forStat         = 8
	MyGrammarParserRULE_loopStat        = 9
	MyGrammarParserRULE_log             = 10
	MyGrammarParserRULE_expr            = 11
	MyGrammarParserRULE_atom            = 12
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
	p.RuleIndex = MyGrammarParserRULE_parse
	return p
}

func InitEmptyParseContext(p *ParseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_parse
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_parse

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
	return s.GetToken(MyGrammarParserEOF, 0)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterParse(s)
	}
}

func (s *ParseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitParse(s)
	}
}

func (s *ParseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitParse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) Parse() (localctx IParseContext) {
	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MyGrammarParserRULE_parse)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.Block()
	}
	{
		p.SetState(27)
		p.Match(MyGrammarParserEOF)
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
	p.RuleIndex = MyGrammarParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_block

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
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MyGrammarParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2266132119552) != 0 {
		{
			p.SetState(29)
			p.Stat()
		}

		p.SetState(34)
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

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_OTHER returns the _OTHER token.
	Get_OTHER() antlr.Token

	// Set_OTHER sets the _OTHER token.
	Set_OTHER(antlr.Token)

	// Getter signatures
	Assignment() IAssignmentContext
	If_stat() IIf_statContext
	WhileStat() IWhileStatContext
	Log() ILogContext
	ForStat() IForStatContext
	LoopStat() ILoopStatContext
	OTHER() antlr.TerminalNode

	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	_OTHER antlr.Token
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_stat
	return p
}

func InitEmptyStatContext(p *StatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_stat
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) Get_OTHER() antlr.Token { return s._OTHER }

func (s *StatContext) Set_OTHER(v antlr.Token) { s._OTHER = v }

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

func (s *StatContext) If_stat() IIf_statContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_statContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_statContext)
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

func (s *StatContext) Log() ILogContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogContext)
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

func (s *StatContext) OTHER() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserOTHER, 0)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterStat(s)
	}
}

func (s *StatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitStat(s)
	}
}

func (s *StatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MyGrammarParserRULE_stat)
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MyGrammarParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(35)
			p.Assignment()
		}

	case MyGrammarParserIF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(36)
			p.If_stat()
		}

	case MyGrammarParserWHILE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(37)
			p.WhileStat()
		}

	case MyGrammarParserLOG:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(38)
			p.Log()
		}

	case MyGrammarParserFOR:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(39)
			p.ForStat()
		}

	case MyGrammarParserLOOP:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(40)
			p.LoopStat()
		}

	case MyGrammarParserOTHER:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(41)

			var _m = p.Match(MyGrammarParserOTHER)

			localctx.(*StatContext)._OTHER = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		fmt.Println("unknown char: " + (func() string {
			if localctx.(*StatContext).Get_OTHER() == nil {
				return ""
			} else {
				return localctx.(*StatContext).Get_OTHER().GetText()
			}
		}()))

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
	p.RuleIndex = MyGrammarParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) ID() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserID, 0)
}

func (s *AssignmentContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserASSIGN, 0)
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
	return s.GetToken(MyGrammarParserSCOL, 0)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MyGrammarParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		p.Match(MyGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(46)
		p.Match(MyGrammarParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(47)
		p.expr(0)
	}
	{
		p.SetState(48)
		p.Match(MyGrammarParserSCOL)
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

// IIf_statContext is an interface to support dynamic dispatch.
type IIf_statContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIF() []antlr.TerminalNode
	IF(i int) antlr.TerminalNode
	AllCondition_block() []ICondition_blockContext
	Condition_block(i int) ICondition_blockContext
	AllELSE() []antlr.TerminalNode
	ELSE(i int) antlr.TerminalNode
	Stat_block() IStat_blockContext

	// IsIf_statContext differentiates from other interfaces.
	IsIf_statContext()
}

type If_statContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_statContext() *If_statContext {
	var p = new(If_statContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_if_stat
	return p
}

func InitEmptyIf_statContext(p *If_statContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_if_stat
}

func (*If_statContext) IsIf_statContext() {}

func NewIf_statContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_statContext {
	var p = new(If_statContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_if_stat

	return p
}

func (s *If_statContext) GetParser() antlr.Parser { return s.parser }

func (s *If_statContext) AllIF() []antlr.TerminalNode {
	return s.GetTokens(MyGrammarParserIF)
}

func (s *If_statContext) IF(i int) antlr.TerminalNode {
	return s.GetToken(MyGrammarParserIF, i)
}

func (s *If_statContext) AllCondition_block() []ICondition_blockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICondition_blockContext); ok {
			len++
		}
	}

	tst := make([]ICondition_blockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICondition_blockContext); ok {
			tst[i] = t.(ICondition_blockContext)
			i++
		}
	}

	return tst
}

func (s *If_statContext) Condition_block(i int) ICondition_blockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICondition_blockContext); ok {
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

	return t.(ICondition_blockContext)
}

func (s *If_statContext) AllELSE() []antlr.TerminalNode {
	return s.GetTokens(MyGrammarParserELSE)
}

func (s *If_statContext) ELSE(i int) antlr.TerminalNode {
	return s.GetToken(MyGrammarParserELSE, i)
}

func (s *If_statContext) Stat_block() IStat_blockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStat_blockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStat_blockContext)
}

func (s *If_statContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_statContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_statContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterIf_stat(s)
	}
}

func (s *If_statContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitIf_stat(s)
	}
}

func (s *If_statContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitIf_stat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) If_stat() (localctx IIf_statContext) {
	localctx = NewIf_statContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MyGrammarParserRULE_if_stat)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Match(MyGrammarParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(51)
		p.Condition_block()
	}
	p.SetState(57)
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
				p.SetState(52)
				p.Match(MyGrammarParserELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(53)
				p.Match(MyGrammarParserIF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(54)
				p.Condition_block()
			}

		}
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(60)
			p.Match(MyGrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(61)
			p.Stat_block()
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

// ICondition_blockContext is an interface to support dynamic dispatch.
type ICondition_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	Stat_block() IStat_blockContext

	// IsCondition_blockContext differentiates from other interfaces.
	IsCondition_blockContext()
}

type Condition_blockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCondition_blockContext() *Condition_blockContext {
	var p = new(Condition_blockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_condition_block
	return p
}

func InitEmptyCondition_blockContext(p *Condition_blockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_condition_block
}

func (*Condition_blockContext) IsCondition_blockContext() {}

func NewCondition_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Condition_blockContext {
	var p = new(Condition_blockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_condition_block

	return p
}

func (s *Condition_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Condition_blockContext) Expr() IExprContext {
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

func (s *Condition_blockContext) Stat_block() IStat_blockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStat_blockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStat_blockContext)
}

func (s *Condition_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Condition_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Condition_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterCondition_block(s)
	}
}

func (s *Condition_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitCondition_block(s)
	}
}

func (s *Condition_blockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitCondition_block(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) Condition_block() (localctx ICondition_blockContext) {
	localctx = NewCondition_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MyGrammarParserRULE_condition_block)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.expr(0)
	}
	{
		p.SetState(65)
		p.Stat_block()
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

// IStat_blockContext is an interface to support dynamic dispatch.
type IStat_blockContext interface {
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

	// IsStat_blockContext differentiates from other interfaces.
	IsStat_blockContext()
}

type Stat_blockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStat_blockContext() *Stat_blockContext {
	var p = new(Stat_blockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_stat_block
	return p
}

func InitEmptyStat_blockContext(p *Stat_blockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_stat_block
}

func (*Stat_blockContext) IsStat_blockContext() {}

func NewStat_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Stat_blockContext {
	var p = new(Stat_blockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_stat_block

	return p
}

func (s *Stat_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Stat_blockContext) OBRACE() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserOBRACE, 0)
}

func (s *Stat_blockContext) Block() IBlockContext {
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

func (s *Stat_blockContext) CBRACE() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserCBRACE, 0)
}

func (s *Stat_blockContext) BEGIN() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserBEGIN, 0)
}

func (s *Stat_blockContext) END() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserEND, 0)
}

func (s *Stat_blockContext) DO() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserDO, 0)
}

func (s *Stat_blockContext) Stat() IStatContext {
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

func (s *Stat_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stat_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Stat_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterStat_block(s)
	}
}

func (s *Stat_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitStat_block(s)
	}
}

func (s *Stat_blockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitStat_block(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) Stat_block() (localctx IStat_blockContext) {
	localctx = NewStat_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MyGrammarParserRULE_stat_block)
	var _la int

	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MyGrammarParserOBRACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(67)
			p.Match(MyGrammarParserOBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(68)
			p.Block()
		}
		{
			p.SetState(69)
			p.Match(MyGrammarParserCBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MyGrammarParserBEGIN, MyGrammarParserDO:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MyGrammarParserDO {
			{
				p.SetState(71)
				p.Match(MyGrammarParserDO)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(74)
			p.Match(MyGrammarParserBEGIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)
			p.Block()
		}
		{
			p.SetState(76)
			p.Match(MyGrammarParserEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MyGrammarParserIF, MyGrammarParserWHILE, MyGrammarParserFOR, MyGrammarParserLOOP, MyGrammarParserLOG, MyGrammarParserID, MyGrammarParserOTHER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(78)
			p.Stat()
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
	Stat_block() IStat_blockContext

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
	p.RuleIndex = MyGrammarParserRULE_whileStat
	return p
}

func InitEmptyWhileStatContext(p *WhileStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_whileStat
}

func (*WhileStatContext) IsWhileStatContext() {}

func NewWhileStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStatContext {
	var p = new(WhileStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_whileStat

	return p
}

func (s *WhileStatContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStatContext) WHILE() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserWHILE, 0)
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

func (s *WhileStatContext) Stat_block() IStat_blockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStat_blockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStat_blockContext)
}

func (s *WhileStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterWhileStat(s)
	}
}

func (s *WhileStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitWhileStat(s)
	}
}

func (s *WhileStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitWhileStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) WhileStat() (localctx IWhileStatContext) {
	localctx = NewWhileStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MyGrammarParserRULE_whileStat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.Match(MyGrammarParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)
		p.expr(0)
	}
	{
		p.SetState(83)
		p.Stat_block()
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
	Stat_block() IStat_blockContext

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
	p.RuleIndex = MyGrammarParserRULE_forStat
	return p
}

func InitEmptyForStatContext(p *ForStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_forStat
}

func (*ForStatContext) IsForStatContext() {}

func NewForStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStatContext {
	var p = new(ForStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_forStat

	return p
}

func (s *ForStatContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStatContext) FOR() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserFOR, 0)
}

func (s *ForStatContext) ID() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserID, 0)
}

func (s *ForStatContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserASSIGN, 0)
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
	return s.GetToken(MyGrammarParserCOL, 0)
}

func (s *ForStatContext) Stat_block() IStat_blockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStat_blockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStat_blockContext)
}

func (s *ForStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterForStat(s)
	}
}

func (s *ForStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitForStat(s)
	}
}

func (s *ForStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitForStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) ForStat() (localctx IForStatContext) {
	localctx = NewForStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MyGrammarParserRULE_forStat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Match(MyGrammarParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(86)
		p.Match(MyGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(87)
		p.Match(MyGrammarParserASSIGN)
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
		p.Match(MyGrammarParserCOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(90)
		p.expr(0)
	}
	{
		p.SetState(91)
		p.Stat_block()
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
	Stat_block() IStat_blockContext

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
	p.RuleIndex = MyGrammarParserRULE_loopStat
	return p
}

func InitEmptyLoopStatContext(p *LoopStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_loopStat
}

func (*LoopStatContext) IsLoopStatContext() {}

func NewLoopStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopStatContext {
	var p = new(LoopStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_loopStat

	return p
}

func (s *LoopStatContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopStatContext) LOOP() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserLOOP, 0)
}

func (s *LoopStatContext) ID() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserID, 0)
}

func (s *LoopStatContext) COL() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserCOL, 0)
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

func (s *LoopStatContext) Stat_block() IStat_blockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStat_blockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStat_blockContext)
}

func (s *LoopStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterLoopStat(s)
	}
}

func (s *LoopStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitLoopStat(s)
	}
}

func (s *LoopStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitLoopStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) LoopStat() (localctx ILoopStatContext) {
	localctx = NewLoopStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MyGrammarParserRULE_loopStat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Match(MyGrammarParserLOOP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)
		p.Match(MyGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(95)
		p.Match(MyGrammarParserCOL)
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
		p.Stat_block()
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

// ILogContext is an interface to support dynamic dispatch.
type ILogContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LOG() antlr.TerminalNode
	Expr() IExprContext
	SCOL() antlr.TerminalNode

	// IsLogContext differentiates from other interfaces.
	IsLogContext()
}

type LogContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogContext() *LogContext {
	var p = new(LogContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_log
	return p
}

func InitEmptyLogContext(p *LogContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_log
}

func (*LogContext) IsLogContext() {}

func NewLogContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogContext {
	var p = new(LogContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_log

	return p
}

func (s *LogContext) GetParser() antlr.Parser { return s.parser }

func (s *LogContext) LOG() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserLOG, 0)
}

func (s *LogContext) Expr() IExprContext {
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

func (s *LogContext) SCOL() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserSCOL, 0)
}

func (s *LogContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterLog(s)
	}
}

func (s *LogContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitLog(s)
	}
}

func (s *LogContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitLog(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) Log() (localctx ILogContext) {
	localctx = NewLogContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MyGrammarParserRULE_log)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Match(MyGrammarParserLOG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(100)
		p.expr(0)
	}
	{
		p.SetState(101)
		p.Match(MyGrammarParserSCOL)
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
	p.RuleIndex = MyGrammarParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_expr

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
	return s.GetToken(MyGrammarParserNOT, 0)
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
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterNotExpr(s)
	}
}

func (s *NotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitNotExpr(s)
	}
}

func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
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
	return s.GetToken(MyGrammarParserMINUS, 0)
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
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterUnaryMinusExpr(s)
	}
}

func (s *UnaryMinusExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitUnaryMinusExpr(s)
	}
}

func (s *UnaryMinusExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
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
	return s.GetToken(MyGrammarParserMULT, 0)
}

func (s *MultiplicationExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserDIV, 0)
}

func (s *MultiplicationExprContext) MOD() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserMOD, 0)
}

func (s *MultiplicationExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterMultiplicationExpr(s)
	}
}

func (s *MultiplicationExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitMultiplicationExpr(s)
	}
}

func (s *MultiplicationExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
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
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterAtomExpr(s)
	}
}

func (s *AtomExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitAtomExpr(s)
	}
}

func (s *AtomExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
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
	return s.GetToken(MyGrammarParserOR, 0)
}

func (s *OrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterOrExpr(s)
	}
}

func (s *OrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitOrExpr(s)
	}
}

func (s *OrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
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
	return s.GetToken(MyGrammarParserPLUS, 0)
}

func (s *AdditiveExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserMINUS, 0)
}

func (s *AdditiveExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterAdditiveExpr(s)
	}
}

func (s *AdditiveExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitAdditiveExpr(s)
	}
}

func (s *AdditiveExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
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
	return s.GetToken(MyGrammarParserPOW, 0)
}

func (s *PowExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterPowExpr(s)
	}
}

func (s *PowExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitPowExpr(s)
	}
}

func (s *PowExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
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
	return s.GetToken(MyGrammarParserLTEQ, 0)
}

func (s *RelationalExprContext) GTEQ() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserGTEQ, 0)
}

func (s *RelationalExprContext) LT() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserLT, 0)
}

func (s *RelationalExprContext) GT() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserGT, 0)
}

func (s *RelationalExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterRelationalExpr(s)
	}
}

func (s *RelationalExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitRelationalExpr(s)
	}
}

func (s *RelationalExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
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
	return s.GetToken(MyGrammarParserEQ, 0)
}

func (s *EqualityExprContext) NEQ() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserNEQ, 0)
}

func (s *EqualityExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterEqualityExpr(s)
	}
}

func (s *EqualityExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitEqualityExpr(s)
	}
}

func (s *EqualityExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
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
	return s.GetToken(MyGrammarParserAND, 0)
}

func (s *AndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterAndExpr(s)
	}
}

func (s *AndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitAndExpr(s)
	}
}

func (s *AndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *MyGrammarParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 22
	p.EnterRecursionRule(localctx, 22, MyGrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MyGrammarParserMINUS:
		localctx = NewUnaryMinusExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(104)
			p.Match(MyGrammarParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)
			p.expr(9)
		}

	case MyGrammarParserNOT:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(106)
			p.Match(MyGrammarParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
			p.expr(8)
		}

	case MyGrammarParserOPAR, MyGrammarParserTRUE, MyGrammarParserFALSE, MyGrammarParserNIL, MyGrammarParserID, MyGrammarParserINT, MyGrammarParserFLOAT, MyGrammarParserSTRING:
		localctx = NewAtomExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(108)
			p.Atom()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(132)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPowExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MyGrammarParserRULE_expr)
				p.SetState(111)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(112)
					p.Match(MyGrammarParserPOW)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(113)
					p.expr(10)
				}

			case 2:
				localctx = NewMultiplicationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MyGrammarParserRULE_expr)
				p.SetState(114)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(115)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MultiplicationExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14336) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MultiplicationExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(116)
					p.expr(8)
				}

			case 3:
				localctx = NewAdditiveExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MyGrammarParserRULE_expr)
				p.SetState(117)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(118)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AdditiveExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MyGrammarParserPLUS || _la == MyGrammarParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AdditiveExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(119)
					p.expr(7)
				}

			case 4:
				localctx = NewRelationalExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MyGrammarParserRULE_expr)
				p.SetState(120)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(121)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*RelationalExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&480) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*RelationalExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(122)
					p.expr(6)
				}

			case 5:
				localctx = NewEqualityExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MyGrammarParserRULE_expr)
				p.SetState(123)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(124)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EqualityExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MyGrammarParserEQ || _la == MyGrammarParserNEQ) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EqualityExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(125)
					p.expr(5)
				}

			case 6:
				localctx = NewAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MyGrammarParserRULE_expr)
				p.SetState(126)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(127)
					p.Match(MyGrammarParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(128)
					p.expr(4)
				}

			case 7:
				localctx = NewOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MyGrammarParserRULE_expr)
				p.SetState(129)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(130)
					p.Match(MyGrammarParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(131)
					p.expr(3)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
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
	p.RuleIndex = MyGrammarParserRULE_atom
	return p
}

func InitEmptyAtomContext(p *AtomContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_atom
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_atom

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
	return s.GetToken(MyGrammarParserOPAR, 0)
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
	return s.GetToken(MyGrammarParserCPAR, 0)
}

func (s *ParExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterParExpr(s)
	}
}

func (s *ParExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitParExpr(s)
	}
}

func (s *ParExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
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
	return s.GetToken(MyGrammarParserTRUE, 0)
}

func (s *BooleanAtomContext) FALSE() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserFALSE, 0)
}

func (s *BooleanAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterBooleanAtom(s)
	}
}

func (s *BooleanAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitBooleanAtom(s)
	}
}

func (s *BooleanAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
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
	return s.GetToken(MyGrammarParserID, 0)
}

func (s *IdAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterIdAtom(s)
	}
}

func (s *IdAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitIdAtom(s)
	}
}

func (s *IdAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
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
	return s.GetToken(MyGrammarParserSTRING, 0)
}

func (s *StringAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterStringAtom(s)
	}
}

func (s *StringAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitStringAtom(s)
	}
}

func (s *StringAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
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
	return s.GetToken(MyGrammarParserNIL, 0)
}

func (s *NilAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterNilAtom(s)
	}
}

func (s *NilAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitNilAtom(s)
	}
}

func (s *NilAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
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
	return s.GetToken(MyGrammarParserINT, 0)
}

func (s *NumberAtomContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserFLOAT, 0)
}

func (s *NumberAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterNumberAtom(s)
	}
}

func (s *NumberAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitNumberAtom(s)
	}
}

func (s *NumberAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitNumberAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MyGrammarParserRULE_atom)
	var _la int

	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MyGrammarParserOPAR:
		localctx = NewParExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(137)
			p.Match(MyGrammarParserOPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.expr(0)
		}
		{
			p.SetState(139)
			p.Match(MyGrammarParserCPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MyGrammarParserINT, MyGrammarParserFLOAT:
		localctx = NewNumberAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(141)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MyGrammarParserINT || _la == MyGrammarParserFLOAT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case MyGrammarParserTRUE, MyGrammarParserFALSE:
		localctx = NewBooleanAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(142)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MyGrammarParserTRUE || _la == MyGrammarParserFALSE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case MyGrammarParserID:
		localctx = NewIdAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(143)
			p.Match(MyGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MyGrammarParserSTRING:
		localctx = NewStringAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(144)
			p.Match(MyGrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MyGrammarParserNIL:
		localctx = NewNilAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(145)
			p.Match(MyGrammarParserNIL)
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

func (p *MyGrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 11:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MyGrammarParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
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
