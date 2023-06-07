// Code generated from MyGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parsing // MyGrammar
import "github.com/antlr4-go/antlr/v4"

type BaseMyGrammarVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMyGrammarVisitor) VisitParse(ctx *ParseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitStat(ctx *StatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitIf_stat(ctx *If_statContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitCondition_block(ctx *Condition_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitStat_block(ctx *Stat_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitWhile_stat(ctx *While_statContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitForStat(ctx *ForStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitLoopStat(ctx *LoopStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitLog(ctx *LogContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitUnaryMinusExpr(ctx *UnaryMinusExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitMultiplicationExpr(ctx *MultiplicationExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitAtomExpr(ctx *AtomExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitOrExpr(ctx *OrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitAdditiveExpr(ctx *AdditiveExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitPowExpr(ctx *PowExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitRelationalExpr(ctx *RelationalExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitEqualityExpr(ctx *EqualityExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitAndExpr(ctx *AndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitParExpr(ctx *ParExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitNumberAtom(ctx *NumberAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitBooleanAtom(ctx *BooleanAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitIdAtom(ctx *IdAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitStringAtom(ctx *StringAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitNilAtom(ctx *NilAtomContext) interface{} {
	return v.VisitChildren(ctx)
}
