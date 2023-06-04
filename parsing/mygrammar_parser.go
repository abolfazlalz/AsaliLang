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
		"program", "statements", "statement", "methodCall", "variableSetterTypes",
		"methodCallArguments", "expression", "condition", "powerExpr", "multipleExpr",
		"sumExpr", "number",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 34, 169, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 29, 8, 1, 1, 1, 1, 1, 5,
		1, 33, 8, 1, 10, 1, 12, 1, 36, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2,
		91, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 3, 4, 101, 8,
		4, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 107, 8, 5, 10, 5, 12, 5, 110, 9, 5, 3,
		5, 112, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 118, 8, 6, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 129, 8, 8, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 140, 8, 9, 10, 9, 12, 9, 143,
		9, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5,
		10, 154, 8, 10, 10, 10, 12, 10, 157, 9, 10, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 167, 8, 11, 1, 11, 0, 2, 18, 20, 12,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 0, 0, 182, 0, 24, 1, 0, 0, 0,
		2, 34, 1, 0, 0, 0, 4, 90, 1, 0, 0, 0, 6, 92, 1, 0, 0, 0, 8, 100, 1, 0,
		0, 0, 10, 111, 1, 0, 0, 0, 12, 117, 1, 0, 0, 0, 14, 119, 1, 0, 0, 0, 16,
		128, 1, 0, 0, 0, 18, 130, 1, 0, 0, 0, 20, 144, 1, 0, 0, 0, 22, 166, 1,
		0, 0, 0, 24, 25, 3, 2, 1, 0, 25, 1, 1, 0, 0, 0, 26, 28, 3, 4, 2, 0, 27,
		29, 5, 32, 0, 0, 28, 27, 1, 0, 0, 0, 28, 29, 1, 0, 0, 0, 29, 30, 1, 0,
		0, 0, 30, 31, 5, 0, 0, 1, 31, 33, 1, 0, 0, 0, 32, 26, 1, 0, 0, 0, 33, 36,
		1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 3, 1, 0, 0, 0,
		36, 34, 1, 0, 0, 0, 37, 38, 5, 14, 0, 0, 38, 39, 5, 16, 0, 0, 39, 91, 3,
		8, 4, 0, 40, 41, 5, 1, 0, 0, 41, 42, 3, 2, 1, 0, 42, 43, 5, 2, 0, 0, 43,
		91, 1, 0, 0, 0, 44, 45, 5, 3, 0, 0, 45, 46, 3, 14, 7, 0, 46, 47, 5, 4,
		0, 0, 47, 48, 3, 4, 2, 0, 48, 91, 1, 0, 0, 0, 49, 50, 5, 3, 0, 0, 50, 51,
		3, 14, 7, 0, 51, 52, 5, 4, 0, 0, 52, 53, 3, 4, 2, 0, 53, 54, 5, 5, 0, 0,
		54, 55, 3, 4, 2, 0, 55, 91, 1, 0, 0, 0, 56, 57, 5, 6, 0, 0, 57, 58, 3,
		14, 7, 0, 58, 59, 5, 17, 0, 0, 59, 60, 3, 4, 2, 0, 60, 91, 1, 0, 0, 0,
		61, 62, 5, 7, 0, 0, 62, 63, 5, 14, 0, 0, 63, 64, 5, 16, 0, 0, 64, 65, 3,
		22, 11, 0, 65, 66, 5, 17, 0, 0, 66, 67, 3, 22, 11, 0, 67, 68, 5, 8, 0,
		0, 68, 69, 3, 4, 2, 0, 69, 91, 1, 0, 0, 0, 70, 71, 5, 7, 0, 0, 71, 72,
		5, 14, 0, 0, 72, 73, 5, 16, 0, 0, 73, 74, 3, 22, 11, 0, 74, 75, 5, 17,
		0, 0, 75, 76, 3, 22, 11, 0, 76, 77, 5, 9, 0, 0, 77, 78, 3, 2, 1, 0, 78,
		79, 5, 10, 0, 0, 79, 91, 1, 0, 0, 0, 80, 81, 5, 11, 0, 0, 81, 82, 5, 14,
		0, 0, 82, 83, 5, 17, 0, 0, 83, 84, 3, 22, 11, 0, 84, 85, 5, 8, 0, 0, 85,
		86, 3, 4, 2, 0, 86, 91, 1, 0, 0, 0, 87, 91, 3, 6, 3, 0, 88, 89, 5, 12,
		0, 0, 89, 91, 3, 10, 5, 0, 90, 37, 1, 0, 0, 0, 90, 40, 1, 0, 0, 0, 90,
		44, 1, 0, 0, 0, 90, 49, 1, 0, 0, 0, 90, 56, 1, 0, 0, 0, 90, 61, 1, 0, 0,
		0, 90, 70, 1, 0, 0, 0, 90, 80, 1, 0, 0, 0, 90, 87, 1, 0, 0, 0, 90, 88,
		1, 0, 0, 0, 91, 5, 1, 0, 0, 0, 92, 93, 5, 14, 0, 0, 93, 94, 5, 29, 0, 0,
		94, 95, 3, 10, 5, 0, 95, 96, 5, 30, 0, 0, 96, 7, 1, 0, 0, 0, 97, 101, 5,
		14, 0, 0, 98, 101, 3, 6, 3, 0, 99, 101, 3, 20, 10, 0, 100, 97, 1, 0, 0,
		0, 100, 98, 1, 0, 0, 0, 100, 99, 1, 0, 0, 0, 101, 9, 1, 0, 0, 0, 102, 112,
		1, 0, 0, 0, 103, 108, 3, 12, 6, 0, 104, 105, 5, 18, 0, 0, 105, 107, 3,
		12, 6, 0, 106, 104, 1, 0, 0, 0, 107, 110, 1, 0, 0, 0, 108, 106, 1, 0, 0,
		0, 108, 109, 1, 0, 0, 0, 109, 112, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 111,
		102, 1, 0, 0, 0, 111, 103, 1, 0, 0, 0, 112, 11, 1, 0, 0, 0, 113, 118, 5,
		13, 0, 0, 114, 118, 5, 14, 0, 0, 115, 118, 3, 6, 3, 0, 116, 118, 5, 15,
		0, 0, 117, 113, 1, 0, 0, 0, 117, 114, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0,
		117, 116, 1, 0, 0, 0, 118, 13, 1, 0, 0, 0, 119, 120, 3, 20, 10, 0, 120,
		121, 5, 31, 0, 0, 121, 122, 3, 20, 10, 0, 122, 15, 1, 0, 0, 0, 123, 129,
		3, 22, 11, 0, 124, 125, 3, 22, 11, 0, 125, 126, 5, 23, 0, 0, 126, 127,
		3, 16, 8, 0, 127, 129, 1, 0, 0, 0, 128, 123, 1, 0, 0, 0, 128, 124, 1, 0,
		0, 0, 129, 17, 1, 0, 0, 0, 130, 131, 6, 9, -1, 0, 131, 132, 3, 16, 8, 0,
		132, 141, 1, 0, 0, 0, 133, 134, 10, 2, 0, 0, 134, 135, 5, 21, 0, 0, 135,
		140, 3, 16, 8, 0, 136, 137, 10, 1, 0, 0, 137, 138, 5, 22, 0, 0, 138, 140,
		3, 16, 8, 0, 139, 133, 1, 0, 0, 0, 139, 136, 1, 0, 0, 0, 140, 143, 1, 0,
		0, 0, 141, 139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 19, 1, 0, 0, 0,
		143, 141, 1, 0, 0, 0, 144, 145, 6, 10, -1, 0, 145, 146, 3, 18, 9, 0, 146,
		155, 1, 0, 0, 0, 147, 148, 10, 2, 0, 0, 148, 149, 5, 19, 0, 0, 149, 154,
		3, 18, 9, 0, 150, 151, 10, 1, 0, 0, 151, 152, 5, 20, 0, 0, 152, 154, 3,
		18, 9, 0, 153, 147, 1, 0, 0, 0, 153, 150, 1, 0, 0, 0, 154, 157, 1, 0, 0,
		0, 155, 153, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 21, 1, 0, 0, 0, 157,
		155, 1, 0, 0, 0, 158, 167, 5, 15, 0, 0, 159, 160, 5, 20, 0, 0, 160, 167,
		3, 22, 11, 0, 161, 167, 5, 14, 0, 0, 162, 163, 5, 29, 0, 0, 163, 164, 3,
		20, 10, 0, 164, 165, 5, 30, 0, 0, 165, 167, 1, 0, 0, 0, 166, 158, 1, 0,
		0, 0, 166, 159, 1, 0, 0, 0, 166, 161, 1, 0, 0, 0, 166, 162, 1, 0, 0, 0,
		167, 23, 1, 0, 0, 0, 13, 28, 34, 90, 100, 108, 111, 117, 128, 139, 141,
		153, 155, 166,
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
	MyGrammarParserEOF          = antlr.TokenEOF
	MyGrammarParserT__0         = 1
	MyGrammarParserT__1         = 2
	MyGrammarParserT__2         = 3
	MyGrammarParserT__3         = 4
	MyGrammarParserT__4         = 5
	MyGrammarParserT__5         = 6
	MyGrammarParserT__6         = 7
	MyGrammarParserT__7         = 8
	MyGrammarParserT__8         = 9
	MyGrammarParserT__9         = 10
	MyGrammarParserT__10        = 11
	MyGrammarParserT__11        = 12
	MyGrammarParserSTRING       = 13
	MyGrammarParserIDENTIFIER   = 14
	MyGrammarParserINTEGER      = 15
	MyGrammarParserEQ           = 16
	MyGrammarParserCOLON        = 17
	MyGrammarParserCOMMA        = 18
	MyGrammarParserPLUS         = 19
	MyGrammarParserMINUS        = 20
	MyGrammarParserMULTI        = 21
	MyGrammarParserDIVIDE       = 22
	MyGrammarParserPOWERBY      = 23
	MyGrammarParserEQUALBY      = 24
	MyGrammarParserNOTEQUALBY   = 25
	MyGrammarParserLT           = 26
	MyGrammarParserGT           = 27
	MyGrammarParserCOT          = 28
	MyGrammarParserOPEN_PARAN   = 29
	MyGrammarParserCLOSE_PARAN  = 30
	MyGrammarParserCONDITION_OP = 31
	MyGrammarParserNEWLINE      = 32
	MyGrammarParserNEXT_PARAM   = 33
	MyGrammarParserEMPTY        = 34
)

// MyGrammarParser rules.
const (
	MyGrammarParserRULE_program             = 0
	MyGrammarParserRULE_statements          = 1
	MyGrammarParserRULE_statement           = 2
	MyGrammarParserRULE_methodCall          = 3
	MyGrammarParserRULE_variableSetterTypes = 4
	MyGrammarParserRULE_methodCallArguments = 5
	MyGrammarParserRULE_expression          = 6
	MyGrammarParserRULE_condition           = 7
	MyGrammarParserRULE_powerExpr           = 8
	MyGrammarParserRULE_multipleExpr        = 9
	MyGrammarParserRULE_sumExpr             = 10
	MyGrammarParserRULE_number              = 11
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Statements() IStatementsContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Statements() IStatementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MyGrammarParserRULE_program)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)
		p.Statements()
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

// IStatementsContext is an interface to support dynamic dispatch.
type IStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	AllEOF() []antlr.TerminalNode
	EOF(i int) antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsStatementsContext differentiates from other interfaces.
	IsStatementsContext()
}

type StatementsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementsContext() *StatementsContext {
	var p = new(StatementsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_statements
	return p
}

func InitEmptyStatementsContext(p *StatementsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_statements
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_statements

	return p
}

func (s *StatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementsContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *StatementsContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *StatementsContext) AllEOF() []antlr.TerminalNode {
	return s.GetTokens(MyGrammarParserEOF)
}

func (s *StatementsContext) EOF(i int) antlr.TerminalNode {
	return s.GetToken(MyGrammarParserEOF, i)
}

func (s *StatementsContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MyGrammarParserNEWLINE)
}

func (s *StatementsContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MyGrammarParserNEWLINE, i)
}

func (s *StatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterStatements(s)
	}
}

func (s *StatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitStatements(s)
	}
}

func (s *StatementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitStatements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) Statements() (localctx IStatementsContext) {
	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MyGrammarParserRULE_statements)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&22730) != 0 {
		{
			p.SetState(26)
			p.Statement()
		}
		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MyGrammarParserNEWLINE {
			{
				p.SetState(27)
				p.Match(MyGrammarParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(30)
			p.Match(MyGrammarParserEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(36)
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyAll(ctx *StatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Statement_begin_endContext struct {
	StatementContext
}

func NewStatement_begin_endContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Statement_begin_endContext {
	var p = new(Statement_begin_endContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *Statement_begin_endContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Statement_begin_endContext) Statements() IStatementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *Statement_begin_endContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterStatement_begin_end(s)
	}
}

func (s *Statement_begin_endContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitStatement_begin_end(s)
	}
}

func (s *Statement_begin_endContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitStatement_begin_end(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatementPrintMethodContext struct {
	StatementContext
}

func NewStatementPrintMethodContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatementPrintMethodContext {
	var p = new(StatementPrintMethodContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *StatementPrintMethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementPrintMethodContext) MethodCallArguments() IMethodCallArgumentsContext {
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

func (s *StatementPrintMethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterStatementPrintMethod(s)
	}
}

func (s *StatementPrintMethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitStatementPrintMethod(s)
	}
}

func (s *StatementPrintMethodContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitStatementPrintMethod(s)

	default:
		return t.VisitChildren(s)
	}
}

type Statement_whileContext struct {
	StatementContext
}

func NewStatement_whileContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Statement_whileContext {
	var p = new(Statement_whileContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *Statement_whileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Statement_whileContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *Statement_whileContext) COLON() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserCOLON, 0)
}

func (s *Statement_whileContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *Statement_whileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterStatement_while(s)
	}
}

func (s *Statement_whileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitStatement_while(s)
	}
}

func (s *Statement_whileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitStatement_while(s)

	default:
		return t.VisitChildren(s)
	}
}

type Statement_loopContext struct {
	StatementContext
}

func NewStatement_loopContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Statement_loopContext {
	var p = new(Statement_loopContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *Statement_loopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Statement_loopContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserIDENTIFIER, 0)
}

func (s *Statement_loopContext) COLON() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserCOLON, 0)
}

func (s *Statement_loopContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *Statement_loopContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *Statement_loopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterStatement_loop(s)
	}
}

func (s *Statement_loopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitStatement_loop(s)
	}
}

func (s *Statement_loopContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitStatement_loop(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatementDefineVariableContext struct {
	StatementContext
}

func NewStatementDefineVariableContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatementDefineVariableContext {
	var p = new(StatementDefineVariableContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *StatementDefineVariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementDefineVariableContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserIDENTIFIER, 0)
}

func (s *StatementDefineVariableContext) EQ() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserEQ, 0)
}

func (s *StatementDefineVariableContext) VariableSetterTypes() IVariableSetterTypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableSetterTypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableSetterTypesContext)
}

func (s *StatementDefineVariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterStatementDefineVariable(s)
	}
}

func (s *StatementDefineVariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitStatementDefineVariable(s)
	}
}

func (s *StatementDefineVariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitStatementDefineVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

type Statement_forContext struct {
	StatementContext
}

func NewStatement_forContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Statement_forContext {
	var p = new(Statement_forContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *Statement_forContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Statement_forContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserIDENTIFIER, 0)
}

func (s *Statement_forContext) EQ() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserEQ, 0)
}

func (s *Statement_forContext) AllNumber() []INumberContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumberContext); ok {
			len++
		}
	}

	tst := make([]INumberContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumberContext); ok {
			tst[i] = t.(INumberContext)
			i++
		}
	}

	return tst
}

func (s *Statement_forContext) Number(i int) INumberContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
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

	return t.(INumberContext)
}

func (s *Statement_forContext) COLON() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserCOLON, 0)
}

func (s *Statement_forContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *Statement_forContext) Statements() IStatementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *Statement_forContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterStatement_for(s)
	}
}

func (s *Statement_forContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitStatement_for(s)
	}
}

func (s *Statement_forContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitStatement_for(s)

	default:
		return t.VisitChildren(s)
	}
}

type Statement_if_elseContext struct {
	StatementContext
}

func NewStatement_if_elseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Statement_if_elseContext {
	var p = new(Statement_if_elseContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *Statement_if_elseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Statement_if_elseContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *Statement_if_elseContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *Statement_if_elseContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *Statement_if_elseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterStatement_if_else(s)
	}
}

func (s *Statement_if_elseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitStatement_if_else(s)
	}
}

func (s *Statement_if_elseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitStatement_if_else(s)

	default:
		return t.VisitChildren(s)
	}
}

type Statement_ifContext struct {
	StatementContext
}

func NewStatement_ifContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Statement_ifContext {
	var p = new(Statement_ifContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *Statement_ifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Statement_ifContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *Statement_ifContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *Statement_ifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterStatement_if(s)
	}
}

func (s *Statement_ifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitStatement_if(s)
	}
}

func (s *Statement_ifContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitStatement_if(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallMethodContext struct {
	StatementContext
}

func NewCallMethodContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallMethodContext {
	var p = new(CallMethodContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *CallMethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallMethodContext) MethodCall() IMethodCallContext {
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

func (s *CallMethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterCallMethod(s)
	}
}

func (s *CallMethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitCallMethod(s)
	}
}

func (s *CallMethodContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitCallMethod(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MyGrammarParserRULE_statement)
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStatementDefineVariableContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(37)
			p.Match(MyGrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(38)
			p.Match(MyGrammarParserEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(39)
			p.VariableSetterTypes()
		}

	case 2:
		localctx = NewStatement_begin_endContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(40)
			p.Match(MyGrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(41)
			p.Statements()
		}
		{
			p.SetState(42)
			p.Match(MyGrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewStatement_ifContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(44)
			p.Match(MyGrammarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(45)
			p.Condition()
		}
		{
			p.SetState(46)
			p.Match(MyGrammarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(47)
			p.Statement()
		}

	case 4:
		localctx = NewStatement_if_elseContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(49)
			p.Match(MyGrammarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.Condition()
		}
		{
			p.SetState(51)
			p.Match(MyGrammarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(52)
			p.Statement()
		}
		{
			p.SetState(53)
			p.Match(MyGrammarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(54)
			p.Statement()
		}

	case 5:
		localctx = NewStatement_whileContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(56)
			p.Match(MyGrammarParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(57)
			p.Condition()
		}
		{
			p.SetState(58)
			p.Match(MyGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(59)
			p.Statement()
		}

	case 6:
		localctx = NewStatement_forContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(61)
			p.Match(MyGrammarParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(62)
			p.Match(MyGrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(63)
			p.Match(MyGrammarParserEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(64)
			p.Number()
		}
		{
			p.SetState(65)
			p.Match(MyGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(66)
			p.Number()
		}
		{
			p.SetState(67)
			p.Match(MyGrammarParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(68)
			p.Statement()
		}

	case 7:
		localctx = NewStatement_forContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(70)
			p.Match(MyGrammarParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			p.Match(MyGrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(72)
			p.Match(MyGrammarParserEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(73)
			p.Number()
		}
		{
			p.SetState(74)
			p.Match(MyGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)
			p.Number()
		}
		{
			p.SetState(76)
			p.Match(MyGrammarParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)
			p.Statements()
		}
		{
			p.SetState(78)
			p.Match(MyGrammarParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewStatement_loopContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(80)
			p.Match(MyGrammarParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(81)
			p.Match(MyGrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(82)
			p.Match(MyGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.Number()
		}
		{
			p.SetState(84)
			p.Match(MyGrammarParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
			p.Statement()
		}

	case 9:
		localctx = NewCallMethodContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(87)
			p.MethodCall()
		}

	case 10:
		localctx = NewStatementPrintMethodContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(88)
			p.Match(MyGrammarParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(89)
			p.MethodCallArguments()
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

// IMethodCallContext is an interface to support dynamic dispatch.
type IMethodCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	OPEN_PARAN() antlr.TerminalNode
	MethodCallArguments() IMethodCallArgumentsContext
	CLOSE_PARAN() antlr.TerminalNode

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
	p.RuleIndex = MyGrammarParserRULE_methodCall
	return p
}

func InitEmptyMethodCallContext(p *MethodCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_methodCall
}

func (*MethodCallContext) IsMethodCallContext() {}

func NewMethodCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodCallContext {
	var p = new(MethodCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_methodCall

	return p
}

func (s *MethodCallContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodCallContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserIDENTIFIER, 0)
}

func (s *MethodCallContext) OPEN_PARAN() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserOPEN_PARAN, 0)
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

func (s *MethodCallContext) CLOSE_PARAN() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserCLOSE_PARAN, 0)
}

func (s *MethodCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterMethodCall(s)
	}
}

func (s *MethodCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitMethodCall(s)
	}
}

func (s *MethodCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitMethodCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) MethodCall() (localctx IMethodCallContext) {
	localctx = NewMethodCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MyGrammarParserRULE_methodCall)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Match(MyGrammarParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)
		p.Match(MyGrammarParserOPEN_PARAN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)
		p.MethodCallArguments()
	}
	{
		p.SetState(95)
		p.Match(MyGrammarParserCLOSE_PARAN)
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

// IVariableSetterTypesContext is an interface to support dynamic dispatch.
type IVariableSetterTypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	MethodCall() IMethodCallContext
	SumExpr() ISumExprContext

	// IsVariableSetterTypesContext differentiates from other interfaces.
	IsVariableSetterTypesContext()
}

type VariableSetterTypesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableSetterTypesContext() *VariableSetterTypesContext {
	var p = new(VariableSetterTypesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_variableSetterTypes
	return p
}

func InitEmptyVariableSetterTypesContext(p *VariableSetterTypesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_variableSetterTypes
}

func (*VariableSetterTypesContext) IsVariableSetterTypesContext() {}

func NewVariableSetterTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableSetterTypesContext {
	var p = new(VariableSetterTypesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_variableSetterTypes

	return p
}

func (s *VariableSetterTypesContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableSetterTypesContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserIDENTIFIER, 0)
}

func (s *VariableSetterTypesContext) MethodCall() IMethodCallContext {
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

func (s *VariableSetterTypesContext) SumExpr() ISumExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISumExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISumExprContext)
}

func (s *VariableSetterTypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableSetterTypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableSetterTypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterVariableSetterTypes(s)
	}
}

func (s *VariableSetterTypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitVariableSetterTypes(s)
	}
}

func (s *VariableSetterTypesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitVariableSetterTypes(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) VariableSetterTypes() (localctx IVariableSetterTypesContext) {
	localctx = NewVariableSetterTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MyGrammarParserRULE_variableSetterTypes)
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(97)
			p.Match(MyGrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(98)
			p.MethodCall()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(99)
			p.sumExpr(0)
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

// IMethodCallArgumentsContext is an interface to support dynamic dispatch.
type IMethodCallArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

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
	p.RuleIndex = MyGrammarParserRULE_methodCallArguments
	return p
}

func InitEmptyMethodCallArgumentsContext(p *MethodCallArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_methodCallArguments
}

func (*MethodCallArgumentsContext) IsMethodCallArgumentsContext() {}

func NewMethodCallArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodCallArgumentsContext {
	var p = new(MethodCallArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_methodCallArguments

	return p
}

func (s *MethodCallArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodCallArgumentsContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MethodCallArgumentsContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *MethodCallArgumentsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MyGrammarParserCOMMA)
}

func (s *MethodCallArgumentsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MyGrammarParserCOMMA, i)
}

func (s *MethodCallArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodCallArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodCallArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterMethodCallArguments(s)
	}
}

func (s *MethodCallArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitMethodCallArguments(s)
	}
}

func (s *MethodCallArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitMethodCallArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) MethodCallArguments() (localctx IMethodCallArgumentsContext) {
	localctx = NewMethodCallArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MyGrammarParserRULE_methodCallArguments)
	var _la int

	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MyGrammarParserEOF, MyGrammarParserT__4, MyGrammarParserCLOSE_PARAN, MyGrammarParserNEWLINE:
		p.EnterOuterAlt(localctx, 1)

	case MyGrammarParserSTRING, MyGrammarParserIDENTIFIER, MyGrammarParserINTEGER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(103)
			p.Expression()
		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MyGrammarParserCOMMA {
			{
				p.SetState(104)
				p.Match(MyGrammarParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(105)
				p.Expression()
			}

			p.SetState(110)
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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	MethodCall() IMethodCallContext
	INTEGER() antlr.TerminalNode

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) STRING() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserSTRING, 0)
}

func (s *ExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserIDENTIFIER, 0)
}

func (s *ExpressionContext) MethodCall() IMethodCallContext {
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

func (s *ExpressionContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserINTEGER, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MyGrammarParserRULE_expression)
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(113)
			p.Match(MyGrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(114)
			p.Match(MyGrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(115)
			p.MethodCall()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(116)
			p.Match(MyGrammarParserINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_condition
	return p
}

func InitEmptyConditionContext(p *ConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_condition
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) CopyAll(ctx *ConditionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Condition_defContext struct {
	ConditionContext
}

func NewCondition_defContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Condition_defContext {
	var p = new(Condition_defContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *Condition_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Condition_defContext) AllSumExpr() []ISumExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISumExprContext); ok {
			len++
		}
	}

	tst := make([]ISumExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISumExprContext); ok {
			tst[i] = t.(ISumExprContext)
			i++
		}
	}

	return tst
}

func (s *Condition_defContext) SumExpr(i int) ISumExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISumExprContext); ok {
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

	return t.(ISumExprContext)
}

func (s *Condition_defContext) CONDITION_OP() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserCONDITION_OP, 0)
}

func (s *Condition_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterCondition_def(s)
	}
}

func (s *Condition_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitCondition_def(s)
	}
}

func (s *Condition_defContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitCondition_def(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MyGrammarParserRULE_condition)
	localctx = NewCondition_defContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.sumExpr(0)
	}
	{
		p.SetState(120)
		p.Match(MyGrammarParserCONDITION_OP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
		p.sumExpr(0)
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

// IPowerExprContext is an interface to support dynamic dispatch.
type IPowerExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPowerExprContext differentiates from other interfaces.
	IsPowerExprContext()
}

type PowerExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPowerExprContext() *PowerExprContext {
	var p = new(PowerExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_powerExpr
	return p
}

func InitEmptyPowerExprContext(p *PowerExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_powerExpr
}

func (*PowerExprContext) IsPowerExprContext() {}

func NewPowerExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PowerExprContext {
	var p = new(PowerExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_powerExpr

	return p
}

func (s *PowerExprContext) GetParser() antlr.Parser { return s.parser }

func (s *PowerExprContext) CopyAll(ctx *PowerExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PowerExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowerExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PowerExprDefaultContext struct {
	PowerExprContext
}

func NewPowerExprDefaultContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowerExprDefaultContext {
	var p = new(PowerExprDefaultContext)

	InitEmptyPowerExprContext(&p.PowerExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*PowerExprContext))

	return p
}

func (s *PowerExprDefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowerExprDefaultContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *PowerExprDefaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterPowerExprDefault(s)
	}
}

func (s *PowerExprDefaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitPowerExprDefault(s)
	}
}

func (s *PowerExprDefaultContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitPowerExprDefault(s)

	default:
		return t.VisitChildren(s)
	}
}

type PowerExprPowerContext struct {
	PowerExprContext
}

func NewPowerExprPowerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowerExprPowerContext {
	var p = new(PowerExprPowerContext)

	InitEmptyPowerExprContext(&p.PowerExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*PowerExprContext))

	return p
}

func (s *PowerExprPowerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowerExprPowerContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *PowerExprPowerContext) POWERBY() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserPOWERBY, 0)
}

func (s *PowerExprPowerContext) PowerExpr() IPowerExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPowerExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPowerExprContext)
}

func (s *PowerExprPowerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterPowerExprPower(s)
	}
}

func (s *PowerExprPowerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitPowerExprPower(s)
	}
}

func (s *PowerExprPowerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitPowerExprPower(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) PowerExpr() (localctx IPowerExprContext) {
	localctx = NewPowerExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MyGrammarParserRULE_powerExpr)
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPowerExprDefaultContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(123)
			p.Number()
		}

	case 2:
		localctx = NewPowerExprPowerContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(124)
			p.Number()
		}
		{
			p.SetState(125)
			p.Match(MyGrammarParserPOWERBY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(126)
			p.PowerExpr()
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

// IMultipleExprContext is an interface to support dynamic dispatch.
type IMultipleExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMultipleExprContext differentiates from other interfaces.
	IsMultipleExprContext()
}

type MultipleExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultipleExprContext() *MultipleExprContext {
	var p = new(MultipleExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_multipleExpr
	return p
}

func InitEmptyMultipleExprContext(p *MultipleExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_multipleExpr
}

func (*MultipleExprContext) IsMultipleExprContext() {}

func NewMultipleExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultipleExprContext {
	var p = new(MultipleExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_multipleExpr

	return p
}

func (s *MultipleExprContext) GetParser() antlr.Parser { return s.parser }

func (s *MultipleExprContext) CopyAll(ctx *MultipleExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *MultipleExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultipleExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MultipleExprDivideContext struct {
	MultipleExprContext
}

func NewMultipleExprDivideContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultipleExprDivideContext {
	var p = new(MultipleExprDivideContext)

	InitEmptyMultipleExprContext(&p.MultipleExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*MultipleExprContext))

	return p
}

func (s *MultipleExprDivideContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultipleExprDivideContext) MultipleExpr() IMultipleExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultipleExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultipleExprContext)
}

func (s *MultipleExprDivideContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserDIVIDE, 0)
}

func (s *MultipleExprDivideContext) PowerExpr() IPowerExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPowerExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPowerExprContext)
}

func (s *MultipleExprDivideContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterMultipleExprDivide(s)
	}
}

func (s *MultipleExprDivideContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitMultipleExprDivide(s)
	}
}

func (s *MultipleExprDivideContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitMultipleExprDivide(s)

	default:
		return t.VisitChildren(s)
	}
}

type MultipleExprMultiContext struct {
	MultipleExprContext
}

func NewMultipleExprMultiContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultipleExprMultiContext {
	var p = new(MultipleExprMultiContext)

	InitEmptyMultipleExprContext(&p.MultipleExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*MultipleExprContext))

	return p
}

func (s *MultipleExprMultiContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultipleExprMultiContext) MultipleExpr() IMultipleExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultipleExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultipleExprContext)
}

func (s *MultipleExprMultiContext) MULTI() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserMULTI, 0)
}

func (s *MultipleExprMultiContext) PowerExpr() IPowerExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPowerExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPowerExprContext)
}

func (s *MultipleExprMultiContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterMultipleExprMulti(s)
	}
}

func (s *MultipleExprMultiContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitMultipleExprMulti(s)
	}
}

func (s *MultipleExprMultiContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitMultipleExprMulti(s)

	default:
		return t.VisitChildren(s)
	}
}

type MultipleExprDefaultContext struct {
	MultipleExprContext
}

func NewMultipleExprDefaultContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultipleExprDefaultContext {
	var p = new(MultipleExprDefaultContext)

	InitEmptyMultipleExprContext(&p.MultipleExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*MultipleExprContext))

	return p
}

func (s *MultipleExprDefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultipleExprDefaultContext) PowerExpr() IPowerExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPowerExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPowerExprContext)
}

func (s *MultipleExprDefaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterMultipleExprDefault(s)
	}
}

func (s *MultipleExprDefaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitMultipleExprDefault(s)
	}
}

func (s *MultipleExprDefaultContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitMultipleExprDefault(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) MultipleExpr() (localctx IMultipleExprContext) {
	return p.multipleExpr(0)
}

func (p *MyGrammarParser) multipleExpr(_p int) (localctx IMultipleExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewMultipleExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMultipleExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, MyGrammarParserRULE_multipleExpr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewMultipleExprDefaultContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(131)
		p.PowerExpr()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(139)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultipleExprMultiContext(p, NewMultipleExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MyGrammarParserRULE_multipleExpr)
				p.SetState(133)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(134)
					p.Match(MyGrammarParserMULTI)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(135)
					p.PowerExpr()
				}

			case 2:
				localctx = NewMultipleExprDivideContext(p, NewMultipleExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MyGrammarParserRULE_multipleExpr)
				p.SetState(136)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(137)
					p.Match(MyGrammarParserDIVIDE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(138)
					p.PowerExpr()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
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

// ISumExprContext is an interface to support dynamic dispatch.
type ISumExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSumExprContext differentiates from other interfaces.
	IsSumExprContext()
}

type SumExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySumExprContext() *SumExprContext {
	var p = new(SumExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_sumExpr
	return p
}

func InitEmptySumExprContext(p *SumExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_sumExpr
}

func (*SumExprContext) IsSumExprContext() {}

func NewSumExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SumExprContext {
	var p = new(SumExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_sumExpr

	return p
}

func (s *SumExprContext) GetParser() antlr.Parser { return s.parser }

func (s *SumExprContext) CopyAll(ctx *SumExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SumExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SumExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SumExprPlusContext struct {
	SumExprContext
}

func NewSumExprPlusContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SumExprPlusContext {
	var p = new(SumExprPlusContext)

	InitEmptySumExprContext(&p.SumExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*SumExprContext))

	return p
}

func (s *SumExprPlusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SumExprPlusContext) SumExpr() ISumExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISumExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISumExprContext)
}

func (s *SumExprPlusContext) PLUS() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserPLUS, 0)
}

func (s *SumExprPlusContext) MultipleExpr() IMultipleExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultipleExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultipleExprContext)
}

func (s *SumExprPlusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterSumExprPlus(s)
	}
}

func (s *SumExprPlusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitSumExprPlus(s)
	}
}

func (s *SumExprPlusContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitSumExprPlus(s)

	default:
		return t.VisitChildren(s)
	}
}

type SumExprMinusContext struct {
	SumExprContext
}

func NewSumExprMinusContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SumExprMinusContext {
	var p = new(SumExprMinusContext)

	InitEmptySumExprContext(&p.SumExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*SumExprContext))

	return p
}

func (s *SumExprMinusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SumExprMinusContext) SumExpr() ISumExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISumExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISumExprContext)
}

func (s *SumExprMinusContext) MINUS() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserMINUS, 0)
}

func (s *SumExprMinusContext) MultipleExpr() IMultipleExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultipleExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultipleExprContext)
}

func (s *SumExprMinusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterSumExprMinus(s)
	}
}

func (s *SumExprMinusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitSumExprMinus(s)
	}
}

func (s *SumExprMinusContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitSumExprMinus(s)

	default:
		return t.VisitChildren(s)
	}
}

type SumExprDefaultContext struct {
	SumExprContext
}

func NewSumExprDefaultContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SumExprDefaultContext {
	var p = new(SumExprDefaultContext)

	InitEmptySumExprContext(&p.SumExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*SumExprContext))

	return p
}

func (s *SumExprDefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SumExprDefaultContext) MultipleExpr() IMultipleExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultipleExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultipleExprContext)
}

func (s *SumExprDefaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterSumExprDefault(s)
	}
}

func (s *SumExprDefaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitSumExprDefault(s)
	}
}

func (s *SumExprDefaultContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitSumExprDefault(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) SumExpr() (localctx ISumExprContext) {
	return p.sumExpr(0)
}

func (p *MyGrammarParser) sumExpr(_p int) (localctx ISumExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewSumExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ISumExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, MyGrammarParserRULE_sumExpr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewSumExprDefaultContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(145)
		p.multipleExpr(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(153)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewSumExprPlusContext(p, NewSumExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MyGrammarParserRULE_sumExpr)
				p.SetState(147)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(148)
					p.Match(MyGrammarParserPLUS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(149)
					p.multipleExpr(0)
				}

			case 2:
				localctx = NewSumExprMinusContext(p, NewSumExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MyGrammarParserRULE_sumExpr)
				p.SetState(150)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(151)
					p.Match(MyGrammarParserMINUS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(152)
					p.multipleExpr(0)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
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

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_number
	return p
}

func InitEmptyNumberContext(p *NumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_number
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) CopyAll(ctx *NumberContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NumberParenthesesContext struct {
	NumberContext
}

func NewNumberParenthesesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberParenthesesContext {
	var p = new(NumberParenthesesContext)

	InitEmptyNumberContext(&p.NumberContext)
	p.parser = parser
	p.CopyAll(ctx.(*NumberContext))

	return p
}

func (s *NumberParenthesesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberParenthesesContext) OPEN_PARAN() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserOPEN_PARAN, 0)
}

func (s *NumberParenthesesContext) SumExpr() ISumExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISumExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISumExprContext)
}

func (s *NumberParenthesesContext) CLOSE_PARAN() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserCLOSE_PARAN, 0)
}

func (s *NumberParenthesesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterNumberParentheses(s)
	}
}

func (s *NumberParenthesesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitNumberParentheses(s)
	}
}

func (s *NumberParenthesesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitNumberParentheses(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberDefaultContext struct {
	NumberContext
}

func NewNumberDefaultContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberDefaultContext {
	var p = new(NumberDefaultContext)

	InitEmptyNumberContext(&p.NumberContext)
	p.parser = parser
	p.CopyAll(ctx.(*NumberContext))

	return p
}

func (s *NumberDefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberDefaultContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserINTEGER, 0)
}

func (s *NumberDefaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterNumberDefault(s)
	}
}

func (s *NumberDefaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitNumberDefault(s)
	}
}

func (s *NumberDefaultContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitNumberDefault(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberMinusContext struct {
	NumberContext
}

func NewNumberMinusContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberMinusContext {
	var p = new(NumberMinusContext)

	InitEmptyNumberContext(&p.NumberContext)
	p.parser = parser
	p.CopyAll(ctx.(*NumberContext))

	return p
}

func (s *NumberMinusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberMinusContext) MINUS() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserMINUS, 0)
}

func (s *NumberMinusContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *NumberMinusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterNumberMinus(s)
	}
}

func (s *NumberMinusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitNumberMinus(s)
	}
}

func (s *NumberMinusContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitNumberMinus(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberIdentifierContext struct {
	NumberContext
}

func NewNumberIdentifierContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberIdentifierContext {
	var p = new(NumberIdentifierContext)

	InitEmptyNumberContext(&p.NumberContext)
	p.parser = parser
	p.CopyAll(ctx.(*NumberContext))

	return p
}

func (s *NumberIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberIdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserIDENTIFIER, 0)
}

func (s *NumberIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterNumberIdentifier(s)
	}
}

func (s *NumberIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitNumberIdentifier(s)
	}
}

func (s *NumberIdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitNumberIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MyGrammarParserRULE_number)
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MyGrammarParserINTEGER:
		localctx = NewNumberDefaultContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(158)
			p.Match(MyGrammarParserINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MyGrammarParserMINUS:
		localctx = NewNumberMinusContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(159)
			p.Match(MyGrammarParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)
			p.Number()
		}

	case MyGrammarParserIDENTIFIER:
		localctx = NewNumberIdentifierContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(161)
			p.Match(MyGrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MyGrammarParserOPEN_PARAN:
		localctx = NewNumberParenthesesContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(162)
			p.Match(MyGrammarParserOPEN_PARAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(163)
			p.sumExpr(0)
		}
		{
			p.SetState(164)
			p.Match(MyGrammarParserCLOSE_PARAN)
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
	case 9:
		var t *MultipleExprContext = nil
		if localctx != nil {
			t = localctx.(*MultipleExprContext)
		}
		return p.MultipleExpr_Sempred(t, predIndex)

	case 10:
		var t *SumExprContext = nil
		if localctx != nil {
			t = localctx.(*SumExprContext)
		}
		return p.SumExpr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MyGrammarParser) MultipleExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *MyGrammarParser) SumExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
