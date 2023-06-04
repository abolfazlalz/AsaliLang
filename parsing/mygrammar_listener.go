// Code generated from MyGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parsing // MyGrammar
import "github.com/antlr4-go/antlr/v4"

// MyGrammarListener is a complete listener for a parse tree produced by MyGrammarParser.
type MyGrammarListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatements is called when entering the statements production.
	EnterStatements(c *StatementsContext)

	// EnterEmptyLine is called when entering the EmptyLine production.
	EnterEmptyLine(c *EmptyLineContext)

	// EnterStatementDefineVariable is called when entering the statementDefineVariable production.
	EnterStatementDefineVariable(c *StatementDefineVariableContext)

	// EnterStatement_begin_end is called when entering the statement_begin_end production.
	EnterStatement_begin_end(c *Statement_begin_endContext)

	// EnterStatement_if is called when entering the statement_if production.
	EnterStatement_if(c *Statement_ifContext)

	// EnterStatement_if_else is called when entering the statement_if_else production.
	EnterStatement_if_else(c *Statement_if_elseContext)

	// EnterStatement_while is called when entering the statement_while production.
	EnterStatement_while(c *Statement_whileContext)

	// EnterStatement_for is called when entering the statement_for production.
	EnterStatement_for(c *Statement_forContext)

	// EnterStatement_loop is called when entering the statement_loop production.
	EnterStatement_loop(c *Statement_loopContext)

	// EnterCallMethod is called when entering the CallMethod production.
	EnterCallMethod(c *CallMethodContext)

	// EnterStatementPrintMethod is called when entering the StatementPrintMethod production.
	EnterStatementPrintMethod(c *StatementPrintMethodContext)

	// EnterMethodCall is called when entering the methodCall production.
	EnterMethodCall(c *MethodCallContext)

	// EnterVariableSetterTypes is called when entering the variableSetterTypes production.
	EnterVariableSetterTypes(c *VariableSetterTypesContext)

	// EnterMethodCallArguments is called when entering the methodCallArguments production.
	EnterMethodCallArguments(c *MethodCallArgumentsContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCondition_def is called when entering the condition_def production.
	EnterCondition_def(c *Condition_defContext)

	// EnterPowerExprDefault is called when entering the powerExprDefault production.
	EnterPowerExprDefault(c *PowerExprDefaultContext)

	// EnterPowerExprPower is called when entering the powerExprPower production.
	EnterPowerExprPower(c *PowerExprPowerContext)

	// EnterMultipleExprDivide is called when entering the multipleExprDivide production.
	EnterMultipleExprDivide(c *MultipleExprDivideContext)

	// EnterMultipleExprMulti is called when entering the multipleExprMulti production.
	EnterMultipleExprMulti(c *MultipleExprMultiContext)

	// EnterMultipleExprDefault is called when entering the multipleExprDefault production.
	EnterMultipleExprDefault(c *MultipleExprDefaultContext)

	// EnterSumExprPlus is called when entering the sumExprPlus production.
	EnterSumExprPlus(c *SumExprPlusContext)

	// EnterSumExprMinus is called when entering the sumExprMinus production.
	EnterSumExprMinus(c *SumExprMinusContext)

	// EnterSumExprDefault is called when entering the sumExprDefault production.
	EnterSumExprDefault(c *SumExprDefaultContext)

	// EnterNumberDefault is called when entering the numberDefault production.
	EnterNumberDefault(c *NumberDefaultContext)

	// EnterNumberMinus is called when entering the NumberMinus production.
	EnterNumberMinus(c *NumberMinusContext)

	// EnterNumberIdentifier is called when entering the NumberIdentifier production.
	EnterNumberIdentifier(c *NumberIdentifierContext)

	// EnterNumberParentheses is called when entering the NumberParentheses production.
	EnterNumberParentheses(c *NumberParenthesesContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatements is called when exiting the statements production.
	ExitStatements(c *StatementsContext)

	// ExitEmptyLine is called when exiting the EmptyLine production.
	ExitEmptyLine(c *EmptyLineContext)

	// ExitStatementDefineVariable is called when exiting the statementDefineVariable production.
	ExitStatementDefineVariable(c *StatementDefineVariableContext)

	// ExitStatement_begin_end is called when exiting the statement_begin_end production.
	ExitStatement_begin_end(c *Statement_begin_endContext)

	// ExitStatement_if is called when exiting the statement_if production.
	ExitStatement_if(c *Statement_ifContext)

	// ExitStatement_if_else is called when exiting the statement_if_else production.
	ExitStatement_if_else(c *Statement_if_elseContext)

	// ExitStatement_while is called when exiting the statement_while production.
	ExitStatement_while(c *Statement_whileContext)

	// ExitStatement_for is called when exiting the statement_for production.
	ExitStatement_for(c *Statement_forContext)

	// ExitStatement_loop is called when exiting the statement_loop production.
	ExitStatement_loop(c *Statement_loopContext)

	// ExitCallMethod is called when exiting the CallMethod production.
	ExitCallMethod(c *CallMethodContext)

	// ExitStatementPrintMethod is called when exiting the StatementPrintMethod production.
	ExitStatementPrintMethod(c *StatementPrintMethodContext)

	// ExitMethodCall is called when exiting the methodCall production.
	ExitMethodCall(c *MethodCallContext)

	// ExitVariableSetterTypes is called when exiting the variableSetterTypes production.
	ExitVariableSetterTypes(c *VariableSetterTypesContext)

	// ExitMethodCallArguments is called when exiting the methodCallArguments production.
	ExitMethodCallArguments(c *MethodCallArgumentsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCondition_def is called when exiting the condition_def production.
	ExitCondition_def(c *Condition_defContext)

	// ExitPowerExprDefault is called when exiting the powerExprDefault production.
	ExitPowerExprDefault(c *PowerExprDefaultContext)

	// ExitPowerExprPower is called when exiting the powerExprPower production.
	ExitPowerExprPower(c *PowerExprPowerContext)

	// ExitMultipleExprDivide is called when exiting the multipleExprDivide production.
	ExitMultipleExprDivide(c *MultipleExprDivideContext)

	// ExitMultipleExprMulti is called when exiting the multipleExprMulti production.
	ExitMultipleExprMulti(c *MultipleExprMultiContext)

	// ExitMultipleExprDefault is called when exiting the multipleExprDefault production.
	ExitMultipleExprDefault(c *MultipleExprDefaultContext)

	// ExitSumExprPlus is called when exiting the sumExprPlus production.
	ExitSumExprPlus(c *SumExprPlusContext)

	// ExitSumExprMinus is called when exiting the sumExprMinus production.
	ExitSumExprMinus(c *SumExprMinusContext)

	// ExitSumExprDefault is called when exiting the sumExprDefault production.
	ExitSumExprDefault(c *SumExprDefaultContext)

	// ExitNumberDefault is called when exiting the numberDefault production.
	ExitNumberDefault(c *NumberDefaultContext)

	// ExitNumberMinus is called when exiting the NumberMinus production.
	ExitNumberMinus(c *NumberMinusContext)

	// ExitNumberIdentifier is called when exiting the NumberIdentifier production.
	ExitNumberIdentifier(c *NumberIdentifierContext)

	// ExitNumberParentheses is called when exiting the NumberParentheses production.
	ExitNumberParentheses(c *NumberParenthesesContext)
}
