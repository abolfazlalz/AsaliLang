package project

import (
	"asalicompiler/parsing"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"log"
	"math"
	"strconv"
)

type Visitor struct {
	*parsing.BaseMyGrammarVisitor
	method *BuildInMethod
	vars   map[string]interface{}
}

func NewVisitor() *Visitor {
	return &Visitor{
		vars: map[string]interface{}{
			"PI": math.Pi,
		},
		method:               NewBuildInMethod(),
		BaseMyGrammarVisitor: &parsing.BaseMyGrammarVisitor{},
	}
}

func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *Visitor) VisitParse(ctx *parsing.ParseContext) interface{} {
	for _, c := range ctx.Block().AllStat() {
		v.Visit(c)
	}
	return nil
}
func (v *Visitor) VisitStat(ctx *parsing.StatContext) interface{} {
	return v.Visit(ctx.GetChild(0).(antlr.ParseTree))
}

func (v *Visitor) VisitAssignment(ctx *parsing.AssignmentContext) interface{} {
	v.vars[ctx.ID().GetText()] = v.Visit(ctx.Expr())
	return nil
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
	case parsing.MuLexerMULT:
		return left * right
	case parsing.MuLexerDIV:
		return left / right
	case parsing.MuLexerMOD:
		return math.Mod(left, right)
	}

	panic("Undefined operator")
}

func (v *Visitor) VisitAdditiveExpr(ctx *parsing.AdditiveExprContext) interface{} {
	left := v.Visit(ctx.Expr(0)).(float64)
	right := v.Visit(ctx.Expr(1)).(float64)

	switch ctx.GetOp().GetTokenType() {
	case parsing.MuLexerPLUS:
		return left + right
	case parsing.MuLexerMINUS:
		return left - right
	}
	panic("Undefined operator")
}

func (v *Visitor) VisitRelationalExpr(ctx *parsing.RelationalExprContext) interface{} {
	left := v.Visit(ctx.Expr(0)).(float64)
	right := v.Visit(ctx.Expr(1)).(float64)

	switch ctx.GetOp().GetTokenType() {
	case parsing.MuLexerLT:
		return left < right
	case parsing.MuLexerLTEQ:
		return left <= right
	case parsing.MuLexerGT:
		return left > right
	case parsing.MuLexerGTEQ:
		return left >= right
	}
	panic("Undefined operator")
}

func (v *Visitor) VisitEqualityExpr(ctx *parsing.EqualityExprContext) interface{} {
	left := v.Visit(ctx.Expr(0))
	right := v.Visit(ctx.Expr(1))

	switch ctx.GetOp().GetTokenType() {
	case parsing.MuLexerEQ:
		return left == right
	case parsing.MuLexerNEQ:
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

func (v *Visitor) VisitLog(ctx *parsing.LogContext) interface{} {
	log.Print(v.Visit(ctx.Expr()))
	return nil
}
