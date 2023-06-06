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

	// EnterAssignmentStatementBlock is called when entering the assignmentStatementBlock production.
	EnterAssignmentStatementBlock(c *AssignmentStatementBlockContext)

	// EnterIfStatementBlock is called when entering the ifStatementBlock production.
	EnterIfStatementBlock(c *IfStatementBlockContext)

	// EnterWhileStatementBlock is called when entering the whileStatementBlock production.
	EnterWhileStatementBlock(c *WhileStatementBlockContext)

	// EnterMethodCallBlock is called when entering the methodCallBlock production.
	EnterMethodCallBlock(c *MethodCallBlockContext)

	// EnterPrintStatementBlock is called when entering the printStatementBlock production.
	EnterPrintStatementBlock(c *PrintStatementBlockContext)

	// EnterAssignmentStatement is called when entering the assignmentStatement production.
	EnterAssignmentStatement(c *AssignmentStatementContext)

	// EnterPrintStatement is called when entering the printStatement production.
	EnterPrintStatement(c *PrintStatementContext)

	// EnterBlockStatement is called when entering the blockStatement production.
	EnterBlockStatement(c *BlockStatementContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterMethodCall is called when entering the methodCall production.
	EnterMethodCall(c *MethodCallContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterVariableSetterTypes is called when entering the variableSetterTypes production.
	EnterVariableSetterTypes(c *VariableSetterTypesContext)

	// EnterMethodCallArguments is called when entering the methodCallArguments production.
	EnterMethodCallArguments(c *MethodCallArgumentsContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

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

	// ExitAssignmentStatementBlock is called when exiting the assignmentStatementBlock production.
	ExitAssignmentStatementBlock(c *AssignmentStatementBlockContext)

	// ExitIfStatementBlock is called when exiting the ifStatementBlock production.
	ExitIfStatementBlock(c *IfStatementBlockContext)

	// ExitWhileStatementBlock is called when exiting the whileStatementBlock production.
	ExitWhileStatementBlock(c *WhileStatementBlockContext)

	// ExitMethodCallBlock is called when exiting the methodCallBlock production.
	ExitMethodCallBlock(c *MethodCallBlockContext)

	// ExitPrintStatementBlock is called when exiting the printStatementBlock production.
	ExitPrintStatementBlock(c *PrintStatementBlockContext)

	// ExitAssignmentStatement is called when exiting the assignmentStatement production.
	ExitAssignmentStatement(c *AssignmentStatementContext)

	// ExitPrintStatement is called when exiting the printStatement production.
	ExitPrintStatement(c *PrintStatementContext)

	// ExitBlockStatement is called when exiting the blockStatement production.
	ExitBlockStatement(c *BlockStatementContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitMethodCall is called when exiting the methodCall production.
	ExitMethodCall(c *MethodCallContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitVariableSetterTypes is called when exiting the variableSetterTypes production.
	ExitVariableSetterTypes(c *VariableSetterTypesContext)

	// ExitMethodCallArguments is called when exiting the methodCallArguments production.
	ExitMethodCallArguments(c *MethodCallArgumentsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

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
