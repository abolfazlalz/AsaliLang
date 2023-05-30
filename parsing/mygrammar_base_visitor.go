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

func (v *BaseMyGrammarVisitor) VisitStatement_def_var(ctx *Statement_def_varContext) interface{} {
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

func (v *BaseMyGrammarVisitor) VisitStatement_print(ctx *Statement_printContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitStatement_print_1(ctx *Statement_print_1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitCondition_def(ctx *Condition_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitPriority1_number(ctx *Priority1_numberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitPriority1_multiple(ctx *Priority1_multipleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitPriority1_divide(ctx *Priority1_divideContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitPriority2_plus(ctx *Priority2_plusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitPriority2_def(ctx *Priority2_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitPriority2_minus(ctx *Priority2_minusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitNumber_def(ctx *Number_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitNumber_minus(ctx *Number_minusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitNumber_identifier(ctx *Number_identifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyGrammarVisitor) VisitNumber_paran(ctx *Number_paranContext) interface{} {
	return v.VisitChildren(ctx)
}
