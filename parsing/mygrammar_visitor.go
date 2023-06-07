// Code generated from MyGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parsing // MyGrammar
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by MyGrammarParser.
type MyGrammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MyGrammarParser#parse.
	VisitParse(ctx *ParseContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#stat.
	VisitStat(ctx *StatContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#if_stat.
	VisitIf_stat(ctx *If_statContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#condition_block.
	VisitCondition_block(ctx *Condition_blockContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#stat_block.
	VisitStat_block(ctx *Stat_blockContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#while_stat.
	VisitWhile_stat(ctx *While_statContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#forStat.
	VisitForStat(ctx *ForStatContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#log.
	VisitLog(ctx *LogContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#notExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#unaryMinusExpr.
	VisitUnaryMinusExpr(ctx *UnaryMinusExprContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#multiplicationExpr.
	VisitMultiplicationExpr(ctx *MultiplicationExprContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#atomExpr.
	VisitAtomExpr(ctx *AtomExprContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#orExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#additiveExpr.
	VisitAdditiveExpr(ctx *AdditiveExprContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#powExpr.
	VisitPowExpr(ctx *PowExprContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#relationalExpr.
	VisitRelationalExpr(ctx *RelationalExprContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#equalityExpr.
	VisitEqualityExpr(ctx *EqualityExprContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#andExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#parExpr.
	VisitParExpr(ctx *ParExprContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#numberAtom.
	VisitNumberAtom(ctx *NumberAtomContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#booleanAtom.
	VisitBooleanAtom(ctx *BooleanAtomContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#idAtom.
	VisitIdAtom(ctx *IdAtomContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#stringAtom.
	VisitStringAtom(ctx *StringAtomContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#nilAtom.
	VisitNilAtom(ctx *NilAtomContext) interface{}
}
