// Code generated from MyGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parsing // MyGrammar
import "github.com/antlr4-go/antlr/v4"

type BaseMyGrammarVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMyGrammarVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitStatements(ctx *StatementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitAssignmentStatementBlock(ctx *AssignmentStatementBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitIfStatementBlock(ctx *IfStatementBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitWhileStatementBlock(ctx *WhileStatementBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitMethodCallBlock(ctx *MethodCallBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitPrintStatementBlock(ctx *PrintStatementBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitAssignmentStatement(ctx *AssignmentStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitPrintStatement(ctx *PrintStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitBlockStatement(ctx *BlockStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitMethodCall(ctx *MethodCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitVariableSetterTypes(ctx *VariableSetterTypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitMethodCallArguments(ctx *MethodCallArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitPowerExprDefault(ctx *PowerExprDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitPowerExprPower(ctx *PowerExprPowerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitMultipleExprDivide(ctx *MultipleExprDivideContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitMultipleExprMulti(ctx *MultipleExprMultiContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitMultipleExprDefault(ctx *MultipleExprDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitSumExprPlus(ctx *SumExprPlusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitSumExprMinus(ctx *SumExprMinusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitSumExprDefault(ctx *SumExprDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitNumberDefault(ctx *NumberDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitNumberMinus(ctx *NumberMinusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitNumberIdentifier(ctx *NumberIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitNumberParentheses(ctx *NumberParenthesesContext) interface{} {
	return v.VisitChildren(ctx)
}
