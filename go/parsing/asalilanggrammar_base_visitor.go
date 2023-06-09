// Code generated from ./AsaliLangGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parsing // AsaliLangGrammar
import "github.com/antlr4-go/antlr/v4"

type BaseAsaliLangGrammarVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseAsaliLangGrammarVisitor) VisitParse(ctx *ParseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitStat(ctx *StatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitIfStat(ctx *IfStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitConditionBlock(ctx *ConditionBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitStatBlock(ctx *StatBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitWhileStat(ctx *WhileStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitForStat(ctx *ForStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitLoopStat(ctx *LoopStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitMethodCallStat(ctx *MethodCallStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitDefineFuncStats(ctx *DefineFuncStatsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitMethodCall(ctx *MethodCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitInlineMethodCall(ctx *InlineMethodCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitMethodCallArguments(ctx *MethodCallArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitDefineFuncArguments(ctx *DefineFuncArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitMethodCallExpr(ctx *MethodCallExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitUnaryMinusExpr(ctx *UnaryMinusExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitMultiplicationExpr(ctx *MultiplicationExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitAtomExpr(ctx *AtomExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitOrExpr(ctx *OrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitAdditiveExpr(ctx *AdditiveExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitPowExpr(ctx *PowExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitRelationalExpr(ctx *RelationalExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitEqualityExpr(ctx *EqualityExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitAndExpr(ctx *AndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitParExpr(ctx *ParExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitNumberAtom(ctx *NumberAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitBooleanAtom(ctx *BooleanAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitIdAtom(ctx *IdAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitStringAtom(ctx *StringAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAsaliLangGrammarVisitor) VisitNilAtom(ctx *NilAtomContext) interface{} {
	return v.VisitChildren(ctx)
}
