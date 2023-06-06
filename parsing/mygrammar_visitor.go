// Code generated from MyGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parsing // MyGrammar
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by MyGrammarParser.
type MyGrammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MyGrammarParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#statements.
	VisitStatements(ctx *StatementsContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#assignmentStatementBlock.
	VisitAssignmentStatementBlock(ctx *AssignmentStatementBlockContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#ifStatementBlock.
	VisitIfStatementBlock(ctx *IfStatementBlockContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#whileStatementBlock.
	VisitWhileStatementBlock(ctx *WhileStatementBlockContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#methodCallBlock.
	VisitMethodCallBlock(ctx *MethodCallBlockContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#printStatementBlock.
	VisitPrintStatementBlock(ctx *PrintStatementBlockContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#assignmentStatement.
	VisitAssignmentStatement(ctx *AssignmentStatementContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#printStatement.
	VisitPrintStatement(ctx *PrintStatementContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#blockStatement.
	VisitBlockStatement(ctx *BlockStatementContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#methodCall.
	VisitMethodCall(ctx *MethodCallContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#variableSetterTypes.
	VisitVariableSetterTypes(ctx *VariableSetterTypesContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#methodCallArguments.
	VisitMethodCallArguments(ctx *MethodCallArgumentsContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#powerExprDefault.
	VisitPowerExprDefault(ctx *PowerExprDefaultContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#powerExprPower.
	VisitPowerExprPower(ctx *PowerExprPowerContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#multipleExprDivide.
	VisitMultipleExprDivide(ctx *MultipleExprDivideContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#multipleExprMulti.
	VisitMultipleExprMulti(ctx *MultipleExprMultiContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#multipleExprDefault.
	VisitMultipleExprDefault(ctx *MultipleExprDefaultContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#sumExprPlus.
	VisitSumExprPlus(ctx *SumExprPlusContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#sumExprMinus.
	VisitSumExprMinus(ctx *SumExprMinusContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#sumExprDefault.
	VisitSumExprDefault(ctx *SumExprDefaultContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#numberDefault.
	VisitNumberDefault(ctx *NumberDefaultContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#NumberMinus.
	VisitNumberMinus(ctx *NumberMinusContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#NumberIdentifier.
	VisitNumberIdentifier(ctx *NumberIdentifierContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#NumberParentheses.
	VisitNumberParentheses(ctx *NumberParenthesesContext) interface{}
}
