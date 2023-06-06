package project

import (
	"asalicompiler/parsing"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"log"
	"math"
	"reflect"
	"strconv"
	"strings"
)

type Visitor struct {
	parsing.BaseMyGrammarVisitor
	method *BuildInMethod
	vars   map[string]interface{}
}

func NewVisitor() *Visitor {
	return &Visitor{
		vars: map[string]interface{}{
			"PI": math.Pi,
		},
		method: NewBuildInMethod(),
	}
}

func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parsing.ProgramContext:
		return v.VisitProgram(val)
	case *parsing.AssignmentStatementContext:
		return v.VisitAssignmentStatement(val)
	case *parsing.NumberDefaultContext:
		return v.VisitNumberDefault(val)
	case *parsing.NumberIdentifierContext:
		return v.VisitNumberIdentifier(val)
	case *parsing.PowerExprPowerContext:
		return v.VisitPowerExprPower(val)
	case *parsing.PowerExprDefaultContext:
		return v.VisitPowerExprDefault(val)
	case *parsing.NumberMinusContext:
		return v.VisitNumberMinus(val)
	case *parsing.MultipleExprDefaultContext:
		return v.VisitMultipleExprDefault(val)
	case *parsing.MultipleExprMultiContext:
		return v.VisitMultipleExprMulti(val)
	case *parsing.MultipleExprDivideContext:
		return v.VisitMultipleExprDivide(val)
	case *parsing.SumExprDefaultContext:
		return v.VisitSumExprDefault(val)
	case *parsing.SumExprPlusContext:
		return v.VisitSumExprPlus(val)
	case *parsing.SumExprMinusContext:
		return v.VisitSumExprMinus(val)
	case *parsing.NumberParenthesesContext:
		return v.VisitNumberParentheses(val)
	case *parsing.PrintStatementContext:
		return v.VisitStatementPrintMethod(val)
	case *parsing.MethodCallArgumentsContext:
		return v.VisitMethodCallArguments(val)
	case *parsing.MethodCallContext:
		return v.VisitMethodCall(val)
	case *parsing.IfStatementContext:
		fmt.Println("e")
		return v.VisitStatementIf(val)
	case *parsing.BlockStatementContext:
		return v.VisitBlockStatement(val)
	case *parsing.ExpressionContext:
		return v.VisitExpression(val)
	default:
		panic(fmt.Sprintf("Invalid context\nReflect Check: %s\nName: %s", reflect.TypeOf(val), val.GetText()))
		//return v.BaseMyGrammarVisitor.Visit(tree)
	}
}

func (v *Visitor) VisitProgram(ctx *parsing.ProgramContext) float64 {
	for _, statement := range ctx.Statements().AllStatement() {
		for _, child := range statement.GetChildren() {
			v.Visit(child.(antlr.ParseTree))
		}
	}
	fmt.Println(v.vars)
	return 1
}

func (v *Visitor) VisitBlockStatement(ctx *parsing.BlockStatementContext) interface{} {
	for _, statement := range ctx.Statements().AllStatement() {
		v.Visit(statement)
	}
	return nil
}

func (v *Visitor) VisitPowerExprDefault(ctx *parsing.PowerExprDefaultContext) float64 {
	return v.Visit(ctx.Number()).(float64)
}

func (v *Visitor) VisitPowerExprPower(ctx *parsing.PowerExprPowerContext) float64 {
	first, _ := v.Visit(ctx.Number()).(float64)
	second, _ := v.Visit(ctx.PowerExpr()).(float64)
	return math.Pow(first, second)
}

func (v *Visitor) VisitMultipleExprDefault(ctx *parsing.MultipleExprDefaultContext) float64 {
	val, ok := v.Visit(ctx.PowerExpr()).(float64)
	if !ok {
		panic("Invalid numeric value")
	}
	return val
}

func (v *Visitor) VisitNumberDefault(ctx *parsing.NumberDefaultContext) float64 {
	n, _ := strconv.ParseFloat(ctx.GetText(), 64)
	return n
}

func (v *Visitor) VisitNumberIdentifier(ctx *parsing.NumberIdentifierContext) float64 {
	variableName := ctx.IDENTIFIER().GetText()
	result := v.getVariableValue(variableName)
	value, ok := result.(float64)
	if !ok {
		panic(fmt.Sprintf("%s does not valid numeric variable", variableName))
	}
	return value
}

func (v *Visitor) VisitNumberMinus(ctx *parsing.NumberMinusContext) float64 {
	n, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		panic(err)
	}
	return -1 * n
}

func (v *Visitor) VisitMultipleExprMulti(ctx *parsing.MultipleExprMultiContext) float64 {
	first, _ := v.Visit(ctx.PowerExpr()).(float64)
	second, _ := v.Visit(ctx.MultipleExpr()).(float64)
	return first * second
}

func (v *Visitor) VisitMultipleExprDivide(ctx *parsing.MultipleExprDivideContext) float64 {
	first, _ := v.Visit(ctx.PowerExpr()).(float64)
	second, _ := v.Visit(ctx.MultipleExpr()).(float64)
	return second / first
}

func (v *Visitor) VisitSumExprDefault(ctx *parsing.SumExprDefaultContext) float64 {
	return v.Visit(ctx.MultipleExpr()).(float64)
}

func (v *Visitor) VisitSumExprPlus(ctx *parsing.SumExprPlusContext) float64 {
	first, _ := v.Visit(ctx.MultipleExpr()).(float64)
	second, _ := v.Visit(ctx.SumExpr()).(float64)

	return first + second
}

func (v *Visitor) VisitSumExprMinus(ctx *parsing.SumExprMinusContext) float64 {
	first, _ := v.Visit(ctx.MultipleExpr()).(float64)
	second, _ := v.Visit(ctx.SumExpr()).(float64)

	return second - first
}

func (v *Visitor) VisitAssignmentStatement(ctx *parsing.AssignmentStatementContext) interface{} {
	var value interface{}
	setter := ctx.VariableSetterTypes()
	if setter.SumExpr() != nil {
		value = v.Visit(setter.SumExpr())
	} else if setter.IDENTIFIER() != nil {
		value = v.Visit(setter.IDENTIFIER())
	} else if setter.MethodCall() != nil {
		value = v.Visit(setter.MethodCall())
	}

	v.vars[ctx.IDENTIFIER().GetText()] = value
	return value
}

func (v *Visitor) VisitNumberParentheses(ctx *parsing.NumberParenthesesContext) float64 {
	return v.Visit(ctx.SumExpr()).(float64)
}

func (v *Visitor) VisitMethodCall(ctx *parsing.MethodCallContext) interface{} {
	methodName := ctx.IDENTIFIER().GetText()
	args := v.Visit(ctx.MethodCallArguments()).([]interface{})

	if methodName == "sin" {
		if len(args) < 1 {
			v.notEnoughArgs(methodName, args, 1)
		}
		return math.Sin(args[0].(float64))
	} else if methodName == "cos" {
		if len(args) < 1 {
			v.notEnoughArgs(methodName, args, 1)
		}
		return math.Cos(args[0].(float64))
	} else if methodName == "log" {
		argsStr := make([]string, len(args))
		for i, arg := range args {
			argsStr[i] = fmt.Sprintf("%v", arg)
		}
		log.Println(strings.Join(argsStr, " "))
	}
	return nil
}

func (v *Visitor) VisitStatementPrintMethod(ctx *parsing.PrintStatementContext) interface{} {
	//args, ok := v.Visit(ctx.MethodCallArguments()).([]interface{})
	//if !ok {
	//	panic("Invalid arguments types")
	//}
	//argsStr := make([]string, len(args))
	//for i, arg := range args {
	//	argsStr[i] = fmt.Sprintf("%v", arg)
	//}
	//fmt.Println(strings.Join(argsStr, " "))
	return 0
}

func (v *Visitor) VisitMethodCallArguments(ctx *parsing.MethodCallArgumentsContext) []interface{} {
	args := make([]interface{}, 0)
	for _, expr := range ctx.AllExpression() {
		args = append(args, v.Visit(expr))
	}
	return args
}

func (v *Visitor) VisitExpression(ctx *parsing.ExpressionContext) interface{} {
	if ctx.IDENTIFIER() != nil {
		return v.getVariableValue(ctx.IDENTIFIER().GetText())
	} else if ctx.STRING() != nil {
		return v.getStringValue(ctx.STRING().GetText())
	} else if ctx.INTEGER() != nil {
		value, err := strconv.ParseFloat(ctx.INTEGER().GetText(), 64)
		if err != nil {
			panic(err)
		}
		return value
	} else if ctx.MethodCall() != nil {
		return v.Visit(ctx.MethodCall())
	} else if ctx.SumExpr() != nil {
		return v.Visit(ctx.SumExpr())
	}
	return nil
}

func (v *Visitor) getVariableValue(key string) interface{} {
	variable, ok := v.vars[key]
	if !ok {
		panic(fmt.Sprintf("\"%s\" does not set as a variable.", key))
	}
	return variable
}

// getStringValue Remove quotation from text string
func (v *Visitor) getStringValue(str string) string {
	str = strings.ReplaceAll(str, "\"", "")
	str = strings.ReplaceAll(str, "\\n", "\n")
	return str
}

// notEnoughArgs panic an error if a method
// methodName
// args list of method arguments
// required Number of required parameters
func (v *Visitor) notEnoughArgs(methodName string, args []interface{}, required int) {
	if required > len(args) {
		panic(fmt.Sprintf("Not enough arguments for \"%s\" method.", methodName))
	}
}

func (v *Visitor) VisitStatementIf(ctx *parsing.IfStatementContext) interface{} {
	return 0
}
