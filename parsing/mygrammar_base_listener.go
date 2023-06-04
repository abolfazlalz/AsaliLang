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

// EnterStatementDefineVariable is called when production statementDefineVariable is entered.
func (s *BaseMyGrammarListener) EnterStatementDefineVariable(ctx *StatementDefineVariableContext) {}

// ExitStatementDefineVariable is called when production statementDefineVariable is exited.
func (s *BaseMyGrammarListener) ExitStatementDefineVariable(ctx *StatementDefineVariableContext) {}

// EnterStatement_begin_end is called when production statement_begin_end is entered.
func (s *BaseMyGrammarListener) EnterStatement_begin_end(ctx *Statement_begin_endContext) {}

// ExitStatement_begin_end is called when production statement_begin_end is exited.
func (s *BaseMyGrammarListener) ExitStatement_begin_end(ctx *Statement_begin_endContext) {}

// EnterStatement_if is called when production statement_if is entered.
func (s *BaseMyGrammarListener) EnterStatement_if(ctx *Statement_ifContext) {}

// ExitStatement_if is called when production statement_if is exited.
func (s *BaseMyGrammarListener) ExitStatement_if(ctx *Statement_ifContext) {}

// EnterStatement_if_else is called when production statement_if_else is entered.
func (s *BaseMyGrammarListener) EnterStatement_if_else(ctx *Statement_if_elseContext) {}

// ExitStatement_if_else is called when production statement_if_else is exited.
func (s *BaseMyGrammarListener) ExitStatement_if_else(ctx *Statement_if_elseContext) {}

// EnterStatement_while is called when production statement_while is entered.
func (s *BaseMyGrammarListener) EnterStatement_while(ctx *Statement_whileContext) {}

// ExitStatement_while is called when production statement_while is exited.
func (s *BaseMyGrammarListener) ExitStatement_while(ctx *Statement_whileContext) {}

// EnterStatement_for is called when production statement_for is entered.
func (s *BaseMyGrammarListener) EnterStatement_for(ctx *Statement_forContext) {}

// ExitStatement_for is called when production statement_for is exited.
func (s *BaseMyGrammarListener) ExitStatement_for(ctx *Statement_forContext) {}

// EnterStatement_loop is called when production statement_loop is entered.
func (s *BaseMyGrammarListener) EnterStatement_loop(ctx *Statement_loopContext) {}

// ExitStatement_loop is called when production statement_loop is exited.
func (s *BaseMyGrammarListener) ExitStatement_loop(ctx *Statement_loopContext) {}

// EnterCallMethod is called when production CallMethod is entered.
func (s *BaseMyGrammarListener) EnterCallMethod(ctx *CallMethodContext) {}

// ExitCallMethod is called when production CallMethod is exited.
func (s *BaseMyGrammarListener) ExitCallMethod(ctx *CallMethodContext) {}

// EnterStatementPrintMethod is called when production StatementPrintMethod is entered.
func (s *BaseMyGrammarListener) EnterStatementPrintMethod(ctx *StatementPrintMethodContext) {}

// ExitStatementPrintMethod is called when production StatementPrintMethod is exited.
func (s *BaseMyGrammarListener) ExitStatementPrintMethod(ctx *StatementPrintMethodContext) {}

// EnterMethodCall is called when production methodCall is entered.
func (s *BaseMyGrammarListener) EnterMethodCall(ctx *MethodCallContext) {}

// ExitMethodCall is called when production methodCall is exited.
func (s *BaseMyGrammarListener) ExitMethodCall(ctx *MethodCallContext) {}

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

// EnterCondition_def is called when production condition_def is entered.
func (s *BaseMyGrammarListener) EnterCondition_def(ctx *Condition_defContext) {}

// ExitCondition_def is called when production condition_def is exited.
func (s *BaseMyGrammarListener) ExitCondition_def(ctx *Condition_defContext) {}

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
