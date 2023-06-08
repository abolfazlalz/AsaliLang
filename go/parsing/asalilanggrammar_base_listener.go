// Code generated from AsaliLangGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parsing // AsaliLangGrammar
import "github.com/antlr4-go/antlr/v4"

// BaseAsaliLangGrammarListener is a complete listener for a parse tree produced by AsaliLangGrammarParser.
type BaseAsaliLangGrammarListener struct{}

var _ AsaliLangGrammarListener = &BaseAsaliLangGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAsaliLangGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAsaliLangGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAsaliLangGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAsaliLangGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BaseAsaliLangGrammarListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BaseAsaliLangGrammarListener) ExitParse(ctx *ParseContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseAsaliLangGrammarListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseAsaliLangGrammarListener) ExitBlock(ctx *BlockContext) {}

// EnterStat is called when production stat is entered.
func (s *BaseAsaliLangGrammarListener) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *BaseAsaliLangGrammarListener) ExitStat(ctx *StatContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseAsaliLangGrammarListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseAsaliLangGrammarListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterIfStat is called when production ifStat is entered.
func (s *BaseAsaliLangGrammarListener) EnterIfStat(ctx *IfStatContext) {}

// ExitIfStat is called when production ifStat is exited.
func (s *BaseAsaliLangGrammarListener) ExitIfStat(ctx *IfStatContext) {}

// EnterCondition_block is called when production condition_block is entered.
func (s *BaseAsaliLangGrammarListener) EnterCondition_block(ctx *Condition_blockContext) {}

// ExitCondition_block is called when production condition_block is exited.
func (s *BaseAsaliLangGrammarListener) ExitCondition_block(ctx *Condition_blockContext) {}

// EnterStat_block is called when production stat_block is entered.
func (s *BaseAsaliLangGrammarListener) EnterStat_block(ctx *Stat_blockContext) {}

// ExitStat_block is called when production stat_block is exited.
func (s *BaseAsaliLangGrammarListener) ExitStat_block(ctx *Stat_blockContext) {}

// EnterWhileStat is called when production whileStat is entered.
func (s *BaseAsaliLangGrammarListener) EnterWhileStat(ctx *WhileStatContext) {}

// ExitWhileStat is called when production whileStat is exited.
func (s *BaseAsaliLangGrammarListener) ExitWhileStat(ctx *WhileStatContext) {}

// EnterForStat is called when production forStat is entered.
func (s *BaseAsaliLangGrammarListener) EnterForStat(ctx *ForStatContext) {}

// ExitForStat is called when production forStat is exited.
func (s *BaseAsaliLangGrammarListener) ExitForStat(ctx *ForStatContext) {}

// EnterLoopStat is called when production loopStat is entered.
func (s *BaseAsaliLangGrammarListener) EnterLoopStat(ctx *LoopStatContext) {}

// ExitLoopStat is called when production loopStat is exited.
func (s *BaseAsaliLangGrammarListener) ExitLoopStat(ctx *LoopStatContext) {}

// EnterMethodCallStat is called when production methodCallStat is entered.
func (s *BaseAsaliLangGrammarListener) EnterMethodCallStat(ctx *MethodCallStatContext) {}

// ExitMethodCallStat is called when production methodCallStat is exited.
func (s *BaseAsaliLangGrammarListener) ExitMethodCallStat(ctx *MethodCallStatContext) {}

// EnterMethodCall is called when production methodCall is entered.
func (s *BaseAsaliLangGrammarListener) EnterMethodCall(ctx *MethodCallContext) {}

// ExitMethodCall is called when production methodCall is exited.
func (s *BaseAsaliLangGrammarListener) ExitMethodCall(ctx *MethodCallContext) {}

// EnterInlineMethodCall is called when production inlineMethodCall is entered.
func (s *BaseAsaliLangGrammarListener) EnterInlineMethodCall(ctx *InlineMethodCallContext) {}

// ExitInlineMethodCall is called when production inlineMethodCall is exited.
func (s *BaseAsaliLangGrammarListener) ExitInlineMethodCall(ctx *InlineMethodCallContext) {}

// EnterMethodCallArguments is called when production methodCallArguments is entered.
func (s *BaseAsaliLangGrammarListener) EnterMethodCallArguments(ctx *MethodCallArgumentsContext) {}

// ExitMethodCallArguments is called when production methodCallArguments is exited.
func (s *BaseAsaliLangGrammarListener) ExitMethodCallArguments(ctx *MethodCallArgumentsContext) {}

// EnterMethodCallExpr is called when production methodCallExpr is entered.
func (s *BaseAsaliLangGrammarListener) EnterMethodCallExpr(ctx *MethodCallExprContext) {}

// ExitMethodCallExpr is called when production methodCallExpr is exited.
func (s *BaseAsaliLangGrammarListener) ExitMethodCallExpr(ctx *MethodCallExprContext) {}

// EnterNotExpr is called when production notExpr is entered.
func (s *BaseAsaliLangGrammarListener) EnterNotExpr(ctx *NotExprContext) {}

// ExitNotExpr is called when production notExpr is exited.
func (s *BaseAsaliLangGrammarListener) ExitNotExpr(ctx *NotExprContext) {}

// EnterUnaryMinusExpr is called when production unaryMinusExpr is entered.
func (s *BaseAsaliLangGrammarListener) EnterUnaryMinusExpr(ctx *UnaryMinusExprContext) {}

// ExitUnaryMinusExpr is called when production unaryMinusExpr is exited.
func (s *BaseAsaliLangGrammarListener) ExitUnaryMinusExpr(ctx *UnaryMinusExprContext) {}

// EnterMultiplicationExpr is called when production multiplicationExpr is entered.
func (s *BaseAsaliLangGrammarListener) EnterMultiplicationExpr(ctx *MultiplicationExprContext) {}

// ExitMultiplicationExpr is called when production multiplicationExpr is exited.
func (s *BaseAsaliLangGrammarListener) ExitMultiplicationExpr(ctx *MultiplicationExprContext) {}

// EnterAtomExpr is called when production atomExpr is entered.
func (s *BaseAsaliLangGrammarListener) EnterAtomExpr(ctx *AtomExprContext) {}

// ExitAtomExpr is called when production atomExpr is exited.
func (s *BaseAsaliLangGrammarListener) ExitAtomExpr(ctx *AtomExprContext) {}

// EnterOrExpr is called when production orExpr is entered.
func (s *BaseAsaliLangGrammarListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production orExpr is exited.
func (s *BaseAsaliLangGrammarListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterAdditiveExpr is called when production additiveExpr is entered.
func (s *BaseAsaliLangGrammarListener) EnterAdditiveExpr(ctx *AdditiveExprContext) {}

// ExitAdditiveExpr is called when production additiveExpr is exited.
func (s *BaseAsaliLangGrammarListener) ExitAdditiveExpr(ctx *AdditiveExprContext) {}

// EnterPowExpr is called when production powExpr is entered.
func (s *BaseAsaliLangGrammarListener) EnterPowExpr(ctx *PowExprContext) {}

// ExitPowExpr is called when production powExpr is exited.
func (s *BaseAsaliLangGrammarListener) ExitPowExpr(ctx *PowExprContext) {}

// EnterRelationalExpr is called when production relationalExpr is entered.
func (s *BaseAsaliLangGrammarListener) EnterRelationalExpr(ctx *RelationalExprContext) {}

// ExitRelationalExpr is called when production relationalExpr is exited.
func (s *BaseAsaliLangGrammarListener) ExitRelationalExpr(ctx *RelationalExprContext) {}

// EnterEqualityExpr is called when production equalityExpr is entered.
func (s *BaseAsaliLangGrammarListener) EnterEqualityExpr(ctx *EqualityExprContext) {}

// ExitEqualityExpr is called when production equalityExpr is exited.
func (s *BaseAsaliLangGrammarListener) ExitEqualityExpr(ctx *EqualityExprContext) {}

// EnterAndExpr is called when production andExpr is entered.
func (s *BaseAsaliLangGrammarListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production andExpr is exited.
func (s *BaseAsaliLangGrammarListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterParExpr is called when production parExpr is entered.
func (s *BaseAsaliLangGrammarListener) EnterParExpr(ctx *ParExprContext) {}

// ExitParExpr is called when production parExpr is exited.
func (s *BaseAsaliLangGrammarListener) ExitParExpr(ctx *ParExprContext) {}

// EnterNumberAtom is called when production numberAtom is entered.
func (s *BaseAsaliLangGrammarListener) EnterNumberAtom(ctx *NumberAtomContext) {}

// ExitNumberAtom is called when production numberAtom is exited.
func (s *BaseAsaliLangGrammarListener) ExitNumberAtom(ctx *NumberAtomContext) {}

// EnterBooleanAtom is called when production booleanAtom is entered.
func (s *BaseAsaliLangGrammarListener) EnterBooleanAtom(ctx *BooleanAtomContext) {}

// ExitBooleanAtom is called when production booleanAtom is exited.
func (s *BaseAsaliLangGrammarListener) ExitBooleanAtom(ctx *BooleanAtomContext) {}

// EnterIdAtom is called when production idAtom is entered.
func (s *BaseAsaliLangGrammarListener) EnterIdAtom(ctx *IdAtomContext) {}

// ExitIdAtom is called when production idAtom is exited.
func (s *BaseAsaliLangGrammarListener) ExitIdAtom(ctx *IdAtomContext) {}

// EnterStringAtom is called when production stringAtom is entered.
func (s *BaseAsaliLangGrammarListener) EnterStringAtom(ctx *StringAtomContext) {}

// ExitStringAtom is called when production stringAtom is exited.
func (s *BaseAsaliLangGrammarListener) ExitStringAtom(ctx *StringAtomContext) {}

// EnterNilAtom is called when production nilAtom is entered.
func (s *BaseAsaliLangGrammarListener) EnterNilAtom(ctx *NilAtomContext) {}

// ExitNilAtom is called when production nilAtom is exited.
func (s *BaseAsaliLangGrammarListener) ExitNilAtom(ctx *NilAtomContext) {}
