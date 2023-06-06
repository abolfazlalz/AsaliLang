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

// EnterProgram is called when production program is entered.
func (s *BaseMyGrammarListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseMyGrammarListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatements is called when production statements is entered.
func (s *BaseMyGrammarListener) EnterStatements(ctx *StatementsContext) {}

// ExitStatements is called when production statements is exited.
func (s *BaseMyGrammarListener) ExitStatements(ctx *StatementsContext) {}

// EnterAssignmentStatementBlock is called when production assignmentStatementBlock is entered.
func (s *BaseMyGrammarListener) EnterAssignmentStatementBlock(ctx *AssignmentStatementBlockContext) {}

// ExitAssignmentStatementBlock is called when production assignmentStatementBlock is exited.
func (s *BaseMyGrammarListener) ExitAssignmentStatementBlock(ctx *AssignmentStatementBlockContext) {}

// EnterIfStatementBlock is called when production ifStatementBlock is entered.
func (s *BaseMyGrammarListener) EnterIfStatementBlock(ctx *IfStatementBlockContext) {}

// ExitIfStatementBlock is called when production ifStatementBlock is exited.
func (s *BaseMyGrammarListener) ExitIfStatementBlock(ctx *IfStatementBlockContext) {}

// EnterWhileStatementBlock is called when production whileStatementBlock is entered.
func (s *BaseMyGrammarListener) EnterWhileStatementBlock(ctx *WhileStatementBlockContext) {}

// ExitWhileStatementBlock is called when production whileStatementBlock is exited.
func (s *BaseMyGrammarListener) ExitWhileStatementBlock(ctx *WhileStatementBlockContext) {}

// EnterMethodCallBlock is called when production methodCallBlock is entered.
func (s *BaseMyGrammarListener) EnterMethodCallBlock(ctx *MethodCallBlockContext) {}

// ExitMethodCallBlock is called when production methodCallBlock is exited.
func (s *BaseMyGrammarListener) ExitMethodCallBlock(ctx *MethodCallBlockContext) {}

// EnterPrintStatementBlock is called when production printStatementBlock is entered.
func (s *BaseMyGrammarListener) EnterPrintStatementBlock(ctx *PrintStatementBlockContext) {}

// ExitPrintStatementBlock is called when production printStatementBlock is exited.
func (s *BaseMyGrammarListener) ExitPrintStatementBlock(ctx *PrintStatementBlockContext) {}

// EnterAssignmentStatement is called when production assignmentStatement is entered.
func (s *BaseMyGrammarListener) EnterAssignmentStatement(ctx *AssignmentStatementContext) {}

// ExitAssignmentStatement is called when production assignmentStatement is exited.
func (s *BaseMyGrammarListener) ExitAssignmentStatement(ctx *AssignmentStatementContext) {}

// EnterPrintStatement is called when production printStatement is entered.
func (s *BaseMyGrammarListener) EnterPrintStatement(ctx *PrintStatementContext) {}

// ExitPrintStatement is called when production printStatement is exited.
func (s *BaseMyGrammarListener) ExitPrintStatement(ctx *PrintStatementContext) {}

// EnterBlockStatement is called when production blockStatement is entered.
func (s *BaseMyGrammarListener) EnterBlockStatement(ctx *BlockStatementContext) {}

// ExitBlockStatement is called when production blockStatement is exited.
func (s *BaseMyGrammarListener) ExitBlockStatement(ctx *BlockStatementContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseMyGrammarListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseMyGrammarListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterMethodCall is called when production methodCall is entered.
func (s *BaseMyGrammarListener) EnterMethodCall(ctx *MethodCallContext) {}

// ExitMethodCall is called when production methodCall is exited.
func (s *BaseMyGrammarListener) ExitMethodCall(ctx *MethodCallContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseMyGrammarListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseMyGrammarListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterVariableSetterTypes is called when production variableSetterTypes is entered.
func (s *BaseMyGrammarListener) EnterVariableSetterTypes(ctx *VariableSetterTypesContext) {}

// ExitVariableSetterTypes is called when production variableSetterTypes is exited.
func (s *BaseMyGrammarListener) ExitVariableSetterTypes(ctx *VariableSetterTypesContext) {}

// EnterMethodCallArguments is called when production methodCallArguments is entered.
func (s *BaseMyGrammarListener) EnterMethodCallArguments(ctx *MethodCallArgumentsContext) {}

// ExitMethodCallArguments is called when production methodCallArguments is exited.
func (s *BaseMyGrammarListener) ExitMethodCallArguments(ctx *MethodCallArgumentsContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMyGrammarListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMyGrammarListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPowerExprDefault is called when production powerExprDefault is entered.
func (s *BaseMyGrammarListener) EnterPowerExprDefault(ctx *PowerExprDefaultContext) {}

// ExitPowerExprDefault is called when production powerExprDefault is exited.
func (s *BaseMyGrammarListener) ExitPowerExprDefault(ctx *PowerExprDefaultContext) {}

// EnterPowerExprPower is called when production powerExprPower is entered.
func (s *BaseMyGrammarListener) EnterPowerExprPower(ctx *PowerExprPowerContext) {}

// ExitPowerExprPower is called when production powerExprPower is exited.
func (s *BaseMyGrammarListener) ExitPowerExprPower(ctx *PowerExprPowerContext) {}

// EnterMultipleExprDivide is called when production multipleExprDivide is entered.
func (s *BaseMyGrammarListener) EnterMultipleExprDivide(ctx *MultipleExprDivideContext) {}

// ExitMultipleExprDivide is called when production multipleExprDivide is exited.
func (s *BaseMyGrammarListener) ExitMultipleExprDivide(ctx *MultipleExprDivideContext) {}

// EnterMultipleExprMulti is called when production multipleExprMulti is entered.
func (s *BaseMyGrammarListener) EnterMultipleExprMulti(ctx *MultipleExprMultiContext) {}

// ExitMultipleExprMulti is called when production multipleExprMulti is exited.
func (s *BaseMyGrammarListener) ExitMultipleExprMulti(ctx *MultipleExprMultiContext) {}

// EnterMultipleExprDefault is called when production multipleExprDefault is entered.
func (s *BaseMyGrammarListener) EnterMultipleExprDefault(ctx *MultipleExprDefaultContext) {}

// ExitMultipleExprDefault is called when production multipleExprDefault is exited.
func (s *BaseMyGrammarListener) ExitMultipleExprDefault(ctx *MultipleExprDefaultContext) {}

// EnterSumExprPlus is called when production sumExprPlus is entered.
func (s *BaseMyGrammarListener) EnterSumExprPlus(ctx *SumExprPlusContext) {}

// ExitSumExprPlus is called when production sumExprPlus is exited.
func (s *BaseMyGrammarListener) ExitSumExprPlus(ctx *SumExprPlusContext) {}

// EnterSumExprMinus is called when production sumExprMinus is entered.
func (s *BaseMyGrammarListener) EnterSumExprMinus(ctx *SumExprMinusContext) {}

// ExitSumExprMinus is called when production sumExprMinus is exited.
func (s *BaseMyGrammarListener) ExitSumExprMinus(ctx *SumExprMinusContext) {}

// EnterSumExprDefault is called when production sumExprDefault is entered.
func (s *BaseMyGrammarListener) EnterSumExprDefault(ctx *SumExprDefaultContext) {}

// ExitSumExprDefault is called when production sumExprDefault is exited.
func (s *BaseMyGrammarListener) ExitSumExprDefault(ctx *SumExprDefaultContext) {}

// EnterNumberDefault is called when production numberDefault is entered.
func (s *BaseMyGrammarListener) EnterNumberDefault(ctx *NumberDefaultContext) {}

// ExitNumberDefault is called when production numberDefault is exited.
func (s *BaseMyGrammarListener) ExitNumberDefault(ctx *NumberDefaultContext) {}

// EnterNumberMinus is called when production NumberMinus is entered.
func (s *BaseMyGrammarListener) EnterNumberMinus(ctx *NumberMinusContext) {}

// ExitNumberMinus is called when production NumberMinus is exited.
func (s *BaseMyGrammarListener) ExitNumberMinus(ctx *NumberMinusContext) {}

// EnterNumberIdentifier is called when production NumberIdentifier is entered.
func (s *BaseMyGrammarListener) EnterNumberIdentifier(ctx *NumberIdentifierContext) {}

// ExitNumberIdentifier is called when production NumberIdentifier is exited.
func (s *BaseMyGrammarListener) ExitNumberIdentifier(ctx *NumberIdentifierContext) {}

// EnterNumberParentheses is called when production NumberParentheses is entered.
func (s *BaseMyGrammarListener) EnterNumberParentheses(ctx *NumberParenthesesContext) {}

// ExitNumberParentheses is called when production NumberParentheses is exited.
func (s *BaseMyGrammarListener) ExitNumberParentheses(ctx *NumberParenthesesContext) {}
