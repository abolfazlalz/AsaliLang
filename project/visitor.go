package project

import (
	"asalicompiler/parsing"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"math"
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
	case *parsing.StatementDefineVariableContext:
		return v.VisitStatementDefineVariable(val)
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
	case *parsing.CallMethodContext:
		return v.VisitCallMethod(val)
	case *parsing.StatementPrintMethodContext:
		return v.VisitStatementPrintMethod(val)
	case *parsing.MethodCallArgumentsContext:
		return v.VisitMethodCallArguments(val)
	case *parsing.MethodCallContext:
		return v.VisitMethodCall(val)
	default:
		panic(fmt.Sprintf("Invalid context\nName: %s", val.GetText()))
	}
}

func (v *Visitor) VisitProgram(ctx *parsing.ProgramContext) float64 {
	for _, statement := range ctx.Statements().AllStatement() {
		v.Visit(statement)
	}
	return 1
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

func (v *Visitor) VisitStatementDefineVariable(ctx *parsing.StatementDefineVariableContext) interface{} {
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

func (v *Visitor) VisitCallMethod(ctx *parsing.CallMethodContext) interface{} {
	return v.Visit(ctx.MethodCall())
}

func (v *Visitor) VisitMethodCall(ctx *parsing.MethodCallContext) interface{} {
	methodName := ctx.IDENTIFIER().GetText()
	args := v.Visit(ctx.MethodCallArguments()).([]interface{})

	if methodName == "sin" {
		if len(args) < 1 {
			v.notEnoughArgs(methodName, args, 1)
		}
		return math.Sin(args[0].(float64))
	}
	if methodName == "cos" {
		if len(args) < 1 {
			v.notEnoughArgs(methodName, args, 1)
		}
		return math.Cos(args[0].(float64))
	}
	return nil
}

func (v *Visitor) VisitStatementPrintMethod(ctx *parsing.StatementPrintMethodContext) interface{} {
	args, ok := v.Visit(ctx.MethodCallArguments()).([]interface{})
	if !ok {
		panic("Invalid arguments types")
	}
	argsStr := make([]string, len(args))
	for i, arg := range args {
		argsStr[i] = fmt.Sprintf("%v", arg)
	}
	fmt.Println(strings.Join(argsStr, " "))
	return 0
}

func (v *Visitor) VisitMethodCallArguments(ctx *parsing.MethodCallArgumentsContext) []interface{} {
	args := make([]interface{}, 0)
	for _, expr := range ctx.AllExpression() {
		if expr.IDENTIFIER() != nil {
			variable := v.getVariableValue(expr.IDENTIFIER().GetText())
			args = append(args, variable)
		} else if expr.STRING() != nil {
			args = append(args, v.getStringValue(expr.STRING().GetText()))
		} else if expr.INTEGER() != nil {
			value, err := strconv.ParseFloat(expr.INTEGER().GetText(), 64)
			if err != nil {
				panic(err)
			}
			args = append(args, value)
		} else if expr.MethodCall() != nil {
			value := v.Visit(expr.MethodCall())
			args = append(args, value)
		}
	}
	return args
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
