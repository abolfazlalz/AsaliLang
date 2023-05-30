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

	// Visit a parse tree produced by MyGrammarParser#statement_def_var.
	VisitStatement_def_var(ctx *Statement_def_varContext) interface{}

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

	// Visit a parse tree produced by MyGrammarParser#statement_print.
	VisitStatement_print(ctx *Statement_printContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#statement_print_1.
	VisitStatement_print_1(ctx *Statement_print_1Context) interface{}

	// Visit a parse tree produced by MyGrammarParser#condition_def.
	VisitCondition_def(ctx *Condition_defContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#priority1_number.
	VisitPriority1_number(ctx *Priority1_numberContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#priority1_multiple.
	VisitPriority1_multiple(ctx *Priority1_multipleContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#priority1_divide.
	VisitPriority1_divide(ctx *Priority1_divideContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#priority2_plus.
	VisitPriority2_plus(ctx *Priority2_plusContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#priority2_def.
	VisitPriority2_def(ctx *Priority2_defContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#priority2_minus.
	VisitPriority2_minus(ctx *Priority2_minusContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#number_def.
	VisitNumber_def(ctx *Number_defContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#number_minus.
	VisitNumber_minus(ctx *Number_minusContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#number_identifier.
	VisitNumber_identifier(ctx *Number_identifierContext) interface{}

	// Visit a parse tree produced by MyGrammarParser#number_paran.
	VisitNumber_paran(ctx *Number_paranContext) interface{}
}
