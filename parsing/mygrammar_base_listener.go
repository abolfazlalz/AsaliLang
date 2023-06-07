// Code generated from MyGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parsing // MyGrammar
import "github.com/antlr4-go/antlr/v4"

// BaseMyGrammarListener is a complete listener for a parse tree produced by MyGrammarParser.
type BaseMyGrammarListener struct{}

var _ MyGrammarListener = &BaseMyGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMyGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMyGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMyGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMyGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BaseMyGrammarListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BaseMyGrammarListener) ExitParse(ctx *ParseContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseMyGrammarListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseMyGrammarListener) ExitBlock(ctx *BlockContext) {}

// EnterStat is called when production stat is entered.
func (s *BaseMyGrammarListener) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *BaseMyGrammarListener) ExitStat(ctx *StatContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseMyGrammarListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseMyGrammarListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterIfStat is called when production ifStat is entered.
func (s *BaseMyGrammarListener) EnterIfStat(ctx *IfStatContext) {}

// ExitIfStat is called when production ifStat is exited.
func (s *BaseMyGrammarListener) ExitIfStat(ctx *IfStatContext) {}

// EnterCondition_block is called when production condition_block is entered.
func (s *BaseMyGrammarListener) EnterCondition_block(ctx *Condition_blockContext) {}

// ExitCondition_block is called when production condition_block is exited.
func (s *BaseMyGrammarListener) ExitCondition_block(ctx *Condition_blockContext) {}

// EnterStat_block is called when production stat_block is entered.
func (s *BaseMyGrammarListener) EnterStat_block(ctx *Stat_blockContext) {}

// ExitStat_block is called when production stat_block is exited.
func (s *BaseMyGrammarListener) ExitStat_block(ctx *Stat_blockContext) {}

// EnterWhileStat is called when production whileStat is entered.
func (s *BaseMyGrammarListener) EnterWhileStat(ctx *WhileStatContext) {}

// ExitWhileStat is called when production whileStat is exited.
func (s *BaseMyGrammarListener) ExitWhileStat(ctx *WhileStatContext) {}

// EnterForStat is called when production forStat is entered.
func (s *BaseMyGrammarListener) EnterForStat(ctx *ForStatContext) {}

// ExitForStat is called when production forStat is exited.
func (s *BaseMyGrammarListener) ExitForStat(ctx *ForStatContext) {}

// EnterLoopStat is called when production loopStat is entered.
func (s *BaseMyGrammarListener) EnterLoopStat(ctx *LoopStatContext) {}

// ExitLoopStat is called when production loopStat is exited.
func (s *BaseMyGrammarListener) ExitLoopStat(ctx *LoopStatContext) {}

// EnterLog is called when production log is entered.
func (s *BaseMyGrammarListener) EnterLog(ctx *LogContext) {}

// ExitLog is called when production log is exited.
func (s *BaseMyGrammarListener) ExitLog(ctx *LogContext) {}

// EnterNotExpr is called when production notExpr is entered.
func (s *BaseMyGrammarListener) EnterNotExpr(ctx *NotExprContext) {}

// ExitNotExpr is called when production notExpr is exited.
func (s *BaseMyGrammarListener) ExitNotExpr(ctx *NotExprContext) {}

// EnterUnaryMinusExpr is called when production unaryMinusExpr is entered.
func (s *BaseMyGrammarListener) EnterUnaryMinusExpr(ctx *UnaryMinusExprContext) {}

// ExitUnaryMinusExpr is called when production unaryMinusExpr is exited.
func (s *BaseMyGrammarListener) ExitUnaryMinusExpr(ctx *UnaryMinusExprContext) {}

// EnterMultiplicationExpr is called when production multiplicationExpr is entered.
func (s *BaseMyGrammarListener) EnterMultiplicationExpr(ctx *MultiplicationExprContext) {}

// ExitMultiplicationExpr is called when production multiplicationExpr is exited.
func (s *BaseMyGrammarListener) ExitMultiplicationExpr(ctx *MultiplicationExprContext) {}

// EnterAtomExpr is called when production atomExpr is entered.
func (s *BaseMyGrammarListener) EnterAtomExpr(ctx *AtomExprContext) {}

// ExitAtomExpr is called when production atomExpr is exited.
func (s *BaseMyGrammarListener) ExitAtomExpr(ctx *AtomExprContext) {}

// EnterOrExpr is called when production orExpr is entered.
func (s *BaseMyGrammarListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production orExpr is exited.
func (s *BaseMyGrammarListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterAdditiveExpr is called when production additiveExpr is entered.
func (s *BaseMyGrammarListener) EnterAdditiveExpr(ctx *AdditiveExprContext) {}

// ExitAdditiveExpr is called when production additiveExpr is exited.
func (s *BaseMyGrammarListener) ExitAdditiveExpr(ctx *AdditiveExprContext) {}

// EnterPowExpr is called when production powExpr is entered.
func (s *BaseMyGrammarListener) EnterPowExpr(ctx *PowExprContext) {}

// ExitPowExpr is called when production powExpr is exited.
func (s *BaseMyGrammarListener) ExitPowExpr(ctx *PowExprContext) {}

// EnterRelationalExpr is called when production relationalExpr is entered.
func (s *BaseMyGrammarListener) EnterRelationalExpr(ctx *RelationalExprContext) {}

// ExitRelationalExpr is called when production relationalExpr is exited.
func (s *BaseMyGrammarListener) ExitRelationalExpr(ctx *RelationalExprContext) {}

// EnterEqualityExpr is called when production equalityExpr is entered.
func (s *BaseMyGrammarListener) EnterEqualityExpr(ctx *EqualityExprContext) {}

// ExitEqualityExpr is called when production equalityExpr is exited.
func (s *BaseMyGrammarListener) ExitEqualityExpr(ctx *EqualityExprContext) {}

// EnterAndExpr is called when production andExpr is entered.
func (s *BaseMyGrammarListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production andExpr is exited.
func (s *BaseMyGrammarListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterParExpr is called when production parExpr is entered.
func (s *BaseMyGrammarListener) EnterParExpr(ctx *ParExprContext) {}

// ExitParExpr is called when production parExpr is exited.
func (s *BaseMyGrammarListener) ExitParExpr(ctx *ParExprContext) {}

// EnterNumberAtom is called when production numberAtom is entered.
func (s *BaseMyGrammarListener) EnterNumberAtom(ctx *NumberAtomContext) {}

// ExitNumberAtom is called when production numberAtom is exited.
func (s *BaseMyGrammarListener) ExitNumberAtom(ctx *NumberAtomContext) {}

// EnterBooleanAtom is called when production booleanAtom is entered.
func (s *BaseMyGrammarListener) EnterBooleanAtom(ctx *BooleanAtomContext) {}

// ExitBooleanAtom is called when production booleanAtom is exited.
func (s *BaseMyGrammarListener) ExitBooleanAtom(ctx *BooleanAtomContext) {}

// EnterIdAtom is called when production idAtom is entered.
func (s *BaseMyGrammarListener) EnterIdAtom(ctx *IdAtomContext) {}

// ExitIdAtom is called when production idAtom is exited.
func (s *BaseMyGrammarListener) ExitIdAtom(ctx *IdAtomContext) {}

// EnterStringAtom is called when production stringAtom is entered.
func (s *BaseMyGrammarListener) EnterStringAtom(ctx *StringAtomContext) {}

// ExitStringAtom is called when production stringAtom is exited.
func (s *BaseMyGrammarListener) ExitStringAtom(ctx *StringAtomContext) {}

// EnterNilAtom is called when production nilAtom is entered.
func (s *BaseMyGrammarListener) EnterNilAtom(ctx *NilAtomContext) {}

// ExitNilAtom is called when production nilAtom is exited.
func (s *BaseMyGrammarListener) ExitNilAtom(ctx *NilAtomContext) {}
