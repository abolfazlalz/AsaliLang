package main

import (
	"asalicompiler/project"
	"fmt"
	"os"

	"asalicompiler/parsing"

	"github.com/antlr4-go/antlr/v4"
)

type TreeShapeListener struct {
	*parsing.BaseMyGrammarVisitor
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parsing.NewMyGrammarLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parsing.NewMyGrammarParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Program()
	//visitor.Visit(tree)
	v := project.NewVisitor()
	tree.Accept(v)
	v.Visit(tree)
}
