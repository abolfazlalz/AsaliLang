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

// EnterStatement_def_var is called when production statement_def_var is entered.
func (s *BaseMyGrammarListener) EnterStatement_def_var(ctx *Statement_def_varContext) {}

// ExitStatement_def_var is called when production statement_def_var is exited.
func (s *BaseMyGrammarListener) ExitStatement_def_var(ctx *Statement_def_varContext) {}

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

// EnterStatement_print is called when production statement_print is entered.
func (s *BaseMyGrammarListener) EnterStatement_print(ctx *Statement_printContext) {}

// ExitStatement_print is called when production statement_print is exited.
func (s *BaseMyGrammarListener) ExitStatement_print(ctx *Statement_printContext) {}

// EnterStatement_print_1 is called when production statement_print_1 is entered.
func (s *BaseMyGrammarListener) EnterStatement_print_1(ctx *Statement_print_1Context) {}

// ExitStatement_print_1 is called when production statement_print_1 is exited.
func (s *BaseMyGrammarListener) ExitStatement_print_1(ctx *Statement_print_1Context) {}

// EnterCondition_def is called when production condition_def is entered.
func (s *BaseMyGrammarListener) EnterCondition_def(ctx *Condition_defContext) {}

// ExitCondition_def is called when production condition_def is exited.
func (s *BaseMyGrammarListener) ExitCondition_def(ctx *Condition_defContext) {}

// EnterPriority1_number is called when production priority1_number is entered.
func (s *BaseMyGrammarListener) EnterPriority1_number(ctx *Priority1_numberContext) {}

// ExitPriority1_number is called when production priority1_number is exited.
func (s *BaseMyGrammarListener) ExitPriority1_number(ctx *Priority1_numberContext) {}

// EnterPriority1_multiple is called when production priority1_multiple is entered.
func (s *BaseMyGrammarListener) EnterPriority1_multiple(ctx *Priority1_multipleContext) {}

// ExitPriority1_multiple is called when production priority1_multiple is exited.
func (s *BaseMyGrammarListener) ExitPriority1_multiple(ctx *Priority1_multipleContext) {}

// EnterPriority1_divide is called when production priority1_divide is entered.
func (s *BaseMyGrammarListener) EnterPriority1_divide(ctx *Priority1_divideContext) {}

// ExitPriority1_divide is called when production priority1_divide is exited.
func (s *BaseMyGrammarListener) ExitPriority1_divide(ctx *Priority1_divideContext) {}

// EnterPriority2_plus is called when production priority2_plus is entered.
func (s *BaseMyGrammarListener) EnterPriority2_plus(ctx *Priority2_plusContext) {}

// ExitPriority2_plus is called when production priority2_plus is exited.
func (s *BaseMyGrammarListener) ExitPriority2_plus(ctx *Priority2_plusContext) {}

// EnterPriority2_def is called when production priority2_def is entered.
func (s *BaseMyGrammarListener) EnterPriority2_def(ctx *Priority2_defContext) {}

// ExitPriority2_def is called when production priority2_def is exited.
func (s *BaseMyGrammarListener) ExitPriority2_def(ctx *Priority2_defContext) {}

// EnterPriority2_minus is called when production priority2_minus is entered.
func (s *BaseMyGrammarListener) EnterPriority2_minus(ctx *Priority2_minusContext) {}

// ExitPriority2_minus is called when production priority2_minus is exited.
func (s *BaseMyGrammarListener) ExitPriority2_minus(ctx *Priority2_minusContext) {}

// EnterNumber_def is called when production number_def is entered.
func (s *BaseMyGrammarListener) EnterNumber_def(ctx *Number_defContext) {}

// ExitNumber_def is called when production number_def is exited.
func (s *BaseMyGrammarListener) ExitNumber_def(ctx *Number_defContext) {}

// EnterNumber_minus is called when production number_minus is entered.
func (s *BaseMyGrammarListener) EnterNumber_minus(ctx *Number_minusContext) {}

// ExitNumber_minus is called when production number_minus is exited.
func (s *BaseMyGrammarListener) ExitNumber_minus(ctx *Number_minusContext) {}

// EnterNumber_identifier is called when production number_identifier is entered.
func (s *BaseMyGrammarListener) EnterNumber_identifier(ctx *Number_identifierContext) {}

// ExitNumber_identifier is called when production number_identifier is exited.
func (s *BaseMyGrammarListener) ExitNumber_identifier(ctx *Number_identifierContext) {}

// EnterNumber_paran is called when production number_paran is entered.
func (s *BaseMyGrammarListener) EnterNumber_paran(ctx *Number_paranContext) {}

// ExitNumber_paran is called when production number_paran is exited.
func (s *BaseMyGrammarListener) ExitNumber_paran(ctx *Number_paranContext) {}
