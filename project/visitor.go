package project

import (
	"asalicompiler/parsing"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"log"
	"strconv"
)

type Visitor struct {
	vars map[string]interface{}
}

func NewVisitor() *Visitor {
	return &Visitor{
		vars: map[string]interface{}{},
	}
}

func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parsing.ProgramContext:
		return v.VisitProgram(val)
	case *parsing.Statement_def_varContext:
		return v.VisitStatement_def_var(val)
	case *parsing.Number_defContext:
		return v.VisitNumber_def(val)
	case *parsing.Priority1_numberContext:
		return v.VisitPriority1_number(val)
	case *parsing.Statement_printContext:
		return v.VisitStatement_print(val)
	case *parsing.Priority1_multipleContext:
		return v.VisitPriority1_multiple(val)
	case *parsing.Priority1_divideContext:
		return v.VisitPriority1_divide(val)
	case *parsing.Priority2_defContext:
		return v.VisitPriority2_def(val)
	case *parsing.Priority2_plusContext:
		return v.VisitPriority2_plus(val)
	default:
		log.Println("Name: ", val.GetText())
		panic("Invalid context")
	}
}

func (v *Visitor) VisitPriority1_number(ctx *parsing.Priority1_numberContext) float64 {
	r, _ := strconv.ParseFloat(ctx.GetText(), 64)
	return r
}

func (v *Visitor) VisitNumber_def(ctx *parsing.Number_defContext) float64 {
	n, _ := strconv.ParseFloat(ctx.GetText(), 64)
	return n
}

func (v *Visitor) VisitProgram(ctx *parsing.ProgramContext) float64 {
	for _, statment := range ctx.Statements().AllStatement() {
		v.Visit(statment)
	}
	return 1
}

func (v *Visitor) VisitPriority1_multiple(ctx *parsing.Priority1_multipleContext) float64 {
	first, _ := v.Visit(ctx.Number()).(float64)
	second, _ := v.Visit(ctx.Priority1()).(float64)
	return first * second
}

func (v *Visitor) VisitPriority1_divide(ctx *parsing.Priority1_divideContext) float64 {
	first, _ := v.Visit(ctx.Priority1()).(float64)
	second, _ := v.Visit(ctx.Number()).(float64)
	return first / second
}

func (v *Visitor) VisitPriority2_plus(ctx *parsing.Priority2_plusContext) float64 {
	first, _ := v.Visit(ctx.Priority1()).(float64)
	second, _ := v.Visit(ctx.Priority2()).(float64)

	return first + second
}

func (v *Visitor) VisitStatement_def_var(ctx *parsing.Statement_def_varContext) interface{} {
	value := v.Visit(ctx.Priority2())
	v.vars[ctx.IDENTIFIER().GetText()] = value
	return value
}

func (v Visitor) VisitPriority2_def(ctx *parsing.Priority2_defContext) float64 {
	return v.Visit(ctx.Priority1()).(float64)
}

func (v *Visitor) VisitStatement_print(ctx *parsing.Statement_printContext) interface{} {
	val := v.vars[ctx.IDENTIFIER().GetText()]
	if val == nil {
		panic("Undefined variable")
	}
	fmt.Println(val)
	return 0
}
