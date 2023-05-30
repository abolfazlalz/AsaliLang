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

	// EnterStatement_def_var is called when entering the statement_def_var production.
	EnterStatement_def_var(c *Statement_def_varContext)

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

	// EnterStatement_print is called when entering the statement_print production.
	EnterStatement_print(c *Statement_printContext)

	// EnterStatement_print_1 is called when entering the statement_print_1 production.
	EnterStatement_print_1(c *Statement_print_1Context)

	// EnterCondition_def is called when entering the condition_def production.
	EnterCondition_def(c *Condition_defContext)

	// EnterPriority1_number is called when entering the priority1_number production.
	EnterPriority1_number(c *Priority1_numberContext)

	// EnterPriority1_multiple is called when entering the priority1_multiple production.
	EnterPriority1_multiple(c *Priority1_multipleContext)

	// EnterPriority1_divide is called when entering the priority1_divide production.
	EnterPriority1_divide(c *Priority1_divideContext)

	// EnterPriority2_plus is called when entering the priority2_plus production.
	EnterPriority2_plus(c *Priority2_plusContext)

	// EnterPriority2_def is called when entering the priority2_def production.
	EnterPriority2_def(c *Priority2_defContext)

	// EnterPriority2_minus is called when entering the priority2_minus production.
	EnterPriority2_minus(c *Priority2_minusContext)

	// EnterNumber_def is called when entering the number_def production.
	EnterNumber_def(c *Number_defContext)

	// EnterNumber_minus is called when entering the number_minus production.
	EnterNumber_minus(c *Number_minusContext)

	// EnterNumber_identifier is called when entering the number_identifier production.
	EnterNumber_identifier(c *Number_identifierContext)

	// EnterNumber_paran is called when entering the number_paran production.
	EnterNumber_paran(c *Number_paranContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatements is called when exiting the statements production.
	ExitStatements(c *StatementsContext)

	// ExitStatement_def_var is called when exiting the statement_def_var production.
	ExitStatement_def_var(c *Statement_def_varContext)

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

	// ExitStatement_print is called when exiting the statement_print production.
	ExitStatement_print(c *Statement_printContext)

	// ExitStatement_print_1 is called when exiting the statement_print_1 production.
	ExitStatement_print_1(c *Statement_print_1Context)

	// ExitCondition_def is called when exiting the condition_def production.
	ExitCondition_def(c *Condition_defContext)

	// ExitPriority1_number is called when exiting the priority1_number production.
	ExitPriority1_number(c *Priority1_numberContext)

	// ExitPriority1_multiple is called when exiting the priority1_multiple production.
	ExitPriority1_multiple(c *Priority1_multipleContext)

	// ExitPriority1_divide is called when exiting the priority1_divide production.
	ExitPriority1_divide(c *Priority1_divideContext)

	// ExitPriority2_plus is called when exiting the priority2_plus production.
	ExitPriority2_plus(c *Priority2_plusContext)

	// ExitPriority2_def is called when exiting the priority2_def production.
	ExitPriority2_def(c *Priority2_defContext)

	// ExitPriority2_minus is called when exiting the priority2_minus production.
	ExitPriority2_minus(c *Priority2_minusContext)

	// ExitNumber_def is called when exiting the number_def production.
	ExitNumber_def(c *Number_defContext)

	// ExitNumber_minus is called when exiting the number_minus production.
	ExitNumber_minus(c *Number_minusContext)

	// ExitNumber_identifier is called when exiting the number_identifier production.
	ExitNumber_identifier(c *Number_identifierContext)

	// ExitNumber_paran is called when exiting the number_paran production.
	ExitNumber_paran(c *Number_paranContext)
}
