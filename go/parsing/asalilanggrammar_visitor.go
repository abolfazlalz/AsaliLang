// Code generated from ./AsaliLangGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parsing // AsaliLangGrammar
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by AsaliLangGrammarParser.
type AsaliLangGrammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by AsaliLangGrammarParser#parse.
	VisitParse(ctx *ParseContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#stat.
	VisitStat(ctx *StatContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#ifStat.
	VisitIfStat(ctx *IfStatContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#conditionBlock.
	VisitConditionBlock(ctx *ConditionBlockContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#statBlock.
	VisitStatBlock(ctx *StatBlockContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#whileStat.
	VisitWhileStat(ctx *WhileStatContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#forStat.
	VisitForStat(ctx *ForStatContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#loopStat.
	VisitLoopStat(ctx *LoopStatContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#methodCallStat.
	VisitMethodCallStat(ctx *MethodCallStatContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#defineFuncStats.
	VisitDefineFuncStats(ctx *DefineFuncStatsContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#returnState.
	VisitReturnState(ctx *ReturnStateContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#methodCall.
	VisitMethodCall(ctx *MethodCallContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#inlineMethodCall.
	VisitInlineMethodCall(ctx *InlineMethodCallContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#methodCallArguments.
	VisitMethodCallArguments(ctx *MethodCallArgumentsContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#defineFuncArguments.
	VisitDefineFuncArguments(ctx *DefineFuncArgumentsContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#methodCallExpr.
	VisitMethodCallExpr(ctx *MethodCallExprContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#notExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#unaryMinusExpr.
	VisitUnaryMinusExpr(ctx *UnaryMinusExprContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#multiplicationExpr.
	VisitMultiplicationExpr(ctx *MultiplicationExprContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#atomExpr.
	VisitAtomExpr(ctx *AtomExprContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#orExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#additiveExpr.
	VisitAdditiveExpr(ctx *AdditiveExprContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#powExpr.
	VisitPowExpr(ctx *PowExprContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#relationalExpr.
	VisitRelationalExpr(ctx *RelationalExprContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#equalityExpr.
	VisitEqualityExpr(ctx *EqualityExprContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#andExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#parExpr.
	VisitParExpr(ctx *ParExprContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#numberAtom.
	VisitNumberAtom(ctx *NumberAtomContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#booleanAtom.
	VisitBooleanAtom(ctx *BooleanAtomContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#idAtom.
	VisitIdAtom(ctx *IdAtomContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#stringAtom.
	VisitStringAtom(ctx *StringAtomContext) interface{}

	// Visit a parse tree produced by AsaliLangGrammarParser#nilAtom.
	VisitNilAtom(ctx *NilAtomContext) interface{}
}
