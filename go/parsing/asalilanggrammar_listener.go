// Code generated from ./AsaliLangGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parsing // AsaliLangGrammar
import "github.com/antlr4-go/antlr/v4"

// AsaliLangGrammarListener is a complete listener for a parse tree produced by AsaliLangGrammarParser.
type AsaliLangGrammarListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStat is called when entering the stat production.
	EnterStat(c *StatContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterIfStat is called when entering the ifStat production.
	EnterIfStat(c *IfStatContext)

	// EnterConditionBlock is called when entering the conditionBlock production.
	EnterConditionBlock(c *ConditionBlockContext)

	// EnterStatBlock is called when entering the statBlock production.
	EnterStatBlock(c *StatBlockContext)

	// EnterWhileStat is called when entering the whileStat production.
	EnterWhileStat(c *WhileStatContext)

	// EnterForStat is called when entering the forStat production.
	EnterForStat(c *ForStatContext)

	// EnterLoopStat is called when entering the loopStat production.
	EnterLoopStat(c *LoopStatContext)

	// EnterMethodCallStat is called when entering the methodCallStat production.
	EnterMethodCallStat(c *MethodCallStatContext)

	// EnterDefineFuncStats is called when entering the defineFuncStats production.
	EnterDefineFuncStats(c *DefineFuncStatsContext)

	// EnterReturnState is called when entering the returnState production.
	EnterReturnState(c *ReturnStateContext)

	// EnterMethodCall is called when entering the methodCall production.
	EnterMethodCall(c *MethodCallContext)

	// EnterInlineMethodCall is called when entering the inlineMethodCall production.
	EnterInlineMethodCall(c *InlineMethodCallContext)

	// EnterMethodCallArguments is called when entering the methodCallArguments production.
	EnterMethodCallArguments(c *MethodCallArgumentsContext)

	// EnterDefineFuncArguments is called when entering the defineFuncArguments production.
	EnterDefineFuncArguments(c *DefineFuncArgumentsContext)

	// EnterMethodCallExpr is called when entering the methodCallExpr production.
	EnterMethodCallExpr(c *MethodCallExprContext)

	// EnterNotExpr is called when entering the notExpr production.
	EnterNotExpr(c *NotExprContext)

	// EnterUnaryMinusExpr is called when entering the unaryMinusExpr production.
	EnterUnaryMinusExpr(c *UnaryMinusExprContext)

	// EnterMultiplicationExpr is called when entering the multiplicationExpr production.
	EnterMultiplicationExpr(c *MultiplicationExprContext)

	// EnterAtomExpr is called when entering the atomExpr production.
	EnterAtomExpr(c *AtomExprContext)

	// EnterOrExpr is called when entering the orExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterAdditiveExpr is called when entering the additiveExpr production.
	EnterAdditiveExpr(c *AdditiveExprContext)

	// EnterPowExpr is called when entering the powExpr production.
	EnterPowExpr(c *PowExprContext)

	// EnterRelationalExpr is called when entering the relationalExpr production.
	EnterRelationalExpr(c *RelationalExprContext)

	// EnterEqualityExpr is called when entering the equalityExpr production.
	EnterEqualityExpr(c *EqualityExprContext)

	// EnterAndExpr is called when entering the andExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterParExpr is called when entering the parExpr production.
	EnterParExpr(c *ParExprContext)

	// EnterNumberAtom is called when entering the numberAtom production.
	EnterNumberAtom(c *NumberAtomContext)

	// EnterBooleanAtom is called when entering the booleanAtom production.
	EnterBooleanAtom(c *BooleanAtomContext)

	// EnterIdAtom is called when entering the idAtom production.
	EnterIdAtom(c *IdAtomContext)

	// EnterStringAtom is called when entering the stringAtom production.
	EnterStringAtom(c *StringAtomContext)

	// EnterNilAtom is called when entering the nilAtom production.
	EnterNilAtom(c *NilAtomContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitIfStat is called when exiting the ifStat production.
	ExitIfStat(c *IfStatContext)

	// ExitConditionBlock is called when exiting the conditionBlock production.
	ExitConditionBlock(c *ConditionBlockContext)

	// ExitStatBlock is called when exiting the statBlock production.
	ExitStatBlock(c *StatBlockContext)

	// ExitWhileStat is called when exiting the whileStat production.
	ExitWhileStat(c *WhileStatContext)

	// ExitForStat is called when exiting the forStat production.
	ExitForStat(c *ForStatContext)

	// ExitLoopStat is called when exiting the loopStat production.
	ExitLoopStat(c *LoopStatContext)

	// ExitMethodCallStat is called when exiting the methodCallStat production.
	ExitMethodCallStat(c *MethodCallStatContext)

	// ExitDefineFuncStats is called when exiting the defineFuncStats production.
	ExitDefineFuncStats(c *DefineFuncStatsContext)

	// ExitReturnState is called when exiting the returnState production.
	ExitReturnState(c *ReturnStateContext)

	// ExitMethodCall is called when exiting the methodCall production.
	ExitMethodCall(c *MethodCallContext)

	// ExitInlineMethodCall is called when exiting the inlineMethodCall production.
	ExitInlineMethodCall(c *InlineMethodCallContext)

	// ExitMethodCallArguments is called when exiting the methodCallArguments production.
	ExitMethodCallArguments(c *MethodCallArgumentsContext)

	// ExitDefineFuncArguments is called when exiting the defineFuncArguments production.
	ExitDefineFuncArguments(c *DefineFuncArgumentsContext)

	// ExitMethodCallExpr is called when exiting the methodCallExpr production.
	ExitMethodCallExpr(c *MethodCallExprContext)

	// ExitNotExpr is called when exiting the notExpr production.
	ExitNotExpr(c *NotExprContext)

	// ExitUnaryMinusExpr is called when exiting the unaryMinusExpr production.
	ExitUnaryMinusExpr(c *UnaryMinusExprContext)

	// ExitMultiplicationExpr is called when exiting the multiplicationExpr production.
	ExitMultiplicationExpr(c *MultiplicationExprContext)

	// ExitAtomExpr is called when exiting the atomExpr production.
	ExitAtomExpr(c *AtomExprContext)

	// ExitOrExpr is called when exiting the orExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitAdditiveExpr is called when exiting the additiveExpr production.
	ExitAdditiveExpr(c *AdditiveExprContext)

	// ExitPowExpr is called when exiting the powExpr production.
	ExitPowExpr(c *PowExprContext)

	// ExitRelationalExpr is called when exiting the relationalExpr production.
	ExitRelationalExpr(c *RelationalExprContext)

	// ExitEqualityExpr is called when exiting the equalityExpr production.
	ExitEqualityExpr(c *EqualityExprContext)

	// ExitAndExpr is called when exiting the andExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitParExpr is called when exiting the parExpr production.
	ExitParExpr(c *ParExprContext)

	// ExitNumberAtom is called when exiting the numberAtom production.
	ExitNumberAtom(c *NumberAtomContext)

	// ExitBooleanAtom is called when exiting the booleanAtom production.
	ExitBooleanAtom(c *BooleanAtomContext)

	// ExitIdAtom is called when exiting the idAtom production.
	ExitIdAtom(c *IdAtomContext)

	// ExitStringAtom is called when exiting the stringAtom production.
	ExitStringAtom(c *StringAtomContext)

	// ExitNilAtom is called when exiting the nilAtom production.
	ExitNilAtom(c *NilAtomContext)
}
