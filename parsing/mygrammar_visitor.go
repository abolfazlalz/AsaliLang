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

	// Visit a parse tree produced by MyGrammarParser#statementDefineVariable.
	VisitStatementDefineVariable(ctx *StatementDefineVariableContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#statement_begin_end.
	VisitStatement_begin_end(ctx *Statement_begin_endContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#statement_if.
	VisitStatement_if(ctx *Statement_ifContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#statement_if_else.
	VisitStatement_if_else(ctx *Statement_if_elseContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#statement_while.
	VisitStatement_while(ctx *Statement_whileContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#statement_for.
	VisitStatement_for(ctx *Statement_forContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#statement_loop.
	VisitStatement_loop(ctx *Statement_loopContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#CallMethod.
	VisitCallMethod(ctx *CallMethodContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#StatementPrintMethod.
	VisitStatementPrintMethod(ctx *StatementPrintMethodContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#methodCall.
	VisitMethodCall(ctx *MethodCallContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#variableSetterTypes.
	VisitVariableSetterTypes(ctx *VariableSetterTypesContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#methodCallArguments.
	VisitMethodCallArguments(ctx *MethodCallArgumentsContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#condition_def.
	VisitCondition_def(ctx *Condition_defContext) interface{}

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
