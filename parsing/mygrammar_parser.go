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
		"program", "statements", "statement", "assignmentStatement", "printStatement",
		"blockStatement", "whileStatement", "methodCall", "ifStatement", "variableSetterTypes",
		"methodCallArguments", "expression", "powerExpr", "multipleExpr", "sumExpr",
		"number",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 36, 146, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 37, 8, 1, 10, 1, 12, 1, 40, 9, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 3, 2, 47, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9,
		1, 9, 1, 9, 3, 9, 81, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 87, 8, 10,
		10, 10, 12, 10, 90, 9, 10, 3, 10, 92, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 3, 11, 99, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 106,
		8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 5,
		13, 117, 8, 13, 10, 13, 12, 13, 120, 9, 13, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 131, 8, 14, 10, 14, 12, 14, 134,
		9, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 144,
		8, 15, 1, 15, 0, 2, 26, 28, 16, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 30, 0, 0, 150, 0, 32, 1, 0, 0, 0, 2, 38, 1, 0, 0, 0, 4,
		46, 1, 0, 0, 0, 6, 48, 1, 0, 0, 0, 8, 53, 1, 0, 0, 0, 10, 56, 1, 0, 0,
		0, 12, 60, 1, 0, 0, 0, 14, 66, 1, 0, 0, 0, 16, 72, 1, 0, 0, 0, 18, 80,
		1, 0, 0, 0, 20, 91, 1, 0, 0, 0, 22, 98, 1, 0, 0, 0, 24, 105, 1, 0, 0, 0,
		26, 107, 1, 0, 0, 0, 28, 121, 1, 0, 0, 0, 30, 143, 1, 0, 0, 0, 32, 33,
		3, 2, 1, 0, 33, 34, 5, 0, 0, 1, 34, 1, 1, 0, 0, 0, 35, 37, 3, 4, 2, 0,
		36, 35, 1, 0, 0, 0, 37, 40, 1, 0, 0, 0, 38, 36, 1, 0, 0, 0, 38, 39, 1,
		0, 0, 0, 39, 3, 1, 0, 0, 0, 40, 38, 1, 0, 0, 0, 41, 47, 3, 6, 3, 0, 42,
		47, 3, 16, 8, 0, 43, 47, 3, 12, 6, 0, 44, 47, 3, 14, 7, 0, 45, 47, 3, 8,
		4, 0, 46, 41, 1, 0, 0, 0, 46, 42, 1, 0, 0, 0, 46, 43, 1, 0, 0, 0, 46, 44,
		1, 0, 0, 0, 46, 45, 1, 0, 0, 0, 47, 5, 1, 0, 0, 0, 48, 49, 5, 2, 0, 0,
		49, 50, 5, 4, 0, 0, 50, 51, 3, 18, 9, 0, 51, 52, 5, 29, 0, 0, 52, 7, 1,
		0, 0, 0, 53, 54, 5, 21, 0, 0, 54, 55, 3, 18, 9, 0, 55, 9, 1, 0, 0, 0, 56,
		57, 5, 15, 0, 0, 57, 58, 3, 2, 1, 0, 58, 59, 5, 16, 0, 0, 59, 11, 1, 0,
		0, 0, 60, 61, 5, 18, 0, 0, 61, 62, 3, 22, 11, 0, 62, 63, 5, 14, 0, 0, 63,
		64, 3, 4, 2, 0, 64, 65, 5, 16, 0, 0, 65, 13, 1, 0, 0, 0, 66, 67, 5, 2,
		0, 0, 67, 68, 5, 30, 0, 0, 68, 69, 3, 20, 10, 0, 69, 70, 5, 31, 0, 0, 70,
		71, 5, 29, 0, 0, 71, 15, 1, 0, 0, 0, 72, 73, 5, 12, 0, 0, 73, 74, 3, 22,
		11, 0, 74, 75, 5, 17, 0, 0, 75, 76, 3, 4, 2, 0, 76, 17, 1, 0, 0, 0, 77,
		81, 5, 2, 0, 0, 78, 81, 3, 14, 7, 0, 79, 81, 3, 28, 14, 0, 80, 77, 1, 0,
		0, 0, 80, 78, 1, 0, 0, 0, 80, 79, 1, 0, 0, 0, 81, 19, 1, 0, 0, 0, 82, 92,
		1, 0, 0, 0, 83, 88, 3, 22, 11, 0, 84, 85, 5, 6, 0, 0, 85, 87, 3, 22, 11,
		0, 86, 84, 1, 0, 0, 0, 87, 90, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 88, 89,
		1, 0, 0, 0, 89, 92, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 91, 82, 1, 0, 0, 0,
		91, 83, 1, 0, 0, 0, 92, 21, 1, 0, 0, 0, 93, 99, 5, 1, 0, 0, 94, 99, 5,
		2, 0, 0, 95, 99, 3, 14, 7, 0, 96, 99, 5, 3, 0, 0, 97, 99, 3, 28, 14, 0,
		98, 93, 1, 0, 0, 0, 98, 94, 1, 0, 0, 0, 98, 95, 1, 0, 0, 0, 98, 96, 1,
		0, 0, 0, 98, 97, 1, 0, 0, 0, 99, 23, 1, 0, 0, 0, 100, 106, 3, 30, 15, 0,
		101, 102, 3, 30, 15, 0, 102, 103, 5, 11, 0, 0, 103, 104, 3, 24, 12, 0,
		104, 106, 1, 0, 0, 0, 105, 100, 1, 0, 0, 0, 105, 101, 1, 0, 0, 0, 106,
		25, 1, 0, 0, 0, 107, 108, 6, 13, -1, 0, 108, 109, 3, 24, 12, 0, 109, 118,
		1, 0, 0, 0, 110, 111, 10, 2, 0, 0, 111, 112, 5, 9, 0, 0, 112, 117, 3, 24,
		12, 0, 113, 114, 10, 1, 0, 0, 114, 115, 5, 10, 0, 0, 115, 117, 3, 24, 12,
		0, 116, 110, 1, 0, 0, 0, 116, 113, 1, 0, 0, 0, 117, 120, 1, 0, 0, 0, 118,
		116, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 27, 1, 0, 0, 0, 120, 118, 1,
		0, 0, 0, 121, 122, 6, 14, -1, 0, 122, 123, 3, 26, 13, 0, 123, 132, 1, 0,
		0, 0, 124, 125, 10, 2, 0, 0, 125, 126, 5, 7, 0, 0, 126, 131, 3, 26, 13,
		0, 127, 128, 10, 1, 0, 0, 128, 129, 5, 8, 0, 0, 129, 131, 3, 26, 13, 0,
		130, 124, 1, 0, 0, 0, 130, 127, 1, 0, 0, 0, 131, 134, 1, 0, 0, 0, 132,
		130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 29, 1, 0, 0, 0, 134, 132, 1,
		0, 0, 0, 135, 144, 5, 3, 0, 0, 136, 137, 5, 8, 0, 0, 137, 144, 3, 30, 15,
		0, 138, 144, 5, 2, 0, 0, 139, 140, 5, 30, 0, 0, 140, 141, 3, 28, 14, 0,
		141, 142, 5, 31, 0, 0, 142, 144, 1, 0, 0, 0, 143, 135, 1, 0, 0, 0, 143,
		136, 1, 0, 0, 0, 143, 138, 1, 0, 0, 0, 143, 139, 1, 0, 0, 0, 144, 31, 1,
		0, 0, 0, 12, 38, 46, 80, 88, 91, 98, 105, 116, 118, 130, 132, 143,
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
	MyGrammarParserSTRING       = 1
	MyGrammarParserIDENTIFIER   = 2
	MyGrammarParserINTEGER      = 3
	MyGrammarParserASSIGN       = 4
	MyGrammarParserCOLON        = 5
	MyGrammarParserCOMMA        = 6
	MyGrammarParserPLUS         = 7
	MyGrammarParserMINUS        = 8
	MyGrammarParserMULTI        = 9
	MyGrammarParserDIVIDE       = 10
	MyGrammarParserPOWERBY      = 11
	MyGrammarParserIF           = 12
	MyGrammarParserElSE         = 13
	MyGrammarParserDO           = 14
	MyGrammarParserBEGIN        = 15
	MyGrammarParserEND          = 16
	MyGrammarParserTHEN         = 17
	MyGrammarParserWHILE        = 18
	MyGrammarParserFOR          = 19
	MyGrammarParserLOOP         = 20
	MyGrammarParserPRINT        = 21
	MyGrammarParserTRUE         = 22
	MyGrammarParserFALSE        = 23
	MyGrammarParserEQ           = 24
	MyGrammarParserNOTEQUALBY   = 25
	MyGrammarParserLT           = 26
	MyGrammarParserGT           = 27
	MyGrammarParserCOT          = 28
	MyGrammarParserSEMICOLON    = 29
	MyGrammarParserOPARAN       = 30
	MyGrammarParserCPARAN       = 31
	MyGrammarParserOBRAC        = 32
	MyGrammarParserCBRAC        = 33
	MyGrammarParserCONDITION_OP = 34
	MyGrammarParserNEXT_PARAM   = 35
	MyGrammarParserEMPTY        = 36
)

// MyGrammarParser rules.
const (
	MyGrammarParserRULE_program             = 0
	MyGrammarParserRULE_statements          = 1
	MyGrammarParserRULE_statement           = 2
	MyGrammarParserRULE_assignmentStatement = 3
	MyGrammarParserRULE_printStatement      = 4
	MyGrammarParserRULE_blockStatement      = 5
	MyGrammarParserRULE_whileStatement      = 6
	MyGrammarParserRULE_methodCall          = 7
	MyGrammarParserRULE_ifStatement         = 8
	MyGrammarParserRULE_variableSetterTypes = 9
	MyGrammarParserRULE_methodCallArguments = 10
	MyGrammarParserRULE_expression          = 11
	MyGrammarParserRULE_powerExpr           = 12
	MyGrammarParserRULE_multipleExpr        = 13
	MyGrammarParserRULE_sumExpr             = 14
	MyGrammarParserRULE_number              = 15
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Statements() IStatementsContext
	EOF() antlr.TerminalNode

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

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserEOF, 0)
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
		p.SetState(32)
		p.Statements()
	}
	{
		p.SetState(33)
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

// IStatementsContext is an interface to support dynamic dispatch.
type IStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

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
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2363396) != 0 {
		{
			p.SetState(35)
			p.Statement()
		}

		p.SetState(40)
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

type WhileStatementBlockContext struct {
	StatementContext
}

func NewWhileStatementBlockContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileStatementBlockContext {
	var p = new(WhileStatementBlockContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *WhileStatementBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementBlockContext) WhileStatement() IWhileStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStatementContext)
}

func (s *WhileStatementBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterWhileStatementBlock(s)
	}
}

func (s *WhileStatementBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitWhileStatementBlock(s)
	}
}

func (s *WhileStatementBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitWhileStatementBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

type MethodCallBlockContext struct {
	StatementContext
}

func NewMethodCallBlockContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MethodCallBlockContext {
	var p = new(MethodCallBlockContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *MethodCallBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodCallBlockContext) MethodCall() IMethodCallContext {
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

func (s *MethodCallBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterMethodCallBlock(s)
	}
}

func (s *MethodCallBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitMethodCallBlock(s)
	}
}

func (s *MethodCallBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitMethodCallBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentStatementBlockContext struct {
	StatementContext
}

func NewAssignmentStatementBlockContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentStatementBlockContext {
	var p = new(AssignmentStatementBlockContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *AssignmentStatementBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementBlockContext) AssignmentStatement() IAssignmentStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentStatementContext)
}

func (s *AssignmentStatementBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterAssignmentStatementBlock(s)
	}
}

func (s *AssignmentStatementBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitAssignmentStatementBlock(s)
	}
}

func (s *AssignmentStatementBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitAssignmentStatementBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfStatementBlockContext struct {
	StatementContext
}

func NewIfStatementBlockContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStatementBlockContext {
	var p = new(IfStatementBlockContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *IfStatementBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementBlockContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *IfStatementBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterIfStatementBlock(s)
	}
}

func (s *IfStatementBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitIfStatementBlock(s)
	}
}

func (s *IfStatementBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitIfStatementBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrintStatementBlockContext struct {
	StatementContext
}

func NewPrintStatementBlockContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintStatementBlockContext {
	var p = new(PrintStatementBlockContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *PrintStatementBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStatementBlockContext) PrintStatement() IPrintStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintStatementContext)
}

func (s *PrintStatementBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterPrintStatementBlock(s)
	}
}

func (s *PrintStatementBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitPrintStatementBlock(s)
	}
}

func (s *PrintStatementBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitPrintStatementBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MyGrammarParserRULE_statement)
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAssignmentStatementBlockContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(41)
			p.AssignmentStatement()
		}

	case 2:
		localctx = NewIfStatementBlockContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(42)
			p.IfStatement()
		}

	case 3:
		localctx = NewWhileStatementBlockContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(43)
			p.WhileStatement()
		}

	case 4:
		localctx = NewMethodCallBlockContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(44)
			p.MethodCall()
		}

	case 5:
		localctx = NewPrintStatementBlockContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(45)
			p.PrintStatement()
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

// IAssignmentStatementContext is an interface to support dynamic dispatch.
type IAssignmentStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	VariableSetterTypes() IVariableSetterTypesContext
	SEMICOLON() antlr.TerminalNode

	// IsAssignmentStatementContext differentiates from other interfaces.
	IsAssignmentStatementContext()
}

type AssignmentStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentStatementContext() *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_assignmentStatement
	return p
}

func InitEmptyAssignmentStatementContext(p *AssignmentStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_assignmentStatement
}

func (*AssignmentStatementContext) IsAssignmentStatementContext() {}

func NewAssignmentStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_assignmentStatement

	return p
}

func (s *AssignmentStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentStatementContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserIDENTIFIER, 0)
}

func (s *AssignmentStatementContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserASSIGN, 0)
}

func (s *AssignmentStatementContext) VariableSetterTypes() IVariableSetterTypesContext {
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

func (s *AssignmentStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserSEMICOLON, 0)
}

func (s *AssignmentStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterAssignmentStatement(s)
	}
}

func (s *AssignmentStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitAssignmentStatement(s)
	}
}

func (s *AssignmentStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitAssignmentStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) AssignmentStatement() (localctx IAssignmentStatementContext) {
	localctx = NewAssignmentStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MyGrammarParserRULE_assignmentStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(MyGrammarParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(49)
		p.Match(MyGrammarParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
		p.VariableSetterTypes()
	}
	{
		p.SetState(51)
		p.Match(MyGrammarParserSEMICOLON)
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

// IPrintStatementContext is an interface to support dynamic dispatch.
type IPrintStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRINT() antlr.TerminalNode
	VariableSetterTypes() IVariableSetterTypesContext

	// IsPrintStatementContext differentiates from other interfaces.
	IsPrintStatementContext()
}

type PrintStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintStatementContext() *PrintStatementContext {
	var p = new(PrintStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_printStatement
	return p
}

func InitEmptyPrintStatementContext(p *PrintStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_printStatement
}

func (*PrintStatementContext) IsPrintStatementContext() {}

func NewPrintStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintStatementContext {
	var p = new(PrintStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_printStatement

	return p
}

func (s *PrintStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintStatementContext) PRINT() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserPRINT, 0)
}

func (s *PrintStatementContext) VariableSetterTypes() IVariableSetterTypesContext {
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

func (s *PrintStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterPrintStatement(s)
	}
}

func (s *PrintStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitPrintStatement(s)
	}
}

func (s *PrintStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitPrintStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) PrintStatement() (localctx IPrintStatementContext) {
	localctx = NewPrintStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MyGrammarParserRULE_printStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Match(MyGrammarParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(54)
		p.VariableSetterTypes()
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

// IBlockStatementContext is an interface to support dynamic dispatch.
type IBlockStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BEGIN() antlr.TerminalNode
	Statements() IStatementsContext
	END() antlr.TerminalNode

	// IsBlockStatementContext differentiates from other interfaces.
	IsBlockStatementContext()
}

type BlockStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockStatementContext() *BlockStatementContext {
	var p = new(BlockStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_blockStatement
	return p
}

func InitEmptyBlockStatementContext(p *BlockStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_blockStatement
}

func (*BlockStatementContext) IsBlockStatementContext() {}

func NewBlockStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockStatementContext {
	var p = new(BlockStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_blockStatement

	return p
}

func (s *BlockStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockStatementContext) BEGIN() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserBEGIN, 0)
}

func (s *BlockStatementContext) Statements() IStatementsContext {
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

func (s *BlockStatementContext) END() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserEND, 0)
}

func (s *BlockStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterBlockStatement(s)
	}
}

func (s *BlockStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitBlockStatement(s)
	}
}

func (s *BlockStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitBlockStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) BlockStatement() (localctx IBlockStatementContext) {
	localctx = NewBlockStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MyGrammarParserRULE_blockStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Match(MyGrammarParserBEGIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(57)
		p.Statements()
	}
	{
		p.SetState(58)
		p.Match(MyGrammarParserEND)
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

// IWhileStatementContext is an interface to support dynamic dispatch.
type IWhileStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHILE() antlr.TerminalNode
	Expression() IExpressionContext
	DO() antlr.TerminalNode
	Statement() IStatementContext
	END() antlr.TerminalNode

	// IsWhileStatementContext differentiates from other interfaces.
	IsWhileStatementContext()
}

type WhileStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStatementContext() *WhileStatementContext {
	var p = new(WhileStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_whileStatement
	return p
}

func InitEmptyWhileStatementContext(p *WhileStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_whileStatement
}

func (*WhileStatementContext) IsWhileStatementContext() {}

func NewWhileStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStatementContext {
	var p = new(WhileStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_whileStatement

	return p
}

func (s *WhileStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStatementContext) WHILE() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserWHILE, 0)
}

func (s *WhileStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhileStatementContext) DO() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserDO, 0)
}

func (s *WhileStatementContext) Statement() IStatementContext {
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

func (s *WhileStatementContext) END() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserEND, 0)
}

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterWhileStatement(s)
	}
}

func (s *WhileStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitWhileStatement(s)
	}
}

func (s *WhileStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitWhileStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) WhileStatement() (localctx IWhileStatementContext) {
	localctx = NewWhileStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MyGrammarParserRULE_whileStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.Match(MyGrammarParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(61)
		p.Expression()
	}
	{
		p.SetState(62)
		p.Match(MyGrammarParserDO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(63)
		p.Statement()
	}
	{
		p.SetState(64)
		p.Match(MyGrammarParserEND)
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
	IDENTIFIER() antlr.TerminalNode
	OPARAN() antlr.TerminalNode
	MethodCallArguments() IMethodCallArgumentsContext
	CPARAN() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode

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

func (s *MethodCallContext) OPARAN() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserOPARAN, 0)
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

func (s *MethodCallContext) CPARAN() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserCPARAN, 0)
}

func (s *MethodCallContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserSEMICOLON, 0)
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
	p.EnterRule(localctx, 14, MyGrammarParserRULE_methodCall)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Match(MyGrammarParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(67)
		p.Match(MyGrammarParserOPARAN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(68)
		p.MethodCallArguments()
	}
	{
		p.SetState(69)
		p.Match(MyGrammarParserCPARAN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(70)
		p.Match(MyGrammarParserSEMICOLON)
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

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	Expression() IExpressionContext
	THEN() antlr.TerminalNode
	Statement() IStatementContext

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyGrammarParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyGrammarParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserIF, 0)
}

func (s *IfStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatementContext) THEN() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserTHEN, 0)
}

func (s *IfStatementContext) Statement() IStatementContext {
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

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyGrammarListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyGrammarVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyGrammarParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MyGrammarParserRULE_ifStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(MyGrammarParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(73)
		p.Expression()
	}
	{
		p.SetState(74)
		p.Match(MyGrammarParserTHEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(75)
		p.Statement()
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
	p.EnterRule(localctx, 18, MyGrammarParserRULE_variableSetterTypes)
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(77)
			p.Match(MyGrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(78)
			p.MethodCall()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(79)
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
	p.EnterRule(localctx, 20, MyGrammarParserRULE_methodCallArguments)
	var _la int

	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MyGrammarParserCPARAN:
		p.EnterOuterAlt(localctx, 1)

	case MyGrammarParserSTRING, MyGrammarParserIDENTIFIER, MyGrammarParserINTEGER, MyGrammarParserMINUS, MyGrammarParserOPARAN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(83)
			p.Expression()
		}
		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MyGrammarParserCOMMA {
			{
				p.SetState(84)
				p.Match(MyGrammarParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(85)
				p.Expression()
			}

			p.SetState(90)
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
	SumExpr() ISumExprContext

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

func (s *ExpressionContext) SumExpr() ISumExprContext {
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
	p.EnterRule(localctx, 22, MyGrammarParserRULE_expression)
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(93)
			p.Match(MyGrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(94)
			p.Match(MyGrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(95)
			p.MethodCall()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(96)
			p.Match(MyGrammarParserINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(97)
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
	p.EnterRule(localctx, 24, MyGrammarParserRULE_powerExpr)
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPowerExprDefaultContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(100)
			p.Number()
		}

	case 2:
		localctx = NewPowerExprPowerContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(101)
			p.Number()
		}
		{
			p.SetState(102)
			p.Match(MyGrammarParserPOWERBY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
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
	_startState := 26
	p.EnterRecursionRule(localctx, 26, MyGrammarParserRULE_multipleExpr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewMultipleExprDefaultContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(108)
		p.PowerExpr()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(118)
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
			p.SetState(116)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultipleExprMultiContext(p, NewMultipleExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MyGrammarParserRULE_multipleExpr)
				p.SetState(110)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(111)
					p.Match(MyGrammarParserMULTI)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(112)
					p.PowerExpr()
				}

			case 2:
				localctx = NewMultipleExprDivideContext(p, NewMultipleExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MyGrammarParserRULE_multipleExpr)
				p.SetState(113)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(114)
					p.Match(MyGrammarParserDIVIDE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(115)
					p.PowerExpr()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(120)
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
	_startState := 28
	p.EnterRecursionRule(localctx, 28, MyGrammarParserRULE_sumExpr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewSumExprDefaultContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(122)
		p.multipleExpr(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(132)
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
			p.SetState(130)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewSumExprPlusContext(p, NewSumExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MyGrammarParserRULE_sumExpr)
				p.SetState(124)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(125)
					p.Match(MyGrammarParserPLUS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(126)
					p.multipleExpr(0)
				}

			case 2:
				localctx = NewSumExprMinusContext(p, NewSumExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MyGrammarParserRULE_sumExpr)
				p.SetState(127)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(128)
					p.Match(MyGrammarParserMINUS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(129)
					p.multipleExpr(0)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(134)
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

func (s *NumberParenthesesContext) OPARAN() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserOPARAN, 0)
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

func (s *NumberParenthesesContext) CPARAN() antlr.TerminalNode {
	return s.GetToken(MyGrammarParserCPARAN, 0)
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
	p.EnterRule(localctx, 30, MyGrammarParserRULE_number)
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MyGrammarParserINTEGER:
		localctx = NewNumberDefaultContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(135)
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
			p.SetState(136)
			p.Match(MyGrammarParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)
			p.Number()
		}

	case MyGrammarParserIDENTIFIER:
		localctx = NewNumberIdentifierContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(138)
			p.Match(MyGrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MyGrammarParserOPARAN:
		localctx = NewNumberParenthesesContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(139)
			p.Match(MyGrammarParserOPARAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.sumExpr(0)
		}
		{
			p.SetState(141)
			p.Match(MyGrammarParserCPARAN)
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
	case 13:
		var t *MultipleExprContext = nil
		if localctx != nil {
			t = localctx.(*MultipleExprContext)
		}
		return p.MultipleExpr_Sempred(t, predIndex)

	case 14:
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
