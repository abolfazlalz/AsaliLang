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
		"program", "statements", "statement", "condition", "priority1", "priority2",
		"number",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 30, 115, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 5, 1, 20, 8, 1,
		10, 1, 12, 1, 23, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 3, 2, 71, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 86, 8, 4, 10, 4, 12, 4, 89, 9, 4,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 100, 8, 5,
		10, 5, 12, 5, 103, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		3, 6, 113, 8, 6, 1, 6, 0, 2, 8, 10, 7, 0, 2, 4, 6, 8, 10, 12, 0, 0, 123,
		0, 14, 1, 0, 0, 0, 2, 21, 1, 0, 0, 0, 4, 70, 1, 0, 0, 0, 6, 72, 1, 0, 0,
		0, 8, 76, 1, 0, 0, 0, 10, 90, 1, 0, 0, 0, 12, 112, 1, 0, 0, 0, 14, 15,
		3, 2, 1, 0, 15, 1, 1, 0, 0, 0, 16, 17, 3, 4, 2, 0, 17, 18, 5, 29, 0, 0,
		18, 20, 1, 0, 0, 0, 19, 16, 1, 0, 0, 0, 20, 23, 1, 0, 0, 0, 21, 19, 1,
		0, 0, 0, 21, 22, 1, 0, 0, 0, 22, 3, 1, 0, 0, 0, 23, 21, 1, 0, 0, 0, 24,
		25, 5, 11, 0, 0, 25, 26, 5, 14, 0, 0, 26, 71, 3, 10, 5, 0, 27, 28, 5, 1,
		0, 0, 28, 29, 3, 2, 1, 0, 29, 30, 5, 2, 0, 0, 30, 71, 1, 0, 0, 0, 31, 32,
		5, 3, 0, 0, 32, 33, 3, 6, 3, 0, 33, 34, 5, 4, 0, 0, 34, 35, 3, 4, 2, 0,
		35, 71, 1, 0, 0, 0, 36, 37, 5, 3, 0, 0, 37, 38, 3, 6, 3, 0, 38, 39, 5,
		4, 0, 0, 39, 40, 3, 4, 2, 0, 40, 41, 5, 5, 0, 0, 41, 42, 3, 4, 2, 0, 42,
		71, 1, 0, 0, 0, 43, 44, 5, 6, 0, 0, 44, 45, 3, 6, 3, 0, 45, 46, 5, 7, 0,
		0, 46, 47, 3, 4, 2, 0, 47, 71, 1, 0, 0, 0, 48, 49, 5, 8, 0, 0, 49, 50,
		5, 11, 0, 0, 50, 51, 5, 14, 0, 0, 51, 52, 3, 12, 6, 0, 52, 53, 5, 15, 0,
		0, 53, 54, 3, 12, 6, 0, 54, 55, 5, 7, 0, 0, 55, 56, 3, 4, 2, 0, 56, 71,
		1, 0, 0, 0, 57, 58, 5, 9, 0, 0, 58, 59, 5, 11, 0, 0, 59, 60, 5, 15, 0,
		0, 60, 61, 3, 12, 6, 0, 61, 62, 5, 7, 0, 0, 62, 63, 3, 4, 2, 0, 63, 71,
		1, 0, 0, 0, 64, 65, 5, 10, 0, 0, 65, 71, 5, 11, 0, 0, 66, 67, 5, 10, 0,
		0, 67, 68, 5, 12, 0, 0, 68, 69, 5, 16, 0, 0, 69, 71, 5, 11, 0, 0, 70, 24,
		1, 0, 0, 0, 70, 27, 1, 0, 0, 0, 70, 31, 1, 0, 0, 0, 70, 36, 1, 0, 0, 0,
		70, 43, 1, 0, 0, 0, 70, 48, 1, 0, 0, 0, 70, 57, 1, 0, 0, 0, 70, 64, 1,
		0, 0, 0, 70, 66, 1, 0, 0, 0, 71, 5, 1, 0, 0, 0, 72, 73, 3, 8, 4, 0, 73,
		74, 5, 28, 0, 0, 74, 75, 3, 8, 4, 0, 75, 7, 1, 0, 0, 0, 76, 77, 6, 4, -1,
		0, 77, 78, 3, 12, 6, 0, 78, 87, 1, 0, 0, 0, 79, 80, 10, 2, 0, 0, 80, 81,
		5, 19, 0, 0, 81, 86, 3, 12, 6, 0, 82, 83, 10, 1, 0, 0, 83, 84, 5, 20, 0,
		0, 84, 86, 3, 12, 6, 0, 85, 79, 1, 0, 0, 0, 85, 82, 1, 0, 0, 0, 86, 89,
		1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 9, 1, 0, 0, 0,
		89, 87, 1, 0, 0, 0, 90, 91, 6, 5, -1, 0, 91, 92, 3, 8, 4, 0, 92, 101, 1,
		0, 0, 0, 93, 94, 10, 2, 0, 0, 94, 95, 5, 17, 0, 0, 95, 100, 3, 8, 4, 0,
		96, 97, 10, 1, 0, 0, 97, 98, 5, 18, 0, 0, 98, 100, 3, 8, 4, 0, 99, 93,
		1, 0, 0, 0, 99, 96, 1, 0, 0, 0, 100, 103, 1, 0, 0, 0, 101, 99, 1, 0, 0,
		0, 101, 102, 1, 0, 0, 0, 102, 11, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 104,
		113, 5, 13, 0, 0, 105, 106, 5, 18, 0, 0, 106, 113, 3, 12, 6, 0, 107, 113,
		5, 11, 0, 0, 108, 109, 5, 26, 0, 0, 109, 110, 3, 8, 4, 0, 110, 111, 5,
		27, 0, 0, 111, 113, 1, 0, 0, 0, 112, 104, 1, 0, 0, 0, 112, 105, 1, 0, 0,
		0, 112, 107, 1, 0, 0, 0, 112, 108, 1, 0, 0, 0, 113, 13, 1, 0, 0, 0, 7,
		21, 70, 85, 87, 99, 101, 112,
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
	MyGrammarParserEOF            = antlr.TokenEOF
	MyGrammarParserT__0           = 1
	MyGrammarParserT__1           = 2
	MyGrammarParserT__2           = 3
	MyGrammarParserT__3           = 4
	MyGrammarParserT__4           = 5
	MyGrammarParserT__5           = 6
	MyGrammarParserT__6           = 7
	MyGrammarParserT__7           = 8
	MyGrammarParserT__8           = 9
	MyGrammarParserT__9           = 10
	MyGrammarParserIDENTIFIER     = 11
	MyGrammarParserSTRING_LITERAL = 12
	MyGrammarParserINTEGER        = 13
	MyGrammarParserEQ             = 14
	MyGrammarParserCOLON          = 15
	MyGrammarParserCOMMA          = 16
	MyGrammarParserPLUS           = 17
	MyGrammarParserMINUS          = 18
	MyGrammarParserMULTI          = 19
	MyGrammarParserDIVIDE         = 20
	MyGrammarParserPOWERBY        = 21
	MyGrammarParserEQUALBY        = 22
	MyGrammarParserNOTEQUALBY     = 23
	MyGrammarParserLT             = 24
	MyGrammarParserGT             = 25
	MyGrammarParserOPEN_PARAN     = 26
	MyGrammarParserCLOSE_PARAN    = 27
	MyGrammarParserCONDITION_OP   = 28
	MyGrammarParserNEWLINE        = 29
	MyGrammarParserEMPTY          = 30
)

// MyGrammarParser rules.
const (
	MyGrammarParserRULE_program    = 0
	MyGrammarParserRULE_statements = 1
	MyGrammarParserRULE_statement  = 2
	MyGrammarParserRULE_condition  = 3
	MyGrammarParserRULE_priority1  = 4
	MyGrammarParserRULE_priority2  = 5
	MyGrammarParserRULE_number     = 6
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
		p.SetState(14)
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
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3914) != 0 {
		{
			p.SetState(16)
			p.Statement()
		}
		{
			p.SetState(17)
			p.Match(MyGrammarParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(23)
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

type Statement_def_varContext struct {
	StatementContext
}

func NewStatement_def_varContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Statement_def_varContext {
	var p = new(Statement_def_varContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *Statement_def_varContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Statement_def_varContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserIDENTIFIER, 0)
}

func (s *Statement_def_varContext) EQ() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserEQ, 0)
}

func (s *Statement_def_varContext) Priority2() IPriority2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPriority2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPriority2Context)
}

func (s *Statement_def_varContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterStatement_def_var(s)
	}
}

func (s *Statement_def_varContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitStatement_def_var(s)
	}
}

func (s *Statement_def_varContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitStatement_def_var(s)

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

type Statement_print_1Context struct {
	StatementContext
}

func NewStatement_print_1Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Statement_print_1Context {
	var p = new(Statement_print_1Context)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *Statement_print_1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Statement_print_1Context) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserSTRING_LITERAL, 0)
}

func (s *Statement_print_1Context) COMMA() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserCOMMA, 0)
}

func (s *Statement_print_1Context) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserIDENTIFIER, 0)
}

func (s *Statement_print_1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterStatement_print_1(s)
	}
}

func (s *Statement_print_1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitStatement_print_1(s)
	}
}

func (s *Statement_print_1Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitStatement_print_1(s)

	default:
		return t.VisitChildren(s)
	}
}

type Statement_printContext struct {
	StatementContext
}

func NewStatement_printContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Statement_printContext {
	var p = new(Statement_printContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *Statement_printContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Statement_printContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserIDENTIFIER, 0)
}

func (s *Statement_printContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterStatement_print(s)
	}
}

func (s *Statement_printContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitStatement_print(s)
	}
}

func (s *Statement_printContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitStatement_print(s)

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

func (p *MyGrammarParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MyGrammarParserRULE_statement)
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStatement_def_varContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(24)
			p.Match(MyGrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(25)
			p.Match(MyGrammarParserEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(26)
			p.priority2(0)
		}

	case 2:
		localctx = NewStatement_begin_endContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(27)
			p.Match(MyGrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(28)
			p.Statements()
		}
		{
			p.SetState(29)
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
			p.SetState(31)
			p.Match(MyGrammarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(32)
			p.Condition()
		}
		{
			p.SetState(33)
			p.Match(MyGrammarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(34)
			p.Statement()
		}

	case 4:
		localctx = NewStatement_if_elseContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(36)
			p.Match(MyGrammarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(37)
			p.Condition()
		}
		{
			p.SetState(38)
			p.Match(MyGrammarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(39)
			p.Statement()
		}
		{
			p.SetState(40)
			p.Match(MyGrammarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(41)
			p.Statement()
		}

	case 5:
		localctx = NewStatement_whileContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(43)
			p.Match(MyGrammarParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(44)
			p.Condition()
		}
		{
			p.SetState(45)
			p.Match(MyGrammarParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(46)
			p.Statement()
		}

	case 6:
		localctx = NewStatement_forContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(48)
			p.Match(MyGrammarParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(49)
			p.Match(MyGrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.Match(MyGrammarParserEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(51)
			p.Number()
		}
		{
			p.SetState(52)
			p.Match(MyGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)
			p.Number()
		}
		{
			p.SetState(54)
			p.Match(MyGrammarParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(55)
			p.Statement()
		}

	case 7:
		localctx = NewStatement_loopContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(57)
			p.Match(MyGrammarParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(58)
			p.Match(MyGrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(59)
			p.Match(MyGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(60)
			p.Number()
		}
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
			p.Statement()
		}

	case 8:
		localctx = NewStatement_printContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(64)
			p.Match(MyGrammarParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)
			p.Match(MyGrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewStatement_print_1Context(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(66)
			p.Match(MyGrammarParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)
			p.Match(MyGrammarParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(68)
			p.Match(MyGrammarParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)
			p.Match(MyGrammarParserIDENTIFIER)
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

func (s *Condition_defContext) AllPriority1() []IPriority1Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPriority1Context); ok {
			len++
		}
	}

	tst := make([]IPriority1Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPriority1Context); ok {
			tst[i] = t.(IPriority1Context)
			i++
		}
	}

	return tst
}

func (s *Condition_defContext) Priority1(i int) IPriority1Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPriority1Context); ok {
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

	return t.(IPriority1Context)
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
	p.EnterRule(localctx, 6, MyGrammarParserRULE_condition)
	localctx = NewCondition_defContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.priority1(0)
	}
	{
		p.SetState(73)
		p.Match(MyGrammarParserCONDITION_OP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(74)
		p.priority1(0)
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

// IPriority1Context is an interface to support dynamic dispatch.
type IPriority1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPriority1Context differentiates from other interfaces.
	IsPriority1Context()
}

type Priority1Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPriority1Context() *Priority1Context {
	var p = new(Priority1Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_priority1
	return p
}

func InitEmptyPriority1Context(p *Priority1Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_priority1
}

func (*Priority1Context) IsPriority1Context() {}

func NewPriority1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Priority1Context {
	var p = new(Priority1Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_priority1

	return p
}

func (s *Priority1Context) GetParser() antlr.Parser { return s.parser }

func (s *Priority1Context) CopyAll(ctx *Priority1Context) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Priority1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Priority1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Priority1_numberContext struct {
	Priority1Context
}

func NewPriority1_numberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Priority1_numberContext {
	var p = new(Priority1_numberContext)

	InitEmptyPriority1Context(&p.Priority1Context)
	p.parser = parser
	p.CopyAll(ctx.(*Priority1Context))

	return p
}

func (s *Priority1_numberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Priority1_numberContext) Number() INumberContext {
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

func (s *Priority1_numberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterPriority1_number(s)
	}
}

func (s *Priority1_numberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitPriority1_number(s)
	}
}

func (s *Priority1_numberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitPriority1_number(s)

	default:
		return t.VisitChildren(s)
	}
}

type Priority1_multipleContext struct {
	Priority1Context
}

func NewPriority1_multipleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Priority1_multipleContext {
	var p = new(Priority1_multipleContext)

	InitEmptyPriority1Context(&p.Priority1Context)
	p.parser = parser
	p.CopyAll(ctx.(*Priority1Context))

	return p
}

func (s *Priority1_multipleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Priority1_multipleContext) Priority1() IPriority1Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPriority1Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPriority1Context)
}

func (s *Priority1_multipleContext) MULTI() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserMULTI, 0)
}

func (s *Priority1_multipleContext) Number() INumberContext {
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

func (s *Priority1_multipleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterPriority1_multiple(s)
	}
}

func (s *Priority1_multipleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitPriority1_multiple(s)
	}
}

func (s *Priority1_multipleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitPriority1_multiple(s)

	default:
		return t.VisitChildren(s)
	}
}

type Priority1_divideContext struct {
	Priority1Context
}

func NewPriority1_divideContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Priority1_divideContext {
	var p = new(Priority1_divideContext)

	InitEmptyPriority1Context(&p.Priority1Context)
	p.parser = parser
	p.CopyAll(ctx.(*Priority1Context))

	return p
}

func (s *Priority1_divideContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Priority1_divideContext) Priority1() IPriority1Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPriority1Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPriority1Context)
}

func (s *Priority1_divideContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserDIVIDE, 0)
}

func (s *Priority1_divideContext) Number() INumberContext {
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

func (s *Priority1_divideContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterPriority1_divide(s)
	}
}

func (s *Priority1_divideContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitPriority1_divide(s)
	}
}

func (s *Priority1_divideContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitPriority1_divide(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) Priority1() (localctx IPriority1Context) {
	return p.priority1(0)
}

func (p *MyGrammarParser) priority1(_p int) (localctx IPriority1Context) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewPriority1Context(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPriority1Context = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, MyGrammarParserRULE_priority1, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewPriority1_numberContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(77)
		p.Number()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(85)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPriority1_multipleContext(p, NewPriority1Context(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MyGrammarParserRULE_priority1)
				p.SetState(79)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(80)
					p.Match(MyGrammarParserMULTI)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(81)
					p.Number()
				}

			case 2:
				localctx = NewPriority1_divideContext(p, NewPriority1Context(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MyGrammarParserRULE_priority1)
				p.SetState(82)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(83)
					p.Match(MyGrammarParserDIVIDE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(84)
					p.Number()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
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

// IPriority2Context is an interface to support dynamic dispatch.
type IPriority2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPriority2Context differentiates from other interfaces.
	IsPriority2Context()
}

type Priority2Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPriority2Context() *Priority2Context {
	var p = new(Priority2Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_priority2
	return p
}

func InitEmptyPriority2Context(p *Priority2Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_priority2
}

func (*Priority2Context) IsPriority2Context() {}

func NewPriority2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Priority2Context {
	var p = new(Priority2Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_priority2

	return p
}

func (s *Priority2Context) GetParser() antlr.Parser { return s.parser }

func (s *Priority2Context) CopyAll(ctx *Priority2Context) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Priority2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Priority2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Priority2_plusContext struct {
	Priority2Context
}

func NewPriority2_plusContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Priority2_plusContext {
	var p = new(Priority2_plusContext)

	InitEmptyPriority2Context(&p.Priority2Context)
	p.parser = parser
	p.CopyAll(ctx.(*Priority2Context))

	return p
}

func (s *Priority2_plusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Priority2_plusContext) Priority2() IPriority2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPriority2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPriority2Context)
}

func (s *Priority2_plusContext) PLUS() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserPLUS, 0)
}

func (s *Priority2_plusContext) Priority1() IPriority1Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPriority1Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPriority1Context)
}

func (s *Priority2_plusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterPriority2_plus(s)
	}
}

func (s *Priority2_plusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitPriority2_plus(s)
	}
}

func (s *Priority2_plusContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitPriority2_plus(s)

	default:
		return t.VisitChildren(s)
	}
}

type Priority2_defContext struct {
	Priority2Context
}

func NewPriority2_defContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Priority2_defContext {
	var p = new(Priority2_defContext)

	InitEmptyPriority2Context(&p.Priority2Context)
	p.parser = parser
	p.CopyAll(ctx.(*Priority2Context))

	return p
}

func (s *Priority2_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Priority2_defContext) Priority1() IPriority1Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPriority1Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPriority1Context)
}

func (s *Priority2_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterPriority2_def(s)
	}
}

func (s *Priority2_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitPriority2_def(s)
	}
}

func (s *Priority2_defContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitPriority2_def(s)

	default:
		return t.VisitChildren(s)
	}
}

type Priority2_minusContext struct {
	Priority2Context
}

func NewPriority2_minusContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Priority2_minusContext {
	var p = new(Priority2_minusContext)

	InitEmptyPriority2Context(&p.Priority2Context)
	p.parser = parser
	p.CopyAll(ctx.(*Priority2Context))

	return p
}

func (s *Priority2_minusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Priority2_minusContext) Priority2() IPriority2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPriority2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPriority2Context)
}

func (s *Priority2_minusContext) MINUS() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserMINUS, 0)
}

func (s *Priority2_minusContext) Priority1() IPriority1Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPriority1Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPriority1Context)
}

func (s *Priority2_minusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterPriority2_minus(s)
	}
}

func (s *Priority2_minusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitPriority2_minus(s)
	}
}

func (s *Priority2_minusContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitPriority2_minus(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) Priority2() (localctx IPriority2Context) {
	return p.priority2(0)
}

func (p *MyGrammarParser) priority2(_p int) (localctx IPriority2Context) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewPriority2Context(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPriority2Context = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, MyGrammarParserRULE_priority2, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewPriority2_defContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(91)
		p.priority1(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(99)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPriority2_plusContext(p, NewPriority2Context(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MyGrammarParserRULE_priority2)
				p.SetState(93)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(94)
					p.Match(MyGrammarParserPLUS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(95)
					p.priority1(0)
				}

			case 2:
				localctx = NewPriority2_minusContext(p, NewPriority2Context(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MyGrammarParserRULE_priority2)
				p.SetState(96)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(97)
					p.Match(MyGrammarParserMINUS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(98)
					p.priority1(0)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
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

type Number_identifierContext struct {
	NumberContext
}

func NewNumber_identifierContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Number_identifierContext {
	var p = new(Number_identifierContext)

	InitEmptyNumberContext(&p.NumberContext)
	p.parser = parser
	p.CopyAll(ctx.(*NumberContext))

	return p
}

func (s *Number_identifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Number_identifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserIDENTIFIER, 0)
}

func (s *Number_identifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterNumber_identifier(s)
	}
}

func (s *Number_identifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitNumber_identifier(s)
	}
}

func (s *Number_identifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitNumber_identifier(s)

	default:
		return t.VisitChildren(s)
	}
}

type Number_minusContext struct {
	NumberContext
}

func NewNumber_minusContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Number_minusContext {
	var p = new(Number_minusContext)

	InitEmptyNumberContext(&p.NumberContext)
	p.parser = parser
	p.CopyAll(ctx.(*NumberContext))

	return p
}

func (s *Number_minusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Number_minusContext) MINUS() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserMINUS, 0)
}

func (s *Number_minusContext) Number() INumberContext {
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

func (s *Number_minusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterNumber_minus(s)
	}
}

func (s *Number_minusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitNumber_minus(s)
	}
}

func (s *Number_minusContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitNumber_minus(s)

	default:
		return t.VisitChildren(s)
	}
}

type Number_defContext struct {
	NumberContext
}

func NewNumber_defContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Number_defContext {
	var p = new(Number_defContext)

	InitEmptyNumberContext(&p.NumberContext)
	p.parser = parser
	p.CopyAll(ctx.(*NumberContext))

	return p
}

func (s *Number_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Number_defContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserINTEGER, 0)
}

func (s *Number_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterNumber_def(s)
	}
}

func (s *Number_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitNumber_def(s)
	}
}

func (s *Number_defContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitNumber_def(s)

	default:
		return t.VisitChildren(s)
	}
}

type Number_paranContext struct {
	NumberContext
}

func NewNumber_paranContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Number_paranContext {
	var p = new(Number_paranContext)

	InitEmptyNumberContext(&p.NumberContext)
	p.parser = parser
	p.CopyAll(ctx.(*NumberContext))

	return p
}

func (s *Number_paranContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Number_paranContext) OPEN_PARAN() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserOPEN_PARAN, 0)
}

func (s *Number_paranContext) Priority1() IPriority1Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPriority1Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPriority1Context)
}

func (s *Number_paranContext) CLOSE_PARAN() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserCLOSE_PARAN, 0)
}

func (s *Number_paranContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterNumber_paran(s)
	}
}

func (s *Number_paranContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitNumber_paran(s)
	}
}

func (s *Number_paranContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitNumber_paran(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MyGrammarParserRULE_number)
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MyGrammarParserINTEGER:
		localctx = NewNumber_defContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(104)
			p.Match(MyGrammarParserINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MyGrammarParserMINUS:
		localctx = NewNumber_minusContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(105)
			p.Match(MyGrammarParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Number()
		}

	case MyGrammarParserIDENTIFIER:
		localctx = NewNumber_identifierContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(107)
			p.Match(MyGrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MyGrammarParserOPEN_PARAN:
		localctx = NewNumber_paranContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(108)
			p.Match(MyGrammarParserOPEN_PARAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)
			p.priority1(0)
		}
		{
			p.SetState(110)
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
	case 4:
		var t *Priority1Context = nil
		if localctx != nil {
			t = localctx.(*Priority1Context)
		}
		return p.Priority1_Sempred(t, predIndex)

	case 5:
		var t *Priority2Context = nil
		if localctx != nil {
			t = localctx.(*Priority2Context)
		}
		return p.Priority2_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MyGrammarParser) Priority1_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *MyGrammarParser) Priority2_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
