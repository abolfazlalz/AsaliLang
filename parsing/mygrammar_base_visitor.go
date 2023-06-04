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

func (v *BaseMyGrammarVisitor) VisitStatementDefineVariable(ctx *StatementDefineVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitStatement_begin_end(ctx *Statement_begin_endContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitStatement_if(ctx *Statement_ifContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitStatement_if_else(ctx *Statement_if_elseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitStatement_while(ctx *Statement_whileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitStatement_for(ctx *Statement_forContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitStatement_loop(ctx *Statement_loopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitCallMethod(ctx *CallMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitStatementPrintMethod(ctx *StatementPrintMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitMethodCall(ctx *MethodCallContext) interface{} {
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

func (v *BaseMyGrammarVisitor) VisitCondition_def(ctx *Condition_defContext) interface{} {
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
