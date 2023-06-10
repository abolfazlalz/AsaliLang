package project

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	parsing "github.com/abolfazlalz/AsaliLang/parsing"
	"github.com/antlr4-go/antlr/v4"
)

type Visitor struct {
	*parsing.BaseAsaliLangGrammarVisitor
	method *BuildInMethod
	vars   map[string]interface{}
}

func NewVisitor() *Visitor {
	return &Visitor{
		vars: map[string]interface{}{
			"PI": math.Pi,
		},
		method:                      NewBuildInMethod(),
		BaseAsaliLangGrammarVisitor: &parsing.BaseAsaliLangGrammarVisitor{},
	}
}

func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *Visitor) VisitBlock(ctx *parsing.BlockContext) interface{} {
	for _, c := range ctx.AllStat() {
		v.Visit(c)
	}
	return nil
}

func (v *Visitor) VisitStatBlock(ctx *parsing.StatBlockContext) interface{} {
	if ctx.Stat() != nil {
		v.Visit(ctx.Stat())
		return nil
	}
	for _, c := range ctx.Block().AllStat() {
		v.Visit(c)
	}
	return nil
}

func (v *Visitor) VisitParse(ctx *parsing.ParseContext) interface{} {
	v.Visit(ctx.Block())
	return nil
}
func (v *Visitor) VisitStat(ctx *parsing.StatContext) interface{} {
	return v.Visit(ctx.GetChild(0).(antlr.ParseTree))
}

func (v *Visitor) VisitAssignment(ctx *parsing.AssignmentContext) interface{} {
	v.vars[ctx.ID().GetText()] = v.Visit(ctx.Expr())
	return nil
}

func (v *Visitor) VisitMethodCallExpr(ctx *parsing.MethodCallExprContext) interface{} {
	return v.Visit(ctx.InlineMethodCall())
}

func (v *Visitor) VisitPowExpr(ctx *parsing.PowExprContext) interface{} {
	left := v.Visit(ctx.Expr(0))
	right := v.Visit(ctx.Expr(1))
	return math.Pow(left.(float64), right.(float64))
}

func (v *Visitor) VisitUnaryMinusExpr(ctx *parsing.UnaryMinusExprContext) interface{} {
	result := v.Visit(ctx.Expr())
	if value, ok := result.(float64); ok {
		return value * -1
	} else if value, ok := result.(int); ok {
		return value * -1
	} else if value, ok := result.(int64); ok {
		return value * -1
	}
	panic(fmt.Sprintf("Value to minus must be integer or float"))
}

func (v *Visitor) VisitMultiplicationExpr(ctx *parsing.MultiplicationExprContext) interface{} {
	left := v.Visit(ctx.Expr(0)).(float64)
	right := v.Visit(ctx.Expr(1)).(float64)

	switch ctx.GetOp().GetTokenType() {
	case parsing.AsaliLangGrammarLexerMULT:
		return left * right
	case parsing.AsaliLangGrammarLexerDIV:
		return left / right
	case parsing.AsaliLangGrammarLexerMOD:
		return math.Mod(left, right)
	}

	panic("Undefined operator")
}

func (v *Visitor) VisitAdditiveExpr(ctx *parsing.AdditiveExprContext) interface{} {
	left := v.Visit(ctx.Expr(0))
	right := v.Visit(ctx.Expr(1))

	switch ctx.GetOp().GetTokenType() {
	case parsing.AsaliLangGrammarLexerPLUS:
		valLeft, okLeft := left.(float64)
		valRight, okRight := right.(float64)
		if okRight && okLeft {
			return valRight + valLeft
		}
		return fmt.Sprintf("%v", left) + fmt.Sprintf("%v", right)
	case parsing.AsaliLangGrammarLexerMINUS:
		return left.(float64) - right.(float64)
	}
	panic("Undefined operator")
}

func (v *Visitor) VisitRelationalExpr(ctx *parsing.RelationalExprContext) interface{} {
	left := v.Visit(ctx.Expr(0)).(float64)
	right := v.Visit(ctx.Expr(1)).(float64)

	switch ctx.GetOp().GetTokenType() {
	case parsing.AsaliLangGrammarLexerLT:
		return left < right
	case parsing.AsaliLangGrammarLexerLTEQ:
		return left <= right
	case parsing.AsaliLangGrammarLexerGT:
		return left > right
	case parsing.AsaliLangGrammarLexerGTEQ:
		return left >= right
	}
	panic("Undefined operator")
}

func (v *Visitor) VisitEqualityExpr(ctx *parsing.EqualityExprContext) interface{} {
	left := v.Visit(ctx.Expr(0))
	right := v.Visit(ctx.Expr(1))

	switch ctx.GetOp().GetTokenType() {
	case parsing.AsaliLangGrammarLexerEQ:
		return left == right
	case parsing.AsaliLangGrammarLexerNEQ:
		return left != right
	}
	panic("Undefined operator")
}

func (v *Visitor) VisitAndExpr(ctx *parsing.AndExprContext) interface{} {
	left := newValue(v.Visit(ctx.Expr(0))).asBoolean()
	right := newValue(v.Visit(ctx.Expr(1))).asBoolean()
	return right && left
}

func (v *Visitor) VisitOrExpr(ctx *parsing.OrExprContext) interface{} {
	left := newValue(v.Visit(ctx.Expr(0))).asBoolean()
	right := newValue(v.Visit(ctx.Expr(1))).asBoolean()
	return right || left
}

func (v *Visitor) VisitAtomExpr(ctx *parsing.AtomExprContext) interface{} {
	return v.Visit(ctx.Atom())
}

func (v *Visitor) VisitIdAtom(ctx *parsing.IdAtomContext) interface{} {
	return v.vars[ctx.ID().GetText()]
}

func (v *Visitor) VisitStringAtom(ctx *parsing.StringAtomContext) interface{} {
	str := ctx.GetText()
	str = str[1 : len(str)-1]
	str = strings.ReplaceAll(str, "\"\"", "\"")
	return str
}

func (v *Visitor) VisitNotExpr(ctx *parsing.NotExprContext) interface{} {
	value := v.Visit(ctx.Expr())
	return !newValue(value).asBoolean()
}

func (v *Visitor) VisitParExpr(ctx *parsing.ParExprContext) interface{} {
	return v.Visit(ctx.Expr())
}

func (v *Visitor) VisitNumberAtom(ctx *parsing.NumberAtomContext) interface{} {
	if ctx.INT() != nil {
		val, _ := strconv.ParseFloat(ctx.INT().GetText(), 64)
		return val
	} else if ctx.FLOAT() != nil {
		val, _ := strconv.ParseFloat(ctx.FLOAT().GetText(), 64)
		return val
	}
	return nil
}

func (v *Visitor) VisitBooleanAtom(ctx *parsing.BooleanAtomContext) interface{} {
	return ctx.TRUE() != nil
}

func (v *Visitor) VisitMethodCall(ctx *parsing.MethodCallContext) interface{} {
	methodName := ctx.ID().GetText()
	return v.handleMethodCall(methodName, ctx.MethodCallArguments())
}

func (v *Visitor) VisitInlineMethodCall(ctx *parsing.InlineMethodCallContext) interface{} {
	id := ctx.ID().GetText()
	return v.handleMethodCall(id, ctx.MethodCallArguments())
}

func (v *Visitor) handleMethodCall(methodName string, argsContext parsing.IMethodCallArgumentsContext) interface{} {
	args := v.Visit(argsContext).([]interface{})
	if methodName == "print" {
		v.method.print(args...)
	} else if methodName == "sin" {
		return v.method.sin(args...)
	}
	return nil
}

func (v *Visitor) VisitMethodCallArguments(ctx *parsing.MethodCallArgumentsContext) interface{} {
	exprList := ctx.AllExpr()
	result := make([]interface{}, len(exprList))
	for i, exprList := range exprList {
		result[i] = v.Visit(exprList)
	}
	return result
}

func (v *Visitor) VisitMethodCallStat(ctx *parsing.MethodCallStatContext) interface{} {
	return v.Visit(ctx.MethodCall())
}

func (v *Visitor) VisitForStat(ctx *parsing.ForStatContext) interface{} {
	varName := ctx.ID().GetText()
	start := newValue(v.Visit(ctx.Expr(0))).asFloat()
	end := newValue(v.Visit(ctx.Expr(1))).asFloat()
	lastVariable := v.vars[varName]
	for v.vars[varName] = start; toFloat(v.vars[varName]) < end; {
		v.Visit(ctx.StatBlock())
		v.vars[varName] = toFloat(v.vars[varName]) + 1
	}
	v.vars[varName] = lastVariable
	return nil
}

func (v *Visitor) VisitLoopStat(ctx *parsing.LoopStatContext) interface{} {
	varName := ctx.ID().GetText()
	end := newValue(v.Visit(ctx.Expr())).asFloat()

	lastVariable := v.vars[varName]

	for v.vars[varName] = 0; toFloat(v.vars[varName]) < end; {
		v.vars[varName] = toFloat(v.vars[varName]) + 1
		v.Visit(ctx.StatBlock())
	}

	v.vars[varName] = lastVariable

	return nil
}

func (v *Visitor) VisitWhileStat(ctx *parsing.WhileStatContext) interface{} {
	result := toBoolean(v.Visit(ctx.Expr()))
	for result {
		result = toBoolean(v.Visit(ctx.Expr()))
		v.Visit(ctx.StatBlock())
	}
	return nil
}

func (v *Visitor) VisitIfStat(ctx *parsing.IfStatContext) interface{} {
	conditions := ctx.AllConditionBlock()
	evaluatedBlock := false
	for _, condition := range conditions {
		evaluated := v.Visit(condition.Expr())
		if toBoolean(evaluated) {
			evaluatedBlock = true
			v.Visit(condition.StatBlock())
			break
		}
	}

	if !evaluatedBlock && ctx.StatBlock() != nil {
		v.Visit(ctx.StatBlock())
	}

	return nil
}
